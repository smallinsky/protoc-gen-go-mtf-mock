package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"

	"github.com/smallinsky/protoc-gen-go-mtf-mock/assets"
)

var (
	PackageGoImport = make(map[string]string)
	NameGoImport    = make(map[string]string)
)

type FileDesc struct {
	Package  string
	Imports  []string
	Services []ServiceDesc
	File     string
}

type ServiceDesc struct {
	Name    string
	Methods []MethodsDesc
}

type MethodsDesc struct {
	Name       string
	InputType  string
	OutputType string
}

func NamesToGoImports(ss []string) []string {
	var out []string
	for _, s := range ss {
		if v, ok := NameGoImport[s]; ok {
			out = append(out, v)
		}
	}
	return out
}

func initializeImports(req *plugin.CodeGeneratorRequest) {
	for _, pf := range req.GetProtoFile() {
		PackageGoImport[pf.GetPackage()] = pf.GetOptions().GetGoPackage()
		NameGoImport[pf.GetName()] = pf.GetOptions().GetGoPackage()
	}
}

func CodeGenerator(req *plugin.CodeGeneratorRequest) *plugin.CodeGeneratorResponse {
	filesDescriptor := BuildFilesDescriptor(req)
	tmpl := template.Must(template.New("").Parse(assets.MustAssetString("mock.tpl")))
	var resp plugin.CodeGeneratorResponse
	for _, desc := range filesDescriptor {
		if len(desc.Services) == 0 {
			continue
		}
		var buff bytes.Buffer
		if err := tmpl.Execute(&buff, &desc); err != nil {
			panic(err)
		}
		content := buff.String()
		resp.File = append(resp.File, &plugin.CodeGeneratorResponse_File{
			Content: &content,
			Name:    &desc.File,
		})
	}
	return &resp
}

func contains(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

func BuildFilesDescriptor(req *plugin.CodeGeneratorRequest) []FileDesc {
	initializeImports(req)

	var filesDesc []FileDesc
	for _, p := range req.GetProtoFile() {
		fileDesc := FileDesc{
			Package: p.GetPackage(),
		}
		fileDesc.Imports = append(fileDesc.Imports, NamesToGoImports(p.GetDependency())...)

		for _, s := range p.GetService() {
			svc := ServiceDesc{
				Name: s.GetName(),
			}
			for _, m := range s.GetMethod() {
				svc.Methods = append(svc.Methods, MethodsDesc{
					Name:       m.GetName(),
					InputType:  MessageType(m.GetInputType(), p.GetDependency(), req.GetProtoFile()),
					OutputType: MessageType(m.GetOutputType(), p.GetDependency(), req.GetProtoFile()),
				})
			}
			fileDesc.Services = append(fileDesc.Services, svc)
			if contains(req.GetFileToGenerate(), p.GetName()) {
				fileDesc.File = p.GetName() + ".mock.go"
			}
		}
		filesDesc = append(filesDesc, fileDesc)
	}
	return filesDesc
}

func MessageType(name string, deps []string, protos []*descriptor.FileDescriptorProto) string {
	ss := strings.Split(name, ".")[1:]
	tp := ss[len(ss)-1]
	pkg := strings.Join(ss[:len(ss)-1], ".")

	for _, p := range protos {
		for _, dep := range deps {
			if *p.Name != dep || *p.Package != pkg {
				continue
			}
			i := strings.LastIndex(p.GetOptions().GetGoPackage(), "/")
			return fmt.Sprintf("%s.%s", p.GetOptions().GetGoPackage()[i+1:], tp)
		}
	}
	return tp
}
