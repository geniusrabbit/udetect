package tools

import (
	_ "embed"
	"net/http"
	"strconv"
	"text/template"

	"go.uber.org/zap"

	swaggerFiles "github.com/swaggo/files"

	"github.com/geniusrabbit/udetect/examples/server/context/ctxlogger"
	"github.com/geniusrabbit/udetect/protocol"
)

// SwaggerServer returns swagger specification files located under "/swagger/"
func SwaggerServer(url string, deepLinking bool) http.HandlerFunc {

	//create a template with name
	t := template.New("swagger_index.html")
	index, _ := t.Parse(swaggerIndexTempl)

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "index.html", "":
			err := index.Execute(w, &swaggerUIBundle{
				URL:         url,
				DeepLinking: deepLinking,
			})
			if err != nil {
				ctxlogger.Get(r.Context()).Error("write HTTP template response", zap.Error(err))
			}
		case "swagger.json":
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Length", strconv.Itoa(len(protocol.SwaggerDefenition)))
			_, err := w.Write([]byte(protocol.SwaggerDefenition))
			if err != nil {
				ctxlogger.Get(r.Context()).Error("write HTTP response", zap.Error(err))
			}
		default:
			swaggerFiles.Handler.ServeHTTP(w, r)
		}
	}
}

type swaggerUIBundle struct {
	URL         string
	DeepLinking bool
}

//go:embed swagger.gotmpl
var swaggerIndexTempl string
