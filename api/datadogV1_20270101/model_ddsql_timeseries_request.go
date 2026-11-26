// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DDSQLTimeseriesRequest Test-only generated schema for the dashboard versioning proof of concept.
type DDSQLTimeseriesRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DdsqlQuery string `json:"ddsql_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType DDSQLTimeseriesRequestRequestType `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat DDSQLTimeseriesRequestResponseFormat `json:"response_format"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDDSQLTimeseriesRequest instantiates a new DDSQLTimeseriesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDDSQLTimeseriesRequest(ddsqlQuery string, requestType DDSQLTimeseriesRequestRequestType, responseFormat DDSQLTimeseriesRequestResponseFormat) *DDSQLTimeseriesRequest {
	this := DDSQLTimeseriesRequest{}
	this.DdsqlQuery = ddsqlQuery
	this.RequestType = requestType
	this.ResponseFormat = responseFormat
	return &this
}

// NewDDSQLTimeseriesRequestWithDefaults instantiates a new DDSQLTimeseriesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDDSQLTimeseriesRequestWithDefaults() *DDSQLTimeseriesRequest {
	this := DDSQLTimeseriesRequest{}
	return &this
}

// GetDdsqlQuery returns the DdsqlQuery field value.
func (o *DDSQLTimeseriesRequest) GetDdsqlQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DdsqlQuery
}

// GetDdsqlQueryOk returns a tuple with the DdsqlQuery field value
// and a boolean to check if the value has been set.
func (o *DDSQLTimeseriesRequest) GetDdsqlQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DdsqlQuery, true
}

// SetDdsqlQuery sets field value.
func (o *DDSQLTimeseriesRequest) SetDdsqlQuery(v string) {
	o.DdsqlQuery = v
}

// GetRequestType returns the RequestType field value.
func (o *DDSQLTimeseriesRequest) GetRequestType() DDSQLTimeseriesRequestRequestType {
	if o == nil {
		var ret DDSQLTimeseriesRequestRequestType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *DDSQLTimeseriesRequest) GetRequestTypeOk() (*DDSQLTimeseriesRequestRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *DDSQLTimeseriesRequest) SetRequestType(v DDSQLTimeseriesRequestRequestType) {
	o.RequestType = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *DDSQLTimeseriesRequest) GetResponseFormat() DDSQLTimeseriesRequestResponseFormat {
	if o == nil {
		var ret DDSQLTimeseriesRequestResponseFormat
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *DDSQLTimeseriesRequest) GetResponseFormatOk() (*DDSQLTimeseriesRequestResponseFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *DDSQLTimeseriesRequest) SetResponseFormat(v DDSQLTimeseriesRequestResponseFormat) {
	o.ResponseFormat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DDSQLTimeseriesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["ddsql_query"] = o.DdsqlQuery
	toSerialize["request_type"] = o.RequestType
	toSerialize["response_format"] = o.ResponseFormat
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DDSQLTimeseriesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DdsqlQuery     *string                               `json:"ddsql_query"`
		RequestType    *DDSQLTimeseriesRequestRequestType    `json:"request_type"`
		ResponseFormat *DDSQLTimeseriesRequestResponseFormat `json:"response_format"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DdsqlQuery == nil {
		return fmt.Errorf("required field ddsql_query missing")
	}
	if all.RequestType == nil {
		return fmt.Errorf("required field request_type missing")
	}
	if all.ResponseFormat == nil {
		return fmt.Errorf("required field response_format missing")
	}

	hasInvalidField := false
	o.DdsqlQuery = *all.DdsqlQuery
	if !all.RequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.RequestType = *all.RequestType
	}
	if !all.ResponseFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.ResponseFormat = *all.ResponseFormat
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
