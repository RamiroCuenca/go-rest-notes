package handler

import "net/http"

// Send the response to an http request.
func SendResponse(w http.ResponseWriter, status int, data []byte) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)

}

// Send an error over an http request. Writes an empty json as response.
func SendError(w http.ResponseWriter, status int) {

	data := []byte(`{}`)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)

}
