package installer

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

const (
	DARWIN_NODE        = "node-v16.15.0-darwin-x64"
	WINDOWS_32BIT_NODE = "node-v16.15.0-win-x86"
	WINDOWS_64BIT_NODE = "node-v16.15.0-win-x64"
	LINUX_NODE         = "node-v16.15.0-linux-x64"
)

func Exec(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Node() string {
	switch runtime.GOOS {
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			return WINDOWS_64BIT_NODE
		case "386":
			return WINDOWS_32BIT_NODE
		default:
			panic("Installer doesn't support Windows Arch: " + runtime.GOARCH)
		}
	case "darwin":
		return DARWIN_NODE
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			return LINUX_NODE
		default:
			panic("Installer doesn't support Linux Arch: " + runtime.GOARCH)
		}
	default:
		panic("Installer doesn't support OS: " + runtime.GOOS)
	}
}

// Install nodejs to that path, returns the path to NodeJS
func InstallNode(path string) {
	pathToNode := path + string(os.PathSeparator) + "node"

	pathToActualNode := ""

	switch runtime.GOOS {
	case "windows":
		nodeName := ""

		switch runtime.GOARCH {
		case "amd64":
			nodeName = WINDOWS_64BIT_NODE
		case "386":
			nodeName = WINDOWS_32BIT_NODE
		default:
			panic("Unsupported Windows Arch: " + runtime.GOARCH)
		}

		Install(fmt.Sprint("https://nodejs.org/dist/latest-v16.x/", nodeName, ".zip"), pathToNode)

		pathToActualNode = path + string(os.PathSeparator) + nodeName

		if err := Unzip(pathToNode, pathToActualNode); err != nil {
			log.Fatalln("Failed to unzip file '"+pathToNode+":", err)
		}
	case "darwin", "linux":
		installArgs := make([]interface{}, 0, 3) // Allocates, size of 3
		installArgs = append(installArgs, "https://nodejs.org/dist/latest-v16.x/")
		if runtime.GOOS == "linux" {
			installArgs = append(installArgs, LINUX_NODE)
			installArgs = append(installArgs, ".tar.xz")
		} else if runtime.GOOS == "darwin" {
			installArgs = append(installArgs, DARWIN_NODE)
			installArgs = append(installArgs, ".tar.gz")
		}
		Install(fmt.Sprint(installArgs...), pathToNode)
		if err := exec.Command("tar", "-xvf", pathToNode, "-C", path).Run(); err != nil {
			log.Fatalln("Error executing 'tar -xvf "+pathToNode+" -C "+path+"' command on "+runtime.GOOS+":", err)
		}
		pathToActualNode = path + string(os.PathSeparator) + installArgs[1].(string)
	}

	if err := os.Remove(pathToNode); err != nil {
		log.Fatalln("Failed to remove path '"+pathToNode+"' from system:", err)
	}

	if pathToActualNode == "" {
		panic("Bug in the Wrapper Offline Electron Universal Installer - There should have been a path returned for NodeJS")
	}

	if err := os.Chdir(path); err != nil {
		log.Fatalln("Failed to change to directory '"+path+"':", err)
	}

	npmPath := pathToActualNode + string(os.PathSeparator) + "bin" + string(os.PathSeparator) + "npm"
	if runtime.GOOS == "windows" {
		npmPath += ".exe"
	}

	fmt.Println()
	fmt.Println("Installing NPM (NodeJS) packages for Wrapper Offline Electron...")
	fmt.Println()

	npmInstallArguments := []string{"install"}

	if runtime.GOOS == "darwin" {
		npmInstallArguments = append(npmInstallArguments, "--arch=x64")
	}

	if err := Exec(npmPath, npmInstallArguments...); err != nil {
		log.Fatalln("Failed to install NPM packages for Wrapper Offline Electron:", err)
	}

	if err := os.Chdir(path + string(os.PathSeparator) + "wrapper"); err != nil {
		log.Fatalln("Failed to change to directory up a level 'wrapper' of '"+path+"':", err)
	}

	fmt.Println()
	fmt.Println("Installing NPM (NodeJS) packages FOR Wrapper itself...")
	fmt.Println()

	if err := Exec(npmPath, npmInstallArguments...); err != nil {
		log.Fatalln("Failed to install NPM packages for Wrapper Offline Electron the Wrapper itself:", err)
	}

	fmt.Println()
	fmt.Println("Successfully installed Wrapper Offline Electron!")
}
