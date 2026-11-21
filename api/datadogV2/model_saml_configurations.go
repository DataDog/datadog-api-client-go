// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SamlConfigurations A JSON array of saml configurations.
type SamlConfigurations struct {
	// The content of metadata XML file.
	IdpMetadata *UploadIdPMetadataOptionalParameters `json:"idp_metadata,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSamlConfigurations instantiates a new SamlConfigurations object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSamlConfigurations() *SamlConfigurations {
	this := SamlConfigurations{}
	return &this
}

// NewSamlConfigurationsWithDefaults instantiates a new SamlConfigurations object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSamlConfigurationsWithDefaults() *SamlConfigurations {
	this := SamlConfigurations{}
	return &this
}

// GetIdpMetadata returns the IdpMetadata field value if set, empty structure otherwise.
func (o *SamlConfigurations) GetIdpMetadata() UploadIdPMetadataOptionalParameters {
	if o == nil || o.IdpMetadata == nil {
		var ret UploadIdPMetadataOptionalParameters
		return ret
	}
	return *o.IdpMetadata
}

// GetIdpMetadataOk returns a tuple with the IdpMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfigurations) GetIdpMetadataOk() (*UploadIdPMetadataOptionalParameters, bool) {
	if o == nil || o.IdpMetadata == nil {
		return nil, false
	}
	return o.IdpMetadata, true
}

// HasIdpMetadata returns a boolean if a field has been set.
func (o *SamlConfigurations) HasIdpMetadata() bool {
	return o != nil && o.IdpMetadata != nil
}

// SetIdpMetadata gets a reference to the given structure and assigns it to the IdpMetadata field.
func (o *SamlConfigurations) SetIdpMetadata(v *UploadIdPMetadataOptionalParameters) {
	o.IdpMetadata = v
}

// MarshalJSON serializes the struct using spec logic.
func (o *SamlConfigurations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IdpMetadata != nil {
		toSerialize["idp_metadata"] = o.IdpMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SamlConfigurations) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IdpMetadata *UploadIdPMetadataOptionalParameters `json:"idp_metadata,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"idp_metadata"})
	} else {
		return err
	}

	o.IdpMetadata = all.IdpMetadata

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
