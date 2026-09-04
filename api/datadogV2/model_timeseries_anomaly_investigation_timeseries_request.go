// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationTimeseriesRequest Metrics timeseries request to investigate.
type TimeseriesAnomalyInvestigationTimeseriesRequest struct {
	// Formulas to evaluate. Each formula may contain an explicit `anomalies()` call or a supported metrics expression.
	Formulas []TimeseriesAnomalyInvestigationFormula `json:"formulas"`
	// Start of the investigation time window in milliseconds since the Unix epoch.
	From int64 `json:"from"`
	// Optional requested aggregation interval in milliseconds.
	Interval *int64 `json:"interval,omitempty"`
	// Metrics queries referenced by the formulas.
	Queries []TimeseriesAnomalyInvestigationMetricQuery `json:"queries"`
	// End of the investigation time window in milliseconds since the Unix epoch. Must be later than `from`.
	To int64 `json:"to"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationTimeseriesRequest instantiates a new TimeseriesAnomalyInvestigationTimeseriesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationTimeseriesRequest(formulas []TimeseriesAnomalyInvestigationFormula, from int64, queries []TimeseriesAnomalyInvestigationMetricQuery, to int64) *TimeseriesAnomalyInvestigationTimeseriesRequest {
	this := TimeseriesAnomalyInvestigationTimeseriesRequest{}
	this.Formulas = formulas
	this.From = from
	this.Queries = queries
	this.To = to
	return &this
}

// NewTimeseriesAnomalyInvestigationTimeseriesRequestWithDefaults instantiates a new TimeseriesAnomalyInvestigationTimeseriesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationTimeseriesRequestWithDefaults() *TimeseriesAnomalyInvestigationTimeseriesRequest {
	this := TimeseriesAnomalyInvestigationTimeseriesRequest{}
	return &this
}

// GetFormulas returns the Formulas field value.
func (o *TimeseriesAnomalyInvestigationTimeseriesRequest) GetFormulas() []TimeseriesAnomalyInvestigationFormula {
	if o == nil {
		var ret []TimeseriesAnomalyInvestigationFormula
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationTimeseriesRequest) GetFormulasOk() (*[]TimeseriesAnomalyInvestigationFormula, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// SetFormulas sets field value.
func (o *TimeseriesAnomalyInvestigationTimeseriesRequest) SetFormulas(v []TimeseriesAnomalyInvestigationFormula) {
	o.Formulas = v
}

// GetFrom returns the From field value.
func (o *TimeseriesAnomalyInvestigationTimeseriesRequest) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationTimeseriesRequest) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *TimeseriesAnomalyInvestigationTimeseriesRequest) SetFrom(v int64) {
	o.From = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *TimeseriesAnomalyInvestigationTimeseriesRequest) GetInterval() int64 {
	if o == nil || o.Interval == nil {
		var ret int64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationTimeseriesRequest) GetIntervalOk() (*int64, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *TimeseriesAnomalyInvestigationTimeseriesRequest) HasInterval() bool {
	return o != nil && o.Interval != nil
}

// SetInterval gets a reference to the given int64 and assigns it to the Interval field.
func (o *TimeseriesAnomalyInvestigationTimeseriesRequest) SetInterval(v int64) {
	o.Interval = &v
}

// GetQueries returns the Queries field value.
func (o *TimeseriesAnomalyInvestigationTimeseriesRequest) GetQueries() []TimeseriesAnomalyInvestigationMetricQuery {
	if o == nil {
		var ret []TimeseriesAnomalyInvestigationMetricQuery
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationTimeseriesRequest) GetQueriesOk() (*[]TimeseriesAnomalyInvestigationMetricQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *TimeseriesAnomalyInvestigationTimeseriesRequest) SetQueries(v []TimeseriesAnomalyInvestigationMetricQuery) {
	o.Queries = v
}

// GetTo returns the To field value.
func (o *TimeseriesAnomalyInvestigationTimeseriesRequest) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationTimeseriesRequest) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *TimeseriesAnomalyInvestigationTimeseriesRequest) SetTo(v int64) {
	o.To = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationTimeseriesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["formulas"] = o.Formulas
	toSerialize["from"] = o.From
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	toSerialize["queries"] = o.Queries
	toSerialize["to"] = o.To

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesAnomalyInvestigationTimeseriesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Formulas *[]TimeseriesAnomalyInvestigationFormula     `json:"formulas"`
		From     *int64                                       `json:"from"`
		Interval *int64                                       `json:"interval,omitempty"`
		Queries  *[]TimeseriesAnomalyInvestigationMetricQuery `json:"queries"`
		To       *int64                                       `json:"to"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Formulas == nil {
		return fmt.Errorf("required field formulas missing")
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	if all.Queries == nil {
		return fmt.Errorf("required field queries missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"formulas", "from", "interval", "queries", "to"})
	} else {
		return err
	}
	o.Formulas = *all.Formulas
	o.From = *all.From
	o.Interval = all.Interval
	o.Queries = *all.Queries
	o.To = *all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
