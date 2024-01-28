package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	berkatbepkg "github.com/berkatauto/berkat-bepkg"
)

func init() {
	functions.HTTP("HelloPostArticle", HelloPostArticle)
}

func HelloPostArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Max-Age", "3600")
	if r.Method == "OPTIONS" {
		// Jika permintaan adalah preflight OPTIONS, Anda hanya perlu mengirimkan header CORS.
		return
	}

	// Tulis respons Anda ke Writer seperti yang Anda lakukan sebelumnya.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response := berkatbepkg.GCFPostArticle("MONGOSTRING", "berkatauto", "articleSet", r)
	fmt.Fprintf(w, response)
}
