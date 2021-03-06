package glab

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"glab/cmd/glab/utils"
	"glab/commands"
	"strings"
)

var (
	version string
	build   string
	commit  string
)

func PrintVersion(_ map[string]string, _ map[int]string) {
	fmt.Println()
	fmt.Println("GLab version", version)
	fmt.Println("Build:", build)
	fmt.Println("Commit:", commit)
	fmt.Println("https://github.com/profclems/glab")
	fmt.Println()
	fmt.Println("Made with ❤ by Clement Sam <https://clementsam.tech>")
	fmt.Println()
}

func Issue(cmdArgs map[string]string, arrCmd map[int]string) {
	commands.ExecIssue(cmdArgs, arrCmd)
}

func MergeRequest(cmdArgs map[string]string, arrCmd map[int]string) {
	commands.ExecMergeRequest(cmdArgs, arrCmd)
}

func Help(args map[string]string, arrCmd map[int]string) {
	//OpenFile("./utils/help.go")
	utils.PrintHelpHelp()
}

func Config(cmdArgs map[string]string, arrCmd map[int]string) {
	cmdHelpList := map[string]string{
		"uri":   "GITLAB_URI",
		"url":   "GITLAB_URI",
		"token": "GITLAB_TOKEN",
		"repo":  "GITLAB_REPO",
		"pid":   "GITLAB_PROJECT_ID",
	}
	isUpdated := false
	if arrCmd[0] == "global" {
		commands.UseGlobalConfig = true
	}
	fmt.Println() //Upper Space
	for i := 0; i < len(arrCmd); i++ {
		if commands.CommandArgExists(cmdArgs, arrCmd[i]) && commands.CommandArgExists(cmdHelpList, arrCmd[i]) {
			commands.SetEnv(cmdHelpList[arrCmd[i]], cmdArgs[arrCmd[i]])
			isUpdated = true
		} else {
			if arrCmd[0] != "global" {
				fmt.Println(aurora.Red(arrCmd[i] + ": invalid flag"))
			}
		}
	}

	if isUpdated {
		fmt.Println(aurora.Green("Environment variable(s) updated"))
	}
	fmt.Println() //ending space
}

func Exec(cmd string, cmdArgs map[string]string, arrCmd map[int]string) {
	commandList := map[string]func(map[string]string, map[int]string){
		"issue":     Issue,
		"mr":        MergeRequest,
		"help":      Help,
		"config":    Config,
		"version":   PrintVersion,
		"--version": PrintVersion,
		"-v":        PrintVersion,
	}
	cmd = strings.Trim(cmd, " ")
	if cmd == "" {
		Help(cmdArgs, arrCmd)
	}
	if commands.CommandExists(commandList, cmd) {

		if len(cmdArgs) > 0 {
			if cmdArgs["help"] == "true" {
				cmdHelpList := map[string]func(){
					"help":  utils.PrintHelpHelp,
					"issue": utils.PrintHelpIssue,
					"mr":    utils.PrintHelpMr,
					"repo":  utils.PrintHelpRepo,
				}
				//OpenFile("./utils/"+cmd+".txt")
				cmdHelpList[cmd]()
			}
		}
		commandList[cmd](cmdArgs, arrCmd)
	} else {
		fmt.Println(cmd + ": command not found")
		fmt.Println()
		Help(cmdArgs, arrCmd)
	}
}
