FROM golang

COPY . /build
WORKDIR /build

RUN go build -o prime .
CMD ["prime"]