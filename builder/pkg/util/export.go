package util

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/kjk/notionapi"
)

func NotionExport(client *notionapi.Client, pageID string) {
	createContentDir()
	zipByte, err := client.ExportPages(
		pageID,
		notionapi.ExportTypeMarkdown,
		true,
	)
	if err != nil {
		panic(err)
	}

	zipReader, err := zip.NewReader(
		bytes.NewReader(zipByte),
		int64(len(zipByte)),
	)
	if err != nil {
		panic(err)
	}

	for _, zipFile := range zipReader.File {
		unzippedFileBytes, err := readZipFile(zipFile)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(zipFile.Name)
		err = ioutil.WriteFile(
			filepath.Join(contentPath, zipFile.Name),
			unzippedFileBytes,
			0644,
		)
		if err != nil {
			log.Println(err)
			continue
		}
	}

}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
