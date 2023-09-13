// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HttpDestination The HTTP destination.
type HttpDestination struct {
	// The authentication method used for HTTP destinations.
	Auth HttpDestinationAuth `json:"auth"`
	// The intake URL to forward events to.
	Endpoint string `json:"endpoint"`
	// The HTTP destination type.
	Type HttpDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewHttpDestination instantiates a new HttpDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHttpDestination(auth HttpDestinationAuth, endpoint string, typeVar HttpDestinationType) *HttpDestination {
	this := HttpDestination{}
	this.Auth = auth
	this.Endpoint = endpoint
	this.Type = typeVar
	return &this
}

// NewHttpDestinationWithDefaults instantiates a new HttpDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHttpDestinationWithDefaults() *HttpDestination {
	this := HttpDestination{}
	var typeVar HttpDestinationType = HTTPDESTINATIONTYPE_HTTP
	this.Type = typeVar
	return &this
}

// GetAuth returns the Auth field value.
func (o *HttpDestination) GetAuth() HttpDestinationAuth {
	if o == nil {
		var ret HttpDestinationAuth
		return ret
	}
	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *HttpDestination) GetAuthOk() (*HttpDestinationAuth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value.
func (o *HttpDestination) SetAuth(v HttpDestinationAuth) {
	o.Auth = v
}

// GetEndpoint returns the Endpoint field value.
func (o *HttpDestination) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *HttpDestination) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value.
func (o *HttpDestination) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetType returns the Type field value.
func (o *HttpDestination) GetType() HttpDestinationType {
	if o == nil {
		var ret HttpDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HttpDestination) GetTypeOk() (*HttpDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *HttpDestination) SetType(v HttpDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HttpDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["auth"] = o.Auth
	toSerialize["endpoint"] = o.Endpoint
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HttpDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Auth     *HttpDestinationAuth `json:"auth"`
		Endpoint *string              `json:"endpoint"`
		Type     *HttpDestinationType `json:"type"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Auth == nil {
		return fmt.Errorf("required field auth missing")
	}
	if all.Endpoint == nil {
		return fmt.Errorf("required field endpoint missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth", "endpoint", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Auth = *all.Auth
	o.Endpoint = *all.Endpoint
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
