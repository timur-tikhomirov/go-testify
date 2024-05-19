package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// проверим, что при корректном запросе сервер вернет 200 и тело ответа не будет пустым
func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1&city=moscow", nil) // создаем запрос к сервису
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	// проверим, что сервер возвращает 200
	require.Equal(t, responseRecorder.Code, http.StatusOK)
	// проверим, что тело ответа не пустое
	require.NotEmpty(t, responseRecorder.Body.String())
}

// проверим,что если город в запросе не поддерживается, сервер вернет 400 и ошибку wrong city value
func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1&city=wrong", nil) // создаем запрос к сервису
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	// проверим, что сервер возвращает 400
	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	// проверим, что в теле ответа "wrong city value"
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}

// проверим, что если запросить больше элементов, чем есть в базе, метод вернет все элементы
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // создаем запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// проверим, что сервер возвращает 200
	require.Equal(t, responseRecorder.Code, http.StatusOK)
	// проверим, что тело ответа не пустое
	require.NotEmpty(t, responseRecorder.Body.String())
	// проверим, что если запросить больше элементов, чем есть в базе, метод вернет все элементы
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	assert.Equal(t, len(list), totalCount)

}
