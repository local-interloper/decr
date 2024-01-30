package types

import (
	"decr/lib/util"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type DesktopEntry struct {
	Name     string
	Version  string
	Encoding string
	Type     string
	Path     string
	Exec     string
	Terminal string
}

func (e DesktopEntry) ToString() string {
	rows := []string{
		"[Desktop Entry]",
		fmt.Sprintf("Name=%s", e.Name),
		fmt.Sprintf("Version=%s", e.Version),
		fmt.Sprintf("Encoding=%s", e.Encoding),
		fmt.Sprintf("Type=%s", e.Type),
		fmt.Sprintf("Path=%s", e.Path),
		fmt.Sprintf("Exec=%s", e.Exec),
		fmt.Sprintf("Terminal=%s", e.Terminal),
		"",
	}

	return strings.Join(rows, "\n")
}

func (e DesktopEntry) Write(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)

	if err != nil {
		return err
	}

	if _, err = file.WriteString(e.ToString()); err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	return nil
}

func (e DesktopEntry) Add() {
	home, err := os.UserHomeDir()
	util.Check(err)
	applications := home + "/.local/share/applications/"
	path := applications +
		regexp.MustCompile("[-_. ]").ReplaceAllString(strings.ToLower(e.Name), "-") +
		".desktop"

	util.Check(e.Write(path))
}
