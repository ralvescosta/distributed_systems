FROM golang:alpine as build

WORKDIR /build
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build --ldflags "-s -w" -o exec main.go

FROM scratch

WORKDIR /app
USER app
COPY --from=build ./build/exec ./
COPY ./.env.* ./
EXPOSE 3333

CMD ["./exec"]