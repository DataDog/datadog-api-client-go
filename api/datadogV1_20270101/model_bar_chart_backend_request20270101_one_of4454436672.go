// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartBackendRequest20270101OneOf4454436672 Test-only generated schema for the dashboard versioning proof of concept.
type BarChartBackendRequest20270101OneOf4454436672 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Projection *TabularProjection `json:"projection,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query BarChartBackendRequest20270101OneOfQuery `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544 `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864 `json:"response_format"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewBarChartBackendRequest20270101OneOf4454436672 instantiates a new BarChartBackendRequest20270101OneOf4454436672 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBarChartBackendRequest20270101OneOf4454436672(query BarChartBackendRequest20270101OneOfQuery, requestType BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544, responseFormat BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864) *BarChartBackendRequest20270101OneOf4454436672 {
	this := BarChartBackendRequest20270101OneOf4454436672{}
	this.Query = query
	this.RequestType = requestType
	this.ResponseFormat = responseFormat
	return &this
}

// NewBarChartBackendRequest20270101OneOf4454436672WithDefaults instantiates a new BarChartBackendRequest20270101OneOf4454436672 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBarChartBackendRequest20270101OneOf4454436672WithDefaults() *BarChartBackendRequest20270101OneOf4454436672 {
	this := BarChartBackendRequest20270101OneOf4454436672{}
	return &this
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *BarChartBackendRequest20270101OneOf4454436672) GetConditionalFormats() []ThresholdFormattingRule {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartBackendRequest20270101OneOf4454436672) GetConditionalFormatsOk() (*[]ThresholdFormattingRule, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *BarChartBackendRequest20270101OneOf4454436672) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule and assigns it to the ConditionalFormats field.
func (o *BarChartBackendRequest20270101OneOf4454436672) SetConditionalFormats(v []ThresholdFormattingRule) {
	o.ConditionalFormats = v
}

// GetProjection returns the Projection field value if set, zero value otherwise.
func (o *BarChartBackendRequest20270101OneOf4454436672) GetProjection() TabularProjection {
	if o == nil || o.Projection == nil {
		var ret TabularProjection
		return ret
	}
	return *o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartBackendRequest20270101OneOf4454436672) GetProjectionOk() (*TabularProjection, bool) {
	if o == nil || o.Projection == nil {
		return nil, false
	}
	return o.Projection, true
}

// HasProjection returns a boolean if a field has been set.
func (o *BarChartBackendRequest20270101OneOf4454436672) HasProjection() bool {
	return o != nil && o.Projection != nil
}

// SetProjection gets a reference to the given TabularProjection and assigns it to the Projection field.
func (o *BarChartBackendRequest20270101OneOf4454436672) SetProjection(v TabularProjection) {
	o.Projection = &v
}

// GetQuery returns the Query field value.
func (o *BarChartBackendRequest20270101OneOf4454436672) GetQuery() BarChartBackendRequest20270101OneOfQuery {
	if o == nil {
		var ret BarChartBackendRequest20270101OneOfQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *BarChartBackendRequest20270101OneOf4454436672) GetQueryOk() (*BarChartBackendRequest20270101OneOfQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *BarChartBackendRequest20270101OneOf4454436672) SetQuery(v BarChartBackendRequest20270101OneOfQuery) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *BarChartBackendRequest20270101OneOf4454436672) GetRequestType() BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544 {
	if o == nil {
		var ret BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *BarChartBackendRequest20270101OneOf4454436672) GetRequestTypeOk() (*BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *BarChartBackendRequest20270101OneOf4454436672) SetRequestType(v BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544) {
	o.RequestType = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *BarChartBackendRequest20270101OneOf4454436672) GetResponseFormat() BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864 {
	if o == nil {
		var ret BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *BarChartBackendRequest20270101OneOf4454436672) GetResponseFormatOk() (*BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *BarChartBackendRequest20270101OneOf4454436672) SetResponseFormat(v BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864) {
	o.ResponseFormat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BarChartBackendRequest20270101OneOf4454436672) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	if o.Projection != nil {
		toSerialize["projection"] = o.Projection
	}
	toSerialize["query"] = o.Query
	toSerialize["request_type"] = o.RequestType
	toSerialize["response_format"] = o.ResponseFormat
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BarChartBackendRequest20270101OneOf4454436672) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConditionalFormats []ThresholdFormattingRule                                     `json:"conditional_formats,omitempty"`
		Projection         *TabularProjection                                            `json:"projection,omitempty"`
		Query              *BarChartBackendRequest20270101OneOfQuery                     `json:"query"`
		RequestType        *BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544    `json:"request_type"`
		ResponseFormat     *BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864 `json:"response_format"`
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
	o.ConditionalFormats = all.ConditionalFormats
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

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
