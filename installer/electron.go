package installer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

func InstallWrapperOfflineElectron(path string) {
	pathToZip := path + string(os.PathSeparator) + "WrapperOfflineElectron.zip"
	pathToDir := path + string(os.PathSeparator) + "WrapperOfflineElectron"

	fmt.Println("Installing Wrapper Offline Electron... (This may take a while, please wait.)")

	Install("https://github.com/jackprogramsjp/Wrapper-Offline-Electron/archive/refs/heads/main.zip", pathToZip)

	if err := Unzip(pathToZip, pathToDir); err != nil {
		panic("A bug in the universal installer: " + err.Error())
	}

	pathToFiles := pathToDir + string(os.PathSeparator) + "Wrapper-Offline-Electron-main"
	files, err := ioutil.ReadDir(pathToFiles)
	if err != nil {
		log.Fatalln("Failed to read directory '"+pathToFiles+"':", err)
	}

	for _, f := range files {
		if err := os.Rename(pathToFiles+string(os.PathSeparator)+f.Name(), pathToDir+string(os.PathSeparator)+f.Name()); err != nil {
			log.Fatalln("Failed to uncompress Wrapper Offline Electron main directory:", err)
		}
	}

	if err := os.RemoveAll(pathToFiles); err != nil {
		log.Fatalln("Failed to remove whole directory '"+pathToFiles+"':", err)
	}

	if err := os.Remove(pathToZip); err != nil {
		log.Fatalln("Failed to remove zip file '"+pathToZip+"':", err)
	}

	// WINDOWS USERS EXPECTED TO INSTALL NODEJS FOR NOW
	if runtime.GOOS == "windows" {
		// temporarily
		if err := os.Chdir(path); err != nil {
			log.Fatalln("Chdir path '"+path+"' error:", err)
		}

		if err := Exec("npm", "install"); err != nil {
			log.Fatalln("'npm install' failed on path '"+path+"':", err)
		}

		if err := os.Chdir(path + string(os.PathSeparator) + "wrapper"); err != nil {
			log.Fatalln("Chdir path '"+path+string(os.PathSeparator)+"wrapper"+"' error:", err)
		}

		if err := Exec("npm", "install"); err != nil {
			log.Fatalln("'npm install' failed on path '"+path+string(os.PathSeparator)+"wrapper"+"':", err)
		}
	} else {
		InstallNode(pathToDir)
	}
}
