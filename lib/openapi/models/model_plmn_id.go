/*
 * Nsmf_EventExposure
 *
 * Session Management Event Exposure Service API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PlmnId struct {
	Mcc string `json:"mcc" yaml:"mcc" bson:"mcc" mapstructure:"Mcc"`
	Mnc string `json:"mnc" yaml:"mnc" bson:"mnc" mapstructure:"Mnc"`
}
