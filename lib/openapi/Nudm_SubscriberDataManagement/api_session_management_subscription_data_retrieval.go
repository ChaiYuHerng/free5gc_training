/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_SubscriberDataManagement

import (
	"free5gc/lib/openapi"
	"free5gc/lib/openapi/models"

	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type SessionManagementSubscriptionDataRetrievalApiService service

/*
SessionManagementSubscriptionDataRetrievalApiService retrieve a UE's Session Management Subscription Data
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param supi Identifier of the UE
 * @param optional nil or *GetSmDataParamOpts - Optional Parameters:
 * @param "SupportedFeatures" (optional.String) -  Supported Features
 * @param "SingleNssai" (optional.Interface of models.Snssai) -
 * @param "Dnn" (optional.String) -
 * @param "PlmnId" (optional.Interface of models.PlmnId) -
 * @param "IfNoneMatch" (optional.String) -  Validator for conditional requests, as described in RFC 7232, 3.2
 * @param "IfModifiedSince" (optional.String) -  Validator for conditional requests, as described in RFC 7232, 3.3
@return []models.SessionManagementSubscriptionData
*/

type GetSmDataParamOpts struct {
	SupportedFeatures optional.String
	SingleNssai       optional.Interface
	Dnn               optional.String
	PlmnId            optional.Interface
	IfNoneMatch       optional.String
	IfModifiedSince   optional.String
}

func (a *SessionManagementSubscriptionDataRetrievalApiService) GetSmData(ctx context.Context, supi string, localVarOptionals *GetSmDataParamOpts) ([]models.SessionManagementSubscriptionData, *http.Response, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []models.SessionManagementSubscriptionData
	)

	// create path and map variables
	//fmt.Printf("smf get udm uri:%s\n",a.client.cfg.BasePath())
	//a.client.cfg.BasePath() = "http://192.168.2.238:29503/nudm-sdm/v1"
	localVarPath := "http://192.168.2.106:29503/nudm-sdm/v1" + "/{supi}/sm-data"
	localVarPath = strings.Replace(localVarPath, "{"+"supi"+"}", fmt.Sprintf("%v", supi), -1)

	fmt.Printf("localVarPath is %s\n",localVarPath)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.SupportedFeatures.IsSet() {
		fmt.Printf("GetSmData test1~~~~~~~~~~\n")
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(localVarOptionals.SupportedFeatures.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SingleNssai.IsSet() {
		fmt.Printf("GetSmData test2~~~~~~~~~~\n")
		localVarQueryParams.Add("single-nssai", openapi.ParameterToString(localVarOptionals.SingleNssai.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Dnn.IsSet() {
		fmt.Printf("GetSmData test3~~~~~~~~~~\n")
		localVarQueryParams.Add("dnn", openapi.ParameterToString(localVarOptionals.Dnn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PlmnId.IsSet() {
		fmt.Printf("GetSmData test4~~~~~~~~~~\n")
		localVarQueryParams.Add("plmn-id", openapi.ParameterToString(localVarOptionals.PlmnId.Value(), ""))
	}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}
	fmt.Printf("GetSmData localVarHTTPHeaderAccepts is %s\n\n",localVarHTTPHeaderAccepts)

	// set Accept header
	localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	fmt.Printf("GetSmData localVarHTTPHeaderAccept is %s\n\n",localVarHTTPHeaderAccept)
	if localVarHTTPHeaderAccept != "" {
		fmt.Printf("GetSmData test5~~~~~~~~~~\n")
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if localVarOptionals != nil && localVarOptionals.IfNoneMatch.IsSet() {
		fmt.Printf("GetSmData test6~~~~~~~~~~\n")
		localVarHeaderParams["If-None-Match"] = openapi.ParameterToString(localVarOptionals.IfNoneMatch.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.IfModifiedSince.IsSet() {
		fmt.Printf("GetSmData test7~~~~~~~~~~\n")
		localVarHeaderParams["If-Modified-Since"] = openapi.ParameterToString(localVarOptionals.IfModifiedSince.Value(), "")
	}

	fmt.Printf("ctx is %s\n\n",ctx)
	fmt.Printf("a.client.cfg is %s\n\n",a.client.cfg)
	fmt.Printf("localVarPath is %s\n\n",localVarPath)
	fmt.Printf("localVarHTTPMethod is %s\n\n",localVarHTTPMethod)
	fmt.Printf("localVarPostBody is %s\n\n",localVarPostBody)
	fmt.Printf("localVarHeaderParams is %s\n\n",localVarHeaderParams)
	fmt.Printf("localVarQueryParams is %s\n\n",localVarQueryParams)
	fmt.Printf("localVarFormParams is %s\n\n",localVarFormParams)
	fmt.Printf("localVarFormFileName is %s\n\n",localVarFormFileName)
	fmt.Printf("localVarFileName is %s\n\n",localVarFileName)
	fmt.Printf("localVarFileBytes is %s\n\n",localVarFileBytes)
	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	fmt.Printf("GetSmData r is %s\n\n",r)
	if err != nil {
		fmt.Printf("GetSmData test8~~~~~~~~~~\n")
		return localVarReturnValue, nil, err
	}


	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	fmt.Printf("a.client.cfg is %s\n\n",a.client.cfg)
	fmt.Printf("GetSmData localVarHTTPResponser is %s\n",localVarHTTPResponse)
	if err != nil || localVarHTTPResponse == nil {
		fmt.Printf("GetSmData test9~~~~~~~~~~\n")
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	fmt.Printf("GetSmData localVarBody is %s\n",localVarBody)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		fmt.Printf("GetSmData test10~~~~~~~~~~\n")
		return localVarReturnValue, localVarHTTPResponse, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.Status,
	}

	switch localVarHTTPResponse.StatusCode {
	case 200:
		fmt.Printf("GetSmData case 200\n")
		err = openapi.Deserialize(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
		}
		return localVarReturnValue, localVarHTTPResponse, nil
	case 400:
		fmt.Printf("GetSmData case 400\n")
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 404:
		fmt.Printf("GetSmData case 404\n")
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 500:
		fmt.Printf("GetSmData case 500\n")
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 503:
		fmt.Printf("GetSmData case 503\n")
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	default:
		fmt.Printf("GetSmData case default\n")
		return localVarReturnValue, localVarHTTPResponse, nil
	}
}
