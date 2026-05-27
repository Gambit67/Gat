package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	"github.com/alecthomas/chroma/v2/quick"
)

func main() {
	checkConfig()
	if len(os.Args) < 2 {
		fmt.Println("Usage: gat <file or directory>")
		fmt.Println("Provide file or directory")
		os.Exit(1)
	}
	fileName := os.Args[1]
	if listFiles(fileName) {
		return
	}
	configFile := os.ExpandEnv("$HOME/.config/gat/config")  // Get config file to get theme
	themeBytes, err := os.ReadFile(configFile)
	theme := "dracula"  // Use default
	if err == nil {
		theme = strings.TrimSpace(string(themeBytes)) // Overwrite default if there's no err
	}
	contents, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	quick.Highlight(os.Stdout, string(contents), fileName, "terminal256", theme)
}

func checkConfig () {
	configDir := os.ExpandEnv("$HOME/.config/gat")
	configFile := configDir + "/config"

	if _, err := os.Stat(configFile); os.IsNotExist (err) {
		os.MkdirAll(configDir, 0755)
		os.WriteFile(configFile, []byte("dracula"), 0644)
	}
}

func listFiles (fileName string) bool {
	info, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}

	if info.IsDir() {
		files, err := os.ReadDir(fileName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(" \uf07b " + fileName)
		fmt.Println("─────────────────")
			for _, file := range files {
    			if file.IsDir() {
        			fmt.Println("  \uf07b " + file.Name() + "/")
       			} else {
          			fmt.Println("  \uf15b " + file.Name())
          		}
			}
		return true
	}
	return false
}
