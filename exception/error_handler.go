package exception

import (
	"auth/helper"
	"auth/model/web"
	"log"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if badRequest(writer, request, err) {
		return
	}

	if notFound(writer, request, err) {
		return
	}

	if unauthorized(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func badRequest(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(BadRequestError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		response := web.Response{
			Status: "BAD_REQUEST",
			Data:   exception.Error,
		}
		// Example of logging the request path or method for debugging
		log.Printf("Internal server error at: %s %s", request.Method, request.URL.Path)
		helper.WriteToBody(writer, response)
		return true
	} else {
		return false
	}
}

func notFound(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		response := web.Response{
			Status: "NOT_FOUND",
			Data:   exception.Error,
		}
		// Example of logging the request path or method for debugging
		log.Printf("Internal server error at: %s %s", request.Method, request.URL.Path)
		helper.WriteToBody(writer, response)
		return true
	} else {
		return false
	}
}

func unauthorized(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(UnauthorizedError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		response := web.Response{
			Status: "UNAUTHORIZED",
			Data:   exception.Error,
		}
		// Example of logging the request path or method for debugging
		log.Printf("Internal server error at: %s %s", request.Method, request.URL.Path)
		helper.WriteToBody(writer, response)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)

	response := web.Response{
		Status: "BAD_REQUEST",
		Data:   err,
	}
	// Example of logging the request path or method for debugging
	log.Printf("Internal server error at: %s %s", request.Method, request.URL.Path)
	helper.WriteToBody(writer, response)

}
