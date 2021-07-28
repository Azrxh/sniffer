build:
	@go build -ldflags="-w -s" main.go
	upx main
