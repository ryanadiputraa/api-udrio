package cmd

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ryanadiputraa/api-udrio/pkg/database"
	"github.com/spf13/viper"
)

func serveHTTP() {
	_ = database.GetConnection()

	// router
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})

	// setup server port & handler
	port := viper.GetString("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}
	fmt.Printf("server running on port %s", port)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
