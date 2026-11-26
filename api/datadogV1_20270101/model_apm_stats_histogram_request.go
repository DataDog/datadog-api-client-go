// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmStatsHistogramRequest Test-only generated schema for the dashboard versioning proof of concept.
type ApmStatsHistogramRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	NumberFormat *WidgetNumberFormat `json:"number_format,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query FormulaApmStatsQuery `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RequestType ApmStatsHistogramRequestRequestType `json:"request_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewApmStatsHistogramRequest instantiates a new ApmStatsHistogramRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApmStatsHistogramRequest(query FormulaApmStatsQuery, requestType ApmStatsHistogramRequestRequestType) *ApmStatsHistogramRequest {
	this := ApmStatsHistogramRequest{}
	this.Query = query
	this.RequestType = requestType
	return &this
}

// NewApmStatsHistogramRequestWithDefaults instantiates a new ApmStatsHistogramRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApmStatsHistogramRequestWithDefaults() *ApmStatsHistogramRequest {
	this := ApmStatsHistogramRequest{}
	return &this
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *ApmStatsHistogramRequest) GetNumberFormat() WidgetNumberFormat {
	if o == nil || o.NumberFormat == nil {
		var ret WidgetNumberFormat
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmStatsHistogramRequest) GetNumberFormatOk() (*WidgetNumberFormat, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *ApmStatsHistogramRequest) HasNumberFormat() bool {
	return o != nil && o.NumberFormat != nil
}

// SetNumberFormat gets a reference to the given WidgetNumberFormat and assigns it to the NumberFormat field.
func (o *ApmStatsHistogramRequest) SetNumberFormat(v WidgetNumberFormat) {
	o.NumberFormat = &v
}

// GetQuery returns the Query field value.
func (o *ApmStatsHistogramRequest) GetQuery() FormulaApmStatsQuery {
	if o == nil {
		var ret FormulaApmStatsQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ApmStatsHistogramRequest) GetQueryOk() (*FormulaApmStatsQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ApmStatsHistogramRequest) SetQuery(v FormulaApmStatsQuery) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *ApmStatsHistogramRequest) GetRequestType() ApmStatsHistogramRequestRequestType {
	if o == nil {
		var ret ApmStatsHistogramRequestRequestType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *ApmStatsHistogramRequest) GetRequestTypeOk() (*ApmStatsHistogramRequestRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *ApmStatsHistogramRequest) SetRequestType(v ApmStatsHistogramRequestRequestType) {
	o.RequestType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApmStatsHistogramRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.NumberFormat != nil {
		toSerialize["number_format"] = o.NumberFormat
	}
	toSerialize["query"] = o.Query
	toSerialize["request_type"] = o.RequestType
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApmStatsHistogramRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NumberFormat *WidgetNumberFormat                  `json:"number_format,omitempty"`
		Query        *FormulaApmStatsQuery                `json:"query"`
		RequestType  *ApmStatsHistogramRequestRequestType `json:"request_type"`
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

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
