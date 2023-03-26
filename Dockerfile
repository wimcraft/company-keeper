FROM golang:1.19-alpine

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update && apk add --no-cache git make
COPY . ./
RUN go mod download
EXPOSE 8080

CMD ["go", "run", "main.go"]