// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentMetricGroupV2DTODataAttributes Name, purpose, and selected metrics of an experiment metric group.
type ExperimentsExperimentMetricGroupV2DTODataAttributes struct {
	// Whether this group contains the experiment decision metrics.
	IsDecision *bool `json:"is_decision,omitempty"`
	// Metrics selected for this experiment metric group.
	Metrics []ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems `json:"metrics,omitempty"`
	// Metadata associated with migration of this resource.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Display name of the experiment metric group.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsExperimentMetricGroupV2DTODataAttributes instantiates a new ExperimentsExperimentMetricGroupV2DTODataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsExperimentMetricGroupV2DTODataAttributes() *ExperimentsExperimentMetricGroupV2DTODataAttributes {
	this := ExperimentsExperimentMetricGroupV2DTODataAttributes{}
	return &this
}

// NewExperimentsExperimentMetricGroupV2DTODataAttributesWithDefaults instantiates a new ExperimentsExperimentMetricGroupV2DTODataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsExperimentMetricGroupV2DTODataAttributesWithDefaults() *ExperimentsExperimentMetricGroupV2DTODataAttributes {
	this := ExperimentsExperimentMetricGroupV2DTODataAttributes{}
	return &this
}

// GetIsDecision returns the IsDecision field value if set, zero value otherwise.
func (o *ExperimentsExperimentMetricGroupV2DTODataAttributes) GetIsDecision() bool {
	if o == nil || o.IsDecision == nil {
		var ret bool
		return ret
	}
	return *o.IsDecision
}

// GetIsDecisionOk returns a tuple with the IsDecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentMetricGroupV2DTODataAttributes) GetIsDecisionOk() (*bool, bool) {
	if o == nil || o.IsDecision == nil {
		return nil, false
	}
	return o.IsDecision, true
}

// HasIsDecision returns a boolean if a field has been set.
func (o *ExperimentsExperimentMetricGroupV2DTODataAttributes) HasIsDecision() bool {
	return o != nil && o.IsDecision != nil
}

// SetIsDecision gets a reference to the given bool and assigns it to the IsDecision field.
func (o *ExperimentsExperimentMetricGroupV2DTODataAttributes) SetIsDecision(v bool) {
	o.IsDecision = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *ExperimentsExperimentMetricGroupV2DTODataAttributes) GetMetrics() []ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems {
	if o == nil || o.Metrics == nil {
		var ret []ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentMetricGroupV2DTODataAttributes) GetMetricsOk() (*[]ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *ExperimentsExperimentMetricGroupV2DTODataAttributes) HasMetrics() bool {
	return o != nil && o.Metrics != nil
}

// SetMetrics gets a reference to the given []ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems and assigns it to the Metrics field.
func (o *ExperimentsExperimentMetricGroupV2DTODataAttributes) SetMetrics(v []ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems) {
	o.Metrics = v
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsExperimentMetricGroupV2DTODataAttributes) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentMetricGroupV2DTODataAttributes) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsExperimentMetricGroupV2DTODataAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsExperimentMetricGroupV2DTODataAttributes) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsExperimentMetricGroupV2DTODataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentMetricGroupV2DTODataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsExperimentMetricGroupV2DTODataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsExperimentMetricGroupV2DTODataAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsExperimentMetricGroupV2DTODataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IsDecision != nil {
		toSerialize["is_decision"] = o.IsDecision
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
func (o *ExperimentsExperimentMetricGroupV2DTODataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsDecision        *bool                                                                  `json:"is_decision,omitempty"`
		Metrics           []ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems `json:"metrics,omitempty"`
		MigrationMetadata interface{}                                                            `json:"migration_metadata,omitempty"`
		Name              *string                                                                `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_decision", "metrics", "migration_metadata", "name"})
	} else {
		return err
	}
	o.IsDecision = all.IsDecision
	o.Metrics = all.Metrics
	o.MigrationMetadata = all.MigrationMetadata
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
