// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendPointsRequest20270101OneOf4457230272 Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendPointsRequest20270101OneOf4457230272 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *float64 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Projection GeomapPointsProjection `json:"projection"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query DataProjectionQuery `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616 `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *GeomapBackendPointsRequestAnyOfAnyOfItemStyle `json:"style,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TextFormats []TextFormattingRule `json:"text_formats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGeomapBackendPointsRequest20270101OneOf4457230272 instantiates a new GeomapBackendPointsRequest20270101OneOf4457230272 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeomapBackendPointsRequest20270101OneOf4457230272(projection GeomapPointsProjection, query DataProjectionQuery, requestType GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616) *GeomapBackendPointsRequest20270101OneOf4457230272 {
	this := GeomapBackendPointsRequest20270101OneOf4457230272{}
	this.Projection = projection
	this.Query = query
	this.RequestType = requestType
	return &this
}

// NewGeomapBackendPointsRequest20270101OneOf4457230272WithDefaults instantiates a new GeomapBackendPointsRequest20270101OneOf4457230272 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeomapBackendPointsRequest20270101OneOf4457230272WithDefaults() *GeomapBackendPointsRequest20270101OneOf4457230272 {
	this := GeomapBackendPointsRequest20270101OneOf4457230272{}
	return &this
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) GetConditionalFormats() []ThresholdFormattingRule {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) GetConditionalFormatsOk() (*[]ThresholdFormattingRule, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule and assigns it to the ConditionalFormats field.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) SetConditionalFormats(v []ThresholdFormattingRule) {
	o.ConditionalFormats = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) GetLimit() float64 {
	if o == nil || o.Limit == nil {
		var ret float64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) GetLimitOk() (*float64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given float64 and assigns it to the Limit field.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) SetLimit(v float64) {
	o.Limit = &v
}

// GetProjection returns the Projection field value.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) GetProjection() GeomapPointsProjection {
	if o == nil {
		var ret GeomapPointsProjection
		return ret
	}
	return o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value
// and a boolean to check if the value has been set.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) GetProjectionOk() (*GeomapPointsProjection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Projection, true
}

// SetProjection sets field value.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) SetProjection(v GeomapPointsProjection) {
	o.Projection = v
}

// GetQuery returns the Query field value.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) GetQuery() DataProjectionQuery {
	if o == nil {
		var ret DataProjectionQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) GetQueryOk() (*DataProjectionQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) SetQuery(v DataProjectionQuery) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) GetRequestType() GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616 {
	if o == nil {
		var ret GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) GetRequestTypeOk() (*GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) SetRequestType(v GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616) {
	o.RequestType = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) GetStyle() GeomapBackendPointsRequestAnyOfAnyOfItemStyle {
	if o == nil || o.Style == nil {
		var ret GeomapBackendPointsRequestAnyOfAnyOfItemStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) GetStyleOk() (*GeomapBackendPointsRequestAnyOfAnyOfItemStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given GeomapBackendPointsRequestAnyOfAnyOfItemStyle and assigns it to the Style field.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) SetStyle(v GeomapBackendPointsRequestAnyOfAnyOfItemStyle) {
	o.Style = &v
}

// GetTextFormats returns the TextFormats field value if set, zero value otherwise.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) GetTextFormats() []TextFormattingRule {
	if o == nil || o.TextFormats == nil {
		var ret []TextFormattingRule
		return ret
	}
	return o.TextFormats
}

// GetTextFormatsOk returns a tuple with the TextFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) GetTextFormatsOk() (*[]TextFormattingRule, bool) {
	if o == nil || o.TextFormats == nil {
		return nil, false
	}
	return &o.TextFormats, true
}

// HasTextFormats returns a boolean if a field has been set.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) HasTextFormats() bool {
	return o != nil && o.TextFormats != nil
}

// SetTextFormats gets a reference to the given []TextFormattingRule and assigns it to the TextFormats field.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) SetTextFormats(v []TextFormattingRule) {
	o.TextFormats = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeomapBackendPointsRequest20270101OneOf4457230272) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	toSerialize["projection"] = o.Projection
	toSerialize["query"] = o.Query
	toSerialize["request_type"] = o.RequestType
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	if o.TextFormats != nil {
		toSerialize["text_formats"] = o.TextFormats
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeomapBackendPointsRequest20270101OneOf4457230272) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConditionalFormats []ThresholdFormattingRule                                      `json:"conditional_formats,omitempty"`
		Limit              *float64                                                       `json:"limit,omitempty"`
		Projection         *GeomapPointsProjection                                        `json:"projection"`
		Query              *DataProjectionQuery                                           `json:"query"`
		RequestType        *GeomapBackendPointsRequestAnyOfAnyOfItemRequestType4447003616 `json:"request_type"`
		Style              *GeomapBackendPointsRequestAnyOfAnyOfItemStyle                 `json:"style,omitempty"`
		TextFormats        []TextFormattingRule                                           `json:"text_formats,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Projection == nil {
		return fmt.Errorf("required field projection missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.RequestType == nil {
		return fmt.Errorf("required field request_type missing")
	}

	hasInvalidField := false
	o.ConditionalFormats = all.ConditionalFormats
	o.Limit = all.Limit
	if all.Projection.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Projection = *all.Projection
	o.Query = *all.Query
	if !all.RequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.RequestType = *all.RequestType
	}
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style
	o.TextFormats = all.TextFormats

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
