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

dist: build
	rm -rf dist/win
	rm -f dist/*.zip
	mkdir dist/win
	cp ${BINARY_NAME} dist/win
	cp README.md dist/win
	cp LICENSE.md dist/win
	7z a dist/mouse-plunger.win64.zip ./dist/win/*
