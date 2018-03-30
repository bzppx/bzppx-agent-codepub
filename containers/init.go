package containers

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	version = "v0.8.1"
)

// init container
func init()  {
	poster()
	initFlag()
	initConfig()
	initLog()
	go Workers.Task()
}

// poster logo
func poster() {
	fg := color.New(color.FgHiYellow).SprintFunc()
	logo := fg(`
                    __                     __                                     __
  _____ ____   ____/ /___   ____   __  __ / /_       ____    ____   ___   ____   / /_
 / ___// __ \ / __  // _ \ / __ \ / / / // __ \     / __ \  / __ \ / _ \ / __ \ / __/
/ /__ / /_/ // /_/ //  __// /_/ // /_/ // /_/ / -- / /_/ / / /_/ //  __// / / // /_
\___/ \____/ \__,_/ \___ / .___/ \__,_//_.___/     \__/\_\ \___,/ \___//_/ /_/ \__/
                        /_/                               /____/
`+
"Author: bzppx\r\n"+
"Vserion: "+version+"\r\n"+
"Link: github.com/bzppx/bzppx-agent-codepub")
	fmt.Println(logo)
}