package http

type HttpResponse struct {
	Message string
	Status  int
	Data    any
}

func JSONResponse(message string, status int, data any) HttpResponse {
	return HttpResponse{Message: message, Status: status, Data: data}
}
