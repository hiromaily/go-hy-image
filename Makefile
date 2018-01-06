genjson:
	gogentype -file $(PWD)/jsons/config.json

bld:
	go build -i -race -v -o ${GOPATH}/bin/gohy main.go

run:
	go run -race main.go -j $(PWD)/jsons/config.json
