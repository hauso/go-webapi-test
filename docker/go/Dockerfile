FROM golang:1.22-alpine3.19

RUN mkdir /app
WORKDIR /app
COPY /api/. .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o webapi

# not running, require login to console and run manualy
CMD ["/webapi"]
