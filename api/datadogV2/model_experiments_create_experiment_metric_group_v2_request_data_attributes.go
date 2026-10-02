// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes Name and metric selection for the new experiment metric group.
type ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes struct {
	// Metrics to include in the experiment metric group.
	Metrics []ExperimentsCreateExperimentMetricGroupV2RequestDataAttributesMetricsItems `json:"metrics,omitempty"`
	// Metadata associated with migration of this resource.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Display name of the experiment metric group.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsCreateExperimentMetricGroupV2RequestDataAttributes instantiates a new ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateExperimentMetricGroupV2RequestDataAttributes(name string) *ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes {
	this := ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes{}
	this.Name = name
	return &this
}

// NewExperimentsCreateExperimentMetricGroupV2RequestDataAttributesWithDefaults instantiates a new ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateExperimentMetricGroupV2RequestDataAttributesWithDefaults() *ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes {
	this := ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes{}
	return &this
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes) GetMetrics() []ExperimentsCreateExperimentMetricGroupV2RequestDataAttributesMetricsItems {
	if o == nil || o.Metrics == nil {
		var ret []ExperimentsCreateExperimentMetricGroupV2RequestDataAttributesMetricsItems
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes) GetMetricsOk() (*[]ExperimentsCreateExperimentMetricGroupV2RequestDataAttributesMetricsItems, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes) HasMetrics() bool {
	return o != nil && o.Metrics != nil
}

// SetMetrics gets a reference to the given []ExperimentsCreateExperimentMetricGroupV2RequestDataAttributesMetricsItems and assigns it to the Metrics field.
func (o *ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes) SetMetrics(v []ExperimentsCreateExperimentMetricGroupV2RequestDataAttributesMetricsItems) {
	o.Metrics = v
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value.
func (o *ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateExperimentMetricGroupV2RequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Metrics           []ExperimentsCreateExperimentMetricGroupV2RequestDataAttributesMetricsItems `json:"metrics,omitempty"`
		MigrationMetadata interface{}                                                                 `json:"migration_metadata,omitempty"`
		Name              *string                                                                     `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"metrics", "migration_metadata", "name"})
	} else {
		return err
	}
	o.Metrics = all.Metrics
	o.MigrationMetadata = all.MigrationMetadata
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
