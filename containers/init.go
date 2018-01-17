package containers

import (
	"fmt"
	"github.com/fatih/color"
)

// init container
func init()  {
	poster()
	initConfig()
	initLog()
	go Workers.Task()
}

// poster logo
func poster() {
	fg := color.New(color.FgHiYellow).SprintFunc()
	logo := fg(`
 █████╗      ██████╗     ███████╗    ███╗   ██╗    ████████╗    ██╗  ██╗
██╔══██╗    ██╔════╝     ██╔════╝    ████╗  ██║    ╚══██╔══╝    ╚██╗██╔╝
███████║    ██║  ███╗    █████╗      ██╔██╗ ██║       ██║        ╚███╔╝
██╔══██║    ██║   ██║    ██╔══╝      ██║╚██╗██║       ██║        ██╔██╗
██║  ██║    ╚██████╔╝    ███████╗    ██║ ╚████║       ██║       ██╔╝ ██╗
╚═╝  ╚═╝     ╚═════╝     ╚══════╝    ╚═╝  ╚═══╝       ╚═╝       ╚═╝  ╚═╝
Author: bzppx
Link  : https://github.com/bzppx/bzppx-codepub
`)
	fmt.Println(logo)
}