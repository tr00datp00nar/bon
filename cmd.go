package bon

import (
	_ "embed"

	"fmt"
	"strings"
	"text/template"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:        `bon`,
	Aliases:     []string{`bonzai`},
	Summary:     help.S(_bon),
	Description: help.D(_bon),
	Version:     `v0.1.3`,
	Copyright:   `Copyright 2022 Rob Muhlestein`,
	Site:        `rwxrob.tv`,
	Source:      `git@github.com:rwxrob/bon.git`,
	Issues:      `https://github.com/rwxrob/bon/issues`,
	License:     `Apache-2.0`,

	Commands: []*Z.Cmd{
		branchreadmeCmd,
		depsCmd,
		fullCommandCmd,
		gobadgesCmd,
		quickCommandCmd,
		help.Cmd,
	},
}

var fullCommandCmd = &Z.Cmd{
	Name:        `full`,
	Summary:     help.S(_full),
	Description: help.D(_full),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {

		if len(args) == 0 {
			args = append(args, `foo`)
		}

		t, err := template.New("t").Parse(fullcommandtmpl)
		if err != nil {
			return err
		}
		var buf strings.Builder
		if err := t.Execute(&buf, struct{ Name string }{args[0]}); err != nil {
			return err
		}
		fmt.Print(buf.String())
		return nil
	},
}

var quickCommandCmd = &Z.Cmd{
	Name:        `quick`,
	Summary:     help.S(_quick),
	Description: help.D(_quick),
	Commands:    []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {

		if len(args) == 0 {
			args = append(args, `foo`)
		}

		t, err := template.New("t").Parse(quickcommandtmpl)
		if err != nil {
			return err
		}
		var buf strings.Builder
		if err := t.Execute(&buf, struct{ Name string }{args[0]}); err != nil {
			return err
		}
		fmt.Print(buf.String())
		return nil
	},
}

var depsCmd = &Z.Cmd{
	Name:        `deps`,
	Summary:     help.S(_deps),
	Description: help.D(_deps),
	Commands:    []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {

		if len(args) == 0 {
			args = append(args, `foo`)
		}

		t, err := template.New("t").Parse(depstmpl)
		if err != nil {
			return err
		}
		var buf strings.Builder
		if err := t.Execute(&buf, struct{ Name string }{args[0]}); err != nil {
			return err
		}
		fmt.Print(buf.String())
		return nil
	},
}

var branchreadmeCmd = &Z.Cmd{
	Name:        `branchreadme`,
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_branchreadme),
	Description: help.D(_branchreadme),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {

		if len(args) == 0 {
			args = append(args, `foo`)
		}

		t, err := template.New("t").Parse(branchreadme)
		if err != nil {
			return err
		}
		var buf strings.Builder
		if err := t.Execute(&buf, struct{ Name string }{args[0]}); err != nil {
			return err
		}
		fmt.Print(buf.String())
		return nil
	},
}

var gobadgesCmd = &Z.Cmd{
	Name:        `gobadges`,
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_gobadges),
	Description: help.D(_gobadges),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {

		if len(args) == 0 {
			args = append(args, `foo`)
		}

		t, err := template.New("t").Parse(gobadges)
		if err != nil {
			return err
		}
		var buf strings.Builder
		if err := t.Execute(&buf, struct{ Name string }{args[0]}); err != nil {
			return err
		}
		fmt.Print(buf.String())
		return nil
	},
}
