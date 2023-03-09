## Build
FROM golang:1.20.1-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build ./cmd/user-crud/main.go ./cmd/user-crud/di.go

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app/main /user-crud

EXPOSE 1234

USER nonroot:nonroot

ENTRYPOINT ["/user-crud"]
