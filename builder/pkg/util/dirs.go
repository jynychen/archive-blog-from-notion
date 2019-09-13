package util

import (
	"os"
	"path/filepath"
	"strings"
)

func createContentDir() {
	os.MkdirAll(
		contentPath,
		os.ModePerm,
	)
}

func createPostDir() {
	os.MkdirAll(
		filepath.Join(contentPath, postPath),
		os.ModePerm,
	)
}

func createMediaDir() {
	os.MkdirAll(
		filepath.Join(contentPath, mediaPath),
		os.ModePerm,
	)
}

func cleanText(in string) string {
	ret := in
	ret = strings.ReplaceAll(
		ret,
		"'",
		"",
	)
	ret = strings.ReplaceAll(
		ret,
		"\"",
		"",
	)
	ret = strings.ReplaceAll(
		ret,
		"\n",
		" ",
	)
	return ret
}
