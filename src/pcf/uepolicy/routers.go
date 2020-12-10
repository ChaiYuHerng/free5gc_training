/*
 * Npcf_UEPolicyControl
 *
 * UE Policy Control Service API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package uepolicy

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

func NewRouter() *gin.Engine {
	router := gin.Default()
	AddService(router)
	return router
}

func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/npcf-ue-policy-control/v1/")

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "PATCH":
			group.PATCH(route.Pattern, route.HandlerFunc)
		}
	}
	return group
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		"GET",
		"",
		Index,
	},

	{
		"PoliciesPolAssoIdDelete",
		strings.ToUpper("Delete"),
		"/policies/{polAssoId}",
		PoliciesPolAssoIdDelete,
	},

	{
		"PoliciesPolAssoIdGet",
		strings.ToUpper("Get"),
		"/policies/{polAssoId}",
		PoliciesPolAssoIdGet,
	},

	{
		"PoliciesPolAssoIdUpdatePost",
		strings.ToUpper("Post"),
		"/policies/{polAssoId}/update",
		PoliciesPolAssoIdUpdatePost,
	},

	{
		"PoliciesPost",
		strings.ToUpper("Post"),
		"/policies",
		PoliciesPost,
	},
}
