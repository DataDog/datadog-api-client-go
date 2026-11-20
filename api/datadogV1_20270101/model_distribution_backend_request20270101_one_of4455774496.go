// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionBackendRequest20270101OneOf4455774496 Test-only generated schema for the dashboard versioning proof of concept.
type DistributionBackendRequest20270101OneOf4455774496 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	NumberFormat *WidgetNumberFormat `json:"number_format,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query FormulaApmStatsQuery `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944 `json:"request_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *RequestStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDistributionBackendRequest20270101OneOf4455774496 instantiates a new DistributionBackendRequest20270101OneOf4455774496 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDistributionBackendRequest20270101OneOf4455774496(query FormulaApmStatsQuery, requestType DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944) *DistributionBackendRequest20270101OneOf4455774496 {
	this := DistributionBackendRequest20270101OneOf4455774496{}
	this.Query = query
	this.RequestType = requestType
	return &this
}

// NewDistributionBackendRequest20270101OneOf4455774496WithDefaults instantiates a new DistributionBackendRequest20270101OneOf4455774496 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDistributionBackendRequest20270101OneOf4455774496WithDefaults() *DistributionBackendRequest20270101OneOf4455774496 {
	this := DistributionBackendRequest20270101OneOf4455774496{}
	return &this
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *DistributionBackendRequest20270101OneOf4455774496) GetNumberFormat() WidgetNumberFormat {
	if o == nil || o.NumberFormat == nil {
		var ret WidgetNumberFormat
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455774496) GetNumberFormatOk() (*WidgetNumberFormat, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *DistributionBackendRequest20270101OneOf4455774496) HasNumberFormat() bool {
	return o != nil && o.NumberFormat != nil
}

// SetNumberFormat gets a reference to the given WidgetNumberFormat and assigns it to the NumberFormat field.
func (o *DistributionBackendRequest20270101OneOf4455774496) SetNumberFormat(v WidgetNumberFormat) {
	o.NumberFormat = &v
}

// GetQuery returns the Query field value.
func (o *DistributionBackendRequest20270101OneOf4455774496) GetQuery() FormulaApmStatsQuery {
	if o == nil {
		var ret FormulaApmStatsQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455774496) GetQueryOk() (*FormulaApmStatsQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *DistributionBackendRequest20270101OneOf4455774496) SetQuery(v FormulaApmStatsQuery) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *DistributionBackendRequest20270101OneOf4455774496) GetRequestType() DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944 {
	if o == nil {
		var ret DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455774496) GetRequestTypeOk() (*DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *DistributionBackendRequest20270101OneOf4455774496) SetRequestType(v DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944) {
	o.RequestType = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *DistributionBackendRequest20270101OneOf4455774496) GetStyle() RequestStyle {
	if o == nil || o.Style == nil {
		var ret RequestStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455774496) GetStyleOk() (*RequestStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *DistributionBackendRequest20270101OneOf4455774496) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given RequestStyle and assigns it to the Style field.
func (o *DistributionBackendRequest20270101OneOf4455774496) SetStyle(v RequestStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DistributionBackendRequest20270101OneOf4455774496) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.NumberFormat != nil {
		toSerialize["number_format"] = o.NumberFormat
	}
	toSerialize["query"] = o.Query
	toSerialize["request_type"] = o.RequestType
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DistributionBackendRequest20270101OneOf4455774496) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NumberFormat *WidgetNumberFormat                                            `json:"number_format,omitempty"`
		Query        *FormulaApmStatsQuery                                          `json:"query"`
		RequestType  *DistributionBackendRequestAnyOfAnyOfItemRequestType4445470944 `json:"request_type"`
		Style        *RequestStyle                                                  `json:"style,omitempty"`
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

	hasInvalidField := false
	if all.NumberFormat != nil && all.NumberFormat.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NumberFormat = all.NumberFormat
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

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
