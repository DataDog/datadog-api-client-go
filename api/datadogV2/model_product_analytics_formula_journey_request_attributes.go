// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFormulaJourneyRequestAttributes Attributes of a journey timeseries request.
type ProductAnalyticsFormulaJourneyRequestAttributes struct {
	// Start of the query window, in epoch milliseconds.
	From int64 `json:"from"`
	// Time bucket interval in milliseconds.
	Interval *int64 `json:"interval,omitempty"`
	// Query definition for a journey timeseries request.
	Query ProductAnalyticsFormulaJourneyQuery `json:"query"`
	// End of the query window, in epoch milliseconds.
	To int64 `json:"to"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsFormulaJourneyRequestAttributes instantiates a new ProductAnalyticsFormulaJourneyRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsFormulaJourneyRequestAttributes(from int64, query ProductAnalyticsFormulaJourneyQuery, to int64) *ProductAnalyticsFormulaJourneyRequestAttributes {
	this := ProductAnalyticsFormulaJourneyRequestAttributes{}
	this.From = from
	this.Query = query
	this.To = to
	return &this
}

// NewProductAnalyticsFormulaJourneyRequestAttributesWithDefaults instantiates a new ProductAnalyticsFormulaJourneyRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsFormulaJourneyRequestAttributesWithDefaults() *ProductAnalyticsFormulaJourneyRequestAttributes {
	this := ProductAnalyticsFormulaJourneyRequestAttributes{}
	return &this
}

// GetFrom returns the From field value.
func (o *ProductAnalyticsFormulaJourneyRequestAttributes) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFormulaJourneyRequestAttributes) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *ProductAnalyticsFormulaJourneyRequestAttributes) SetFrom(v int64) {
	o.From = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *ProductAnalyticsFormulaJourneyRequestAttributes) GetInterval() int64 {
	if o == nil || o.Interval == nil {
		var ret int64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFormulaJourneyRequestAttributes) GetIntervalOk() (*int64, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *ProductAnalyticsFormulaJourneyRequestAttributes) HasInterval() bool {
	return o != nil && o.Interval != nil
}

// SetInterval gets a reference to the given int64 and assigns it to the Interval field.
func (o *ProductAnalyticsFormulaJourneyRequestAttributes) SetInterval(v int64) {
	o.Interval = &v
}

// GetQuery returns the Query field value.
func (o *ProductAnalyticsFormulaJourneyRequestAttributes) GetQuery() ProductAnalyticsFormulaJourneyQuery {
	if o == nil {
		var ret ProductAnalyticsFormulaJourneyQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFormulaJourneyRequestAttributes) GetQueryOk() (*ProductAnalyticsFormulaJourneyQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ProductAnalyticsFormulaJourneyRequestAttributes) SetQuery(v ProductAnalyticsFormulaJourneyQuery) {
	o.Query = v
}

// GetTo returns the To field value.
func (o *ProductAnalyticsFormulaJourneyRequestAttributes) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFormulaJourneyRequestAttributes) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *ProductAnalyticsFormulaJourneyRequestAttributes) SetTo(v int64) {
	o.To = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsFormulaJourneyRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["from"] = o.From
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	toSerialize["query"] = o.Query
	toSerialize["to"] = o.To

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsFormulaJourneyRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From     *int64                               `json:"from"`
		Interval *int64                               `json:"interval,omitempty"`
		Query    *ProductAnalyticsFormulaJourneyQuery `json:"query"`
		To       *int64                               `json:"to"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"from", "interval", "query", "to"})
	} else {
		return err
	}

	hasInvalidField := false
	o.From = *all.From
	o.Interval = all.Interval
	if all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = *all.Query
	o.To = *all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
