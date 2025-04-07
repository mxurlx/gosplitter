package main

import (
	"fmt"
	"os"

	"github.com/mxurlx/gosplitter/cmd"
	"github.com/mxurlx/gosplitter/common"

	"github.com/mxurlx/flagit"
)

func main() {
	//err := flagit.GenFiles(common.Flags)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	subcmd, flags, mandatoryArgs, err := flagit.ParseFlags(os.Args, common.Flags)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = flagit.ExecuteCmd(subcmd, flags, mandatoryArgs, cmd.CmdFuncs)
	if err != nil {
		fmt.Println(err)
		return
	}
}
