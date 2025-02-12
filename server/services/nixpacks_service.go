package services

import (
	"errors"
	"fmt"

	"github.com/samber/lo"
)

type NixpacksCallback func(out *string, err error)

type InstallArgs struct {
	Callback NixpacksCallback
}

type BuildArgs struct {
	AppName      string
	AppDirectory string
	Env          *map[string]string
	Callback     NixpacksCallback
}

type RunArgs struct {
	AppName  string
	Env      *map[string]string
	Ports    *map[string]string
	Callback NixpacksCallback
}

type INixpacksService interface {
	Install(args InstallArgs) (err error)
	Build(args BuildArgs) (err error)
	Run(args RunArgs) (err error)
}

type NixpacksService struct {
	execHelper IExecService
}

func (n *NixpacksService) Install(args InstallArgs) (err error) {
	return n.execHelper.Execute(ExecuteArgs{
		Command: "curl -sSL https://nixpacks.com/install.sh | bash",
		OutputCallback: func(s string) {
			args.Callback(&s, nil)
		},
		ErrorCallback: func(s string) {
			args.Callback(nil, errors.New(s))
		},
	})
}

func (n *NixpacksService) Build(args BuildArgs) (err error) {
	command := fmt.Sprintf("nixpacks build %s --name %s", args.AppDirectory, args.AppName)

	if args.Env != nil {
		env := ""
		lo.ForEach(lo.Keys(*args.Env), func(k string, _ int) {
			env += fmt.Sprintf("%s=%s ", k, (*args.Env)[k])
		})
		command += fmt.Sprintf(` --env "%s"`, env)
	}

	return n.execHelper.Execute(ExecuteArgs{
		Command: command,
		OutputCallback: func(s string) {
			args.Callback(&s, nil)
		},
		ErrorCallback: func(s string) {
			args.Callback(&s, nil)
			// BUG: status code from exec returns 1 even if it's a successful build
			// args.Callback(nil, errors.New(s))
		},
	})
}

func (n *NixpacksService) Run(args RunArgs) (err error) {
	command := "docker run -t -d"

	if args.Env != nil {
		lo.ForEach(lo.Keys(*args.Env), func(k string, _ int) {
			command += fmt.Sprintf(" -e %s=\"%s\"", k, (*args.Env)[k])
		})
	}

	if args.Ports != nil {
		lo.ForEach(lo.Keys(*args.Ports), func(k string, _ int) {
			command += fmt.Sprintf(" -p %s:%s", k, (*args.Ports)[k])
		})
	}

	command += fmt.Sprintf(" %s", args.AppName)

	return n.execHelper.Execute(ExecuteArgs{
		Command: command,
		OutputCallback: func(s string) {
			args.Callback(&s, nil)
		},
		ErrorCallback: func(s string) {
			args.Callback(nil, errors.New(s))
		},
	})
}

func NewNixpacksService(execHelper IExecService) INixpacksService {
	return &NixpacksService{execHelper: execHelper}
}
