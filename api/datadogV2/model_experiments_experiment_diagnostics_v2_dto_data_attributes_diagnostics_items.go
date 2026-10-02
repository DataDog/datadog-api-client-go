// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems Result of one diagnostic check for an experiment.
type ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems struct {
	// Explanation of the diagnostic check result.
	Message datadog.NullableString `json:"message,omitempty"`
	// Identifier of the metric associated with this check.
	MetricId datadog.NullableString `json:"metric_id,omitempty"`
	// Reason the diagnostic check could not be evaluated.
	SkippedReason NullableExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason `json:"skipped_reason,omitempty"`
	// Outcome of an individual diagnostic check.
	Status ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus `json:"status"`
	// Short title of the diagnostic check.
	Title string `json:"title"`
	// Kind of diagnostic check performed.
	Type ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems instantiates a new ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems(status ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus, title string, typeVar ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsType) *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems {
	this := ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems{}
	this.Status = status
	this.Title = title
	this.Type = typeVar
	return &this
}

// NewExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsWithDefaults instantiates a new ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsWithDefaults() *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems {
	this := ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) HasMessage() bool {
	return o != nil && o.Message.IsSet()
}

// SetMessage gets a reference to the given datadog.NullableString and assigns it to the Message field.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) SetMessage(v string) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) UnsetMessage() {
	o.Message.Unset()
}

// GetMetricId returns the MetricId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) GetMetricId() string {
	if o == nil || o.MetricId.Get() == nil {
		var ret string
		return ret
	}
	return *o.MetricId.Get()
}

// GetMetricIdOk returns a tuple with the MetricId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) GetMetricIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricId.Get(), o.MetricId.IsSet()
}

// HasMetricId returns a boolean if a field has been set.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) HasMetricId() bool {
	return o != nil && o.MetricId.IsSet()
}

// SetMetricId gets a reference to the given datadog.NullableString and assigns it to the MetricId field.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) SetMetricId(v string) {
	o.MetricId.Set(&v)
}

// SetMetricIdNil sets the value for MetricId to be an explicit nil.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) SetMetricIdNil() {
	o.MetricId.Set(nil)
}

// UnsetMetricId ensures that no value is present for MetricId, not even an explicit nil.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) UnsetMetricId() {
	o.MetricId.Unset()
}

// GetSkippedReason returns the SkippedReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) GetSkippedReason() ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason {
	if o == nil || o.SkippedReason.Get() == nil {
		var ret ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason
		return ret
	}
	return *o.SkippedReason.Get()
}

// GetSkippedReasonOk returns a tuple with the SkippedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) GetSkippedReasonOk() (*ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkippedReason.Get(), o.SkippedReason.IsSet()
}

// HasSkippedReason returns a boolean if a field has been set.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) HasSkippedReason() bool {
	return o != nil && o.SkippedReason.IsSet()
}

// SetSkippedReason gets a reference to the given NullableExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason and assigns it to the SkippedReason field.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) SetSkippedReason(v ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason) {
	o.SkippedReason.Set(&v)
}

// SetSkippedReasonNil sets the value for SkippedReason to be an explicit nil.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) SetSkippedReasonNil() {
	o.SkippedReason.Set(nil)
}

// UnsetSkippedReason ensures that no value is present for SkippedReason, not even an explicit nil.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) UnsetSkippedReason() {
	o.SkippedReason.Unset()
}

// GetStatus returns the Status field value.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) GetStatus() ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus {
	if o == nil {
		var ret ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) GetStatusOk() (*ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) SetStatus(v ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus) {
	o.Status = v
}

// GetTitle returns the Title field value.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) GetType() ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsType {
	if o == nil {
		var ret ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) GetTypeOk() (*ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) SetType(v ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.MetricId.IsSet() {
		toSerialize["metric_id"] = o.MetricId.Get()
	}
	if o.SkippedReason.IsSet() {
		toSerialize["skipped_reason"] = o.SkippedReason.Get()
	}
	toSerialize["status"] = o.Status
	toSerialize["title"] = o.Title
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message       datadog.NullableString                                                                   `json:"message,omitempty"`
		MetricId      datadog.NullableString                                                                   `json:"metric_id,omitempty"`
		SkippedReason NullableExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason `json:"skipped_reason,omitempty"`
		Status        *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsStatus               `json:"status"`
		Title         *string                                                                                  `json:"title"`
		Type          *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsType                 `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"message", "metric_id", "skipped_reason", "status", "title", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Message = all.Message
	o.MetricId = all.MetricId
	if all.SkippedReason.Get() != nil && !all.SkippedReason.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.SkippedReason = all.SkippedReason
	}
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Title = *all.Title
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
