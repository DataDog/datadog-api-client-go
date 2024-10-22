// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSLogsServicesResponseAttributes AWS Logs Services response body
type AWSLogsServicesResponseAttributes struct {
	// List of AWS services that can send logs to Datadog
	LogsServices []string `json:"logs_services,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSLogsServicesResponseAttributes instantiates a new AWSLogsServicesResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSLogsServicesResponseAttributes() *AWSLogsServicesResponseAttributes {
	this := AWSLogsServicesResponseAttributes{}
	return &this
}

// NewAWSLogsServicesResponseAttributesWithDefaults instantiates a new AWSLogsServicesResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSLogsServicesResponseAttributesWithDefaults() *AWSLogsServicesResponseAttributes {
	this := AWSLogsServicesResponseAttributes{}
	return &this
}

// GetLogsServices returns the LogsServices field value if set, zero value otherwise.
func (o *AWSLogsServicesResponseAttributes) GetLogsServices() []string {
	if o == nil || o.LogsServices == nil {
		var ret []string
		return ret
	}
	return o.LogsServices
}

// GetLogsServicesOk returns a tuple with the LogsServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSLogsServicesResponseAttributes) GetLogsServicesOk() (*[]string, bool) {
	if o == nil || o.LogsServices == nil {
		return nil, false
	}
	return &o.LogsServices, true
}

// HasLogsServices returns a boolean if a field has been set.
func (o *AWSLogsServicesResponseAttributes) HasLogsServices() bool {
	return o != nil && o.LogsServices != nil
}

// SetLogsServices gets a reference to the given []string and assigns it to the LogsServices field.
func (o *AWSLogsServicesResponseAttributes) SetLogsServices(v []string) {
	o.LogsServices = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSLogsServicesResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.LogsServices != nil {
		toSerialize["logs_services"] = o.LogsServices
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSLogsServicesResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LogsServices []string `json:"logs_services,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"logs_services"})
	} else {
		return err
	}
	o.LogsServices = all.LogsServices

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
