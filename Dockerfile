FROM golang

COPY . /build
WORKDIR /build

#RUN go build -o prime .
EXPOSE 6060
#CMD ["prime"]

CMD ["go", "run" , "main.go"]