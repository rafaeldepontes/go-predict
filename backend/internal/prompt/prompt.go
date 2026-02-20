package prompt

import (
	"io"
	"log"
	"os"
)

var prompt []byte

func Get() []byte {
	if len(prompt) < 1 {
		load()
	}
	return prompt
}

func ResetForTest() {
	prompt = []byte{}
}

func load() {
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
	prompt = b
}
