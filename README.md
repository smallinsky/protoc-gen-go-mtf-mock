# protoc-gen-go-mtf-mock

#### Usage:
Install package:
```shell 
go install github.com/smallinsky/protoc-gen-go-mtf-mock
```

Pass  `--go-mtf-mock_out=.` flag to the `protoc` proto generator with proto dir:
```shell 
protoc --go_out=plugins=grpc:. --go-mtf-mock_out=. ./proto/weather/*.proto
```
