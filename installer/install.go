package installer

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/briandowns/spinner"
)

func Install(url, pathOut string) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	fmt.Printf("Installing from '%s'\n", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("Failed to get url '"+url+"':", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalln("Response Status Code for installing/getting '"+url+"' is not successful:", resp.StatusCode)
	}

	// Create file
	out, err := os.Create(pathOut)
	if err != nil {
		log.Fatalln("Failed to create '"+pathOut+"' file:", err)
	}
	defer out.Close()

	// Write body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalln("Failed to write url '"+url+"' response body to '"+pathOut+"' file:", err)
	}

	s.Stop()
}
