// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPublicProtocolListResponseDataAttributes Summary of the protocol and its selected subject type and primary metric.
type ExperimentsPublicProtocolListResponseDataAttributes struct {
	// Text that explains the protocol.
	Description *string `json:"description,omitempty"`
	// Display name of the protocol.
	Name string `json:"name"`
	// Subject type selected by the protocol.
	PrimaryMetric *ExperimentsPublicProtocolResponseDataAttributesSubjectType `json:"primary_metric,omitempty"`
	// ID of the primary metric supplied by the protocol.
	PrimaryMetricId *string `json:"primary_metric_id,omitempty"`
	// Publication status of the protocol.
	Status ExperimentsPublicProtocolResponseDataAttributesStatus `json:"status"`
	// Subject type selected by the protocol.
	SubjectType *ExperimentsPublicProtocolResponseDataAttributesSubjectType `json:"subject_type,omitempty"`
	// ID of the subject type used by this configuration.
	SubjectTypeId *string `json:"subject_type_id,omitempty"`
	// RFC3339 update time. Preserve all fractional seconds when passing this value as expected_updated_at.
	UpdatedAt string `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPublicProtocolListResponseDataAttributes instantiates a new ExperimentsPublicProtocolListResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPublicProtocolListResponseDataAttributes(name string, status ExperimentsPublicProtocolResponseDataAttributesStatus, updatedAt string) *ExperimentsPublicProtocolListResponseDataAttributes {
	this := ExperimentsPublicProtocolListResponseDataAttributes{}
	this.Name = name
	this.Status = status
	this.UpdatedAt = updatedAt
	return &this
}

// NewExperimentsPublicProtocolListResponseDataAttributesWithDefaults instantiates a new ExperimentsPublicProtocolListResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPublicProtocolListResponseDataAttributesWithDefaults() *ExperimentsPublicProtocolListResponseDataAttributes {
	this := ExperimentsPublicProtocolListResponseDataAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) SetName(v string) {
	o.Name = v
}

// GetPrimaryMetric returns the PrimaryMetric field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) GetPrimaryMetric() ExperimentsPublicProtocolResponseDataAttributesSubjectType {
	if o == nil || o.PrimaryMetric == nil {
		var ret ExperimentsPublicProtocolResponseDataAttributesSubjectType
		return ret
	}
	return *o.PrimaryMetric
}

// GetPrimaryMetricOk returns a tuple with the PrimaryMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) GetPrimaryMetricOk() (*ExperimentsPublicProtocolResponseDataAttributesSubjectType, bool) {
	if o == nil || o.PrimaryMetric == nil {
		return nil, false
	}
	return o.PrimaryMetric, true
}

// HasPrimaryMetric returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) HasPrimaryMetric() bool {
	return o != nil && o.PrimaryMetric != nil
}

// SetPrimaryMetric gets a reference to the given ExperimentsPublicProtocolResponseDataAttributesSubjectType and assigns it to the PrimaryMetric field.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) SetPrimaryMetric(v ExperimentsPublicProtocolResponseDataAttributesSubjectType) {
	o.PrimaryMetric = &v
}

// GetPrimaryMetricId returns the PrimaryMetricId field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) GetPrimaryMetricId() string {
	if o == nil || o.PrimaryMetricId == nil {
		var ret string
		return ret
	}
	return *o.PrimaryMetricId
}

// GetPrimaryMetricIdOk returns a tuple with the PrimaryMetricId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) GetPrimaryMetricIdOk() (*string, bool) {
	if o == nil || o.PrimaryMetricId == nil {
		return nil, false
	}
	return o.PrimaryMetricId, true
}

// HasPrimaryMetricId returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) HasPrimaryMetricId() bool {
	return o != nil && o.PrimaryMetricId != nil
}

// SetPrimaryMetricId gets a reference to the given string and assigns it to the PrimaryMetricId field.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) SetPrimaryMetricId(v string) {
	o.PrimaryMetricId = &v
}

// GetStatus returns the Status field value.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) GetStatus() ExperimentsPublicProtocolResponseDataAttributesStatus {
	if o == nil {
		var ret ExperimentsPublicProtocolResponseDataAttributesStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) GetStatusOk() (*ExperimentsPublicProtocolResponseDataAttributesStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) SetStatus(v ExperimentsPublicProtocolResponseDataAttributesStatus) {
	o.Status = v
}

// GetSubjectType returns the SubjectType field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) GetSubjectType() ExperimentsPublicProtocolResponseDataAttributesSubjectType {
	if o == nil || o.SubjectType == nil {
		var ret ExperimentsPublicProtocolResponseDataAttributesSubjectType
		return ret
	}
	return *o.SubjectType
}

// GetSubjectTypeOk returns a tuple with the SubjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) GetSubjectTypeOk() (*ExperimentsPublicProtocolResponseDataAttributesSubjectType, bool) {
	if o == nil || o.SubjectType == nil {
		return nil, false
	}
	return o.SubjectType, true
}

// HasSubjectType returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) HasSubjectType() bool {
	return o != nil && o.SubjectType != nil
}

// SetSubjectType gets a reference to the given ExperimentsPublicProtocolResponseDataAttributesSubjectType and assigns it to the SubjectType field.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) SetSubjectType(v ExperimentsPublicProtocolResponseDataAttributesSubjectType) {
	o.SubjectType = &v
}

// GetSubjectTypeId returns the SubjectTypeId field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) GetSubjectTypeId() string {
	if o == nil || o.SubjectTypeId == nil {
		var ret string
		return ret
	}
	return *o.SubjectTypeId
}

// GetSubjectTypeIdOk returns a tuple with the SubjectTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) GetSubjectTypeIdOk() (*string, bool) {
	if o == nil || o.SubjectTypeId == nil {
		return nil, false
	}
	return o.SubjectTypeId, true
}

// HasSubjectTypeId returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) HasSubjectTypeId() bool {
	return o != nil && o.SubjectTypeId != nil
}

// SetSubjectTypeId gets a reference to the given string and assigns it to the SubjectTypeId field.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) SetSubjectTypeId(v string) {
	o.SubjectTypeId = &v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPublicProtocolListResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name
	if o.PrimaryMetric != nil {
		toSerialize["primary_metric"] = o.PrimaryMetric
	}
	if o.PrimaryMetricId != nil {
		toSerialize["primary_metric_id"] = o.PrimaryMetricId
	}
	toSerialize["status"] = o.Status
	if o.SubjectType != nil {
		toSerialize["subject_type"] = o.SubjectType
	}
	if o.SubjectTypeId != nil {
		toSerialize["subject_type_id"] = o.SubjectTypeId
	}
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPublicProtocolListResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description     *string                                                     `json:"description,omitempty"`
		Name            *string                                                     `json:"name"`
		PrimaryMetric   *ExperimentsPublicProtocolResponseDataAttributesSubjectType `json:"primary_metric,omitempty"`
		PrimaryMetricId *string                                                     `json:"primary_metric_id,omitempty"`
		Status          *ExperimentsPublicProtocolResponseDataAttributesStatus      `json:"status"`
		SubjectType     *ExperimentsPublicProtocolResponseDataAttributesSubjectType `json:"subject_type,omitempty"`
		SubjectTypeId   *string                                                     `json:"subject_type_id,omitempty"`
		UpdatedAt       *string                                                     `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "name", "primary_metric", "primary_metric_id", "status", "subject_type", "subject_type_id", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.Name = *all.Name
	if all.PrimaryMetric != nil && all.PrimaryMetric.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PrimaryMetric = all.PrimaryMetric
	o.PrimaryMetricId = all.PrimaryMetricId
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	if all.SubjectType != nil && all.SubjectType.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SubjectType = all.SubjectType
	o.SubjectTypeId = all.SubjectTypeId
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
