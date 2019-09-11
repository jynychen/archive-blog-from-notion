package util

import (
	"log"

	"github.com/kjk/notionapi"
)

func NotionClient(authToken string) *notionapi.Client {
	client := &notionapi.Client{}
	client.AuthToken = authToken
	return client
}

func getPages(client *notionapi.Client, mainPageID string) []string {
	page, err := client.DownloadPage(mainPageID)
	if err != nil {
		panic(err)
	}
	return page.GetSubPages()
}

func getPage(client *notionapi.Client, pageID string) notionapi.Page {
	page, err := client.DownloadPage(pageID)
	if err != nil {
		panic(err)
	}
	return *page
}

func NotionPages(client *notionapi.Client, mainPageID string) {
	// list subpage in Posts database page
	pages := getPages(client, mainPageID)
	for _, pageID := range pages {
		fm := frontMatter{PageID: pageID}
		page := getPage(client, pageID)
		page.ForEachBlock(func(block *notionapi.Block) {
			fm.Check(block)
		})

		mdfile := fm.GetFile()
		log.Println(mdfile)
		fmtMediaLink(mdfile)
		fmtFrontMatter(mdfile, fm)
	}
}
