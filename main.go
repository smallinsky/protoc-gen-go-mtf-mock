package main

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/codegenerator"
)

var (
	importPrefix = flag.String("import_prefix", "", "prefix to be added to go package paths for imported proto files")
	file         = flag.String("file", "-", "where to load data from")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	f := os.Stdin
	if *file != "-" {
		var err error
		f, err = os.Open(*file)
		if err != nil {
			glog.Fatal(err)
		}
	}
	req, err := codegenerator.ParseRequest(f)
	if err != nil {
		glog.Fatal(err)
	}
	resp := CodeGenerator(req)
	for _, v := range resp.GetFile() {
		err := ioutil.WriteFile(v.GetName(), []byte(v.GetContent()), 0666)
		if err != nil {
			glog.Fatal(err)
		}
	}
}
