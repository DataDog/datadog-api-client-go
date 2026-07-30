// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentOnCallPageTarget The target of an on-call page.
type IncidentOnCallPageTarget struct {
	// The identifier of the page target.
	Identifier string `json:"identifier"`
	// The type of the page target.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentOnCallPageTarget instantiates a new IncidentOnCallPageTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentOnCallPageTarget(identifier string, typeVar string) *IncidentOnCallPageTarget {
	this := IncidentOnCallPageTarget{}
	this.Identifier = identifier
	this.Type = typeVar
	return &this
}

// NewIncidentOnCallPageTargetWithDefaults instantiates a new IncidentOnCallPageTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentOnCallPageTargetWithDefaults() *IncidentOnCallPageTarget {
	this := IncidentOnCallPageTarget{}
	return &this
}

// GetIdentifier returns the Identifier field value.
func (o *IncidentOnCallPageTarget) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *IncidentOnCallPageTarget) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value.
func (o *IncidentOnCallPageTarget) SetIdentifier(v string) {
	o.Identifier = v
}

// GetType returns the Type field value.
func (o *IncidentOnCallPageTarget) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentOnCallPageTarget) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentOnCallPageTarget) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentOnCallPageTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["identifier"] = o.Identifier
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentOnCallPageTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Identifier *string `json:"identifier"`
		Type       *string `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Identifier == nil {
		return fmt.Errorf("required field identifier missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"identifier", "type"})
	} else {
		return err
	}
	o.Identifier = *all.Identifier
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
