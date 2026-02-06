// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ClientTokenUpdateRequest Request to update an existing client token.
type ClientTokenUpdateRequest struct {
	// Hash value of the client token to update.
	Hash string `json:"hash"`
	// New name for the client token.
	Name string `json:"name"`
	// New list of allowed origin URLs. If provided, this will replace the existing list.
	OriginUrls []string `json:"origin_urls,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClientTokenUpdateRequest instantiates a new ClientTokenUpdateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClientTokenUpdateRequest(hash string, name string) *ClientTokenUpdateRequest {
	this := ClientTokenUpdateRequest{}
	this.Hash = hash
	this.Name = name
	return &this
}

// NewClientTokenUpdateRequestWithDefaults instantiates a new ClientTokenUpdateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClientTokenUpdateRequestWithDefaults() *ClientTokenUpdateRequest {
	this := ClientTokenUpdateRequest{}
	return &this
}

// GetHash returns the Hash field value.
func (o *ClientTokenUpdateRequest) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *ClientTokenUpdateRequest) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value.
func (o *ClientTokenUpdateRequest) SetHash(v string) {
	o.Hash = v
}

// GetName returns the Name field value.
func (o *ClientTokenUpdateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClientTokenUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ClientTokenUpdateRequest) SetName(v string) {
	o.Name = v
}

// GetOriginUrls returns the OriginUrls field value if set, zero value otherwise.
func (o *ClientTokenUpdateRequest) GetOriginUrls() []string {
	if o == nil || o.OriginUrls == nil {
		var ret []string
		return ret
	}
	return o.OriginUrls
}

// GetOriginUrlsOk returns a tuple with the OriginUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientTokenUpdateRequest) GetOriginUrlsOk() (*[]string, bool) {
	if o == nil || o.OriginUrls == nil {
		return nil, false
	}
	return &o.OriginUrls, true
}

// HasOriginUrls returns a boolean if a field has been set.
func (o *ClientTokenUpdateRequest) HasOriginUrls() bool {
	return o != nil && o.OriginUrls != nil
}

// SetOriginUrls gets a reference to the given []string and assigns it to the OriginUrls field.
func (o *ClientTokenUpdateRequest) SetOriginUrls(v []string) {
	o.OriginUrls = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClientTokenUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["hash"] = o.Hash
	toSerialize["name"] = o.Name
	if o.OriginUrls != nil {
		toSerialize["origin_urls"] = o.OriginUrls
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClientTokenUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Hash       *string  `json:"hash"`
		Name       *string  `json:"name"`
		OriginUrls []string `json:"origin_urls,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Hash == nil {
		return fmt.Errorf("required field hash missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hash", "name", "origin_urls"})
	} else {
		return err
	}
	o.Hash = *all.Hash
	o.Name = *all.Name
	o.OriginUrls = all.OriginUrls

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
