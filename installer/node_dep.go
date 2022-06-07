package installer

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func NodeExists() error {
	if _, err := exec.LookPath("node"); err != nil {
		return err
	}
	if _, err := exec.LookPath("npm"); err != nil {
		return err
	}
	return nil
}

// path of the WrapperOfflineElectron
func InstallNodeDep(path string) {
	fmt.Println()
	fmt.Println("Installing NodeJS Dependencies...")
	fmt.Println()

	if err := os.Chdir(path); err != nil {
		log.Fatalln("Chdir path '"+path+"' error:", err)
	}

	arguments := []string{"install"}

	if runtime.GOOS == "darwin" {
		arguments = append(arguments, "--arch=x64")
	}

	if err := Exec("npm", arguments...); err != nil {
		log.Fatalln("'npm install' failed on path '"+path+"':", err)
	}

	if err := os.Chdir(path + string(os.PathSeparator) + "wrapper"); err != nil {
		log.Fatalln("Chdir path '"+path+string(os.PathSeparator)+"wrapper"+"' error:", err)
	}

	if err := Exec("npm", "install"); err != nil {
		log.Fatalln("'npm install' failed on path '"+path+string(os.PathSeparator)+"wrapper"+"':", err)
	}
}
