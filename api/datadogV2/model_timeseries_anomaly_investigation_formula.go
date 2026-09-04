// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationFormula Formula evaluated by the timeseries request.
type TimeseriesAnomalyInvestigationFormula struct {
	// Formula expression referencing one or more named queries.
	Formula string `json:"formula"`
	// Optional formula limit accepted for compatibility with Timeseries API requests. Formula limits have no effect on timeseries queries.
	Limit *TimeseriesAnomalyInvestigationFormulaLimit `json:"limit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationFormula instantiates a new TimeseriesAnomalyInvestigationFormula object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationFormula(formula string) *TimeseriesAnomalyInvestigationFormula {
	this := TimeseriesAnomalyInvestigationFormula{}
	this.Formula = formula
	return &this
}

// NewTimeseriesAnomalyInvestigationFormulaWithDefaults instantiates a new TimeseriesAnomalyInvestigationFormula object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationFormulaWithDefaults() *TimeseriesAnomalyInvestigationFormula {
	this := TimeseriesAnomalyInvestigationFormula{}
	return &this
}

// GetFormula returns the Formula field value.
func (o *TimeseriesAnomalyInvestigationFormula) GetFormula() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Formula
}

// GetFormulaOk returns a tuple with the Formula field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationFormula) GetFormulaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formula, true
}

// SetFormula sets field value.
func (o *TimeseriesAnomalyInvestigationFormula) SetFormula(v string) {
	o.Formula = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *TimeseriesAnomalyInvestigationFormula) GetLimit() TimeseriesAnomalyInvestigationFormulaLimit {
	if o == nil || o.Limit == nil {
		var ret TimeseriesAnomalyInvestigationFormulaLimit
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationFormula) GetLimitOk() (*TimeseriesAnomalyInvestigationFormulaLimit, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *TimeseriesAnomalyInvestigationFormula) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given TimeseriesAnomalyInvestigationFormulaLimit and assigns it to the Limit field.
func (o *TimeseriesAnomalyInvestigationFormula) SetLimit(v TimeseriesAnomalyInvestigationFormulaLimit) {
	o.Limit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationFormula) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["formula"] = o.Formula
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesAnomalyInvestigationFormula) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Formula *string                                     `json:"formula"`
		Limit   *TimeseriesAnomalyInvestigationFormulaLimit `json:"limit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Formula == nil {
		return fmt.Errorf("required field formula missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"formula", "limit"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Formula = *all.Formula
	if all.Limit != nil && all.Limit.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Limit = all.Limit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
