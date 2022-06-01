package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/Wrapper-Offline-Electron/Universal-Installer/installer"
)

const (
	INSTALL   = "install"
	UNINSTALL = "uninstall"
	UPDATE    = "update"
	RUN       = "run"
	EXIT      = "exit"
)

func main() {
	fmt.Print("Welcome!\n\n")

	choice := ""
	prompt := &survey.Select{
		Message: "Wrapper Offline Electron Universal Installer",
		Options: []string{INSTALL, UNINSTALL, UPDATE, RUN, EXIT},
	}
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		if err == terminal.InterruptErr {
			log.Fatalln("interrupted")
		} else {
			log.Fatalln("Failed to ask for options:", err)
		}
	}

	fmt.Println()

	path, err := os.UserHomeDir()

	if err != nil {
		log.Fatalln("Failed to get the user's home directory:", err)
	}

	if choice == INSTALL || choice == UNINSTALL || choice == UPDATE {
		fmt.Println("Wrapper Offline Electron will be installed in your home directory:", path)
	}

	wrapperOfflineElectronPath := path + string(os.PathSeparator) + "WrapperOfflineElectron"

	uninstall := func(force bool) bool {
		if _, err := os.Stat(wrapperOfflineElectronPath); err != nil {
			if os.IsNotExist(err) {
				fmt.Println("Wrapper Offline Electron wasn't installed.")
				return false
			}
		}

		if !force {
			uninstall := false
			prompt := &survey.Confirm{
				Message: "Are you sure you want to uninstall?",
			}
			err := survey.AskOne(prompt, &uninstall)
			if err != nil {
				log.Fatalln("Failed to prompt if they want to uninstall:", err)
			}
			if !uninstall {
				return false
			}
		}

		if err := os.RemoveAll(wrapperOfflineElectronPath); err != nil {
			log.Fatalln("Failed to uninstall Wrapper Offline Electron:", err)
		}
		return true
	}

	install := func() {
		if _, err := os.Stat(wrapperOfflineElectronPath); err != nil {
			if os.IsExist(err) {
				uninstall(true)
			}
		}
		installer.InstallWrapperOfflineElectron(path)
	}

	switch choice {
	case INSTALL:
		install()
	case UNINSTALL:
		fmt.Println("Uninstalling Wrapper Offline Electron...")
		if uninstall(false) {
			fmt.Println("Uninstalled.")
		}
	case UPDATE:
		install()
	case RUN:
		npmPath := wrapperOfflineElectronPath + string(os.PathSeparator) + installer.Node() + string(os.PathSeparator) + "bin" + string(os.PathSeparator) + "npm"
		if err := os.Chdir(wrapperOfflineElectronPath); err != nil {
			log.Fatalln("Failed to Chdir (Choose directory) to path '"+path+"':", err)
		}
		fmt.Println(">>> " + npmPath + " start")
		fmt.Println()
		installer.Exec(npmPath, "start")
	}

	fmt.Println("Press 'Enter' to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
