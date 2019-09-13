build: extract
	@echo hugo build site to public/
	@if [ -e export/post/* ]; then mv export/post/* site/content/post/ ; fi
	@if [ -e export/media/* ]; then mv export/media/* site/content/media/ ; fi
	@if [ -e public ]; then rm -r public ; fi
	@cd site ; \
	hugo ; \
	cd .. ; \
	mv site/public ./

extract: builder
	@echo extract from Notion.so
	@if [ -e export ]; then rm -r export ; fi
	@builder/builder

builder: mod fmt
	@cd builder ; \
	go build -o builder . ; \
	cd ..

fmt:
	@cd builder ; \
	gofmt -w ./ ; \
	cd ..

mod:
	@cd builder ; \
	go mod download ; \
	cd ..

.PHONY: clean
clean:
	if [ -e builder/builder ]; then rm builder/builder ; fi
	if [ -e export ]; then rm -r export ; fi
	if [ -e public ]; then rm -r public ; fi
