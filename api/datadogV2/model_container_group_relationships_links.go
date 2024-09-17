// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV2

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"

)


// ContainerGroupRelationshipsLinks Links attributes.
type ContainerGroupRelationshipsLinks struct {
	// Link to related containers.
	Related *string `json:"related,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewContainerGroupRelationshipsLinks instantiates a new ContainerGroupRelationshipsLinks object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewContainerGroupRelationshipsLinks() *ContainerGroupRelationshipsLinks {
	this := ContainerGroupRelationshipsLinks{}
	return &this
}

// NewContainerGroupRelationshipsLinksWithDefaults instantiates a new ContainerGroupRelationshipsLinks object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewContainerGroupRelationshipsLinksWithDefaults() *ContainerGroupRelationshipsLinks {
	this := ContainerGroupRelationshipsLinks{}
	return &this
}
// GetRelated returns the Related field value if set, zero value otherwise.
func (o *ContainerGroupRelationshipsLinks) GetRelated() string {
	if o == nil || o.Related == nil {
		var ret string
		return ret
	}
	return *o.Related
}

// GetRelatedOk returns a tuple with the Related field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroupRelationshipsLinks) GetRelatedOk() (*string, bool) {
	if o == nil || o.Related == nil {
		return nil, false
	}
	return o.Related, true
}

// HasRelated returns a boolean if a field has been set.
func (o *ContainerGroupRelationshipsLinks) HasRelated() bool {
	return o != nil && o.Related != nil
}

// SetRelated gets a reference to the given string and assigns it to the Related field.
func (o *ContainerGroupRelationshipsLinks) SetRelated(v string) {
	o.Related = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o ContainerGroupRelationshipsLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Related != nil {
		toSerialize["related"] = o.Related
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ContainerGroupRelationshipsLinks) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Related *string `json:"related,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{ "related",  })
	} else {
		return err
	}
	o.Related = all.Related

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
