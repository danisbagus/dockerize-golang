package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/danisbagus/dockerize-golang/internal/handler"
	"github.com/danisbagus/dockerize-golang/internal/repo"
	"github.com/danisbagus/dockerize-golang/internal/usecase"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mysqlClient := GetMysqlClient()

	router := mux.NewRouter()

	transactionRepo := repo.NewTransactionRepo(mysqlClient)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionUsecase)

	router.HandleFunc("/api/transactions", transactionHandler.GetAllTransaction).Methods(http.MethodGet)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	appPort := ":8100"
	fmt.Println("Starting the application at:", appPort)
	log.Fatal(http.ListenAndServe(appPort, router))

}

func GetMysqlClient() *sqlx.DB {
	dbHost := "myappdb"
	dbPort := "3306"
	dbUser := "root"
	dbPassword := "mypass"
	dbName := "myapp"

	connection := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", dbUser, dbPassword, dbHost, dbPort, dbName)
	client, err := sqlx.Open("mysql", connection)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
