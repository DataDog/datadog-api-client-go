// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateSnapshotAdditionalConfig Additional configuration options for snapshot creation.
type CreateSnapshotAdditionalConfig struct {
	// List of template variable definitions for snapshot rendering.
	TemplateVariables []CreateSnapshotTemplateVariable `json:"template_variables,omitempty"`
	// The legend display type for timeseries widgets. A value of `none` hides the legend entirely; omitting the field lets the frontend choose automatically.
	TimeseriesLegendType *CreateSnapshotTimeseriesLegendType `json:"timeseries_legend_type,omitempty"`
	// Timezone offset in minutes from UTC. Positive values are west of UTC (for example, `300` for UTC-5). Use `0` for UTC.
	TimezoneOffsetMinutes *int64 `json:"timezone_offset_minutes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateSnapshotAdditionalConfig instantiates a new CreateSnapshotAdditionalConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateSnapshotAdditionalConfig() *CreateSnapshotAdditionalConfig {
	this := CreateSnapshotAdditionalConfig{}
	return &this
}

// NewCreateSnapshotAdditionalConfigWithDefaults instantiates a new CreateSnapshotAdditionalConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateSnapshotAdditionalConfigWithDefaults() *CreateSnapshotAdditionalConfig {
	this := CreateSnapshotAdditionalConfig{}
	return &this
}

// GetTemplateVariables returns the TemplateVariables field value if set, zero value otherwise.
func (o *CreateSnapshotAdditionalConfig) GetTemplateVariables() []CreateSnapshotTemplateVariable {
	if o == nil || o.TemplateVariables == nil {
		var ret []CreateSnapshotTemplateVariable
		return ret
	}
	return o.TemplateVariables
}

// GetTemplateVariablesOk returns a tuple with the TemplateVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotAdditionalConfig) GetTemplateVariablesOk() (*[]CreateSnapshotTemplateVariable, bool) {
	if o == nil || o.TemplateVariables == nil {
		return nil, false
	}
	return &o.TemplateVariables, true
}

// HasTemplateVariables returns a boolean if a field has been set.
func (o *CreateSnapshotAdditionalConfig) HasTemplateVariables() bool {
	return o != nil && o.TemplateVariables != nil
}

// SetTemplateVariables gets a reference to the given []CreateSnapshotTemplateVariable and assigns it to the TemplateVariables field.
func (o *CreateSnapshotAdditionalConfig) SetTemplateVariables(v []CreateSnapshotTemplateVariable) {
	o.TemplateVariables = v
}

// GetTimeseriesLegendType returns the TimeseriesLegendType field value if set, zero value otherwise.
func (o *CreateSnapshotAdditionalConfig) GetTimeseriesLegendType() CreateSnapshotTimeseriesLegendType {
	if o == nil || o.TimeseriesLegendType == nil {
		var ret CreateSnapshotTimeseriesLegendType
		return ret
	}
	return *o.TimeseriesLegendType
}

// GetTimeseriesLegendTypeOk returns a tuple with the TimeseriesLegendType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotAdditionalConfig) GetTimeseriesLegendTypeOk() (*CreateSnapshotTimeseriesLegendType, bool) {
	if o == nil || o.TimeseriesLegendType == nil {
		return nil, false
	}
	return o.TimeseriesLegendType, true
}

// HasTimeseriesLegendType returns a boolean if a field has been set.
func (o *CreateSnapshotAdditionalConfig) HasTimeseriesLegendType() bool {
	return o != nil && o.TimeseriesLegendType != nil
}

// SetTimeseriesLegendType gets a reference to the given CreateSnapshotTimeseriesLegendType and assigns it to the TimeseriesLegendType field.
func (o *CreateSnapshotAdditionalConfig) SetTimeseriesLegendType(v CreateSnapshotTimeseriesLegendType) {
	o.TimeseriesLegendType = &v
}

// GetTimezoneOffsetMinutes returns the TimezoneOffsetMinutes field value if set, zero value otherwise.
func (o *CreateSnapshotAdditionalConfig) GetTimezoneOffsetMinutes() int64 {
	if o == nil || o.TimezoneOffsetMinutes == nil {
		var ret int64
		return ret
	}
	return *o.TimezoneOffsetMinutes
}

// GetTimezoneOffsetMinutesOk returns a tuple with the TimezoneOffsetMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotAdditionalConfig) GetTimezoneOffsetMinutesOk() (*int64, bool) {
	if o == nil || o.TimezoneOffsetMinutes == nil {
		return nil, false
	}
	return o.TimezoneOffsetMinutes, true
}

// HasTimezoneOffsetMinutes returns a boolean if a field has been set.
func (o *CreateSnapshotAdditionalConfig) HasTimezoneOffsetMinutes() bool {
	return o != nil && o.TimezoneOffsetMinutes != nil
}

// SetTimezoneOffsetMinutes gets a reference to the given int64 and assigns it to the TimezoneOffsetMinutes field.
func (o *CreateSnapshotAdditionalConfig) SetTimezoneOffsetMinutes(v int64) {
	o.TimezoneOffsetMinutes = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateSnapshotAdditionalConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TemplateVariables != nil {
		toSerialize["template_variables"] = o.TemplateVariables
	}
	if o.TimeseriesLegendType != nil {
		toSerialize["timeseries_legend_type"] = o.TimeseriesLegendType
	}
	if o.TimezoneOffsetMinutes != nil {
		toSerialize["timezone_offset_minutes"] = o.TimezoneOffsetMinutes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateSnapshotAdditionalConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TemplateVariables     []CreateSnapshotTemplateVariable    `json:"template_variables,omitempty"`
		TimeseriesLegendType  *CreateSnapshotTimeseriesLegendType `json:"timeseries_legend_type,omitempty"`
		TimezoneOffsetMinutes *int64                              `json:"timezone_offset_minutes,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"template_variables", "timeseries_legend_type", "timezone_offset_minutes"})
	} else {
		return err
	}

	hasInvalidField := false
	o.TemplateVariables = all.TemplateVariables
	if all.TimeseriesLegendType != nil && !all.TimeseriesLegendType.IsValid() {
		hasInvalidField = true
	} else {
		o.TimeseriesLegendType = all.TimeseriesLegendType
	}
	o.TimezoneOffsetMinutes = all.TimezoneOffsetMinutes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
