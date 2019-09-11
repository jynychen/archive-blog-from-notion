package main

import (
	"os"

	"github.com/Jyny/blog/builder/pkg/util"
)

func main() {
	mainPageID := os.Getenv("PAGE_ID")
	authToken := os.Getenv("AUTH_TOKEN")
	client := util.NotionClient(authToken)

	util.RmExport()
	util.NotionExport(client, mainPageID)
	util.NotionPages(client, mainPageID)
}
