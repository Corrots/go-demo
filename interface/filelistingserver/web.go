package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/corrots/demo/20200407/filelistingserver/filelisting"
	"github.com/gpmgo/gopm/modules/log"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

type clientError interface {
	error
	Message() string
}

func errWrapper(handler appHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func(w http.ResponseWriter) {
			if r := recover(); r != nil {
				log.Print(log.ERROR, "Panic %v\n", r)
				code := http.StatusInternalServerError
				http.Error(w, http.StatusText(code), code)
			}
		}(w)

		err := handler(w, r)
		if err != nil {
			log.Error("unexpected request err: %v\n", err)

			if userError, ok := err.(clientError); ok {
				http.Error(w, userError.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			//case os.IsTimeout(err):
			//	code = http.StatusGatewayTimeout
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.Handler))
	port := ":8080"
	fmt.Printf("Listening %s, running: %v\n", port, http.ListenAndServe(port, nil))
}
