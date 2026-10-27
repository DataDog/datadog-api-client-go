// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArchiveSearchRehydrateRequestAttributes Attributes accepted when rehydrating an Archive Search.
type ArchiveSearchRehydrateRequestAttributes struct {
	// Rehydration settings to apply.
	Rehydration ArchiveSearchRehydrateRequestRehydration `json:"rehydration"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewArchiveSearchRehydrateRequestAttributes instantiates a new ArchiveSearchRehydrateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewArchiveSearchRehydrateRequestAttributes(rehydration ArchiveSearchRehydrateRequestRehydration) *ArchiveSearchRehydrateRequestAttributes {
	this := ArchiveSearchRehydrateRequestAttributes{}
	this.Rehydration = rehydration
	return &this
}

// NewArchiveSearchRehydrateRequestAttributesWithDefaults instantiates a new ArchiveSearchRehydrateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewArchiveSearchRehydrateRequestAttributesWithDefaults() *ArchiveSearchRehydrateRequestAttributes {
	this := ArchiveSearchRehydrateRequestAttributes{}
	return &this
}

// GetRehydration returns the Rehydration field value.
func (o *ArchiveSearchRehydrateRequestAttributes) GetRehydration() ArchiveSearchRehydrateRequestRehydration {
	if o == nil {
		var ret ArchiveSearchRehydrateRequestRehydration
		return ret
	}
	return o.Rehydration
}

// GetRehydrationOk returns a tuple with the Rehydration field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchRehydrateRequestAttributes) GetRehydrationOk() (*ArchiveSearchRehydrateRequestRehydration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rehydration, true
}

// SetRehydration sets field value.
func (o *ArchiveSearchRehydrateRequestAttributes) SetRehydration(v ArchiveSearchRehydrateRequestRehydration) {
	o.Rehydration = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ArchiveSearchRehydrateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["rehydration"] = o.Rehydration

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ArchiveSearchRehydrateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Rehydration *ArchiveSearchRehydrateRequestRehydration `json:"rehydration"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Rehydration == nil {
		return fmt.Errorf("required field rehydration missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rehydration"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Rehydration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rehydration = *all.Rehydration

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
