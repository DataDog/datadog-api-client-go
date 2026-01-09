// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RevertCustomRuleRevisionRequestDataAttributes
type RevertCustomRuleRevisionRequestDataAttributes struct {
	// Current revision ID
	CurrentRevisionId *string `json:"current_revision_id,omitempty"`
	// Target revision ID to revert to
	RevertToRevisionId *string `json:"revert_to_revision_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRevertCustomRuleRevisionRequestDataAttributes instantiates a new RevertCustomRuleRevisionRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRevertCustomRuleRevisionRequestDataAttributes() *RevertCustomRuleRevisionRequestDataAttributes {
	this := RevertCustomRuleRevisionRequestDataAttributes{}
	return &this
}

// NewRevertCustomRuleRevisionRequestDataAttributesWithDefaults instantiates a new RevertCustomRuleRevisionRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRevertCustomRuleRevisionRequestDataAttributesWithDefaults() *RevertCustomRuleRevisionRequestDataAttributes {
	this := RevertCustomRuleRevisionRequestDataAttributes{}
	return &this
}

// GetCurrentRevisionId returns the CurrentRevisionId field value if set, zero value otherwise.
func (o *RevertCustomRuleRevisionRequestDataAttributes) GetCurrentRevisionId() string {
	if o == nil || o.CurrentRevisionId == nil {
		var ret string
		return ret
	}
	return *o.CurrentRevisionId
}

// GetCurrentRevisionIdOk returns a tuple with the CurrentRevisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevertCustomRuleRevisionRequestDataAttributes) GetCurrentRevisionIdOk() (*string, bool) {
	if o == nil || o.CurrentRevisionId == nil {
		return nil, false
	}
	return o.CurrentRevisionId, true
}

// HasCurrentRevisionId returns a boolean if a field has been set.
func (o *RevertCustomRuleRevisionRequestDataAttributes) HasCurrentRevisionId() bool {
	return o != nil && o.CurrentRevisionId != nil
}

// SetCurrentRevisionId gets a reference to the given string and assigns it to the CurrentRevisionId field.
func (o *RevertCustomRuleRevisionRequestDataAttributes) SetCurrentRevisionId(v string) {
	o.CurrentRevisionId = &v
}

// GetRevertToRevisionId returns the RevertToRevisionId field value if set, zero value otherwise.
func (o *RevertCustomRuleRevisionRequestDataAttributes) GetRevertToRevisionId() string {
	if o == nil || o.RevertToRevisionId == nil {
		var ret string
		return ret
	}
	return *o.RevertToRevisionId
}

// GetRevertToRevisionIdOk returns a tuple with the RevertToRevisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevertCustomRuleRevisionRequestDataAttributes) GetRevertToRevisionIdOk() (*string, bool) {
	if o == nil || o.RevertToRevisionId == nil {
		return nil, false
	}
	return o.RevertToRevisionId, true
}

// HasRevertToRevisionId returns a boolean if a field has been set.
func (o *RevertCustomRuleRevisionRequestDataAttributes) HasRevertToRevisionId() bool {
	return o != nil && o.RevertToRevisionId != nil
}

// SetRevertToRevisionId gets a reference to the given string and assigns it to the RevertToRevisionId field.
func (o *RevertCustomRuleRevisionRequestDataAttributes) SetRevertToRevisionId(v string) {
	o.RevertToRevisionId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RevertCustomRuleRevisionRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CurrentRevisionId != nil {
		toSerialize["current_revision_id"] = o.CurrentRevisionId
	}
	if o.RevertToRevisionId != nil {
		toSerialize["revert_to_revision_id"] = o.RevertToRevisionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RevertCustomRuleRevisionRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CurrentRevisionId  *string `json:"current_revision_id,omitempty"`
		RevertToRevisionId *string `json:"revert_to_revision_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"current_revision_id", "revert_to_revision_id"})
	} else {
		return err
	}
	o.CurrentRevisionId = all.CurrentRevisionId
	o.RevertToRevisionId = all.RevertToRevisionId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
