package main

import (
	//nolint
	"github.com/perfectbuii/gen-go-template/cmd/protoc-gen-custom/internal"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	opt := protogen.Options{}
	internal.Run(opt)
}
