package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("articlePost", articlePost)
}

func articlePost(w http.ResponseWriter, r *http.Request) {
	// Set CORS
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set Cors Header for main
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, peda.GCFPostHandler("MONGODATA", "berkatauto", "articleSet"))
}
