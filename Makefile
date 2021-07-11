run:
	go run main.go

build:
	go build -o prime .

dock:
	sudo docker build -t get-prime .