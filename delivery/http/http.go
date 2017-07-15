package main

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"net/http"
)

func SetupRequestLogger(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	logger := logrus.New()
	context.Set(r, "reqlogger", logger)
	next(w, r)
}

func LogRequest(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	logger, _ := context.Get(r, "reqlogger").(logrus.StdLogger)
	logger.Printf("%+v", r)
	next(w, r)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hey!"))
	})

	n := negroni.New(
		negroni.NewRecovery(),
	)

	n.UseFunc(SetupRequestLogger)
	n.UseFunc(LogRequest)
	n.UseHandler(router)

	http.ListenAndServe(":4444", n)
}
