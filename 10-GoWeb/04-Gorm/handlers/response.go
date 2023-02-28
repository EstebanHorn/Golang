package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sendData(res http.ResponseWriter, data interface{}, status int) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)

	output, _ := json.Marshal(&data)
	fmt.Fprintln(res, string(output))

}

func sendError(res http.ResponseWriter, status int) {
	res.WriteHeader(status)
	fmt.Fprintln(res, "ERROR! resource not found")

}
