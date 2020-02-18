FROM danisbagus/base-go

RUN mkdir -p /go /go/src /go/bin/ go/pkg

COPY /app/ /go/src

RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/joho/godotenv

WORKDIR /go/src

CMD ["go","run", "main.go"]
