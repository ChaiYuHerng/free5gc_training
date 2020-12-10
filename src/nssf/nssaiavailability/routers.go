/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nssaiavailability

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
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

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	AddService(router)
	return router
}

func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/nnssf-nssaiavailability/v1")

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			group.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.DELETE(route.Pattern, route.HandlerFunc)
		case "PATCH":
			group.PATCH(route.Pattern, route.HandlerFunc)
		}
	}
	return group
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
	},

	{
		"NSSAIAvailabilityDelete",
		strings.ToUpper("Delete"),
		"/nssai-availability/:nfId",
		ApiNfInstanceIdDocumentDelete,
	},

	{
		"NSSAIAvailabilityPatch",
		strings.ToUpper("Patch"),
		"/nssai-availability/:nfId",
		ApiNfInstanceIdDocumentPatch,
	},

	{
		"NSSAIAvailabilityPut",
		strings.ToUpper("Put"),
		"/nssai-availability/:nfId",
		ApiNfInstanceIdDocumentPut,
	},

	// Regular expressions for route matching should be unique in Gin package
	// 'subscriptions' would conflict with existing wildcard ':nfId'
	// Simply replace 'subscriptions' with ':nfId' and check if ':nfId' is 'subscriptions' in handler function
	{
		"NSSAIAvailabilityUnsubscribe",
		strings.ToUpper("Delete"),
		// "/nssai-availability/subscriptions/:subscriptionId",
		"/nssai-availability/:nfId/:subscriptionId",
		ApiSubscriptionIdDocument,
	},

	{
		"NSSAIAvailabilityPost",
		strings.ToUpper("Post"),
		"/nssai-availability/subscriptions",
		ApiSubscriptionsCollection,
	},
}
