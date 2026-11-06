// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OIDCCloudAuthProviderCreateData Custom OIDC provider creation data.
type OIDCCloudAuthProviderCreateData struct {
	// Attributes for creating a Custom OIDC provider.
	Attributes OIDCCloudAuthProviderCreateAttributes `json:"attributes"`
	// Resource type.
	Type OIDCCloudAuthProviderType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOIDCCloudAuthProviderCreateData instantiates a new OIDCCloudAuthProviderCreateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOIDCCloudAuthProviderCreateData(attributes OIDCCloudAuthProviderCreateAttributes, typeVar OIDCCloudAuthProviderType) *OIDCCloudAuthProviderCreateData {
	this := OIDCCloudAuthProviderCreateData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewOIDCCloudAuthProviderCreateDataWithDefaults instantiates a new OIDCCloudAuthProviderCreateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOIDCCloudAuthProviderCreateDataWithDefaults() *OIDCCloudAuthProviderCreateData {
	this := OIDCCloudAuthProviderCreateData{}
	var typeVar OIDCCloudAuthProviderType = OIDCCLOUDAUTHPROVIDERTYPE_CUSTOM_OIDC_PROVIDER_CONFIG
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *OIDCCloudAuthProviderCreateData) GetAttributes() OIDCCloudAuthProviderCreateAttributes {
	if o == nil {
		var ret OIDCCloudAuthProviderCreateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderCreateData) GetAttributesOk() (*OIDCCloudAuthProviderCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *OIDCCloudAuthProviderCreateData) SetAttributes(v OIDCCloudAuthProviderCreateAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *OIDCCloudAuthProviderCreateData) GetType() OIDCCloudAuthProviderType {
	if o == nil {
		var ret OIDCCloudAuthProviderType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderCreateData) GetTypeOk() (*OIDCCloudAuthProviderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OIDCCloudAuthProviderCreateData) SetType(v OIDCCloudAuthProviderType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OIDCCloudAuthProviderCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OIDCCloudAuthProviderCreateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *OIDCCloudAuthProviderCreateAttributes `json:"attributes"`
		Type       *OIDCCloudAuthProviderType             `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
