// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// STIXBundleRequest A STIX 2.1 bundle containing threat intelligence indicator objects.
type STIXBundleRequest struct {
	// The STIX bundle identifier.
	Id string `json:"id"`
	// The indicator objects included in the bundle.
	Objects []STIXIndicatorObject `json:"objects"`
	// The supported STIX specification version.
	SpecVersion *STIXSpecVersion `json:"spec_version,omitempty"`
	// The STIX object type for a bundle.
	Type STIXBundleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSTIXBundleRequest instantiates a new STIXBundleRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSTIXBundleRequest(id string, objects []STIXIndicatorObject, typeVar STIXBundleType) *STIXBundleRequest {
	this := STIXBundleRequest{}
	this.Id = id
	this.Objects = objects
	this.Type = typeVar
	return &this
}

// NewSTIXBundleRequestWithDefaults instantiates a new STIXBundleRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSTIXBundleRequestWithDefaults() *STIXBundleRequest {
	this := STIXBundleRequest{}
	return &this
}

// GetId returns the Id field value.
func (o *STIXBundleRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *STIXBundleRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *STIXBundleRequest) SetId(v string) {
	o.Id = v
}

// GetObjects returns the Objects field value.
func (o *STIXBundleRequest) GetObjects() []STIXIndicatorObject {
	if o == nil {
		var ret []STIXIndicatorObject
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *STIXBundleRequest) GetObjectsOk() (*[]STIXIndicatorObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Objects, true
}

// SetObjects sets field value.
func (o *STIXBundleRequest) SetObjects(v []STIXIndicatorObject) {
	o.Objects = v
}

// GetSpecVersion returns the SpecVersion field value if set, zero value otherwise.
func (o *STIXBundleRequest) GetSpecVersion() STIXSpecVersion {
	if o == nil || o.SpecVersion == nil {
		var ret STIXSpecVersion
		return ret
	}
	return *o.SpecVersion
}

// GetSpecVersionOk returns a tuple with the SpecVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXBundleRequest) GetSpecVersionOk() (*STIXSpecVersion, bool) {
	if o == nil || o.SpecVersion == nil {
		return nil, false
	}
	return o.SpecVersion, true
}

// HasSpecVersion returns a boolean if a field has been set.
func (o *STIXBundleRequest) HasSpecVersion() bool {
	return o != nil && o.SpecVersion != nil
}

// SetSpecVersion gets a reference to the given STIXSpecVersion and assigns it to the SpecVersion field.
func (o *STIXBundleRequest) SetSpecVersion(v STIXSpecVersion) {
	o.SpecVersion = &v
}

// GetType returns the Type field value.
func (o *STIXBundleRequest) GetType() STIXBundleType {
	if o == nil {
		var ret STIXBundleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *STIXBundleRequest) GetTypeOk() (*STIXBundleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *STIXBundleRequest) SetType(v STIXBundleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o STIXBundleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["objects"] = o.Objects
	if o.SpecVersion != nil {
		toSerialize["spec_version"] = o.SpecVersion
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *STIXBundleRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id          *string                `json:"id"`
		Objects     *[]STIXIndicatorObject `json:"objects"`
		SpecVersion *STIXSpecVersion       `json:"spec_version,omitempty"`
		Type        *STIXBundleType        `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Objects == nil {
		return fmt.Errorf("required field objects missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "objects", "spec_version", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	o.Objects = *all.Objects
	if all.SpecVersion != nil && !all.SpecVersion.IsValid() {
		hasInvalidField = true
	} else {
		o.SpecVersion = all.SpecVersion
	}
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
