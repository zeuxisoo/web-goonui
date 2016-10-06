RELEASE_ROOT="release"
RELEASE_GOONUI="release/goonui"

NOW = $(shell date -u '+%Y%m%d%I%M%S')

usage:
	@echo "make env"

env:
	@go get github.com/Unknwon/bra

	@glide install

clean:
	@rm -rf go-goonui
	@rm -rf goonui
	@rm -rf $(RELEASE_ROOT)

dev-server: clean
	@bra run

dev-install: clean
	go install
	go build
	./go-goonui install

build:
	go install -v

	rm -rf ./goonui
	cp '$(GOPATH)/bin/go-goonui' ./goonui

	rm -rf $(RELEASE_GOONUI)
	mkdir -p $(RELEASE_GOONUI)
	mkdir -p $(RELEASE_GOONUI)/app/templates

	cp -r goonui conf public storage $(RELEASE_GOONUI)
	cp -r app/templates $(RELEASE_GOONUI)/app

	rm -rf $(RELEASE_GOONUI)/storage/database.sqlite

	cd $(RELEASE_ROOT) && zip -r goonui.$(NOW).zip "goonui"
