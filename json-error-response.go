package jsonhelpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/geekible/jsonhelpers/models"
)

type JsonErrorResponse struct {
	writer     http.ResponseWriter
	err        error
	errSource  string
	httpStatus int
}

func NewJsonErrorResponse(writer http.ResponseWriter, err error, errSource string, httpStatus int) *JsonErrorResponse {
	return &JsonErrorResponse{
		writer:     writer,
		err:        err,
		errSource:  errSource,
		httpStatus: httpStatus,
	}
}

func (h *JsonErrorResponse) ReturnError() error {
	if h.httpStatus == 0 {
		h.httpStatus = http.StatusInternalServerError
	}

	resp := models.JsonResponse{
		IsSuccessful: false,
		Message:      fmt.Sprintf("error occurred at %s: err: %s", h.errSource, h.err.Error()),
	}

	out, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	h.writer.Header().Set("Content-Type", "application/json")
	h.writer.WriteHeader(h.httpStatus)

	if _, err := h.writer.Write(out); err != nil {
		return err
	}

	return nil
}
