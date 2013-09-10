// restish
//

package restish

// Describe all resource. This is the generic internal resource representation. It should be capable of creating
// output in the same structure as the following "example" JSON.
//
//    {
//        "@xmlns": "urn:com.restish.example",
//        "@xmlns:atom": "http://www.w3.org/2005/Atom",
//        "@xmlns:another": "urn:com.restish.another",
//        "atom:link": {
//            "@rel": "self",
//            "@href": "http://.../example/123",
//            "@type": "application/vnd.com.restish+json"
//        },
//        "title": "Example Resource",
//        "description": "An example to describe the intended resource structure",
//        "another:another": {
//            "@count": "2",
//            "atom:link": {
//                "@rel": "self",
//                "@href": "http://.../example/123/another",
//                "@type": "application/vnd.com.restish+json"
//            }
//            "title": "More Resources",
//            "description": "More resources linked from the example resource."
//        }
//    }

// A REST Resource. A Resource can be dispatched to Handlers and Controllers for processing. A Resource is always
// returned from a successful Dispatch
type Resource struct {
	Status     StatusCode
	Links      []Link
	Type       string
	Properties map[string]string
	Resources  []Resource

	self *Link
}

// A resource link. Defines the properties required to link to resources.
type Link struct {
	Rel  string
	Href string
	Type string
}

// Create a new Resource
func NewResource(href string) *Resource {
	var resource *Resource

	resource = new(Resource)
	resource.Status = StatusOk
	resource.AddLink("self", href, "application/vnd.com.restish")

	return resource
}

// Add a link to the resource.
func (resource *Resource) AddLink(rel string, href string, linkType string) {
	var link Link

	link = Link{rel, href, linkType}
	resource.Links = append(resource.Links, link)

	if "self" == rel {
		resource.self = &link
	}
}

// Return the link to the resource "self"
func (resource *Resource) Self() *Link {
	return resource.self
}
