BINARY_NAME=mouse-plunger.exe

build:
	go mod tidy
	rsrc -ico build/icon.ico -o rsrc.syso
	CGO_ENABLED=1 GOARCH=amd64 GOOS=windows go build -ldflags "-s -w -H=windowsgui" -o ${BINARY_NAME}

run:
	CGO_ENABLED=1 GOARCH=amd64 GOOS=windows go build ./main.go

clean:
	go clean
	rm -f ${BINARY_NAME}
	rm -f rsrc.syso