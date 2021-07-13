run:
	go run main.go

build:
	go build -o prime .

test:
	go test ./...

dock-ui:
	sudo docker build -t get-prime-ui ./web/get-prime/

dock-server:
	sudo docker build -t get-prime .