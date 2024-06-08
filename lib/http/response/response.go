package response

import (
	"encoding/json"
	"net/http"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/errors"
)

type Response struct {
	Code    int          `json:"code"`
	Status  string       `json:"status"`
	Message *string      `json:"message,omitempty"`
	Data    *interface{} `json:"data,omitempty"`
}

func Json(w http.ResponseWriter, httpCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	res := Response{
		Code:    httpCode,
		Status:  http.StatusText(httpCode),
		Message: &message,
		Data:    &data,
	}
	json.NewEncoder(w).Encode(res)
}

func Text(w http.ResponseWriter, httpCode int, message string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(httpCode)
	w.Write([]byte(message))
}

func Err(w http.ResponseWriter, err error) {
	_, ok := err.(*errors.RespError)
	if !ok {
		err = errors.InternalServerError(err.Error())
	}

	er, _ := err.(*errors.RespError)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(er.Code)
	res := Response{
		Code:    er.Code,
		Status:  http.StatusText(er.Code),
		Message: &er.Message,
	}
	json.NewEncoder(w).Encode(res)
}
