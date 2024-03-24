FROM golang:1.22.1
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "executable" ]
WORKDIR /
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod tidy && go mod download && go mod verify

COPY . .
RUN go build -v -o / / ...

CMD ["maps _cron"]