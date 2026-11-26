// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreemapBackendRequest20270101OneOf4463752384 Test-only generated schema for the dashboard versioning proof of concept.
type TreemapBackendRequest20270101OneOf4463752384 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Projection *TabularProjection_20270101 `json:"projection,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query TreemapBackendRequest20270101OneOfQuery `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208_20270101 `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528_20270101 `json:"response_format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *TreemapRequestStyle_20270101 `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTreemapBackendRequest20270101OneOf4463752384 instantiates a new TreemapBackendRequest20270101OneOf4463752384 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTreemapBackendRequest20270101OneOf4463752384(query TreemapBackendRequest20270101OneOfQuery, requestType TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208_20270101, responseFormat TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528_20270101) *TreemapBackendRequest20270101OneOf4463752384 {
	this := TreemapBackendRequest20270101OneOf4463752384{}
	this.Query = query
	this.RequestType = requestType
	this.ResponseFormat = responseFormat
	return &this
}

// NewTreemapBackendRequest20270101OneOf4463752384WithDefaults instantiates a new TreemapBackendRequest20270101OneOf4463752384 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTreemapBackendRequest20270101OneOf4463752384WithDefaults() *TreemapBackendRequest20270101OneOf4463752384 {
	this := TreemapBackendRequest20270101OneOf4463752384{}
	return &this
}

// GetProjection returns the Projection field value if set, zero value otherwise.
func (o *TreemapBackendRequest20270101OneOf4463752384) GetProjection() TabularProjection_20270101 {
	if o == nil || o.Projection == nil {
		var ret TabularProjection_20270101
		return ret
	}
	return *o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreemapBackendRequest20270101OneOf4463752384) GetProjectionOk() (*TabularProjection_20270101, bool) {
	if o == nil || o.Projection == nil {
		return nil, false
	}
	return o.Projection, true
}

// HasProjection returns a boolean if a field has been set.
func (o *TreemapBackendRequest20270101OneOf4463752384) HasProjection() bool {
	return o != nil && o.Projection != nil
}

// SetProjection gets a reference to the given TabularProjection_20270101 and assigns it to the Projection field.
func (o *TreemapBackendRequest20270101OneOf4463752384) SetProjection(v TabularProjection_20270101) {
	o.Projection = &v
}

// GetQuery returns the Query field value.
func (o *TreemapBackendRequest20270101OneOf4463752384) GetQuery() TreemapBackendRequest20270101OneOfQuery {
	if o == nil {
		var ret TreemapBackendRequest20270101OneOfQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *TreemapBackendRequest20270101OneOf4463752384) GetQueryOk() (*TreemapBackendRequest20270101OneOfQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *TreemapBackendRequest20270101OneOf4463752384) SetQuery(v TreemapBackendRequest20270101OneOfQuery) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *TreemapBackendRequest20270101OneOf4463752384) GetRequestType() TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208_20270101 {
	if o == nil {
		var ret TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208_20270101
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *TreemapBackendRequest20270101OneOf4463752384) GetRequestTypeOk() (*TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *TreemapBackendRequest20270101OneOf4463752384) SetRequestType(v TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208_20270101) {
	o.RequestType = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *TreemapBackendRequest20270101OneOf4463752384) GetResponseFormat() TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528_20270101 {
	if o == nil {
		var ret TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528_20270101
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *TreemapBackendRequest20270101OneOf4463752384) GetResponseFormatOk() (*TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *TreemapBackendRequest20270101OneOf4463752384) SetResponseFormat(v TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528_20270101) {
	o.ResponseFormat = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *TreemapBackendRequest20270101OneOf4463752384) GetStyle() TreemapRequestStyle_20270101 {
	if o == nil || o.Style == nil {
		var ret TreemapRequestStyle_20270101
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreemapBackendRequest20270101OneOf4463752384) GetStyleOk() (*TreemapRequestStyle_20270101, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *TreemapBackendRequest20270101OneOf4463752384) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given TreemapRequestStyle_20270101 and assigns it to the Style field.
func (o *TreemapBackendRequest20270101OneOf4463752384) SetStyle(v TreemapRequestStyle_20270101) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TreemapBackendRequest20270101OneOf4463752384) MarshalJSON() ([]byte, error) {
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
func (o *TreemapBackendRequest20270101OneOf4463752384) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Projection     *TabularProjection_20270101                                           `json:"projection,omitempty"`
		Query          *TreemapBackendRequest20270101OneOfQuery                              `json:"query"`
		RequestType    *TreemapBackendRequestAnyOfAnyOfItemRequestType4453530208_20270101    `json:"request_type"`
		ResponseFormat *TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453530528_20270101 `json:"response_format"`
		Style          *TreemapRequestStyle_20270101                                         `json:"style,omitempty"`
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
	o.Style = all.Style

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
