/*
 * Nudm_UECM
 *
 * Nudm Context Management Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package uecontextmanagement

import (
	"free5gc/lib/http_wrapper"
	"free5gc/lib/openapi/models"
	"free5gc/src/udm/logger"
	"free5gc/src/udm/handler"
	udm_message "free5gc/src/udm/handler/message"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateAmfNon3gppAccess - update a parameter in the AMF registration for non-3GPP access
func UpdateAmfNon3gppAccess(c *gin.Context) {
	var amfNon3GppAccessRegistrationModification models.AmfNon3GppAccessRegistrationModification
	if err := c.ShouldBindJSON(&amfNon3GppAccessRegistrationModification); err != nil {
		logger.UeauLog.Errorln(err)
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := http_wrapper.NewRequest(c.Request, amfNon3GppAccessRegistrationModification)
	req.Params["ueId"] = c.Param("ueId")

	handlerMsg := udm_message.NewHandlerMessage(udm_message.EventUpdateAmfNon3gppAccess, req)
	handler.SendMessage(handlerMsg)
	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
	return
}
