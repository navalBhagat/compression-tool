build-windows: 
	go build -o compression-tool.exe cmd/compression-tool/main.go
	
build-mac: 
	go build -o compression-tool cmd/compression-tool/main.go

test: 
	go test ./tests/step1

run: 
	# go run cmd/compression-tool/main.go ${file} - TODO