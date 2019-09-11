build: extract
	mv export/post/* site/content/post/
	mv export/media/* site/content/media/
	@cd site ; \
	hugo ; \
	cd .. ; \
	mv site/public ./

extract: builder
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
	rm builder/builder
	rm -r export
	rm -r public