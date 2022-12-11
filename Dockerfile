FROM golang:1.19.4-alpine
COPY . ~/app
WORKDIR ~/app
RUN go get github.com/labstack/echo/v4
RUN go get github.com/maxatome/go-testdeep
RUN go get github.com/go-sql-driver/mysql
RUN go mod tidy
ENTRYPOINT go run main.go
