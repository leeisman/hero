FROM golang:1.14 as builder

WORKDIR /hero
COPY . /hero

ENV CGO_ENABLED=0
ENV GO111MODULE=on
#RUN go mod tidy && go mod download
RUN go build -mod=vendor -o hero
FROM alpine:3

RUN apk add --no-cache ca-certificates

RUN mkdir -p /hero
WORKDIR /hero
COPY --from=builder /hero/hero /hero
COPY --from=builder /hero/configs /hero/configs
ENV env=dev
# Run the web service on container startup.
RUN ls -l
CMD ["/hero/hero", "server"]

