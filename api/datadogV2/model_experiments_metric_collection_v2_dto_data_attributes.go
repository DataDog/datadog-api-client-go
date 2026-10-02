// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsMetricCollectionV2DTODataAttributes Details of the metric collection.
type ExperimentsMetricCollectionV2DTODataAttributes struct {
	// Time when this resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Text that explains the metric collection.
	Description datadog.NullableString `json:"description,omitempty"`
	// Whether the collection is used for guardrail metrics.
	IsGuardrail *bool `json:"is_guardrail,omitempty"`
	// Number of metrics in this collection.
	MetricCount *int64 `json:"metric_count,omitempty"`
	// Metrics included in this collection.
	Metrics []ExperimentsMetricCollectionV2DTODataAttributesMetricsItems `json:"metrics,omitempty"`
	// Display name of the metric collection.
	Name *string `json:"name,omitempty"`
	// Time when this resource was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsMetricCollectionV2DTODataAttributes instantiates a new ExperimentsMetricCollectionV2DTODataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsMetricCollectionV2DTODataAttributes() *ExperimentsMetricCollectionV2DTODataAttributes {
	this := ExperimentsMetricCollectionV2DTODataAttributes{}
	return &this
}

// NewExperimentsMetricCollectionV2DTODataAttributesWithDefaults instantiates a new ExperimentsMetricCollectionV2DTODataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsMetricCollectionV2DTODataAttributesWithDefaults() *ExperimentsMetricCollectionV2DTODataAttributes {
	this := ExperimentsMetricCollectionV2DTODataAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsMetricCollectionV2DTODataAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetIsGuardrail returns the IsGuardrail field value if set, zero value otherwise.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) GetIsGuardrail() bool {
	if o == nil || o.IsGuardrail == nil {
		var ret bool
		return ret
	}
	return *o.IsGuardrail
}

// GetIsGuardrailOk returns a tuple with the IsGuardrail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) GetIsGuardrailOk() (*bool, bool) {
	if o == nil || o.IsGuardrail == nil {
		return nil, false
	}
	return o.IsGuardrail, true
}

// HasIsGuardrail returns a boolean if a field has been set.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) HasIsGuardrail() bool {
	return o != nil && o.IsGuardrail != nil
}

// SetIsGuardrail gets a reference to the given bool and assigns it to the IsGuardrail field.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) SetIsGuardrail(v bool) {
	o.IsGuardrail = &v
}

// GetMetricCount returns the MetricCount field value if set, zero value otherwise.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) GetMetricCount() int64 {
	if o == nil || o.MetricCount == nil {
		var ret int64
		return ret
	}
	return *o.MetricCount
}

// GetMetricCountOk returns a tuple with the MetricCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) GetMetricCountOk() (*int64, bool) {
	if o == nil || o.MetricCount == nil {
		return nil, false
	}
	return o.MetricCount, true
}

// HasMetricCount returns a boolean if a field has been set.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) HasMetricCount() bool {
	return o != nil && o.MetricCount != nil
}

// SetMetricCount gets a reference to the given int64 and assigns it to the MetricCount field.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) SetMetricCount(v int64) {
	o.MetricCount = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) GetMetrics() []ExperimentsMetricCollectionV2DTODataAttributesMetricsItems {
	if o == nil || o.Metrics == nil {
		var ret []ExperimentsMetricCollectionV2DTODataAttributesMetricsItems
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) GetMetricsOk() (*[]ExperimentsMetricCollectionV2DTODataAttributesMetricsItems, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) HasMetrics() bool {
	return o != nil && o.Metrics != nil
}

// SetMetrics gets a reference to the given []ExperimentsMetricCollectionV2DTODataAttributesMetricsItems and assigns it to the Metrics field.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) SetMetrics(v []ExperimentsMetricCollectionV2DTODataAttributesMetricsItems) {
	o.Metrics = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) SetName(v string) {
	o.Name = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsMetricCollectionV2DTODataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.IsGuardrail != nil {
		toSerialize["is_guardrail"] = o.IsGuardrail
	}
	if o.MetricCount != nil {
		toSerialize["metric_count"] = o.MetricCount
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsMetricCollectionV2DTODataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt   *time.Time                                                   `json:"created_at,omitempty"`
		Description datadog.NullableString                                       `json:"description,omitempty"`
		IsGuardrail *bool                                                        `json:"is_guardrail,omitempty"`
		MetricCount *int64                                                       `json:"metric_count,omitempty"`
		Metrics     []ExperimentsMetricCollectionV2DTODataAttributesMetricsItems `json:"metrics,omitempty"`
		Name        *string                                                      `json:"name,omitempty"`
		UpdatedAt   *time.Time                                                   `json:"updated_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "description", "is_guardrail", "metric_count", "metrics", "name", "updated_at"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.Description = all.Description
	o.IsGuardrail = all.IsGuardrail
	o.MetricCount = all.MetricCount
	o.Metrics = all.Metrics
	o.Name = all.Name
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
