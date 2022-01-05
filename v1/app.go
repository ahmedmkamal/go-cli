package cli

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/akaahmedkamal/go-args"
)

var AppName string
var AppVersion string
var AppBuild string

type App struct {
	args *args.ArgsParser
	cmds []Command
}

func initAppInfo() {
	if AppName == "" {
		exe, err := os.Executable()
		if err != nil {
			panic(err)
		}
		AppName = filepath.Base(exe)
	}

	if AppVersion == "" {
		AppVersion = "Unspecified"
	}

	if AppBuild == "" {
		AppBuild = "Unspecified"
	}
}

func NewApp(rawArgs []string) *App {
	initAppInfo()

	parser := args.NewParser(rawArgs)
	if err := parser.Parse(); err != nil {
		log.Fatal(err)
	}

	return &App{
		args: parser,
		cmds: make([]Command, 0),
	}
}

func (a *App) Args() *args.ArgsParser {
	return a.args
}

func (a *App) Commands() []Command {
	return a.cmds
}

func (a *App) Command(name string) Command {
	for _, cmd := range a.cmds {
		if cmd.Name() == name {
			return cmd
		}
	}
	return nil
}

func (a *App) Register(cmd Command) {
	a.cmds = append(a.cmds, cmd)
}

func (a *App) Run() {
	cmdName := strings.Join(a.args.Positional(), "/")

	cmd := a.Command(cmdName)

	if cmd == nil {
		log.Fatal("nothing to do!")
	}

	cmd.Run(a)
}
