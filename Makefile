usage:
	@echo "make env"

env:
	@go get github.com/Unknwon/bra

	@glide install

clean:
	@rm -rf go-goonui

dev-server: clean
	@bra run

dev-install: clean
	go install
	go build
	./go-goonui install
