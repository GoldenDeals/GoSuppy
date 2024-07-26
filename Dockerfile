FROM golang:1.22.5-alpine AS build
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates

WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build -ldflags="-extldflags=-static -s -w" -o /app .

FROM scratch AS final
COPY --from=build /app /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY ./cert.pem /cert.pem
COPY ./key.pem /key.pem

ENTRYPOINT ["/app"]
