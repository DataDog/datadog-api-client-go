// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionConnectionAWSSecretsManager AWS Secrets Manager settings shared by the credential references on a `DATACENTER` connection. These settings are required whenever a field uses an AWS Secrets Manager reference.
type ActionConnectionAWSSecretsManager struct {
	// ID of the AWS connection used to retrieve secrets. This connection must be in the same Datadog organization.
	ConnectionId string `json:"connection_id"`
	// AWS region containing the referenced secrets.
	Region string `json:"region"`
	// AWS Secrets Manager secret backend.
	Type ActionConnectionAWSSecretsManagerType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewActionConnectionAWSSecretsManager instantiates a new ActionConnectionAWSSecretsManager object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewActionConnectionAWSSecretsManager(connectionId string, region string, typeVar ActionConnectionAWSSecretsManagerType) *ActionConnectionAWSSecretsManager {
	this := ActionConnectionAWSSecretsManager{}
	this.ConnectionId = connectionId
	this.Region = region
	this.Type = typeVar
	return &this
}

// NewActionConnectionAWSSecretsManagerWithDefaults instantiates a new ActionConnectionAWSSecretsManager object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewActionConnectionAWSSecretsManagerWithDefaults() *ActionConnectionAWSSecretsManager {
	this := ActionConnectionAWSSecretsManager{}
	return &this
}

// GetConnectionId returns the ConnectionId field value.
func (o *ActionConnectionAWSSecretsManager) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *ActionConnectionAWSSecretsManager) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value.
func (o *ActionConnectionAWSSecretsManager) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetRegion returns the Region field value.
func (o *ActionConnectionAWSSecretsManager) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ActionConnectionAWSSecretsManager) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *ActionConnectionAWSSecretsManager) SetRegion(v string) {
	o.Region = v
}

// GetType returns the Type field value.
func (o *ActionConnectionAWSSecretsManager) GetType() ActionConnectionAWSSecretsManagerType {
	if o == nil {
		var ret ActionConnectionAWSSecretsManagerType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ActionConnectionAWSSecretsManager) GetTypeOk() (*ActionConnectionAWSSecretsManagerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ActionConnectionAWSSecretsManager) SetType(v ActionConnectionAWSSecretsManagerType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ActionConnectionAWSSecretsManager) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["connection_id"] = o.ConnectionId
	toSerialize["region"] = o.Region
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ActionConnectionAWSSecretsManager) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConnectionId *string                                `json:"connection_id"`
		Region       *string                                `json:"region"`
		Type         *ActionConnectionAWSSecretsManagerType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConnectionId == nil {
		return fmt.Errorf("required field connection_id missing")
	}
	if all.Region == nil {
		return fmt.Errorf("required field region missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"connection_id", "region", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConnectionId = *all.ConnectionId
	o.Region = *all.Region
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

// NullableActionConnectionAWSSecretsManager handles when a null is used for ActionConnectionAWSSecretsManager.
type NullableActionConnectionAWSSecretsManager struct {
	value *ActionConnectionAWSSecretsManager
	isSet bool
}

// Get returns the associated value.
func (v NullableActionConnectionAWSSecretsManager) Get() *ActionConnectionAWSSecretsManager {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableActionConnectionAWSSecretsManager) Set(val *ActionConnectionAWSSecretsManager) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableActionConnectionAWSSecretsManager) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableActionConnectionAWSSecretsManager) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableActionConnectionAWSSecretsManager initializes the struct as if Set has been called.
func NewNullableActionConnectionAWSSecretsManager(val *ActionConnectionAWSSecretsManager) *NullableActionConnectionAWSSecretsManager {
	return &NullableActionConnectionAWSSecretsManager{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableActionConnectionAWSSecretsManager) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableActionConnectionAWSSecretsManager) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
