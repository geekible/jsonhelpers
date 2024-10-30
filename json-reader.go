package jsonhelpers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type JsonReader struct {
	writer   http.ResponseWriter
	request  *http.Request
	data     any
	maxBytes int64
}

func NewJsonReader(responseWriter http.ResponseWriter, request *http.Request, data any) *JsonReader {
	return &JsonReader{
		writer:   responseWriter,
		request:  request,
		data:     data,
		maxBytes: 1048576,
	}
}

func (h *JsonReader) Read(data any) error {
	h.request.Body = http.MaxBytesReader(h.writer, h.request.Body, int64(h.maxBytes))
	dec := json.NewDecoder(h.request.Body)
	err := dec.Decode(data)

	if err != nil {
		return err
	}
	err = dec.Decode(&struct{}{})

	if err != io.EOF {
		return errors.New("body must have only a single JSON value")
	}

	return nil
}
