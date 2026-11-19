// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionDataQualityPlatformSettings Platform-specific settings for a data quality monitor.
type MonitorFormulaAndFunctionDataQualityPlatformSettings struct {
	// Snowflake-specific settings for a data quality monitor.
	Snowflake *MonitorFormulaAndFunctionDataQualitySnowflakeSettings `json:"snowflake,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorFormulaAndFunctionDataQualityPlatformSettings instantiates a new MonitorFormulaAndFunctionDataQualityPlatformSettings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorFormulaAndFunctionDataQualityPlatformSettings() *MonitorFormulaAndFunctionDataQualityPlatformSettings {
	this := MonitorFormulaAndFunctionDataQualityPlatformSettings{}
	return &this
}

// NewMonitorFormulaAndFunctionDataQualityPlatformSettingsWithDefaults instantiates a new MonitorFormulaAndFunctionDataQualityPlatformSettings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorFormulaAndFunctionDataQualityPlatformSettingsWithDefaults() *MonitorFormulaAndFunctionDataQualityPlatformSettings {
	this := MonitorFormulaAndFunctionDataQualityPlatformSettings{}
	return &this
}

// GetSnowflake returns the Snowflake field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualityPlatformSettings) GetSnowflake() MonitorFormulaAndFunctionDataQualitySnowflakeSettings {
	if o == nil || o.Snowflake == nil {
		var ret MonitorFormulaAndFunctionDataQualitySnowflakeSettings
		return ret
	}
	return *o.Snowflake
}

// GetSnowflakeOk returns a tuple with the Snowflake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityPlatformSettings) GetSnowflakeOk() (*MonitorFormulaAndFunctionDataQualitySnowflakeSettings, bool) {
	if o == nil || o.Snowflake == nil {
		return nil, false
	}
	return o.Snowflake, true
}

// HasSnowflake returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualityPlatformSettings) HasSnowflake() bool {
	return o != nil && o.Snowflake != nil
}

// SetSnowflake gets a reference to the given MonitorFormulaAndFunctionDataQualitySnowflakeSettings and assigns it to the Snowflake field.
func (o *MonitorFormulaAndFunctionDataQualityPlatformSettings) SetSnowflake(v MonitorFormulaAndFunctionDataQualitySnowflakeSettings) {
	o.Snowflake = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorFormulaAndFunctionDataQualityPlatformSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Snowflake != nil {
		toSerialize["snowflake"] = o.Snowflake
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorFormulaAndFunctionDataQualityPlatformSettings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Snowflake *MonitorFormulaAndFunctionDataQualitySnowflakeSettings `json:"snowflake,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"snowflake"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Snowflake != nil && all.Snowflake.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Snowflake = all.Snowflake

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
