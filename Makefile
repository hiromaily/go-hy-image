.PHONY: gen-json
gen-json:
	gogentype -file $(PWD)/jsons/config.json

.PHONY: build
build:
	go build -v -o ${GOPATH}/bin/gohy main.go

.PHONY: run
run:
	go run main.go -j $(PWD)/jsons/config.json
