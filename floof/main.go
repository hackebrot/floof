package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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

//createGist create a Gist based on the given data
func createGist(fileNames []string, description string, public bool) *Gist {
	c := make(chan *GistFile)
	go loadFiles(fileNames, c)

	gistFiles := make(map[string]GistFile)
	for g := range c {
		gistFiles[g.Name] = *g
	}

	gist := &Gist{Description: description, Public: public, Files: gistFiles}
	return gist
}

func main() {
	description := flag.String(
		"description",
		"Floof Gist",
		"A description of the gist.")

	public := flag.Bool(
		"public",
		false,
		"Indicates whether the gist is public. (default false)")

	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		log.Fatal("No files given")
		return
	}

	userConfig, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	gist := createGist(files, *description, *public)

	gistURL, err := gist.Post(userConfig.Gist.Username, userConfig.Gist.Token)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(gistURL)
}
