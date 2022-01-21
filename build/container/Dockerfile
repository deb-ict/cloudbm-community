ARG RUNTIME_BASE_IMAGE=scratch

# Build
FROM golang:1.17.6-alpine3.15 as build
WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/cloudbm cmd/webhost/main.go

# Runtime
FROM scratch
WORKDIR /app
COPY --from=build /app/cloudbm .
COPY configs/*.yaml /etc/cloudbm/

EXPOSE 80
ENTRYPOINT [ "/app/cloudbm" ]