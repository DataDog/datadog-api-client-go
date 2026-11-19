// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionDataQualitySnowflakeSettings Snowflake-specific settings for a data quality monitor.
type MonitorFormulaAndFunctionDataQualitySnowflakeSettings struct {
	// Snowflake warehouse to run this monitor's queries on instead of the integration user's
	// default warehouse. The name is case-sensitive, and the Datadog role must have `USAGE` on
	// the warehouse directly. Ignored for monitors that read catalog metadata instead of
	// running queries (such as default row count and freshness on base tables).
	Warehouse *string `json:"warehouse,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorFormulaAndFunctionDataQualitySnowflakeSettings instantiates a new MonitorFormulaAndFunctionDataQualitySnowflakeSettings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorFormulaAndFunctionDataQualitySnowflakeSettings() *MonitorFormulaAndFunctionDataQualitySnowflakeSettings {
	this := MonitorFormulaAndFunctionDataQualitySnowflakeSettings{}
	return &this
}

// NewMonitorFormulaAndFunctionDataQualitySnowflakeSettingsWithDefaults instantiates a new MonitorFormulaAndFunctionDataQualitySnowflakeSettings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorFormulaAndFunctionDataQualitySnowflakeSettingsWithDefaults() *MonitorFormulaAndFunctionDataQualitySnowflakeSettings {
	this := MonitorFormulaAndFunctionDataQualitySnowflakeSettings{}
	return &this
}

// GetWarehouse returns the Warehouse field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualitySnowflakeSettings) GetWarehouse() string {
	if o == nil || o.Warehouse == nil {
		var ret string
		return ret
	}
	return *o.Warehouse
}

// GetWarehouseOk returns a tuple with the Warehouse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualitySnowflakeSettings) GetWarehouseOk() (*string, bool) {
	if o == nil || o.Warehouse == nil {
		return nil, false
	}
	return o.Warehouse, true
}

// HasWarehouse returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualitySnowflakeSettings) HasWarehouse() bool {
	return o != nil && o.Warehouse != nil
}

// SetWarehouse gets a reference to the given string and assigns it to the Warehouse field.
func (o *MonitorFormulaAndFunctionDataQualitySnowflakeSettings) SetWarehouse(v string) {
	o.Warehouse = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorFormulaAndFunctionDataQualitySnowflakeSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Warehouse != nil {
		toSerialize["warehouse"] = o.Warehouse
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorFormulaAndFunctionDataQualitySnowflakeSettings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Warehouse *string `json:"warehouse,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"warehouse"})
	} else {
		return err
	}
	o.Warehouse = all.Warehouse

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
