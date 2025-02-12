package main

import (
	"fmt"

	"github.com/overal-x/formatio/services"
)

func main() {
	execService := services.NewExecService()
	nixpacksService := services.NewNixpacksService(execService)
	fileService := services.NewFileService()

	err := fileService.Unzip(services.UnzipArgs{
		ZipFile:     "<path/to/zip/file.zip",
		Destination: "./_temp",
	})
	if err != nil {
		panic(err)
	}
	defer func() {
		err := fileService.Remove(services.RemoveArgs{File: "./_temp"})
		if err != nil {
			panic(err)
		}
	}()

	err = nixpacksService.Install(services.InstallArgs{
		Callback: func(out *string, err error) {
			if err != nil {
				panic(err)
			}
			fmt.Println(*out)
		},
	})
	if err != nil {
		panic(err)
	}

	appDirectory := "./_temp"
	err = nixpacksService.Build(services.BuildArgs{
		AppName:      "my-app",
		AppDirectory: appDirectory,
		Callback: func(out *string, err error) {
			if err != nil {
				panic(err)
			}
			if out != nil {
				fmt.Println(*out)
			}
		},
	})
	if err != nil {
		panic(err)
	}

	err = nixpacksService.Run(services.RunArgs{
		AppName: "my-app",
		Env:     &map[string]string{"PORT": "3000"},
		Ports:   &map[string]string{"3000": "3000"},
		Callback: func(out *string, err error) {
			if err != nil {
				panic(err)
			}
			fmt.Println(*out)
		},
	})
	if err != nil {
		panic(err)
	}
}
