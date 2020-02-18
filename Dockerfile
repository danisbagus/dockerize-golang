FROM danisbagus/base-go

RUN mkdir -p /go

COPY /app /go

RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/joho/godotenv

WORKDIR /go/src

CMD ["go","run", "main.go"]
