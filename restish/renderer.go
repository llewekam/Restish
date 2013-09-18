package restish

import (
	"github.com/llewekam/restish/renderer"
)

// A Renderer used to Render resources into interface formats such as JSON and XML
type Renderer interface {
	Render(map[string]interface {}) []byte
	MimeType() string
}

// Create a new renderer
func NewRenderer() Renderer {
	var jsonRenderer Renderer

	jsonRenderer = &renderer.JsonRenderer{}

	return jsonRenderer
}
