package handler

import (
	"fmt"
	"net/http"

	"github.com/yodfhafx/go-bank/errs"
)

func handleError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case errs.AppError:
		w.WriteHeader(e.Code)
		fmt.Fprintln(w, e)
	case error:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, e)
	}
}
