package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/danisbagus/dockerize-golang/internal/handler"
	"github.com/danisbagus/dockerize-golang/internal/repo"
	"github.com/danisbagus/dockerize-golang/internal/usecase"
	"github.com/spf13/viper"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	InitializeViper()

	mysqlClient := GetMysqlClient()

	router := mux.NewRouter()

	transactionRepo := repo.NewTransactionRepo(mysqlClient)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionUsecase)

	router.HandleFunc("/api/transactions", transactionHandler.GetAllTransaction).Methods(http.MethodGet)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	appAdrr := fmt.Sprintf("%s:%s", viper.GetString("app.host"), viper.GetString("app.port"))

	fmt.Println("Starting the application at:", appAdrr)
	log.Fatal(http.ListenAndServe(appAdrr, router))

}

func GetMysqlClient() *sqlx.DB {
	dbHost := viper.GetString("db.host")
	dbPort := viper.GetString("db.port")
	dbUser := viper.GetString("db.user")
	dbPassword := viper.GetString("db.password")
	dbName := viper.GetString("db.name")

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

func InitializeViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(fmt.Sprintf("Error while reading file %s", err))
	}
}
