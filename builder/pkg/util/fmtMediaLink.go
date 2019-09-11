package util

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func fmtMediaLink(file string) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	// regexp find path `!()[path]`
	reg := regexp.MustCompile(`\!\[\]\((.+)\)`)

	// mv media to folder
	createMediaDir()
	for _, m := range reg.FindAllString(string(bytes), -1) {
		media := strings.TrimLeft(m, "![](")
		media = strings.TrimRight(media, ")")
		extractMedia(media)
	}

	// fix media link to MD
	md := reg.ReplaceAllString(string(bytes), "![](/"+mediaPath+"/$1)")
	ioutil.WriteFile(file, []byte(md), os.ModePerm)
}

func extractMedia(file string) {
	err := os.Rename(
		filepath.Join(contentPath, file),
		filepath.Join(contentPath, mediaPath, file),
	)
	if err != nil {
		panic(err)
	}
	log.Println(filepath.Join(contentPath, mediaPath, file))
}
