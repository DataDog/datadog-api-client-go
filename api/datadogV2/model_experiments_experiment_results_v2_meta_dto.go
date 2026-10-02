// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentResultsV2MetaDTO Information about when experiment results were updated and whether they are stale.
type ExperimentsExperimentResultsV2MetaDTO struct {
	// Whether the saved results require a refresh or their freshness cannot be confirmed. See stale_reasons for
	// details.
	IsStale *bool `json:"is_stale,omitempty"`
	// Time when the experiment results were last updated.
	ResultsLastUpdated datadog.NullableTime `json:"results_last_updated,omitempty"`
	// Reasons the saved experiment results are stale.
	StaleReasons []string `json:"stale_reasons,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsExperimentResultsV2MetaDTO instantiates a new ExperimentsExperimentResultsV2MetaDTO object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsExperimentResultsV2MetaDTO() *ExperimentsExperimentResultsV2MetaDTO {
	this := ExperimentsExperimentResultsV2MetaDTO{}
	return &this
}

// NewExperimentsExperimentResultsV2MetaDTOWithDefaults instantiates a new ExperimentsExperimentResultsV2MetaDTO object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsExperimentResultsV2MetaDTOWithDefaults() *ExperimentsExperimentResultsV2MetaDTO {
	this := ExperimentsExperimentResultsV2MetaDTO{}
	return &this
}

// GetIsStale returns the IsStale field value if set, zero value otherwise.
func (o *ExperimentsExperimentResultsV2MetaDTO) GetIsStale() bool {
	if o == nil || o.IsStale == nil {
		var ret bool
		return ret
	}
	return *o.IsStale
}

// GetIsStaleOk returns a tuple with the IsStale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentResultsV2MetaDTO) GetIsStaleOk() (*bool, bool) {
	if o == nil || o.IsStale == nil {
		return nil, false
	}
	return o.IsStale, true
}

// HasIsStale returns a boolean if a field has been set.
func (o *ExperimentsExperimentResultsV2MetaDTO) HasIsStale() bool {
	return o != nil && o.IsStale != nil
}

// SetIsStale gets a reference to the given bool and assigns it to the IsStale field.
func (o *ExperimentsExperimentResultsV2MetaDTO) SetIsStale(v bool) {
	o.IsStale = &v
}

// GetResultsLastUpdated returns the ResultsLastUpdated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsExperimentResultsV2MetaDTO) GetResultsLastUpdated() time.Time {
	if o == nil || o.ResultsLastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ResultsLastUpdated.Get()
}

// GetResultsLastUpdatedOk returns a tuple with the ResultsLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsExperimentResultsV2MetaDTO) GetResultsLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResultsLastUpdated.Get(), o.ResultsLastUpdated.IsSet()
}

// HasResultsLastUpdated returns a boolean if a field has been set.
func (o *ExperimentsExperimentResultsV2MetaDTO) HasResultsLastUpdated() bool {
	return o != nil && o.ResultsLastUpdated.IsSet()
}

// SetResultsLastUpdated gets a reference to the given datadog.NullableTime and assigns it to the ResultsLastUpdated field.
func (o *ExperimentsExperimentResultsV2MetaDTO) SetResultsLastUpdated(v time.Time) {
	o.ResultsLastUpdated.Set(&v)
}

// SetResultsLastUpdatedNil sets the value for ResultsLastUpdated to be an explicit nil.
func (o *ExperimentsExperimentResultsV2MetaDTO) SetResultsLastUpdatedNil() {
	o.ResultsLastUpdated.Set(nil)
}

// UnsetResultsLastUpdated ensures that no value is present for ResultsLastUpdated, not even an explicit nil.
func (o *ExperimentsExperimentResultsV2MetaDTO) UnsetResultsLastUpdated() {
	o.ResultsLastUpdated.Unset()
}

// GetStaleReasons returns the StaleReasons field value if set, zero value otherwise.
func (o *ExperimentsExperimentResultsV2MetaDTO) GetStaleReasons() []string {
	if o == nil || o.StaleReasons == nil {
		var ret []string
		return ret
	}
	return o.StaleReasons
}

// GetStaleReasonsOk returns a tuple with the StaleReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentResultsV2MetaDTO) GetStaleReasonsOk() (*[]string, bool) {
	if o == nil || o.StaleReasons == nil {
		return nil, false
	}
	return &o.StaleReasons, true
}

// HasStaleReasons returns a boolean if a field has been set.
func (o *ExperimentsExperimentResultsV2MetaDTO) HasStaleReasons() bool {
	return o != nil && o.StaleReasons != nil
}

// SetStaleReasons gets a reference to the given []string and assigns it to the StaleReasons field.
func (o *ExperimentsExperimentResultsV2MetaDTO) SetStaleReasons(v []string) {
	o.StaleReasons = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsExperimentResultsV2MetaDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IsStale != nil {
		toSerialize["is_stale"] = o.IsStale
	}
	if o.ResultsLastUpdated.IsSet() {
		toSerialize["results_last_updated"] = o.ResultsLastUpdated.Get()
	}
	if o.StaleReasons != nil {
		toSerialize["stale_reasons"] = o.StaleReasons
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsExperimentResultsV2MetaDTO) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsStale            *bool                `json:"is_stale,omitempty"`
		ResultsLastUpdated datadog.NullableTime `json:"results_last_updated,omitempty"`
		StaleReasons       []string             `json:"stale_reasons,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_stale", "results_last_updated", "stale_reasons"})
	} else {
		return err
	}
	o.IsStale = all.IsStale
	o.ResultsLastUpdated = all.ResultsLastUpdated
	o.StaleReasons = all.StaleReasons

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
