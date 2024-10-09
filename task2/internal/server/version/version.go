package version

import (
	"fmt"
	"hsse_go_homework/task2/tools/version_tools"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	version := fmt.Sprintf("v%d.%d.%d", version_tools.VERSION.Major, version_tools.VERSION.Minor, version_tools.VERSION.Patch)
	_, err := w.Write([]byte(version))
	if err != nil {
		http.Error(w, "500 Internal Server Error : (Response Write Problem)", http.StatusInternalServerError)
		return
	}
}
