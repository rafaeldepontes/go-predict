package tool

import (
	"log"
	"os"
)

func ChecksEnv(src *string) {
	if _, err := os.Stat(*src); err == nil {
		return
	}
	log.Printf("[WARN] File %s not found, using .env.example instead.\n", *src)
	*src = ".env.example"
	if _, err := os.Stat(*src); err == nil {
		return
	}
	log.Println("[WARN] Both files don't exist in the current directory.")
}
