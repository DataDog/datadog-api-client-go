// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSSecretManager AWS Secrets Manager configuration.
type AWSSecretManager struct {
	// The ID of the connection used to access AWS Secrets Manager.
	ConnectionId uuid.UUID `json:"connection_id"`
	// The AWS region where the secret is stored.
	Region string `json:"region"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSSecretManager instantiates a new AWSSecretManager object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSSecretManager(connectionId uuid.UUID, region string) *AWSSecretManager {
	this := AWSSecretManager{}
	this.ConnectionId = connectionId
	this.Region = region
	return &this
}

// NewAWSSecretManagerWithDefaults instantiates a new AWSSecretManager object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSSecretManagerWithDefaults() *AWSSecretManager {
	this := AWSSecretManager{}
	return &this
}

// GetConnectionId returns the ConnectionId field value.
func (o *AWSSecretManager) GetConnectionId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *AWSSecretManager) GetConnectionIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value.
func (o *AWSSecretManager) SetConnectionId(v uuid.UUID) {
	o.ConnectionId = v
}

// GetRegion returns the Region field value.
func (o *AWSSecretManager) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *AWSSecretManager) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *AWSSecretManager) SetRegion(v string) {
	o.Region = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSSecretManager) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["connection_id"] = o.ConnectionId
	toSerialize["region"] = o.Region

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSSecretManager) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConnectionId *uuid.UUID `json:"connection_id"`
		Region       *string    `json:"region"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"connection_id", "region"})
	} else {
		return err
	}
	o.ConnectionId = *all.ConnectionId
	o.Region = *all.Region

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
