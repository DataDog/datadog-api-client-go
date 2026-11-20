// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotBackendTableRequest20270101OneOf Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotBackendTableRequest20270101OneOf struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DdsqlQuery string `json:"ddsql_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType_20270101 `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat_20270101 `json:"response_format"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewScatterplotBackendTableRequest20270101OneOf instantiates a new ScatterplotBackendTableRequest20270101OneOf object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScatterplotBackendTableRequest20270101OneOf(ddsqlQuery string, requestType ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType_20270101, responseFormat ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat_20270101) *ScatterplotBackendTableRequest20270101OneOf {
	this := ScatterplotBackendTableRequest20270101OneOf{}
	this.DdsqlQuery = ddsqlQuery
	this.RequestType = requestType
	this.ResponseFormat = responseFormat
	return &this
}

// NewScatterplotBackendTableRequest20270101OneOfWithDefaults instantiates a new ScatterplotBackendTableRequest20270101OneOf object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScatterplotBackendTableRequest20270101OneOfWithDefaults() *ScatterplotBackendTableRequest20270101OneOf {
	this := ScatterplotBackendTableRequest20270101OneOf{}
	return &this
}

// GetDdsqlQuery returns the DdsqlQuery field value.
func (o *ScatterplotBackendTableRequest20270101OneOf) GetDdsqlQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DdsqlQuery
}

// GetDdsqlQueryOk returns a tuple with the DdsqlQuery field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf) GetDdsqlQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DdsqlQuery, true
}

// SetDdsqlQuery sets field value.
func (o *ScatterplotBackendTableRequest20270101OneOf) SetDdsqlQuery(v string) {
	o.DdsqlQuery = v
}

// GetRequestType returns the RequestType field value.
func (o *ScatterplotBackendTableRequest20270101OneOf) GetRequestType() ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType_20270101 {
	if o == nil {
		var ret ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType_20270101
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf) GetRequestTypeOk() (*ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *ScatterplotBackendTableRequest20270101OneOf) SetRequestType(v ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType_20270101) {
	o.RequestType = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *ScatterplotBackendTableRequest20270101OneOf) GetResponseFormat() ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat_20270101 {
	if o == nil {
		var ret ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat_20270101
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf) GetResponseFormatOk() (*ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *ScatterplotBackendTableRequest20270101OneOf) SetResponseFormat(v ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat_20270101) {
	o.ResponseFormat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScatterplotBackendTableRequest20270101OneOf) MarshalJSON() ([]byte, error) {
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
func (o *ScatterplotBackendTableRequest20270101OneOf) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DdsqlQuery     *string                                                              `json:"ddsql_query"`
		RequestType    *ScatterplotBackendTableRequestAnyOfAnyOfItemRequestType_20270101    `json:"request_type"`
		ResponseFormat *ScatterplotBackendTableRequestAnyOfAnyOfItemResponseFormat_20270101 `json:"response_format"`
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
