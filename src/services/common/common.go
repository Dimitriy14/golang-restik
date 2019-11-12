package common

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dimitriy14/golang-restik/src/logger"
)

const (
	contentType     = "Content-Type"
	applicationJSON = "application/json; charset=utf-8"
)

//Error contains the message about error
type Error struct {
	Message string `json:"message"`
}

//ErrorMessage contains the error
type ErrorMessage struct {
	Error Error `json:"error"`
}

// renderJSON is used for rendering JSON response
func renderJSON(w http.ResponseWriter, status int, response interface{}) {
	data, err := json.Marshal(response)
	if err != nil {
		SendInternalServerError(w, "", fmt.Errorf("helpers.renderJSON: %s", err))
		return
	}
	render(w, status, data)
}

// RenderJSON is used for rendering JSON response body with appropriate headers
func RenderJSON(w http.ResponseWriter, response interface{}) {
	renderJSON(w, http.StatusOK, response)
}

func render(w http.ResponseWriter, code int, response []byte) {
	w.Header().Set(contentType, applicationJSON)
	w.WriteHeader(code)
	w.Write(response) //nolint
}

var createError = func(msg string) interface{} {
	return ErrorMessage{Error{msg}}
}

// SendError writes a defined string as an error message
// with appropriate headers to the HTTP response
func SendError(w http.ResponseWriter, code int, message string, err error) {
	if err != nil {
		logger.Log.Errorf("", message, "%v", err)
	}
	if message == "" {
		message = http.StatusText(code)
	}
	data, err := json.Marshal(createError(message))
	if err != nil {
		logger.Log.Errorf("", message, "helpers.SendError: %v", err)
	}
	render(w, code, data)
}

// SendInternalServerError sends Internal Server Error Status and logs an error if it exists
func SendInternalServerError(w http.ResponseWriter, message string, err error) {
	SendError(w, http.StatusInternalServerError, message, err)
}

// RenderJSONCreated is used for rendering JSON response body when new resource has been created
func RenderJSONCreated(w http.ResponseWriter, response interface{}) {
	data, err := json.Marshal(response)
	if err != nil {
		SendInternalServerError(w, "failed to marshal response", err)
		return
	}
	render(w, http.StatusCreated, data)
}
