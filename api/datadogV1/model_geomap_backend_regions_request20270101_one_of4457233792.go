// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendRegionsRequest20270101OneOf4457233792 Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendRegionsRequest20270101OneOf4457233792 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule_20270101 `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DdsqlQuery string `json:"ddsql_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType_20270101 `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447008416_20270101 `json:"response_format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *GeomapRequestStyle_20270101 `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGeomapBackendRegionsRequest20270101OneOf4457233792 instantiates a new GeomapBackendRegionsRequest20270101OneOf4457233792 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeomapBackendRegionsRequest20270101OneOf4457233792(ddsqlQuery string, requestType GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType_20270101, responseFormat GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447008416_20270101) *GeomapBackendRegionsRequest20270101OneOf4457233792 {
	this := GeomapBackendRegionsRequest20270101OneOf4457233792{}
	this.DdsqlQuery = ddsqlQuery
	this.RequestType = requestType
	this.ResponseFormat = responseFormat
	return &this
}

// NewGeomapBackendRegionsRequest20270101OneOf4457233792WithDefaults instantiates a new GeomapBackendRegionsRequest20270101OneOf4457233792 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeomapBackendRegionsRequest20270101OneOf4457233792WithDefaults() *GeomapBackendRegionsRequest20270101OneOf4457233792 {
	this := GeomapBackendRegionsRequest20270101OneOf4457233792{}
	return &this
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *GeomapBackendRegionsRequest20270101OneOf4457233792) GetConditionalFormats() []ThresholdFormattingRule_20270101 {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule_20270101
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457233792) GetConditionalFormatsOk() (*[]ThresholdFormattingRule_20270101, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457233792) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule_20270101 and assigns it to the ConditionalFormats field.
func (o *GeomapBackendRegionsRequest20270101OneOf4457233792) SetConditionalFormats(v []ThresholdFormattingRule_20270101) {
	o.ConditionalFormats = v
}

// GetDdsqlQuery returns the DdsqlQuery field value.
func (o *GeomapBackendRegionsRequest20270101OneOf4457233792) GetDdsqlQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DdsqlQuery
}

// GetDdsqlQueryOk returns a tuple with the DdsqlQuery field value
// and a boolean to check if the value has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457233792) GetDdsqlQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DdsqlQuery, true
}

// SetDdsqlQuery sets field value.
func (o *GeomapBackendRegionsRequest20270101OneOf4457233792) SetDdsqlQuery(v string) {
	o.DdsqlQuery = v
}

// GetRequestType returns the RequestType field value.
func (o *GeomapBackendRegionsRequest20270101OneOf4457233792) GetRequestType() GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType_20270101 {
	if o == nil {
		var ret GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType_20270101
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457233792) GetRequestTypeOk() (*GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *GeomapBackendRegionsRequest20270101OneOf4457233792) SetRequestType(v GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType_20270101) {
	o.RequestType = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *GeomapBackendRegionsRequest20270101OneOf4457233792) GetResponseFormat() GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447008416_20270101 {
	if o == nil {
		var ret GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447008416_20270101
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457233792) GetResponseFormatOk() (*GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447008416_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *GeomapBackendRegionsRequest20270101OneOf4457233792) SetResponseFormat(v GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447008416_20270101) {
	o.ResponseFormat = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *GeomapBackendRegionsRequest20270101OneOf4457233792) GetStyle() GeomapRequestStyle_20270101 {
	if o == nil || o.Style == nil {
		var ret GeomapRequestStyle_20270101
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457233792) GetStyleOk() (*GeomapRequestStyle_20270101, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457233792) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given GeomapRequestStyle_20270101 and assigns it to the Style field.
func (o *GeomapBackendRegionsRequest20270101OneOf4457233792) SetStyle(v GeomapRequestStyle_20270101) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeomapBackendRegionsRequest20270101OneOf4457233792) MarshalJSON() ([]byte, error) {
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
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeomapBackendRegionsRequest20270101OneOf4457233792) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConditionalFormats []ThresholdFormattingRule_20270101                                          `json:"conditional_formats,omitempty"`
		DdsqlQuery         *string                                                                     `json:"ddsql_query"`
		RequestType        *GeomapBackendRegionsRequestAnyOfAnyOfItemRequestType_20270101              `json:"request_type"`
		ResponseFormat     *GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447008416_20270101 `json:"response_format"`
		Style              *GeomapRequestStyle_20270101                                                `json:"style,omitempty"`
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
	o.Style = all.Style

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
