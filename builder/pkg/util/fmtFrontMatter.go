package util

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/kjk/notionapi"
)

const descLen = 120

type frontMatter struct {
	PageID      string
	Title       string
	Description string
	Time        int64
}

func (fm *frontMatter) Check(block *notionapi.Block) {
	switch block.Type {
	case notionapi.BlockPage:
		// frontMatter.Date
		fm.Time = block.CreatedTime
		// frontMatter.Title
		fm.Title = block.Title
	case notionapi.BlockText:
		// frontMatter.Description
		ts := block.GetProperty("title")
		for _, text := range ts {
			if len(fm.Description)+len(text.Text) > descLen {
				break
			}
			fm.Description += cleanText(text.Text)
		}
	}
}

func (fm frontMatter) Print() {
	log.Println(fm.PageID)
	log.Println(fm.Title)
	log.Println(fm.Date())
	log.Println(fm.Description)
}

func (fm frontMatter) Date() string {
	return time.Unix(fm.Time/1000, 0).Format("2006-01-02")
}

func (fm frontMatter) GetFile() string {
	var ret string
	err := filepath.Walk(contentPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, fm.PageID) {
			ret = path
		}
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return ret
}

func (fm frontMatter) AddFrontMatter(in string) string {
	md := in

	reg := regexp.MustCompile(`{{\ title\ }}`)
	md = reg.ReplaceAllString(md, "\""+fm.Title+"\"")

	reg = regexp.MustCompile(`{{\ description\ }}`)
	md = reg.ReplaceAllString(md, "\""+fm.Description+"\"")

	reg = regexp.MustCompile(`{{\ date\ }}`)
	md = reg.ReplaceAllString(md, "\""+fm.Date()+"\"")

	return md
}

func fmtFrontMatter(file string, fm frontMatter) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	// add frontmatter
	noHeadMD := rmFirstH1(bytes)
	doneMD := fm.AddFrontMatter(noHeadMD)
	ioutil.WriteFile(file, []byte(doneMD), os.ModePerm)

	// mv post to folder
	createPostDir()
	extractMD(file)
}

func rmFirstH1(in []byte) string {
	buf := bytes.NewBuffer(in)
	for {
		line, err := buf.ReadBytes('\n')
		if err != nil {
			panic(err)
		}
		if string(line) == "---\n" {
			return string(line) + buf.String()
		}
	}
}

func extractMD(pathwithfile string) {
	_, file := filepath.Split(pathwithfile)
	err := os.Rename(
		filepath.Join(contentPath, file),
		filepath.Join(contentPath, postPath, file),
	)
	if err != nil {
		panic(err)
	}
	log.Println(filepath.Join(contentPath, postPath, file))
}
