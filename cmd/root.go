package cmd

import (
	"fmt"

	"github.com/mxurlx/gosplitter/common"

	"github.com/mxurlx/flagit"
)

func printVersion() {
	fmt.Println("gosplitter: v0.1")
}

func Root(flags map[string]any, mandatoryArgs []string) error {
	if flags["help"] == true {
		flagit.PrintHelp(".", common.Flags)
	}
	if flags["version"] == true {
		printVersion()
	}
	return nil
}
