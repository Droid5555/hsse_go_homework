package decode

import (
	"encoding/base64"
	"encoding/json"
	"hsse_go_homework/task2/tools/decode_tools"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var input decode_tools.Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	decodedString, err := base64.StdEncoding.DecodeString(input.InputString)
	if err != nil {
		http.Error(w, "400 Bad Request : (Decode Error)", http.StatusBadRequest)
		return
	}

	output := decode_tools.Output{
		OutputString: string(decodedString),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, "500 Internal Server Error : (Json Encode Problem)", http.StatusInternalServerError)
		return
	}
}
