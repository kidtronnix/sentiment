package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func InternalError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "{\"statusCode\":%d,\"error\":\"%s\"}", http.StatusInternalServerError, "Internal Error!")
}

func BadRequest(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "{\"statusCode\":%d,\"error\":\"%s\",\"message\":\"%s\"}", http.StatusBadRequest, "Bad Request!", msg)
}

func JSON(w http.ResponseWriter, j interface{}) {
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(j)
	if err != nil {
		fmt.Println(err)
		InternalError(w)
		return
	}
	fmt.Fprintf(w, "%s", jsonResp)
}
