/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datarepository

import (
	"free5gc/lib/http_wrapper"
	"free5gc/lib/openapi/models"
	"free5gc/src/udr/handler/message"
	"free5gc/src/udr/logger"

	"github.com/gin-gonic/gin"
)

// CreateAuthenticationSoR - To store the SoR acknowledgement information of a UE
func CreateAuthenticationSoR(c *gin.Context) {
	var sorData models.SorData
	if err := c.ShouldBindJSON(&sorData); err != nil {
		logger.DataRepoLog.Panic(err.Error())
	}

	req := http_wrapper.NewRequest(c.Request, sorData)
	req.Params["ueId"] = c.Params.ByName("ueId")

	handlerMsg := message.NewHandlerMessage(message.EventCreateAuthenticationSoR, req)
	message.SendMessage(handlerMsg)

	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}

// QueryAuthSoR - Retrieves the SoR acknowledgement information of a UE
func QueryAuthSoR(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["ueId"] = c.Params.ByName("ueId")

	handlerMsg := message.NewHandlerMessage(message.EventQueryAuthSoR, req)
	message.SendMessage(handlerMsg)

	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}
