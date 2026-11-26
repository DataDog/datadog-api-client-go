// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreemapBackendRequest20270101OneOf4463751424 Test-only generated schema for the dashboard versioning proof of concept.
type TreemapBackendRequest20270101OneOf4463751424 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DdsqlQuery string `json:"ddsql_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768_20270101 `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453529088_20270101 `json:"response_format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *TreemapRequestStyle_20270101 `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTreemapBackendRequest20270101OneOf4463751424 instantiates a new TreemapBackendRequest20270101OneOf4463751424 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTreemapBackendRequest20270101OneOf4463751424(ddsqlQuery string, requestType TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768_20270101, responseFormat TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453529088_20270101) *TreemapBackendRequest20270101OneOf4463751424 {
	this := TreemapBackendRequest20270101OneOf4463751424{}
	this.DdsqlQuery = ddsqlQuery
	this.RequestType = requestType
	this.ResponseFormat = responseFormat
	return &this
}

// NewTreemapBackendRequest20270101OneOf4463751424WithDefaults instantiates a new TreemapBackendRequest20270101OneOf4463751424 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTreemapBackendRequest20270101OneOf4463751424WithDefaults() *TreemapBackendRequest20270101OneOf4463751424 {
	this := TreemapBackendRequest20270101OneOf4463751424{}
	return &this
}

// GetDdsqlQuery returns the DdsqlQuery field value.
func (o *TreemapBackendRequest20270101OneOf4463751424) GetDdsqlQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DdsqlQuery
}

// GetDdsqlQueryOk returns a tuple with the DdsqlQuery field value
// and a boolean to check if the value has been set.
func (o *TreemapBackendRequest20270101OneOf4463751424) GetDdsqlQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DdsqlQuery, true
}

// SetDdsqlQuery sets field value.
func (o *TreemapBackendRequest20270101OneOf4463751424) SetDdsqlQuery(v string) {
	o.DdsqlQuery = v
}

// GetRequestType returns the RequestType field value.
func (o *TreemapBackendRequest20270101OneOf4463751424) GetRequestType() TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768_20270101 {
	if o == nil {
		var ret TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768_20270101
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *TreemapBackendRequest20270101OneOf4463751424) GetRequestTypeOk() (*TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *TreemapBackendRequest20270101OneOf4463751424) SetRequestType(v TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768_20270101) {
	o.RequestType = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *TreemapBackendRequest20270101OneOf4463751424) GetResponseFormat() TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453529088_20270101 {
	if o == nil {
		var ret TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453529088_20270101
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *TreemapBackendRequest20270101OneOf4463751424) GetResponseFormatOk() (*TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453529088_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *TreemapBackendRequest20270101OneOf4463751424) SetResponseFormat(v TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453529088_20270101) {
	o.ResponseFormat = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *TreemapBackendRequest20270101OneOf4463751424) GetStyle() TreemapRequestStyle_20270101 {
	if o == nil || o.Style == nil {
		var ret TreemapRequestStyle_20270101
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreemapBackendRequest20270101OneOf4463751424) GetStyleOk() (*TreemapRequestStyle_20270101, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *TreemapBackendRequest20270101OneOf4463751424) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given TreemapRequestStyle_20270101 and assigns it to the Style field.
func (o *TreemapBackendRequest20270101OneOf4463751424) SetStyle(v TreemapRequestStyle_20270101) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TreemapBackendRequest20270101OneOf4463751424) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
func (o *TreemapBackendRequest20270101OneOf4463751424) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DdsqlQuery     *string                                                               `json:"ddsql_query"`
		RequestType    *TreemapBackendRequestAnyOfAnyOfItemRequestType4453528768_20270101    `json:"request_type"`
		ResponseFormat *TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453529088_20270101 `json:"response_format"`
		Style          *TreemapRequestStyle_20270101                                         `json:"style,omitempty"`
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
	o.Style = all.Style

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
