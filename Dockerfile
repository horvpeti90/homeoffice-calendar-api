FROM golang:1.16 AS builder

ENV GO111MODULE="on"
ENV CGO_ENABLED=0
ENV GOOS="linux"
ENV GOARCH="amd64"
ENV GOSUMDB="off"

RUN apt-get update
RUN apt-get install git

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "10001" \
    "user"

WORKDIR $GOPATH/src/github.mpi-internal.com/hu/homeoffice-calendar-api
COPY . .

RUN go mod vendor
RUN go mod verify
RUN go build \
	--ldflags "-X main.build=$(git rev-parse HEAD) -X main.version=$(git rev-parse --abbrev-ref HEAD) -w -s" \
	-o /go/bin/homeoffice-calendar-api \
	cmd/main.go

FROM scratch

COPY --from=builder /etc/passwd /etc/group /etc/
USER user:user

COPY --from=builder /go/bin/homeoffice-calendar-api /go/bin/homeoffice-calendar-api

ENTRYPOINT ["/go/bin/homeoffice-calendar-api"]
