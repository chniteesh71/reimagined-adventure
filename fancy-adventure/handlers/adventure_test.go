package handlers

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// minimal template used in tests (keeps tests self-contained)
const testIndexTmpl = `
{{- if .story -}}
  <html><body><h2>{{ .name }} {{ .story }}</h2></body></html>
{{- else -}}
  <html><body><h1>{{ .title }}</h1></body></html>
{{- end -}}
`

func setupRouterWithTemplate() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// parse the small template above and attach to router
	t := template.Must(template.New("index.html").Parse(testIndexTmpl))
	r.SetHTMLTemplate(t)

	// register handlers exactly as in main.go
	r.GET("/", Home)
	r.GET("/adventure/:name", Adventure)

	return r
}

func TestHome(t *testing.T) {
	r := setupRouterWithTemplate()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Fancy Adventure Go App")
}

func TestAdventure(t *testing.T) {
	r := setupRouterWithTemplate()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/adventure/Niteesh", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// Ensure response contains the name. Adventure is random; checking name is deterministic.
	assert.Contains(t, w.Body.String(), "Niteesh")
}
