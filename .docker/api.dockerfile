FROM golang as builder

WORKDIR /go/src/github.com/looped-dev/cms

COPY . .

RUN go get .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# deployment image
FROM alpine:latest
RUN apk --no-cache add ca-certificates

LABEL author="Maina Wycliffe"

WORKDIR /root/
COPY --from=builder /go/src/github.com/looped-dev/cms/app .

CMD [ "./app" ]

EXPOSE 8080
