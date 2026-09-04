// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationDetection Anomaly detection configuration used for the result.
type TimeseriesAnomalyInvestigationDetection struct {
	// Source of the anomaly detection configuration.
	ConfigurationSource TimeseriesAnomalyInvestigationConfigurationSource `json:"configuration_source"`
	// Applied Watchdog Explains profile, or null when the request supplied an explicit `anomalies()` formula. The current Watchdog profile is `watchdog_explains_v1`.
	Profile datadog.NullableString `json:"profile"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationDetection instantiates a new TimeseriesAnomalyInvestigationDetection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationDetection(configurationSource TimeseriesAnomalyInvestigationConfigurationSource, profile datadog.NullableString) *TimeseriesAnomalyInvestigationDetection {
	this := TimeseriesAnomalyInvestigationDetection{}
	this.ConfigurationSource = configurationSource
	this.Profile = profile
	return &this
}

// NewTimeseriesAnomalyInvestigationDetectionWithDefaults instantiates a new TimeseriesAnomalyInvestigationDetection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationDetectionWithDefaults() *TimeseriesAnomalyInvestigationDetection {
	this := TimeseriesAnomalyInvestigationDetection{}
	return &this
}

// GetConfigurationSource returns the ConfigurationSource field value.
func (o *TimeseriesAnomalyInvestigationDetection) GetConfigurationSource() TimeseriesAnomalyInvestigationConfigurationSource {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationConfigurationSource
		return ret
	}
	return o.ConfigurationSource
}

// GetConfigurationSourceOk returns a tuple with the ConfigurationSource field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationDetection) GetConfigurationSourceOk() (*TimeseriesAnomalyInvestigationConfigurationSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationSource, true
}

// SetConfigurationSource sets field value.
func (o *TimeseriesAnomalyInvestigationDetection) SetConfigurationSource(v TimeseriesAnomalyInvestigationConfigurationSource) {
	o.ConfigurationSource = v
}

// GetProfile returns the Profile field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *TimeseriesAnomalyInvestigationDetection) GetProfile() string {
	if o == nil || o.Profile.Get() == nil {
		var ret string
		return ret
	}
	return *o.Profile.Get()
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TimeseriesAnomalyInvestigationDetection) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Profile.Get(), o.Profile.IsSet()
}

// SetProfile sets field value.
func (o *TimeseriesAnomalyInvestigationDetection) SetProfile(v string) {
	o.Profile.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationDetection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["configuration_source"] = o.ConfigurationSource
	toSerialize["profile"] = o.Profile.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesAnomalyInvestigationDetection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigurationSource *TimeseriesAnomalyInvestigationConfigurationSource `json:"configuration_source"`
		Profile             datadog.NullableString                             `json:"profile"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConfigurationSource == nil {
		return fmt.Errorf("required field configuration_source missing")
	}
	if !all.Profile.IsSet() {
		return fmt.Errorf("required field profile missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"configuration_source", "profile"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.ConfigurationSource.IsValid() {
		hasInvalidField = true
	} else {
		o.ConfigurationSource = *all.ConfigurationSource
	}
	o.Profile = all.Profile

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
