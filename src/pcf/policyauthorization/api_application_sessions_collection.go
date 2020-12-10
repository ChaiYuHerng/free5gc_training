/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package policyauthorization

import (
	"free5gc/lib/http_wrapper"
	"free5gc/lib/openapi/models"
	"free5gc/src/pcf/handler/message"
	"free5gc/src/pcf/logger"
	"free5gc/src/pcf/util"

	"github.com/gin-gonic/gin"
)

// PostAppSessions - Creates a new Individual Application Session Context resource
func PostAppSessions(c *gin.Context) {
	var appSessionContext models.AppSessionContext
	err := c.ShouldBindJSON(&appSessionContext)
	if err != nil {
		rsp := util.GetProblemDetail("Malformed request syntax", util.ERROR_INITIAL_PARAMETERS)
		logger.HandlerLog.Errorln(rsp.Detail)
		c.JSON(int(rsp.Status), rsp)
		return
	}
	ascReqData := appSessionContext.AscReqData
	if ascReqData == nil || ascReqData.SuppFeat == "" || ascReqData.NotifUri == "" {
		// Check Mandatory IEs
		rsp := util.GetProblemDetail("Errorneous/Missing Mandotory IE", util.ERROR_INITIAL_PARAMETERS)
		logger.HandlerLog.Errorln(rsp.Detail)
		c.JSON(int(rsp.Status), rsp)
		return
	}

	req := http_wrapper.NewRequest(c.Request, appSessionContext)
	channelMsg := message.NewHttpChannelMessage(message.EventPostAppSessions, req)

	message.SendMessage(channelMsg)
	recvMsg := <-channelMsg.HttpChannel
	HTTPResponse := recvMsg.HTTPResponse

	for key, val := range HTTPResponse.Header {
		c.Header(key, val[0])
	}
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}