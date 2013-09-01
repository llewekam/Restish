package restish

import (
	"encoding/json"
	"net/http"
)

// A Renderer used to Render resources into interface formats such as JSON and XML
type Renderer interface {
	Render(*Resource) []byte
	MimeType() string
}

// Create a new renderer
func NewRenderer() Renderer {
	renderer := JsonRenderer{}

	return renderer
}

// Convert an HTTP request into a rest resource ready for processing
func FormatRequest(request *http.Request) *Resource {
	resource := NewResource(request.RequestURI)

	return resource
}

// A JSOM Renderer. Converts rest resources into JSON
type JsonRenderer struct {
}

func (_ JsonRenderer) Render(resource *Resource) []byte {
	response, _ := json.Marshal(resource.Properties)

	return response
}

func (renderer JsonRenderer) MimeType() string {
	return "application/json"
}
