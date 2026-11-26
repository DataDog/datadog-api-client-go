// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstBackendRequest20270101OneOf4462424800 Test-only generated schema for the dashboard versioning proof of concept.
type SunburstBackendRequest20270101OneOf4462424800 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Projection *TabularProjection `json:"projection,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query SunburstBackendRequest20270101OneOfQuery `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072 `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392 `json:"response_format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *SunburstRequestStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSunburstBackendRequest20270101OneOf4462424800 instantiates a new SunburstBackendRequest20270101OneOf4462424800 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSunburstBackendRequest20270101OneOf4462424800(query SunburstBackendRequest20270101OneOfQuery, requestType SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072, responseFormat SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392) *SunburstBackendRequest20270101OneOf4462424800 {
	this := SunburstBackendRequest20270101OneOf4462424800{}
	this.Query = query
	this.RequestType = requestType
	this.ResponseFormat = responseFormat
	return &this
}

// NewSunburstBackendRequest20270101OneOf4462424800WithDefaults instantiates a new SunburstBackendRequest20270101OneOf4462424800 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSunburstBackendRequest20270101OneOf4462424800WithDefaults() *SunburstBackendRequest20270101OneOf4462424800 {
	this := SunburstBackendRequest20270101OneOf4462424800{}
	return &this
}

// GetProjection returns the Projection field value if set, zero value otherwise.
func (o *SunburstBackendRequest20270101OneOf4462424800) GetProjection() TabularProjection {
	if o == nil || o.Projection == nil {
		var ret TabularProjection
		return ret
	}
	return *o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SunburstBackendRequest20270101OneOf4462424800) GetProjectionOk() (*TabularProjection, bool) {
	if o == nil || o.Projection == nil {
		return nil, false
	}
	return o.Projection, true
}

// HasProjection returns a boolean if a field has been set.
func (o *SunburstBackendRequest20270101OneOf4462424800) HasProjection() bool {
	return o != nil && o.Projection != nil
}

// SetProjection gets a reference to the given TabularProjection and assigns it to the Projection field.
func (o *SunburstBackendRequest20270101OneOf4462424800) SetProjection(v TabularProjection) {
	o.Projection = &v
}

// GetQuery returns the Query field value.
func (o *SunburstBackendRequest20270101OneOf4462424800) GetQuery() SunburstBackendRequest20270101OneOfQuery {
	if o == nil {
		var ret SunburstBackendRequest20270101OneOfQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *SunburstBackendRequest20270101OneOf4462424800) GetQueryOk() (*SunburstBackendRequest20270101OneOfQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *SunburstBackendRequest20270101OneOf4462424800) SetQuery(v SunburstBackendRequest20270101OneOfQuery) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *SunburstBackendRequest20270101OneOf4462424800) GetRequestType() SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072 {
	if o == nil {
		var ret SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *SunburstBackendRequest20270101OneOf4462424800) GetRequestTypeOk() (*SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *SunburstBackendRequest20270101OneOf4462424800) SetRequestType(v SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072) {
	o.RequestType = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *SunburstBackendRequest20270101OneOf4462424800) GetResponseFormat() SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392 {
	if o == nil {
		var ret SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *SunburstBackendRequest20270101OneOf4462424800) GetResponseFormatOk() (*SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *SunburstBackendRequest20270101OneOf4462424800) SetResponseFormat(v SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392) {
	o.ResponseFormat = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *SunburstBackendRequest20270101OneOf4462424800) GetStyle() SunburstRequestStyle {
	if o == nil || o.Style == nil {
		var ret SunburstRequestStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SunburstBackendRequest20270101OneOf4462424800) GetStyleOk() (*SunburstRequestStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *SunburstBackendRequest20270101OneOf4462424800) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given SunburstRequestStyle and assigns it to the Style field.
func (o *SunburstBackendRequest20270101OneOf4462424800) SetStyle(v SunburstRequestStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SunburstBackendRequest20270101OneOf4462424800) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Projection != nil {
		toSerialize["projection"] = o.Projection
	}
	toSerialize["query"] = o.Query
	toSerialize["request_type"] = o.RequestType
	toSerialize["response_format"] = o.ResponseFormat
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SunburstBackendRequest20270101OneOf4462424800) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Projection     *TabularProjection                                            `json:"projection,omitempty"`
		Query          *SunburstBackendRequest20270101OneOfQuery                     `json:"query"`
		RequestType    *SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072    `json:"request_type"`
		ResponseFormat *SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392 `json:"response_format"`
		Style          *SunburstRequestStyle                                         `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.RequestType == nil {
		return fmt.Errorf("required field request_type missing")
	}
	if all.ResponseFormat == nil {
		return fmt.Errorf("required field response_format missing")
	}

	hasInvalidField := false
	o.Projection = all.Projection
	o.Query = *all.Query
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
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
