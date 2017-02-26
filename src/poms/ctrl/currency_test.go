package ctrl

import (
	"testing"
	"encoding/json"
	"net/http"
	"poms/model"
)

var capturedData []byte

func TestGetCurrencies(t *testing.T) {
	// arrange
	var controller CurrencyController
	var responseWriter mockResponseWriter
	responseWriter.header = make(map[string][]string)

	// act
	controller.GetCurrencies(responseWriter, &http.Request{})

	// assert
	// sets content-type Header
	if responseWriter.Header().Get("Content-Type") != "application/json" {
		t.Error("Missing or incorrect Content-Type header")
	}
	// writes correct data to client
	currencies := model.GetCurrencies()
	currencyData, _ := json.Marshal(currencies)
	if string(capturedData) != string(currencyData) {
		t.Error("Incorrect data sent to client")
	}
}

type mockResponseWriter struct {
	header http.Header
}

func (mrw mockResponseWriter) Header() http.Header {
	return mrw.header
}

func (mrw mockResponseWriter) Write(data []byte) (int, error) {
	capturedData = data[:]
	return len(data), nil
}

func (mrw mockResponseWriter) WriteHeader(code int) {
	// note: need by http response interface
}
