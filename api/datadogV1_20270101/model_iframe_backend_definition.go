// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IframeBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type IframeBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type IframeBackendDefinitionType `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Url string `json:"url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewIframeBackendDefinition instantiates a new IframeBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIframeBackendDefinition(typeVar IframeBackendDefinitionType, url string) *IframeBackendDefinition {
	this := IframeBackendDefinition{}
	this.Type = typeVar
	this.Url = url
	return &this
}

// NewIframeBackendDefinitionWithDefaults instantiates a new IframeBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIframeBackendDefinitionWithDefaults() *IframeBackendDefinition {
	this := IframeBackendDefinition{}
	return &this
}

// GetType returns the Type field value.
func (o *IframeBackendDefinition) GetType() IframeBackendDefinitionType {
	if o == nil {
		var ret IframeBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IframeBackendDefinition) GetTypeOk() (*IframeBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IframeBackendDefinition) SetType(v IframeBackendDefinitionType) {
	o.Type = v
}

// GetUrl returns the Url field value.
func (o *IframeBackendDefinition) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *IframeBackendDefinition) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *IframeBackendDefinition) SetUrl(v string) {
	o.Url = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IframeBackendDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IframeBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type *IframeBackendDefinitionType `json:"type"`
		Url  *string                      `json:"url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Url = *all.Url

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
