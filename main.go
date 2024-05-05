package main

import (
	"context"
	"crud-golang/constrain"
	"crud-golang/customers"
	"crud-golang/database"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		constrain.ConsoleLog.Info().Msgf("Connect DB Error: %v", err.Error())
	}
	customers := customers.NewICustomer(db)
	port := 4000
	server := &http.Server{
		Handler:      jsonResponseMiddleware(routes(customers)),
		Addr:         fmt.Sprintf(":%v", port),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	done := make(chan os.Signal, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			constrain.ConsoleLog.Fatal().Msgf("listen: %s\n", err)

		}
	}()
	constrain.ConsoleLog.Info().Msgf("Server Started with port %v", port)
	<-done
	constrain.ConsoleLog.Info().Msg("Server Stoped")
	defer func() {
		db.Close()
		cancel()
	}()
	if err := server.Shutdown(ctx); err != nil {
		constrain.ConsoleLog.Fatal().Msgf("Server Shutdown Failed:%+v", err)
	}
	constrain.ConsoleLog.Info().Msg("Server Exited Properly")
}

func routes(customers *customers.Customers) *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/customers", customers.Create).Methods(http.MethodPost)
	route.HandleFunc("/customers", customers.Update).Methods(http.MethodPut)
	route.HandleFunc("/customers/{userId}", customers.GetById).Methods(http.MethodGet)
	route.HandleFunc("/customers/{userId}", customers.DeleteById).Methods(http.MethodDelete)
	return route
}

func jsonResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
