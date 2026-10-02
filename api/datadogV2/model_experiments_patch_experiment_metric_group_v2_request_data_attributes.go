// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes Fields supplied to update the experiment metric group.
type ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes struct {
	// Metrics included in this group.
	Metrics []ExperimentsCreateExperimentMetricGroupV2RequestDataAttributesMetricsItems `json:"metrics,omitempty"`
	// Metadata retained for resources imported from another system.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Display name of the experiment metric group.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPatchExperimentMetricGroupV2RequestDataAttributes instantiates a new ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPatchExperimentMetricGroupV2RequestDataAttributes() *ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes {
	this := ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes{}
	return &this
}

// NewExperimentsPatchExperimentMetricGroupV2RequestDataAttributesWithDefaults instantiates a new ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPatchExperimentMetricGroupV2RequestDataAttributesWithDefaults() *ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes {
	this := ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes{}
	return &this
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes) GetMetrics() []ExperimentsCreateExperimentMetricGroupV2RequestDataAttributesMetricsItems {
	if o == nil || o.Metrics == nil {
		var ret []ExperimentsCreateExperimentMetricGroupV2RequestDataAttributesMetricsItems
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes) GetMetricsOk() (*[]ExperimentsCreateExperimentMetricGroupV2RequestDataAttributesMetricsItems, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes) HasMetrics() bool {
	return o != nil && o.Metrics != nil
}

// SetMetrics gets a reference to the given []ExperimentsCreateExperimentMetricGroupV2RequestDataAttributesMetricsItems and assigns it to the Metrics field.
func (o *ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes) SetMetrics(v []ExperimentsCreateExperimentMetricGroupV2RequestDataAttributesMetricsItems) {
	o.Metrics = v
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	if o.MigrationMetadata != nil {
		toSerialize["migration_metadata"] = o.MigrationMetadata
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPatchExperimentMetricGroupV2RequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Metrics           []ExperimentsCreateExperimentMetricGroupV2RequestDataAttributesMetricsItems `json:"metrics,omitempty"`
		MigrationMetadata interface{}                                                                 `json:"migration_metadata,omitempty"`
		Name              *string                                                                     `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"metrics", "migration_metadata", "name"})
	} else {
		return err
	}
	o.Metrics = all.Metrics
	o.MigrationMetadata = all.MigrationMetadata
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
