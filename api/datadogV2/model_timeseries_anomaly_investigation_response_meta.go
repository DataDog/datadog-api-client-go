// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationResponseMeta Timeseries execution metadata for the single request accepted by this API version.
type TimeseriesAnomalyInvestigationResponseMeta struct {
	// Effective start of the timeseries query in milliseconds since the Unix epoch.
	FromDate int64 `json:"from_date"`
	// Effective timeseries interval in milliseconds.
	Interval int64 `json:"interval"`
	// Execution status for the request's queries.
	Queries []TimeseriesAnomalyInvestigationQueryStatus `json:"queries"`
	// Response metadata type for a timeseries anomaly investigation.
	ResType TimeseriesAnomalyInvestigationMetaType `json:"res_type"`
	// Non-fatal warnings produced while executing the investigation.
	ResultsWarnings []TimeseriesAnomalyInvestigationResultsWarning `json:"results_warnings"`
	// Effective end of the timeseries query in milliseconds since the Unix epoch.
	ToDate int64 `json:"to_date"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationResponseMeta instantiates a new TimeseriesAnomalyInvestigationResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationResponseMeta(fromDate int64, interval int64, queries []TimeseriesAnomalyInvestigationQueryStatus, resType TimeseriesAnomalyInvestigationMetaType, resultsWarnings []TimeseriesAnomalyInvestigationResultsWarning, toDate int64) *TimeseriesAnomalyInvestigationResponseMeta {
	this := TimeseriesAnomalyInvestigationResponseMeta{}
	this.FromDate = fromDate
	this.Interval = interval
	this.Queries = queries
	this.ResType = resType
	this.ResultsWarnings = resultsWarnings
	this.ToDate = toDate
	return &this
}

// NewTimeseriesAnomalyInvestigationResponseMetaWithDefaults instantiates a new TimeseriesAnomalyInvestigationResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationResponseMetaWithDefaults() *TimeseriesAnomalyInvestigationResponseMeta {
	this := TimeseriesAnomalyInvestigationResponseMeta{}
	return &this
}

// GetFromDate returns the FromDate field value.
func (o *TimeseriesAnomalyInvestigationResponseMeta) GetFromDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FromDate
}

// GetFromDateOk returns a tuple with the FromDate field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationResponseMeta) GetFromDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromDate, true
}

// SetFromDate sets field value.
func (o *TimeseriesAnomalyInvestigationResponseMeta) SetFromDate(v int64) {
	o.FromDate = v
}

// GetInterval returns the Interval field value.
func (o *TimeseriesAnomalyInvestigationResponseMeta) GetInterval() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationResponseMeta) GetIntervalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value.
func (o *TimeseriesAnomalyInvestigationResponseMeta) SetInterval(v int64) {
	o.Interval = v
}

// GetQueries returns the Queries field value.
func (o *TimeseriesAnomalyInvestigationResponseMeta) GetQueries() []TimeseriesAnomalyInvestigationQueryStatus {
	if o == nil {
		var ret []TimeseriesAnomalyInvestigationQueryStatus
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationResponseMeta) GetQueriesOk() (*[]TimeseriesAnomalyInvestigationQueryStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *TimeseriesAnomalyInvestigationResponseMeta) SetQueries(v []TimeseriesAnomalyInvestigationQueryStatus) {
	o.Queries = v
}

// GetResType returns the ResType field value.
func (o *TimeseriesAnomalyInvestigationResponseMeta) GetResType() TimeseriesAnomalyInvestigationMetaType {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationMetaType
		return ret
	}
	return o.ResType
}

// GetResTypeOk returns a tuple with the ResType field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationResponseMeta) GetResTypeOk() (*TimeseriesAnomalyInvestigationMetaType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResType, true
}

// SetResType sets field value.
func (o *TimeseriesAnomalyInvestigationResponseMeta) SetResType(v TimeseriesAnomalyInvestigationMetaType) {
	o.ResType = v
}

// GetResultsWarnings returns the ResultsWarnings field value.
func (o *TimeseriesAnomalyInvestigationResponseMeta) GetResultsWarnings() []TimeseriesAnomalyInvestigationResultsWarning {
	if o == nil {
		var ret []TimeseriesAnomalyInvestigationResultsWarning
		return ret
	}
	return o.ResultsWarnings
}

// GetResultsWarningsOk returns a tuple with the ResultsWarnings field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationResponseMeta) GetResultsWarningsOk() (*[]TimeseriesAnomalyInvestigationResultsWarning, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultsWarnings, true
}

// SetResultsWarnings sets field value.
func (o *TimeseriesAnomalyInvestigationResponseMeta) SetResultsWarnings(v []TimeseriesAnomalyInvestigationResultsWarning) {
	o.ResultsWarnings = v
}

// GetToDate returns the ToDate field value.
func (o *TimeseriesAnomalyInvestigationResponseMeta) GetToDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ToDate
}

// GetToDateOk returns a tuple with the ToDate field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationResponseMeta) GetToDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToDate, true
}

// SetToDate sets field value.
func (o *TimeseriesAnomalyInvestigationResponseMeta) SetToDate(v int64) {
	o.ToDate = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["from_date"] = o.FromDate
	toSerialize["interval"] = o.Interval
	toSerialize["queries"] = o.Queries
	toSerialize["res_type"] = o.ResType
	toSerialize["results_warnings"] = o.ResultsWarnings
	toSerialize["to_date"] = o.ToDate

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesAnomalyInvestigationResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FromDate        *int64                                          `json:"from_date"`
		Interval        *int64                                          `json:"interval"`
		Queries         *[]TimeseriesAnomalyInvestigationQueryStatus    `json:"queries"`
		ResType         *TimeseriesAnomalyInvestigationMetaType         `json:"res_type"`
		ResultsWarnings *[]TimeseriesAnomalyInvestigationResultsWarning `json:"results_warnings"`
		ToDate          *int64                                          `json:"to_date"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FromDate == nil {
		return fmt.Errorf("required field from_date missing")
	}
	if all.Interval == nil {
		return fmt.Errorf("required field interval missing")
	}
	if all.Queries == nil {
		return fmt.Errorf("required field queries missing")
	}
	if all.ResType == nil {
		return fmt.Errorf("required field res_type missing")
	}
	if all.ResultsWarnings == nil {
		return fmt.Errorf("required field results_warnings missing")
	}
	if all.ToDate == nil {
		return fmt.Errorf("required field to_date missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"from_date", "interval", "queries", "res_type", "results_warnings", "to_date"})
	} else {
		return err
	}

	hasInvalidField := false
	o.FromDate = *all.FromDate
	o.Interval = *all.Interval
	o.Queries = *all.Queries
	if !all.ResType.IsValid() {
		hasInvalidField = true
	} else {
		o.ResType = *all.ResType
	}
	o.ResultsWarnings = *all.ResultsWarnings
	o.ToDate = *all.ToDate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
