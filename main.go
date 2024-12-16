package main

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	annotationspb "github.com/gitsang/protoc-gen-example/proto/protoc-gen-example/options"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			// 处理文件级别的 Options
			if fileOpts := proto.GetExtension(f.Desc.Options(), annotationspb.E_FileOptions); fileOpts != nil {
				fmt.Printf("File %s Options: %s\n", f.Desc.Path(), fileOpts)
			}

			// 处理 Messages
			for _, msg := range f.Messages {
				// Message Options
				if msgOpts := proto.GetExtension(msg.Desc.Options(), annotationspb.E_MessageOptions); msgOpts != nil {
					fmt.Printf("Message %s Options: %s\n", msg.Desc.Name(), msgOpts)
				}

				// Fields Options
				for _, field := range msg.Fields {
					if fieldOpts := proto.GetExtension(field.Desc.Options(), annotationspb.E_FieldOptions); fieldOpts != nil {
						fmt.Printf("Field %s Options: %s\n", field.Desc.Name(), fieldOpts)
					}
				}
			}

			// 处理 Services
			for _, service := range f.Services {
				// Service Options
				if svcOpts := proto.GetExtension(service.Desc.Options(), annotationspb.E_ServiceOptions); svcOpts != nil {
					fmt.Printf("Service %s Options: %s\n", service.Desc.Name(), svcOpts)
				}

				// Method Options
				for _, method := range service.Methods {
					if methodOpts := proto.GetExtension(method.Desc.Options(), annotationspb.E_MethodOptions); methodOpts != nil {
						fmt.Printf("Method %s Options: %s\n", method.Desc.Name(), methodOpts)
					}
				}
			}
		}
		return nil
	})
}

// GenerateFile 是一个可选的方法，用于生成输出文件（如果需要）
func generateOutputFile(gen *protogen.Plugin, filename string, content []byte) *protogen.GeneratedFile {
	g := gen.NewGeneratedFile(filename, protogen.FilePath(""))
	g.Write(content)
	return g
}
