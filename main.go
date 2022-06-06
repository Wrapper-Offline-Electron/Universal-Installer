package main

import (
	"bufio"
	_ "embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/Wrapper-Offline-Electron/Universal-Installer/installer"
	"github.com/inancgumus/screen"
)

const (
	INSTALL   = "install"
	UNINSTALL = "uninstall"
	UPDATE    = "update"
	RUN       = "run"
	VERSION   = "version"
	EXIT      = "exit"
)

//go:embed VERSION
var UNIVERSAL_INSTALLER_VERSION string

func checkSurvey(err error) {
	if err != nil {
		if err == terminal.InterruptErr {
			log.Fatalln("interrupted")
		} else {
			log.Fatalln("Failed to ask for survey:", err)
		}
	}
}

func main() {
	fmt.Print("Welcome!\n\n")

	choice := ""
	prompt := &survey.Select{
		Message: "Wrapper Offline Electron Universal Installer",
		Options: []string{INSTALL, UNINSTALL, UPDATE, RUN, VERSION, EXIT},
	}
	err := survey.AskOne(prompt, &choice)
	checkSurvey(err)

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

		if _, err := os.Stat(wrapperOfflineElectronPath + ".zip"); err == nil {
			if err := os.Remove(wrapperOfflineElectronPath + ".zip"); err != nil {
				log.Fatalln("Failed to remove '"+wrapperOfflineElectronPath+".zip':", err)
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
		// check line 150 too
		// WINDOWS ONLY FOR NOW
		if runtime.GOOS == "windows" {
			if _, err := exec.LookPath("node"); err != nil {
				fmt.Fprintln(os.Stderr, "Please install NodeJS, error:", err)
				return
			}
			if _, err := exec.LookPath("npm"); err != nil {
				fmt.Fprintln(os.Stderr, "NPM (Node Package Manager) comes along with NodeJS, it's not in the PATH, error:", err)
				return
			}
		}
		if _, err := os.Stat(wrapperOfflineElectronPath); err == nil {
			uninstall(true)
		}
		installer.InstallWrapperOfflineElectron(path)
	}

	switch choice {
	case INSTALL:
		ifInstall := false
		prompt := &survey.Confirm{
			Message: "Are you sure you want to install?",
		}
		err := survey.AskOne(prompt, &ifInstall)
		if err != nil {
			log.Fatalln("Failed to prompt if they want to install:", err)
		}
		if ifInstall {
			install()
		}
	case UNINSTALL:
		fmt.Println("Uninstalling Wrapper Offline Electron...")
		if uninstall(false) {
			fmt.Println("Uninstalled.")
		}
	case UPDATE:
		update := false
		prompt := &survey.Confirm{
			Message: "Are you sure you want to update?",
		}
		err := survey.AskOne(prompt, &update)
		if err != nil {
			log.Fatalln("Failed to prompt if they want to update:", err)
		}
		if update {
			install()
		}
	case RUN:
		// WINDOWS ONLY FOR NOW
		if runtime.GOOS == "windows" {
			if _, err := exec.LookPath("node"); err != nil {
				fmt.Fprintln(os.Stderr, "Please install NodeJS, error:", err)
				break
			}
			if _, err := exec.LookPath("npm"); err != nil {
				fmt.Fprintln(os.Stderr, "NPM (Node Package Manager) comes along with NodeJS, it's not in the PATH, error:", err)
				break
			}
		}

		_, err := os.Stat(wrapperOfflineElectronPath)
		if os.IsNotExist(err) {
			fmt.Println("Wrapper Offline Electron isn't installed. Please install it to run the app.")
		} else {
			npmPath := wrapperOfflineElectronPath + string(os.PathSeparator) + installer.Node()
			if runtime.GOOS != "windows" {
				npmPath += string(os.PathSeparator) + "bin"
			}
			npmPath += string(os.PathSeparator) + "npm"
			if err := os.Chdir(wrapperOfflineElectronPath); err != nil {
				log.Fatalln("Failed to Chdir (Choose directory) to path '"+path+"':", err)
			}
			// WINDOWS USERS EXPECTED TO INSTALL NODEJS FOR NOW
			if runtime.GOOS == "windows" {
				npmPath = "npm"
			}
			fmt.Println(">>> " + npmPath + " start")
			fmt.Println()
			installer.Exec(npmPath, "start")
		}
	case VERSION:
		choices := []string{"Wrapper Offline Electron (installed)", "Universal Installer"}
		choice := ""
		prompt := &survey.Select{
			Message: "Which version do you want to check?",
			Options: choices,
		}
		err := survey.AskOne(prompt, &choice)
		if err != nil {
			log.Fatalln("Failed to ask prompt for the version:", err)
		}

		if choice == choices[0] {
			_, err := os.Stat(wrapperOfflineElectronPath)
			if os.IsNotExist(err) {
				fmt.Println("Wrapper Offline Electron isn't installed. Please install it to check the version for the Wrapper Offline Electron.")
			} else {
				packageJsonPath := wrapperOfflineElectronPath + string(os.PathSeparator) + "package.json"
				packageContents, err := ioutil.ReadFile(packageJsonPath)
				if err != nil {
					log.Fatalln("Error when opening '"+packageJsonPath+"' package.json file:", err)
				}

				var packageJson map[string]interface{}
				err = json.Unmarshal(packageContents, &packageJson)
				if err != nil {
					log.Fatalln("Error when using json.Unmarshal(), parsing packageContents of '"+packageJsonPath+"' path:", err)
				}

				fmt.Println("Wrapper Offline Electron installed version", packageJson["version"].(string))
			}
		} else if choice == choices[1] {
			fmt.Println("Universal Installer version", UNIVERSAL_INSTALLER_VERSION)
		} else {
			fmt.Println("Invalid choice for checking version")
		}
		fmt.Println()

	case EXIT:
		os.Exit(0)
	}

	fmt.Println("Press 'Enter' to go back...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	screen.Clear()
	main()
}
