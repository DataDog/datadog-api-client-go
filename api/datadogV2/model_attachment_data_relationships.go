// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AttachmentDataRelationships
type AttachmentDataRelationships struct {
	//
	LastModifiedByUser *AttachmentDataRelationshipsLastModifiedByUser `json:"last_modified_by_user,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAttachmentDataRelationships instantiates a new AttachmentDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAttachmentDataRelationships() *AttachmentDataRelationships {
	this := AttachmentDataRelationships{}
	return &this
}

// NewAttachmentDataRelationshipsWithDefaults instantiates a new AttachmentDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAttachmentDataRelationshipsWithDefaults() *AttachmentDataRelationships {
	this := AttachmentDataRelationships{}
	return &this
}

// GetLastModifiedByUser returns the LastModifiedByUser field value if set, zero value otherwise.
func (o *AttachmentDataRelationships) GetLastModifiedByUser() AttachmentDataRelationshipsLastModifiedByUser {
	if o == nil || o.LastModifiedByUser == nil {
		var ret AttachmentDataRelationshipsLastModifiedByUser
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentDataRelationships) GetLastModifiedByUserOk() (*AttachmentDataRelationshipsLastModifiedByUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		return nil, false
	}
	return o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *AttachmentDataRelationships) HasLastModifiedByUser() bool {
	return o != nil && o.LastModifiedByUser != nil
}

// SetLastModifiedByUser gets a reference to the given AttachmentDataRelationshipsLastModifiedByUser and assigns it to the LastModifiedByUser field.
func (o *AttachmentDataRelationships) SetLastModifiedByUser(v AttachmentDataRelationshipsLastModifiedByUser) {
	o.LastModifiedByUser = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AttachmentDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.LastModifiedByUser != nil {
		toSerialize["last_modified_by_user"] = o.LastModifiedByUser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AttachmentDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LastModifiedByUser *AttachmentDataRelationshipsLastModifiedByUser `json:"last_modified_by_user,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"last_modified_by_user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.LastModifiedByUser != nil && all.LastModifiedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LastModifiedByUser = all.LastModifiedByUser

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
