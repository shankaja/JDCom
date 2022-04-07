package wrap

import (
	"encoding/json"
	"log"
	"net/http"
)

// http responses wrapping
type FileResponseWrapper struct {
	FileText string
	Response http.Response
}

func WrapResponse(w http.ResponseWriter, statusCode int, status string, resp map[string]string) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	resp["status"] = status

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func WrapFileResponse(w http.ResponseWriter, text string) {
	resp := make(map[string]string)
	resp["file_text"] = text
	WrapResponse(w, http.StatusOK, "success", resp)
}
