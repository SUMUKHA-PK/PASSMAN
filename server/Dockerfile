FROM docker.io/golang:1.12.5-alpine3.9

LABEL  maintainer = "Sumukha PK"

WORKDIR github.com/SUMUKHA-PK/PASSMAN/server

COPY . .

RUN apk add git

RUN go build -o main -v ./cmd/... 

EXPOSE 8008

ENTRYPOINT ["./main"]

CMD ["passmanServer"]