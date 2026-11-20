// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryValueBackendRequest20270101OneOf4460685056 Test-only generated schema for the dashboard versioning proof of concept.
type QueryValueBackendRequest20270101OneOf4460685056 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregator *ScalarAggregator `json:"aggregator,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Comparison *ComparisonRequest `json:"comparison,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DdsqlQuery string `json:"ddsql_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776 `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096 `json:"response_format"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewQueryValueBackendRequest20270101OneOf4460685056 instantiates a new QueryValueBackendRequest20270101OneOf4460685056 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryValueBackendRequest20270101OneOf4460685056(ddsqlQuery string, requestType QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776, responseFormat QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096) *QueryValueBackendRequest20270101OneOf4460685056 {
	this := QueryValueBackendRequest20270101OneOf4460685056{}
	this.DdsqlQuery = ddsqlQuery
	this.RequestType = requestType
	this.ResponseFormat = responseFormat
	return &this
}

// NewQueryValueBackendRequest20270101OneOf4460685056WithDefaults instantiates a new QueryValueBackendRequest20270101OneOf4460685056 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryValueBackendRequest20270101OneOf4460685056WithDefaults() *QueryValueBackendRequest20270101OneOf4460685056 {
	this := QueryValueBackendRequest20270101OneOf4460685056{}
	return &this
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *QueryValueBackendRequest20270101OneOf4460685056) GetAggregator() ScalarAggregator {
	if o == nil || o.Aggregator == nil {
		var ret ScalarAggregator
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460685056) GetAggregatorOk() (*ScalarAggregator, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *QueryValueBackendRequest20270101OneOf4460685056) HasAggregator() bool {
	return o != nil && o.Aggregator != nil
}

// SetAggregator gets a reference to the given ScalarAggregator and assigns it to the Aggregator field.
func (o *QueryValueBackendRequest20270101OneOf4460685056) SetAggregator(v ScalarAggregator) {
	o.Aggregator = &v
}

// GetComparison returns the Comparison field value if set, zero value otherwise.
func (o *QueryValueBackendRequest20270101OneOf4460685056) GetComparison() ComparisonRequest {
	if o == nil || o.Comparison == nil {
		var ret ComparisonRequest
		return ret
	}
	return *o.Comparison
}

// GetComparisonOk returns a tuple with the Comparison field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460685056) GetComparisonOk() (*ComparisonRequest, bool) {
	if o == nil || o.Comparison == nil {
		return nil, false
	}
	return o.Comparison, true
}

// HasComparison returns a boolean if a field has been set.
func (o *QueryValueBackendRequest20270101OneOf4460685056) HasComparison() bool {
	return o != nil && o.Comparison != nil
}

// SetComparison gets a reference to the given ComparisonRequest and assigns it to the Comparison field.
func (o *QueryValueBackendRequest20270101OneOf4460685056) SetComparison(v ComparisonRequest) {
	o.Comparison = &v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *QueryValueBackendRequest20270101OneOf4460685056) GetConditionalFormats() []ThresholdFormattingRule {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460685056) GetConditionalFormatsOk() (*[]ThresholdFormattingRule, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *QueryValueBackendRequest20270101OneOf4460685056) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule and assigns it to the ConditionalFormats field.
func (o *QueryValueBackendRequest20270101OneOf4460685056) SetConditionalFormats(v []ThresholdFormattingRule) {
	o.ConditionalFormats = v
}

// GetDdsqlQuery returns the DdsqlQuery field value.
func (o *QueryValueBackendRequest20270101OneOf4460685056) GetDdsqlQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DdsqlQuery
}

// GetDdsqlQueryOk returns a tuple with the DdsqlQuery field value
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460685056) GetDdsqlQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DdsqlQuery, true
}

// SetDdsqlQuery sets field value.
func (o *QueryValueBackendRequest20270101OneOf4460685056) SetDdsqlQuery(v string) {
	o.DdsqlQuery = v
}

// GetRequestType returns the RequestType field value.
func (o *QueryValueBackendRequest20270101OneOf4460685056) GetRequestType() QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776 {
	if o == nil {
		var ret QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460685056) GetRequestTypeOk() (*QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *QueryValueBackendRequest20270101OneOf4460685056) SetRequestType(v QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776) {
	o.RequestType = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *QueryValueBackendRequest20270101OneOf4460685056) GetResponseFormat() QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096 {
	if o == nil {
		var ret QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460685056) GetResponseFormatOk() (*QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *QueryValueBackendRequest20270101OneOf4460685056) SetResponseFormat(v QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096) {
	o.ResponseFormat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryValueBackendRequest20270101OneOf4460685056) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aggregator != nil {
		toSerialize["aggregator"] = o.Aggregator
	}
	if o.Comparison != nil {
		toSerialize["comparison"] = o.Comparison
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
func (o *QueryValueBackendRequest20270101OneOf4460685056) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregator         *ScalarAggregator                                               `json:"aggregator,omitempty"`
		Comparison         *ComparisonRequest                                              `json:"comparison,omitempty"`
		ConditionalFormats []ThresholdFormattingRule                                       `json:"conditional_formats,omitempty"`
		DdsqlQuery         *string                                                         `json:"ddsql_query"`
		RequestType        *QueryValueBackendRequestAnyOfAnyOfItemRequestType4450439776    `json:"request_type"`
		ResponseFormat     *QueryValueBackendRequestAnyOfAnyOfItemResponseFormat4450440096 `json:"response_format"`
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
	if all.Aggregator != nil && !all.Aggregator.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregator = all.Aggregator
	}
	if all.Comparison != nil && all.Comparison.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Comparison = all.Comparison
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
