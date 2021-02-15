package {{ .Package }}

import (
    "runtime"
    "net"
    "time"
    "context"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

    {{range .Imports}}"{{.}}"{{end}}
)

type grpcHandler func(context.Context, interface{}) (interface{}, error)
type grpcResp struct {
    msg interface{}
    err error
}

func handle(ctx context.Context, req interface{}, handler grpcHandler) (interface{}, error) {
    resultC := make(chan grpcResp)
    exitC := make(chan struct{})
    go func() {
        defer func() { exitC <- struct{}{} }()
        resp, err := handler(ctx, req)
        resultC <- grpcResp{msg: resp, err: err}
    }()

    select {
        case r := <-resultC:
            return r.msg, r.err
        case <-exitC:
            return nil, status.New(codes.Internal, "test failed").Err()
        case <-time.Tick(time.Second * 5):
            return nil, status.New(codes.Unavailable, "test timeout").Err()
    }
}

func returnIfTestFailed(err error) {
    if err == nil {
        return
    }
    st, ok := status.FromError(err)
    if !ok {
        return
    }
    if st.Err().Error() == "rpc error: code = Internal desc = test failed" {
        runtime.Goexit()
    }
}
{{ range .Services }}
    {{ $ServiceName := .Name}}
func New{{ .Name }}ClientMock(target string) {{ .Name }}Client {
    conn, err := grpc.Dial(target, grpc.WithInsecure())
    if err != nil {
        panic(err)
    }
    return &{{ .Name}}ClientMock{
        cc: New{{ $ServiceName }}Client(conn),
    }
}

func New{{ $ServiceName }}ServerMock() *{{ $ServiceName }}ServiceMock {
    ln, err := net.Listen("tcp", ":0")
    if err != nil {
        panic(err)
    }

    s := grpc.NewServer()

    mock := &{{ $ServiceName}}ServiceMock{
        addr: ln.Addr().String(),
        server: &Mock{{ $ServiceName }}Server{
{{- range .Methods }}
            {{ .Name }}Handler: make(chan Handle{{ .Name }}),
{{- end }}
        },
    }

    Register{{ $ServiceName }}Server(s, mock.server)
    go func() {
        if err := s.Serve(ln); err != nil {
            panic(err)
        }
    }()
    return mock
}

type Mock{{ $ServiceName }}Server struct {
{{- range .Methods }}
    {{ .Name }}Handler chan Handle{{ .Name}}
{{- end }}
}

{{ range .Methods }}
type Handle{{ .Name}} func(ctx context.Context, req *{{ .InputType }}) (*{{ .OutputType}}, error)

func (s *Mock{{ $ServiceName}}Server) {{.Name}} (ctx context.Context, req *{{ .InputType }}) (*{{ .OutputType}}, error) {
    select {
        case handler := <-s.{{ .Name }}Handler:
            fn  := func(ctx context.Context, req interface{}) (interface{}, error) {
                return handler(ctx, req.(*{{.InputType}}))
            }
            resp, err := handle(ctx, req, fn)
            if err != nil {
                return nil, err
            }
            return resp.(*{{ .OutputType}}), nil
        case <-time.Tick(time.Second * 5):
            return nil, status.New(codes.Unavailable, "{{ .Name }} timeout").Err()
    }
}
{{- end }}

type {{ $ServiceName }}ServiceMock struct {
    addr   string
    server *Mock{{ $ServiceName }}Server
}

func (s *{{ $ServiceName }}ServiceMock) Addr() string {
    return s.addr
}

{{ range .Methods }}
func (s *{{ $ServiceName }}ServiceMock) {{ .Name }}(fn Handle{{ .Name }}) {
    go func() {
        select {
            case s.server.{{ .Name }}Handler <- fn:
        return
            case <-time.Tick(time.Second * 5):
        return
        }
    }()
}
{{end}}

type {{ $ServiceName }}ClientMock struct {
    cc {{ $ServiceName}}Client
}
{{ range .Methods }}
func (s *{{ $ServiceName }}ClientMock) {{ .Name }}(ctx context.Context, in *{{ .InputType }}, opts ...grpc.CallOption) (*{{ .OutputType}}, error) {
    resp, err := s.cc.{{ .Name }}(ctx, in, grpc.WaitForReady(true))
    returnIfTestFailed(err)
    return resp, err
}
{{end}}
{{end}}
