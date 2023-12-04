package model

import (
	"fmt"
	"log"
	"net/http"
)

type MessageBoardError struct {
	Message string
	Code    int
}

func (e MessageBoardError) Error() string {
	return fmt.Sprintf("%s (Code: %d)", e.Message, e.Code)
}

func FatalErrorHandler(err error, operationName string) MessageBoardError {
	log.Fatalf("Error during handling the : %s %s", operationName, err.Error())
	return MessageBoardError{
		Message: "internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func NotFoundErrHanlder(keyInof string, valueInfo string, operationName string) MessageBoardError {
	log.Printf("no record found for %s: %s during handling %s ", keyInof, valueInfo, operationName)
	return MessageBoardError{
		Message: "record not found",
		Code:    http.StatusNotFound,
	}
}

func DuplicateHanlder(info string, operationName string) MessageBoardError {
	log.Printf("duplicate recored found for %s during handling %s ", info, operationName)
	return MessageBoardError{
		Message: "duplicate record",
		Code:    http.StatusBadRequest,
	}
}
