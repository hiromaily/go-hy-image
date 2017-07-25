genjson:
	gogentype -file $(PWD)/jsons/config.json

bld:
	go build -i -v -o ${GOPATH}/bin/gohy main.go

run:
	go run main.go -j $(PWD)/jsons/config.json
