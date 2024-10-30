package jsonhelpers

import (
	"encoding/json"
	"net/http"

	"github.com/geekible/jsonhelpers/models"
)

type JsonWriter struct {
	writer  http.ResponseWriter
	request *http.Request
	data    any
}

func NewJsonWriter(responseWriter http.ResponseWriter, request *http.Request, data any) *JsonWriter {
	return &JsonWriter{
		writer:  responseWriter,
		request: request,
		data:    data,
	}
}

func (h *JsonWriter) Write(httpStatus int) error {
	resp := models.JsonResponse{
		Data:    h.data,
		Message: "OK",
	}

	if httpStatus >= 200 || httpStatus < 300 {
		resp.IsSuccessful = true
	}

	out, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	h.writer.Header().Set("Content-Type", "application/json")
	h.writer.WriteHeader(httpStatus)

	if _, err := h.writer.Write(out); err != nil {
		return err
	}

	return nil
}
