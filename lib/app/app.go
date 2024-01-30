package app

import (
	"decr/lib/types"
	"github.com/teris-io/cli"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"regexp"
	"strings"
)

var App cli.App

func init() {
	App = cli.New("Tool for creating .desktop files").
		WithArg(cli.NewArg("path", "Path to the executable")).
		WithOption(cli.NewOption("type", "Entry type")).
		WithOption(cli.NewOption("version", "Program version")).
		WithOption(cli.NewOption("name", "Program display name")).
		WithOption(cli.NewOption("terminal", "Is a terminal program").WithType(cli.TypeBool)).
		WithAction(command)
}
func command(args []string, options map[string]string) int {
	suffixRegExp := regexp.MustCompile("\\..*")
	spaceRegExp := regexp.MustCompile("[-_. ]]")
	argv := strings.Split(args[0], "/")
	caser := cases.Title(language.English)

	exec := argv[len(argv)-1]
	path := strings.Join(argv[:len(argv)-1], "/")
	name := options["name"]
	version := options["version"]
	entryType := options["type"]
	encoding := options["encoding"]
	terminal := options["terminal"]

	if len(version) == 0 {
		version = "1.0"
	}

	if len(entryType) == 0 {
		entryType = "Application"
	}

	if len(name) == 0 {
		name = caser.String(spaceRegExp.ReplaceAllString(suffixRegExp.ReplaceAllString(exec, ""), " "))
	}

	if len(encoding) == 0 {
		encoding = "UTF-8"
	}

	if len(terminal) == 0 {
		terminal = "false"
	}

	entry := types.DesktopEntry{
		Name:     name,
		Exec:     exec,
		Version:  version,
		Path:     path,
		Encoding: encoding,
		Type:     entryType,
		Terminal: terminal,
	}

	entry.Add()

	return 0
}
