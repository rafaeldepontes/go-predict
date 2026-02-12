package prompt

import (
	"io"
	"log"
	"os"
)

var prompt string

func Get() *string {
	if prompt == "" {
		load()
	}
	return &prompt
}

func load() {
	// For now this is only a test... this should be inside the OS etc...
	//
	// And also should be a environment variable instead of hard coded.
	f, err := os.Open(os.Getenv("FILE_PATH"))
	if err != nil {
		log.Println("[ERROR] Could not load the prompt...", err)
		return
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		log.Println("[ERROR] Could not read the prompt...", err)
		return
	}
	prompt = string(b)
}
