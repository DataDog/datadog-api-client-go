// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ClientTokenRevokeRequest Request to revoke a client token.
type ClientTokenRevokeRequest struct {
	// Hash value of the client token to revoke.
	Hash string `json:"hash"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClientTokenRevokeRequest instantiates a new ClientTokenRevokeRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClientTokenRevokeRequest(hash string) *ClientTokenRevokeRequest {
	this := ClientTokenRevokeRequest{}
	this.Hash = hash
	return &this
}

// NewClientTokenRevokeRequestWithDefaults instantiates a new ClientTokenRevokeRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClientTokenRevokeRequestWithDefaults() *ClientTokenRevokeRequest {
	this := ClientTokenRevokeRequest{}
	return &this
}

// GetHash returns the Hash field value.
func (o *ClientTokenRevokeRequest) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *ClientTokenRevokeRequest) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value.
func (o *ClientTokenRevokeRequest) SetHash(v string) {
	o.Hash = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClientTokenRevokeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["hash"] = o.Hash

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClientTokenRevokeRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Hash *string `json:"hash"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Hash == nil {
		return fmt.Errorf("required field hash missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hash"})
	} else {
		return err
	}
	o.Hash = *all.Hash

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
