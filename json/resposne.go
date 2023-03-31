package json

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ResponseSuccess for success message response
type ResponseSuccess struct {
	Status  int         `json:"status"`  // status code
	Message string      `json:"message"` // respone message
	Data    interface{} `json:"data"`    // response data
}

type ResponseCreateSuccess struct {
	Status  int    `json:"status"`  // status code
	Message string `json:"message"` // respone message
	Id      int    `json:"id"`      // response data
}

// ResponseSuccessHealth for success health response
type ResponseHealth struct {
	Health string `json:"health"`
}

// ResponseError for error message response
type ResponseError struct {
	Status  int    `json:"status"`  // status code
	Message string `json:"message"` // respone message
	Error   string `json:"error"`   // response error message
}

// Success writes success header and send json response
func Success(c echo.Context, data interface{}) error {
	response := ResponseSuccess{Status: http.StatusOK, Message: "success", Data: data}
	return c.JSON(http.StatusOK, response)
}

// Success writes success header and send json response
func SuccessCreate(c echo.Context, id int) error {
	response := ResponseCreateSuccess{Status: http.StatusOK, Message: "success", Id: id}
	return c.JSON(http.StatusOK, response)
}

// Error writes error header and send json response
func Error(c echo.Context, err error) error {
	if err.Error() == "validate: Token is expired" || err.Error() == "user_not_logged_in" {
		response := ResponseError{Status: 403, Message: "error", Error: err.Error()}
		return c.JSON(403, response)
	}
	if err.Error() == "user_not_signed_in" {
		response := ResponseError{Status: http.StatusUnauthorized, Message: "error", Error: err.Error()}
		return c.JSON(http.StatusUnauthorized, response)
	}
	if err.Error() == "user not found" {
		response := ResponseError{Status: http.StatusNotFound, Message: "error", Error: err.Error()}
		return c.JSON(http.StatusNotFound, response)
	}
	response := ResponseError{Status: http.StatusInternalServerError, Message: "error", Error: err.Error()}
	return c.JSON(http.StatusInternalServerError, response)
}

// Health serves server health
func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseHealth{Health: "good"})
}
