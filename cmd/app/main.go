package main

import (
	"car-catalog/internal/db"
	"car-catalog/internal/handler"
	"fmt"
	"github.com/MadAppGang/httplog"
	"net/http"
	"os"
)

var (
	carsHandler   http.Handler = http.HandlerFunc(handler.InsertCars)
	deleteHandler http.Handler = http.HandlerFunc(handler.DeleteCar)
	getHandler    http.Handler = http.HandlerFunc(handler.GetWFilters)
	updateHandler http.Handler = http.HandlerFunc(handler.UpdateCars)
)

func main() {
	db.InitDb()
	mux := http.NewServeMux()

	loggerWithFormatter := httplog.LoggerWithFormatter(httplog.DefaultLogFormatterWithRequestHeader)
	mux.Handle("POST /cars", loggerWithFormatter(carsHandler))
	mux.Handle("DELETE /cars", loggerWithFormatter(deleteHandler))
	mux.Handle("GET /cars", loggerWithFormatter(getHandler))
	mux.Handle("PUT /cars", loggerWithFormatter(updateHandler))

	corsHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	err := http.ListenAndServe(":8080", corsHandler(mux))
	if err != nil {
		if err == http.ErrServerClosed {
			fmt.Println("server closed")
		} else {
			fmt.Printf("error starting server: %s\n", err)
			os.Exit(1)
		}
	}
}
