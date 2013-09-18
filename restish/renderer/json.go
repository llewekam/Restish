// json renderer
//
//

package renderer

import (
	"encoding/json"
)

// A JSON Renderer. Converts rest resources into JSON
type JsonRenderer struct {
}

func (_ JsonRenderer) Render(response map[string]interface {}) []byte {
	responseString, _ := json.Marshal(response)

	return responseString
}

func (renderer JsonRenderer) MimeType() string {
	return "application/json"
}
