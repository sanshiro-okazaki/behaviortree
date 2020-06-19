FROM golang AS builder

WORKDIR /go/src/app
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 go build -o bt ./

FROM alpine
COPY --from=builder /go/src/app/bt /usr/local/bin/bt
ENTRYPOINT ["/usr/local/bin/bt"]
