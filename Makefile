run:
	go run main.go

build:
	go build -o prime .

test:
	go test ./...

dock:
	sudo docker build -t get-prime .