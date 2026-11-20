package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/Geekzix/calorie-snap-commons/common/constant"
)

type ApiHandler func(w http.ResponseWriter, r *http.Request) error

func GlobalHandler(fn ApiHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			handleResErr(w, err)
		}
	}
}

func handleResErr(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	msg := "internal server error"

	switch {
	case errors.Is(err, constant.ErrNotFound):
		status = http.StatusNotFound
		msg = err.Error()
	case errors.Is(err, constant.ErrAlreadyExit):
		status = http.StatusConflict
		msg = err.Error()
	case errors.Is(err, constant.ErrInvalidInput):
		status = http.StatusBadRequest
		msg = err.Error()
	default:
		slog.Error("unhandled error", "err", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func RespondJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
