// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ClientTokenCreateRequest Request to create a new client token.
type ClientTokenCreateRequest struct {
	// Name of the client token.
	Name string `json:"name"`
	// List of allowed origin URLs for browser-based applications. Requests from URLs
	// not in this list will be rejected.
	OriginUrls []string `json:"origin_urls"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClientTokenCreateRequest instantiates a new ClientTokenCreateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClientTokenCreateRequest(name string, originUrls []string) *ClientTokenCreateRequest {
	this := ClientTokenCreateRequest{}
	this.Name = name
	this.OriginUrls = originUrls
	return &this
}

// NewClientTokenCreateRequestWithDefaults instantiates a new ClientTokenCreateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClientTokenCreateRequestWithDefaults() *ClientTokenCreateRequest {
	this := ClientTokenCreateRequest{}
	return &this
}

// GetName returns the Name field value.
func (o *ClientTokenCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClientTokenCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ClientTokenCreateRequest) SetName(v string) {
	o.Name = v
}

// GetOriginUrls returns the OriginUrls field value.
func (o *ClientTokenCreateRequest) GetOriginUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OriginUrls
}

// GetOriginUrlsOk returns a tuple with the OriginUrls field value
// and a boolean to check if the value has been set.
func (o *ClientTokenCreateRequest) GetOriginUrlsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginUrls, true
}

// SetOriginUrls sets field value.
func (o *ClientTokenCreateRequest) SetOriginUrls(v []string) {
	o.OriginUrls = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClientTokenCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["origin_urls"] = o.OriginUrls

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClientTokenCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name       *string   `json:"name"`
		OriginUrls *[]string `json:"origin_urls"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OriginUrls == nil {
		return fmt.Errorf("required field origin_urls missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "origin_urls"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.OriginUrls = *all.OriginUrls

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
