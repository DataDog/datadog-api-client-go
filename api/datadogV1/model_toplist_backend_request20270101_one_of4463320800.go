// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendRequest20270101OneOf4463320800 Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendRequest20270101OneOf4463320800 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule_20270101 `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DdsqlQuery string `json:"ddsql_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848_20270101 `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168_20270101 `json:"response_format"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewToplistBackendRequest20270101OneOf4463320800 instantiates a new ToplistBackendRequest20270101OneOf4463320800 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewToplistBackendRequest20270101OneOf4463320800(ddsqlQuery string, requestType ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848_20270101, responseFormat ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168_20270101) *ToplistBackendRequest20270101OneOf4463320800 {
	this := ToplistBackendRequest20270101OneOf4463320800{}
	this.DdsqlQuery = ddsqlQuery
	this.RequestType = requestType
	this.ResponseFormat = responseFormat
	return &this
}

// NewToplistBackendRequest20270101OneOf4463320800WithDefaults instantiates a new ToplistBackendRequest20270101OneOf4463320800 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewToplistBackendRequest20270101OneOf4463320800WithDefaults() *ToplistBackendRequest20270101OneOf4463320800 {
	this := ToplistBackendRequest20270101OneOf4463320800{}
	return &this
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *ToplistBackendRequest20270101OneOf4463320800) GetConditionalFormats() []ThresholdFormattingRule_20270101 {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule_20270101
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistBackendRequest20270101OneOf4463320800) GetConditionalFormatsOk() (*[]ThresholdFormattingRule_20270101, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *ToplistBackendRequest20270101OneOf4463320800) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule_20270101 and assigns it to the ConditionalFormats field.
func (o *ToplistBackendRequest20270101OneOf4463320800) SetConditionalFormats(v []ThresholdFormattingRule_20270101) {
	o.ConditionalFormats = v
}

// GetDdsqlQuery returns the DdsqlQuery field value.
func (o *ToplistBackendRequest20270101OneOf4463320800) GetDdsqlQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DdsqlQuery
}

// GetDdsqlQueryOk returns a tuple with the DdsqlQuery field value
// and a boolean to check if the value has been set.
func (o *ToplistBackendRequest20270101OneOf4463320800) GetDdsqlQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DdsqlQuery, true
}

// SetDdsqlQuery sets field value.
func (o *ToplistBackendRequest20270101OneOf4463320800) SetDdsqlQuery(v string) {
	o.DdsqlQuery = v
}

// GetRequestType returns the RequestType field value.
func (o *ToplistBackendRequest20270101OneOf4463320800) GetRequestType() ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848_20270101 {
	if o == nil {
		var ret ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848_20270101
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *ToplistBackendRequest20270101OneOf4463320800) GetRequestTypeOk() (*ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *ToplistBackendRequest20270101OneOf4463320800) SetRequestType(v ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848_20270101) {
	o.RequestType = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *ToplistBackendRequest20270101OneOf4463320800) GetResponseFormat() ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168_20270101 {
	if o == nil {
		var ret ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168_20270101
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *ToplistBackendRequest20270101OneOf4463320800) GetResponseFormatOk() (*ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *ToplistBackendRequest20270101OneOf4463320800) SetResponseFormat(v ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168_20270101) {
	o.ResponseFormat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ToplistBackendRequest20270101OneOf4463320800) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	toSerialize["ddsql_query"] = o.DdsqlQuery
	toSerialize["request_type"] = o.RequestType
	toSerialize["response_format"] = o.ResponseFormat
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ToplistBackendRequest20270101OneOf4463320800) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConditionalFormats []ThresholdFormattingRule_20270101                                    `json:"conditional_formats,omitempty"`
		DdsqlQuery         *string                                                               `json:"ddsql_query"`
		RequestType        *ToplistBackendRequestAnyOfAnyOfItemRequestType4453034848_20270101    `json:"request_type"`
		ResponseFormat     *ToplistBackendRequestAnyOfAnyOfItemResponseFormat4453035168_20270101 `json:"response_format"`
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
	o.ConditionalFormats = all.ConditionalFormats
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
