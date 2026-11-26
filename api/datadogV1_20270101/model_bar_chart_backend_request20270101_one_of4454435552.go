// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartBackendRequest20270101OneOf4454435552 Test-only generated schema for the dashboard versioning proof of concept.
type BarChartBackendRequest20270101OneOf4454435552 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DdsqlQuery string `json:"ddsql_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944 `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264 `json:"response_format"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewBarChartBackendRequest20270101OneOf4454435552 instantiates a new BarChartBackendRequest20270101OneOf4454435552 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBarChartBackendRequest20270101OneOf4454435552(ddsqlQuery string, requestType BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944, responseFormat BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264) *BarChartBackendRequest20270101OneOf4454435552 {
	this := BarChartBackendRequest20270101OneOf4454435552{}
	this.DdsqlQuery = ddsqlQuery
	this.RequestType = requestType
	this.ResponseFormat = responseFormat
	return &this
}

// NewBarChartBackendRequest20270101OneOf4454435552WithDefaults instantiates a new BarChartBackendRequest20270101OneOf4454435552 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBarChartBackendRequest20270101OneOf4454435552WithDefaults() *BarChartBackendRequest20270101OneOf4454435552 {
	this := BarChartBackendRequest20270101OneOf4454435552{}
	return &this
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *BarChartBackendRequest20270101OneOf4454435552) GetConditionalFormats() []ThresholdFormattingRule {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartBackendRequest20270101OneOf4454435552) GetConditionalFormatsOk() (*[]ThresholdFormattingRule, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *BarChartBackendRequest20270101OneOf4454435552) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule and assigns it to the ConditionalFormats field.
func (o *BarChartBackendRequest20270101OneOf4454435552) SetConditionalFormats(v []ThresholdFormattingRule) {
	o.ConditionalFormats = v
}

// GetDdsqlQuery returns the DdsqlQuery field value.
func (o *BarChartBackendRequest20270101OneOf4454435552) GetDdsqlQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DdsqlQuery
}

// GetDdsqlQueryOk returns a tuple with the DdsqlQuery field value
// and a boolean to check if the value has been set.
func (o *BarChartBackendRequest20270101OneOf4454435552) GetDdsqlQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DdsqlQuery, true
}

// SetDdsqlQuery sets field value.
func (o *BarChartBackendRequest20270101OneOf4454435552) SetDdsqlQuery(v string) {
	o.DdsqlQuery = v
}

// GetRequestType returns the RequestType field value.
func (o *BarChartBackendRequest20270101OneOf4454435552) GetRequestType() BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944 {
	if o == nil {
		var ret BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *BarChartBackendRequest20270101OneOf4454435552) GetRequestTypeOk() (*BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *BarChartBackendRequest20270101OneOf4454435552) SetRequestType(v BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944) {
	o.RequestType = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *BarChartBackendRequest20270101OneOf4454435552) GetResponseFormat() BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264 {
	if o == nil {
		var ret BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *BarChartBackendRequest20270101OneOf4454435552) GetResponseFormatOk() (*BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *BarChartBackendRequest20270101OneOf4454435552) SetResponseFormat(v BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264) {
	o.ResponseFormat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BarChartBackendRequest20270101OneOf4454435552) MarshalJSON() ([]byte, error) {
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
func (o *BarChartBackendRequest20270101OneOf4454435552) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConditionalFormats []ThresholdFormattingRule                                     `json:"conditional_formats,omitempty"`
		DdsqlQuery         *string                                                       `json:"ddsql_query"`
		RequestType        *BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944    `json:"request_type"`
		ResponseFormat     *BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264 `json:"response_format"`
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
