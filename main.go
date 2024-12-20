package main

import (
	"log/slog"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	"github.com/gitsang/protoc-gen-example/proto/protoc_gen_example/options/v1"
)

func generateOutputFile(gen *protogen.Plugin, file *protogen.File) {
	filename := file.GeneratedFilenamePrefix + ".go"

	f := gen.NewGeneratedFile(filename, file.GoImportPath)
	f.P("// Code generated by protoc-gen-example. DO NOT EDIT.")
	f.P()

	f.P("package ", file.GoPackageName)
	f.P()
}

func main() {
	slog.SetDefault(
		slog.New(
			slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}),
		))

	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			// File Options
			if fileOpts := proto.GetExtension(f.Desc.Options(), options.E_FileOptions); fileOpts != nil {
				slog.Info("File", slog.Any("path", f.Desc.Path()), slog.Any("options", fileOpts))
			}

			// Messages
			for _, msg := range f.Messages {
				// Message Options
				if msgOpts := proto.GetExtension(msg.Desc.Options(), options.E_MessageOptions); msgOpts != nil {
					slog.Info("Message", slog.Any("name", msg.Desc.Name()), slog.Any("options", msgOpts))
				}

				// Fields Options
				for _, field := range msg.Fields {
					if fieldOpts := proto.GetExtension(field.Desc.Options(), options.E_FieldOptions); fieldOpts != nil {
						slog.Info("Field", slog.Any("name", field.Desc.Name()), slog.Any("options", fieldOpts))
					}
				}
			}

			// Services
			for _, service := range f.Services {
				// Service Options
				if svcOpts := proto.GetExtension(service.Desc.Options(), options.E_ServiceOptions); svcOpts != nil {
					slog.Info("Service", slog.Any("name", service.Desc.Name()), slog.Any("options", svcOpts))
				}

				// Method Options
				for _, method := range service.Methods {
					if methodOpts := proto.GetExtension(method.Desc.Options(), options.E_MethodOptions); methodOpts != nil {
						slog.Info("Method", slog.Any("name", method.Desc.Name()), slog.Any("options", methodOpts))
					}
				}
			}

			generateOutputFile(gen, f)
		}
		return nil
	})
}
