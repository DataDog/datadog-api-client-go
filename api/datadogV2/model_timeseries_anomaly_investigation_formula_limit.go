// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationFormulaLimit Optional formula limit accepted for compatibility with Timeseries API requests. Formula limits have no effect on timeseries queries.
type TimeseriesAnomalyInvestigationFormulaLimit struct {
	// Requested result limit. This field has no effect on a timeseries anomaly investigation.
	Count *int64 `json:"count,omitempty"`
	// Sort order used when applying a formula series limit.
	Order *TimeseriesAnomalyInvestigationFormulaLimitOrder `json:"order,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationFormulaLimit instantiates a new TimeseriesAnomalyInvestigationFormulaLimit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationFormulaLimit() *TimeseriesAnomalyInvestigationFormulaLimit {
	this := TimeseriesAnomalyInvestigationFormulaLimit{}
	return &this
}

// NewTimeseriesAnomalyInvestigationFormulaLimitWithDefaults instantiates a new TimeseriesAnomalyInvestigationFormulaLimit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationFormulaLimitWithDefaults() *TimeseriesAnomalyInvestigationFormulaLimit {
	this := TimeseriesAnomalyInvestigationFormulaLimit{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *TimeseriesAnomalyInvestigationFormulaLimit) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationFormulaLimit) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *TimeseriesAnomalyInvestigationFormulaLimit) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *TimeseriesAnomalyInvestigationFormulaLimit) SetCount(v int64) {
	o.Count = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *TimeseriesAnomalyInvestigationFormulaLimit) GetOrder() TimeseriesAnomalyInvestigationFormulaLimitOrder {
	if o == nil || o.Order == nil {
		var ret TimeseriesAnomalyInvestigationFormulaLimitOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationFormulaLimit) GetOrderOk() (*TimeseriesAnomalyInvestigationFormulaLimitOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *TimeseriesAnomalyInvestigationFormulaLimit) HasOrder() bool {
	return o != nil && o.Order != nil
}

// SetOrder gets a reference to the given TimeseriesAnomalyInvestigationFormulaLimitOrder and assigns it to the Order field.
func (o *TimeseriesAnomalyInvestigationFormulaLimit) SetOrder(v TimeseriesAnomalyInvestigationFormulaLimitOrder) {
	o.Order = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationFormulaLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesAnomalyInvestigationFormulaLimit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count *int64                                           `json:"count,omitempty"`
		Order *TimeseriesAnomalyInvestigationFormulaLimitOrder `json:"order,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count", "order"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Count = all.Count
	if all.Order != nil && !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = all.Order
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
