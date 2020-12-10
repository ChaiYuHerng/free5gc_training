/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package producer

import (
	"net/http"

	"free5gc/lib/http_wrapper"
	. "free5gc/lib/openapi/models"
	"free5gc/src/nssf/handler/message"
	"free5gc/src/nssf/logger"
)

// NSSAIAvailabilityPost - Creates subscriptions for notification about updates to NSSAI availability information
func NSSAIAvailabilityPost(responseChan chan message.HandlerResponseMessage, createData NssfEventSubscriptionCreateData) {

	logger.Nssaiavailability.Infof("Request received - NSSAIAvailabilityPost")

	var (
		isValidRequest bool = true
		status         int
		createdData    NssfEventSubscriptionCreatedData
		problemDetail  ProblemDetails
	)

	// TODO: If NF consumer is not authorized to update NSSAI availability, return ProblemDetails with code 403 Forbidden

	if isValidRequest {
		status = subscriptionPost(createData, &createdData, &problemDetail)
	}

	if status == http.StatusCreated {
		responseChan <- message.HandlerResponseMessage{
			HttpResponse: &http_wrapper.Response{
				Header: nil,
				Status: status,
				Body:   createData,
			},
		}
	} else {
		responseChan <- message.HandlerResponseMessage{
			HttpResponse: &http_wrapper.Response{
				Header: nil,
				Status: status,
				Body:   problemDetail,
			},
		}
	}
}
