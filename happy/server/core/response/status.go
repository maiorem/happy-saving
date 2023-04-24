package response

import "net/http"

type Status struct {
	Status  int
	Message string
}

const (
	STATUS_OK            = http.StatusOK
	UNAUTHORIZATION      = http.StatusUnauthorized
	NOT_FOUND            = http.StatusNotFound
	UNPROCESSABLE_ENTITY = http.StatusUnprocessableEntity
	SERVER_ERROR         = http.StatusInternalServerError
)

var msg = map[int]string{
	STATUS_OK:            "OK!",
	UNAUTHORIZATION:      "Unanthorized",
	NOT_FOUND:            "Not Found",
	UNPROCESSABLE_ENTITY: "Unprocessable Entity",
	SERVER_ERROR:         "Server error",
}

func Error(err int) (int, *Status) {
	status, msg := getErrorMsg(err)
	resp := NewError(status, msg)
	return status, resp
}

func getErrorMsg(code int) (int, string) {
	switch code {
	case UNAUTHORIZATION:
		status, msg := unauthorized()
		return status, msg
	case NOT_FOUND:
		status, msg := notFound()
		return status, msg
	case UNPROCESSABLE_ENTITY:
		status, msg := unprocessableEntity()
		return status, msg
	case SERVER_ERROR:
		status, msg := serverError()
		return status, msg
	default:
		status, msg := OK()
		return status, msg
	}
}
func NewError(code int, msg string) *Status {
	return &Status{
		Status:  code,
		Message: msg,
	}
}

func OK() (int, string) {
	return STATUS_OK, msg[STATUS_OK]
}

func unauthorized() (int, string) {
	return UNAUTHORIZATION, msg[UNAUTHORIZATION]
}

func notFound() (int, string) {
	return NOT_FOUND, msg[NOT_FOUND]
}

func unprocessableEntity() (int, string) {
	return UNPROCESSABLE_ENTITY, msg[UNPROCESSABLE_ENTITY]
}

func serverError() (int, string) {
	return SERVER_ERROR, msg[SERVER_ERROR]
}
