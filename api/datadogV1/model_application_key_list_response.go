// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV1

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"

)


// ApplicationKeyListResponse An application key response.
type ApplicationKeyListResponse struct {
	// Array of application keys.
	ApplicationKeys []ApplicationKey `json:"application_keys,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewApplicationKeyListResponse instantiates a new ApplicationKeyListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationKeyListResponse() *ApplicationKeyListResponse {
	this := ApplicationKeyListResponse{}
	return &this
}

// NewApplicationKeyListResponseWithDefaults instantiates a new ApplicationKeyListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationKeyListResponseWithDefaults() *ApplicationKeyListResponse {
	this := ApplicationKeyListResponse{}
	return &this
}
// GetApplicationKeys returns the ApplicationKeys field value if set, zero value otherwise.
func (o *ApplicationKeyListResponse) GetApplicationKeys() []ApplicationKey {
	if o == nil || o.ApplicationKeys == nil {
		var ret []ApplicationKey
		return ret
	}
	return o.ApplicationKeys
}

// GetApplicationKeysOk returns a tuple with the ApplicationKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationKeyListResponse) GetApplicationKeysOk() (*[]ApplicationKey, bool) {
	if o == nil || o.ApplicationKeys == nil {
		return nil, false
	}
	return &o.ApplicationKeys, true
}

// HasApplicationKeys returns a boolean if a field has been set.
func (o *ApplicationKeyListResponse) HasApplicationKeys() bool {
	return o != nil && o.ApplicationKeys != nil
}

// SetApplicationKeys gets a reference to the given []ApplicationKey and assigns it to the ApplicationKeys field.
func (o *ApplicationKeyListResponse) SetApplicationKeys(v []ApplicationKey) {
	o.ApplicationKeys = v
}



// MarshalJSON serializes the struct using spec logic.
func (o ApplicationKeyListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApplicationKeys != nil {
		toSerialize["application_keys"] = o.ApplicationKeys
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationKeyListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationKeys []ApplicationKey `json:"application_keys,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{ "application_keys",  })
	} else {
		return err
	}
	o.ApplicationKeys = all.ApplicationKeys

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
