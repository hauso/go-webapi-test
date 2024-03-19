FROM golang:latest

#RUN apt-get update
#RUN apt-get install -y htop
#
#RUN apt update
#RUN apt install net-tools

RUN mkdir /app
WORKDIR /app


COPY /api/. .
#COPY .env .

#ADD ./api/src/go.mod .
#ADD ./api/src/go.sum .
#ADD .env .

RUN go mod download

RUN go env GOPATH=/app
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
RUN go install gorm.io/driver/mysql gorm.io/gorm
RUN go install github.com/joho/godotenv
RUN go install github.com/golang-jwt/jwt

#EXPOSE 8000

ENTRYPOINT CompileDaemon --build="go build webapi" --command=./webapi
#RUN #CGO_ENABLED=0 go build -o /app .

#go build webapi
