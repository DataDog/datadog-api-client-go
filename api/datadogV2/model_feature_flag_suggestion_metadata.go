// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagSuggestionMetadata Metadata associated with a suggestion. Fields other than `variant_id` are set by the server.
type FeatureFlagSuggestionMetadata struct {
	// Whether the server allows this allocation suggestion to reorder persisted drafts.
	AllowDraftReorder *bool `json:"allow_draft_reorder,omitempty"`
	// Server-generated reference to an experiment action requiring approval.
	ExperimentAction *FeatureFlagSuggestionExperimentAction `json:"experiment_action,omitempty"`
	// Server-generated fingerprint of the experiment action request.
	ExperimentActionRequestFingerprint *string `json:"experiment_action_request_fingerprint,omitempty"`
	// The ID of the variant targeted by the suggestion.
	VariantId *uuid.UUID `json:"variant_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFeatureFlagSuggestionMetadata instantiates a new FeatureFlagSuggestionMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFeatureFlagSuggestionMetadata() *FeatureFlagSuggestionMetadata {
	this := FeatureFlagSuggestionMetadata{}
	return &this
}

// NewFeatureFlagSuggestionMetadataWithDefaults instantiates a new FeatureFlagSuggestionMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFeatureFlagSuggestionMetadataWithDefaults() *FeatureFlagSuggestionMetadata {
	this := FeatureFlagSuggestionMetadata{}
	return &this
}

// GetAllowDraftReorder returns the AllowDraftReorder field value if set, zero value otherwise.
func (o *FeatureFlagSuggestionMetadata) GetAllowDraftReorder() bool {
	if o == nil || o.AllowDraftReorder == nil {
		var ret bool
		return ret
	}
	return *o.AllowDraftReorder
}

// GetAllowDraftReorderOk returns a tuple with the AllowDraftReorder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionMetadata) GetAllowDraftReorderOk() (*bool, bool) {
	if o == nil || o.AllowDraftReorder == nil {
		return nil, false
	}
	return o.AllowDraftReorder, true
}

// HasAllowDraftReorder returns a boolean if a field has been set.
func (o *FeatureFlagSuggestionMetadata) HasAllowDraftReorder() bool {
	return o != nil && o.AllowDraftReorder != nil
}

// SetAllowDraftReorder gets a reference to the given bool and assigns it to the AllowDraftReorder field.
func (o *FeatureFlagSuggestionMetadata) SetAllowDraftReorder(v bool) {
	o.AllowDraftReorder = &v
}

// GetExperimentAction returns the ExperimentAction field value if set, zero value otherwise.
func (o *FeatureFlagSuggestionMetadata) GetExperimentAction() FeatureFlagSuggestionExperimentAction {
	if o == nil || o.ExperimentAction == nil {
		var ret FeatureFlagSuggestionExperimentAction
		return ret
	}
	return *o.ExperimentAction
}

// GetExperimentActionOk returns a tuple with the ExperimentAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionMetadata) GetExperimentActionOk() (*FeatureFlagSuggestionExperimentAction, bool) {
	if o == nil || o.ExperimentAction == nil {
		return nil, false
	}
	return o.ExperimentAction, true
}

// HasExperimentAction returns a boolean if a field has been set.
func (o *FeatureFlagSuggestionMetadata) HasExperimentAction() bool {
	return o != nil && o.ExperimentAction != nil
}

// SetExperimentAction gets a reference to the given FeatureFlagSuggestionExperimentAction and assigns it to the ExperimentAction field.
func (o *FeatureFlagSuggestionMetadata) SetExperimentAction(v FeatureFlagSuggestionExperimentAction) {
	o.ExperimentAction = &v
}

// GetExperimentActionRequestFingerprint returns the ExperimentActionRequestFingerprint field value if set, zero value otherwise.
func (o *FeatureFlagSuggestionMetadata) GetExperimentActionRequestFingerprint() string {
	if o == nil || o.ExperimentActionRequestFingerprint == nil {
		var ret string
		return ret
	}
	return *o.ExperimentActionRequestFingerprint
}

// GetExperimentActionRequestFingerprintOk returns a tuple with the ExperimentActionRequestFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionMetadata) GetExperimentActionRequestFingerprintOk() (*string, bool) {
	if o == nil || o.ExperimentActionRequestFingerprint == nil {
		return nil, false
	}
	return o.ExperimentActionRequestFingerprint, true
}

// HasExperimentActionRequestFingerprint returns a boolean if a field has been set.
func (o *FeatureFlagSuggestionMetadata) HasExperimentActionRequestFingerprint() bool {
	return o != nil && o.ExperimentActionRequestFingerprint != nil
}

// SetExperimentActionRequestFingerprint gets a reference to the given string and assigns it to the ExperimentActionRequestFingerprint field.
func (o *FeatureFlagSuggestionMetadata) SetExperimentActionRequestFingerprint(v string) {
	o.ExperimentActionRequestFingerprint = &v
}

// GetVariantId returns the VariantId field value if set, zero value otherwise.
func (o *FeatureFlagSuggestionMetadata) GetVariantId() uuid.UUID {
	if o == nil || o.VariantId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.VariantId
}

// GetVariantIdOk returns a tuple with the VariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionMetadata) GetVariantIdOk() (*uuid.UUID, bool) {
	if o == nil || o.VariantId == nil {
		return nil, false
	}
	return o.VariantId, true
}

// HasVariantId returns a boolean if a field has been set.
func (o *FeatureFlagSuggestionMetadata) HasVariantId() bool {
	return o != nil && o.VariantId != nil
}

// SetVariantId gets a reference to the given uuid.UUID and assigns it to the VariantId field.
func (o *FeatureFlagSuggestionMetadata) SetVariantId(v uuid.UUID) {
	o.VariantId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FeatureFlagSuggestionMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowDraftReorder != nil {
		toSerialize["allow_draft_reorder"] = o.AllowDraftReorder
	}
	if o.ExperimentAction != nil {
		toSerialize["experiment_action"] = o.ExperimentAction
	}
	if o.ExperimentActionRequestFingerprint != nil {
		toSerialize["experiment_action_request_fingerprint"] = o.ExperimentActionRequestFingerprint
	}
	if o.VariantId != nil {
		toSerialize["variant_id"] = o.VariantId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FeatureFlagSuggestionMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowDraftReorder                  *bool                                  `json:"allow_draft_reorder,omitempty"`
		ExperimentAction                   *FeatureFlagSuggestionExperimentAction `json:"experiment_action,omitempty"`
		ExperimentActionRequestFingerprint *string                                `json:"experiment_action_request_fingerprint,omitempty"`
		VariantId                          *uuid.UUID                             `json:"variant_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allow_draft_reorder", "experiment_action", "experiment_action_request_fingerprint", "variant_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AllowDraftReorder = all.AllowDraftReorder
	if all.ExperimentAction != nil && all.ExperimentAction.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ExperimentAction = all.ExperimentAction
	o.ExperimentActionRequestFingerprint = all.ExperimentActionRequestFingerprint
	o.VariantId = all.VariantId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
