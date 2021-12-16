.PHONY: clear build
build:
	go build -trimpath -ldflags "-w -s" -o go-ssh main.go
optime:
	upx ./go-ssh
