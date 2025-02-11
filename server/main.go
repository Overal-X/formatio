package main

import (
	"fmt"

	"github.com/overal-x/formatio/services"
)

func main() {
	execService := services.NewExecService()
	nixpacksService := services.NewNixpacksService(execService)

	err := nixpacksService.Install(services.InstallArgs{
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

	appDirectory := "<test-project-directory>"
	err = nixpacksService.Build(services.BuildArgs{
		AppName:      "shortlet",
		AppDirectory: appDirectory,
		Env: &map[string]string{
			"NIXPACKS_NODE_VERSION": "20",
		},
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
		AppName: "shortlet",
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
