package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//loadFiles creates a GistFile each of which holding the contents of a file
func loadFiles(fileNames []string, c chan *GistFile) {
	for _, name := range fileNames {
		content, err := ioutil.ReadFile(name)
		if err == nil {
			c <- &GistFile{Name: name, Content: string(content)}
		}
	}
	close(c)
}

func main() {
	userConfig, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userConfig.Gist.Username)
	fmt.Println(userConfig.Gist.Token)

	description := os.Args[1]
	public := os.Args[2]
	files := os.Args[3:]
	fmt.Println(description)
	fmt.Println(public)
	fmt.Println(files)

	c := make(chan *GistFile)
	go loadFiles(files, c)

	gistFiles := make(map[string]GistFile)
	for g := range c {
		gistFiles[g.Name] = *g
	}

	gist := &Gist{Description: description, Public: false, Files: gistFiles}

	gistURL, err := gist.Post(userConfig.Gist.Username, userConfig.Gist.Token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(gistURL)
}
