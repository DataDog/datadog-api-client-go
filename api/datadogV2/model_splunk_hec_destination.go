// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplunkHecDestination The Splunk destination.
type SplunkHecDestination struct {
	// The access token of the Splunk destination.
	AccessToken string `json:"accessToken"`
	// The intake URL to forward events to.
	Endpoint string `json:"endpoint"`
	// The Splunk destination type.
	Type SplunkHecDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSplunkHecDestination instantiates a new SplunkHecDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSplunkHecDestination(accessToken string, endpoint string, typeVar SplunkHecDestinationType) *SplunkHecDestination {
	this := SplunkHecDestination{}
	this.AccessToken = accessToken
	this.Endpoint = endpoint
	this.Type = typeVar
	return &this
}

// NewSplunkHecDestinationWithDefaults instantiates a new SplunkHecDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSplunkHecDestinationWithDefaults() *SplunkHecDestination {
	this := SplunkHecDestination{}
	var typeVar SplunkHecDestinationType = SPLUNKHECDESTINATIONTYPE_SPLUNK_HEC
	this.Type = typeVar
	return &this
}

// GetAccessToken returns the AccessToken field value.
func (o *SplunkHecDestination) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *SplunkHecDestination) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value.
func (o *SplunkHecDestination) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetEndpoint returns the Endpoint field value.
func (o *SplunkHecDestination) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *SplunkHecDestination) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value.
func (o *SplunkHecDestination) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetType returns the Type field value.
func (o *SplunkHecDestination) GetType() SplunkHecDestinationType {
	if o == nil {
		var ret SplunkHecDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SplunkHecDestination) GetTypeOk() (*SplunkHecDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SplunkHecDestination) SetType(v SplunkHecDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SplunkHecDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["accessToken"] = o.AccessToken
	toSerialize["endpoint"] = o.Endpoint
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SplunkHecDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessToken *string                   `json:"accessToken"`
		Endpoint    *string                   `json:"endpoint"`
		Type        *SplunkHecDestinationType `json:"type"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccessToken == nil {
		return fmt.Errorf("required field accessToken missing")
	}
	if all.Endpoint == nil {
		return fmt.Errorf("required field endpoint missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"accessToken", "endpoint", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccessToken = *all.AccessToken
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
