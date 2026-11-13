// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagApprovalRequiredAttributes The change to submit for approval.
type FeatureFlagApprovalRequiredAttributes struct {
	// The proposed action.
	Action string `json:"action"`
	// The ID of the environment targeted by the change.
	EnvironmentId *uuid.UUID `json:"environment_id,omitempty"`
	// The feature flag property to change.
	Property string `json:"property"`
	// The proposed value, serialized as a string. Allocation sets are JSON-encoded strings.
	Suggestion string `json:"suggestion"`
	// Metadata associated with a suggestion. Fields other than `variant_id` are set by the server.
	SuggestionMetadata *FeatureFlagSuggestionMetadata `json:"suggestion_metadata,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFeatureFlagApprovalRequiredAttributes instantiates a new FeatureFlagApprovalRequiredAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFeatureFlagApprovalRequiredAttributes(action string, property string, suggestion string) *FeatureFlagApprovalRequiredAttributes {
	this := FeatureFlagApprovalRequiredAttributes{}
	this.Action = action
	this.Property = property
	this.Suggestion = suggestion
	return &this
}

// NewFeatureFlagApprovalRequiredAttributesWithDefaults instantiates a new FeatureFlagApprovalRequiredAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFeatureFlagApprovalRequiredAttributesWithDefaults() *FeatureFlagApprovalRequiredAttributes {
	this := FeatureFlagApprovalRequiredAttributes{}
	return &this
}

// GetAction returns the Action field value.
func (o *FeatureFlagApprovalRequiredAttributes) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagApprovalRequiredAttributes) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *FeatureFlagApprovalRequiredAttributes) SetAction(v string) {
	o.Action = v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *FeatureFlagApprovalRequiredAttributes) GetEnvironmentId() uuid.UUID {
	if o == nil || o.EnvironmentId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagApprovalRequiredAttributes) GetEnvironmentIdOk() (*uuid.UUID, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *FeatureFlagApprovalRequiredAttributes) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given uuid.UUID and assigns it to the EnvironmentId field.
func (o *FeatureFlagApprovalRequiredAttributes) SetEnvironmentId(v uuid.UUID) {
	o.EnvironmentId = &v
}

// GetProperty returns the Property field value.
func (o *FeatureFlagApprovalRequiredAttributes) GetProperty() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagApprovalRequiredAttributes) GetPropertyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value.
func (o *FeatureFlagApprovalRequiredAttributes) SetProperty(v string) {
	o.Property = v
}

// GetSuggestion returns the Suggestion field value.
func (o *FeatureFlagApprovalRequiredAttributes) GetSuggestion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Suggestion
}

// GetSuggestionOk returns a tuple with the Suggestion field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagApprovalRequiredAttributes) GetSuggestionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Suggestion, true
}

// SetSuggestion sets field value.
func (o *FeatureFlagApprovalRequiredAttributes) SetSuggestion(v string) {
	o.Suggestion = v
}

// GetSuggestionMetadata returns the SuggestionMetadata field value if set, zero value otherwise.
func (o *FeatureFlagApprovalRequiredAttributes) GetSuggestionMetadata() FeatureFlagSuggestionMetadata {
	if o == nil || o.SuggestionMetadata == nil {
		var ret FeatureFlagSuggestionMetadata
		return ret
	}
	return *o.SuggestionMetadata
}

// GetSuggestionMetadataOk returns a tuple with the SuggestionMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagApprovalRequiredAttributes) GetSuggestionMetadataOk() (*FeatureFlagSuggestionMetadata, bool) {
	if o == nil || o.SuggestionMetadata == nil {
		return nil, false
	}
	return o.SuggestionMetadata, true
}

// HasSuggestionMetadata returns a boolean if a field has been set.
func (o *FeatureFlagApprovalRequiredAttributes) HasSuggestionMetadata() bool {
	return o != nil && o.SuggestionMetadata != nil
}

// SetSuggestionMetadata gets a reference to the given FeatureFlagSuggestionMetadata and assigns it to the SuggestionMetadata field.
func (o *FeatureFlagApprovalRequiredAttributes) SetSuggestionMetadata(v FeatureFlagSuggestionMetadata) {
	o.SuggestionMetadata = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FeatureFlagApprovalRequiredAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	if o.EnvironmentId != nil {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	toSerialize["property"] = o.Property
	toSerialize["suggestion"] = o.Suggestion
	if o.SuggestionMetadata != nil {
		toSerialize["suggestion_metadata"] = o.SuggestionMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FeatureFlagApprovalRequiredAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action             *string                        `json:"action"`
		EnvironmentId      *uuid.UUID                     `json:"environment_id,omitempty"`
		Property           *string                        `json:"property"`
		Suggestion         *string                        `json:"suggestion"`
		SuggestionMetadata *FeatureFlagSuggestionMetadata `json:"suggestion_metadata,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.Property == nil {
		return fmt.Errorf("required field property missing")
	}
	if all.Suggestion == nil {
		return fmt.Errorf("required field suggestion missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "environment_id", "property", "suggestion", "suggestion_metadata"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Action = *all.Action
	o.EnvironmentId = all.EnvironmentId
	o.Property = *all.Property
	o.Suggestion = *all.Suggestion
	if all.SuggestionMetadata != nil && all.SuggestionMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SuggestionMetadata = all.SuggestionMetadata

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
