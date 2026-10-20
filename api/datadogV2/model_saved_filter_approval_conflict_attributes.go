// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SavedFilterApprovalConflictAttributes One or more feature flags that reference this saved filter require approval for the rule update.
type SavedFilterApprovalConflictAttributes struct {
	// Feature flags preventing this update.
	ApprovalRequiredFlags []SavedFilterFlagReference `json:"approval_required_flags"`
	// The reason the saved filter update was rejected.
	Message string `json:"message"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSavedFilterApprovalConflictAttributes instantiates a new SavedFilterApprovalConflictAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSavedFilterApprovalConflictAttributes(approvalRequiredFlags []SavedFilterFlagReference, message string) *SavedFilterApprovalConflictAttributes {
	this := SavedFilterApprovalConflictAttributes{}
	this.ApprovalRequiredFlags = approvalRequiredFlags
	this.Message = message
	return &this
}

// NewSavedFilterApprovalConflictAttributesWithDefaults instantiates a new SavedFilterApprovalConflictAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSavedFilterApprovalConflictAttributesWithDefaults() *SavedFilterApprovalConflictAttributes {
	this := SavedFilterApprovalConflictAttributes{}
	return &this
}

// GetApprovalRequiredFlags returns the ApprovalRequiredFlags field value.
func (o *SavedFilterApprovalConflictAttributes) GetApprovalRequiredFlags() []SavedFilterFlagReference {
	if o == nil {
		var ret []SavedFilterFlagReference
		return ret
	}
	return o.ApprovalRequiredFlags
}

// GetApprovalRequiredFlagsOk returns a tuple with the ApprovalRequiredFlags field value
// and a boolean to check if the value has been set.
func (o *SavedFilterApprovalConflictAttributes) GetApprovalRequiredFlagsOk() (*[]SavedFilterFlagReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovalRequiredFlags, true
}

// SetApprovalRequiredFlags sets field value.
func (o *SavedFilterApprovalConflictAttributes) SetApprovalRequiredFlags(v []SavedFilterFlagReference) {
	o.ApprovalRequiredFlags = v
}

// GetMessage returns the Message field value.
func (o *SavedFilterApprovalConflictAttributes) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SavedFilterApprovalConflictAttributes) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *SavedFilterApprovalConflictAttributes) SetMessage(v string) {
	o.Message = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SavedFilterApprovalConflictAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["approval_required_flags"] = o.ApprovalRequiredFlags
	toSerialize["message"] = o.Message

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SavedFilterApprovalConflictAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApprovalRequiredFlags *[]SavedFilterFlagReference `json:"approval_required_flags"`
		Message               *string                     `json:"message"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApprovalRequiredFlags == nil {
		return fmt.Errorf("required field approval_required_flags missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"approval_required_flags", "message"})
	} else {
		return err
	}
	o.ApprovalRequiredFlags = *all.ApprovalRequiredFlags
	o.Message = *all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
