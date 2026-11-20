// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WildcardBackendDefinitionSpecification Test-only generated schema for the dashboard versioning proof of concept.
type WildcardBackendDefinitionSpecification struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Contents map[string]interface{} `json:"contents"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type WildcardSpecificationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewWildcardBackendDefinitionSpecification instantiates a new WildcardBackendDefinitionSpecification object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWildcardBackendDefinitionSpecification(contents map[string]interface{}, typeVar WildcardSpecificationType) *WildcardBackendDefinitionSpecification {
	this := WildcardBackendDefinitionSpecification{}
	this.Contents = contents
	this.Type = typeVar
	return &this
}

// NewWildcardBackendDefinitionSpecificationWithDefaults instantiates a new WildcardBackendDefinitionSpecification object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWildcardBackendDefinitionSpecificationWithDefaults() *WildcardBackendDefinitionSpecification {
	this := WildcardBackendDefinitionSpecification{}
	return &this
}

// GetContents returns the Contents field value.
func (o *WildcardBackendDefinitionSpecification) GetContents() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value
// and a boolean to check if the value has been set.
func (o *WildcardBackendDefinitionSpecification) GetContentsOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contents, true
}

// SetContents sets field value.
func (o *WildcardBackendDefinitionSpecification) SetContents(v map[string]interface{}) {
	o.Contents = v
}

// GetType returns the Type field value.
func (o *WildcardBackendDefinitionSpecification) GetType() WildcardSpecificationType {
	if o == nil {
		var ret WildcardSpecificationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WildcardBackendDefinitionSpecification) GetTypeOk() (*WildcardSpecificationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *WildcardBackendDefinitionSpecification) SetType(v WildcardSpecificationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WildcardBackendDefinitionSpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["contents"] = o.Contents
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WildcardBackendDefinitionSpecification) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Contents *map[string]interface{}    `json:"contents"`
		Type     *WildcardSpecificationType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Contents == nil {
		return fmt.Errorf("required field contents missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Contents = *all.Contents
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
