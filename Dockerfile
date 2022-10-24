FROM golang

WORKDIR /app

COPY . . 

RUN go build -o TCPChat cmd/server/main.go

EXPOSE 8989

CMD ["/app/TCPChat"]



