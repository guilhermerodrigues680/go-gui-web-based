default:
	./pkger
	GOOS=linux go build -v -o go-gui-web-based-linux
	docker build -t pkger:go-gui-web-based .
	docker run -p 3000:3000 pkger:go-gui-web-based

build-mac:
	./pkger
	GOOS=darwin go build -v -o go-gui-web-based-mac

build-linux:
	./pkger
	GOOS=linux go build -v -o go-gui-web-based-linux

build-windows:
	./pkger
	GOOS=windows GOARCH=amd64 go build -v -o go-gui-web-based.exe
