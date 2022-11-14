FROM golang:1.19-alpine as builder

RUN apk add git

LABEL maintainer="<>"

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

RUN mkdir /app

WORKDIR /app/

COPY --from=builder /app/main .
COPY --from=builder /app/app.env .

EXPOSE 8000

CMD ["./main"]