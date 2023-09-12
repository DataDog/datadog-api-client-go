// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HttpDestinationCustomHeaderAuth The HTTP destination custom header auth.
type HttpDestinationCustomHeaderAuth struct {
	// The name of the header to use for the HTTP request.
	HeaderName string `json:"headerName"`
	// The value of the header to use for the HTTP request.
	HeaderValue string `json:"headerValue"`
	// The HTTP destination custom header auth type.
	Type HttpDestinationCustomHeaderAuthType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewHttpDestinationCustomHeaderAuth instantiates a new HttpDestinationCustomHeaderAuth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHttpDestinationCustomHeaderAuth(headerName string, headerValue string, typeVar HttpDestinationCustomHeaderAuthType) *HttpDestinationCustomHeaderAuth {
	this := HttpDestinationCustomHeaderAuth{}
	this.HeaderName = headerName
	this.HeaderValue = headerValue
	this.Type = typeVar
	return &this
}

// NewHttpDestinationCustomHeaderAuthWithDefaults instantiates a new HttpDestinationCustomHeaderAuth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHttpDestinationCustomHeaderAuthWithDefaults() *HttpDestinationCustomHeaderAuth {
	this := HttpDestinationCustomHeaderAuth{}
	var typeVar HttpDestinationCustomHeaderAuthType = HTTPDESTINATIONCUSTOMHEADERAUTHTYPE_CUSTOM_HEADER
	this.Type = typeVar
	return &this
}

// GetHeaderName returns the HeaderName field value.
func (o *HttpDestinationCustomHeaderAuth) GetHeaderName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.HeaderName
}

// GetHeaderNameOk returns a tuple with the HeaderName field value
// and a boolean to check if the value has been set.
func (o *HttpDestinationCustomHeaderAuth) GetHeaderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeaderName, true
}

// SetHeaderName sets field value.
func (o *HttpDestinationCustomHeaderAuth) SetHeaderName(v string) {
	o.HeaderName = v
}

// GetHeaderValue returns the HeaderValue field value.
func (o *HttpDestinationCustomHeaderAuth) GetHeaderValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.HeaderValue
}

// GetHeaderValueOk returns a tuple with the HeaderValue field value
// and a boolean to check if the value has been set.
func (o *HttpDestinationCustomHeaderAuth) GetHeaderValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeaderValue, true
}

// SetHeaderValue sets field value.
func (o *HttpDestinationCustomHeaderAuth) SetHeaderValue(v string) {
	o.HeaderValue = v
}

// GetType returns the Type field value.
func (o *HttpDestinationCustomHeaderAuth) GetType() HttpDestinationCustomHeaderAuthType {
	if o == nil {
		var ret HttpDestinationCustomHeaderAuthType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HttpDestinationCustomHeaderAuth) GetTypeOk() (*HttpDestinationCustomHeaderAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *HttpDestinationCustomHeaderAuth) SetType(v HttpDestinationCustomHeaderAuthType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HttpDestinationCustomHeaderAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["headerName"] = o.HeaderName
	toSerialize["headerValue"] = o.HeaderValue
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HttpDestinationCustomHeaderAuth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HeaderName  *string                              `json:"headerName"`
		HeaderValue *string                              `json:"headerValue"`
		Type        *HttpDestinationCustomHeaderAuthType `json:"type"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.HeaderName == nil {
		return fmt.Errorf("required field headerName missing")
	}
	if all.HeaderValue == nil {
		return fmt.Errorf("required field headerValue missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"headerName", "headerValue", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.HeaderName = *all.HeaderName
	o.HeaderValue = *all.HeaderValue
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
