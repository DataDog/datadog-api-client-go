// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationMetadata The metadata relevant to your configuration or recent request.
type CustomDestinationMetadata struct {
	// The amount of custom destinations your organization is able to create.
	CustomDestinationsLimit *int64 `json:"customDestinationsLimit,omitempty"`
	// The managed enclave IDs of the resource that is deleted.
	DeletedEnclaveIds []string `json:"deletedEnclaveIds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewCustomDestinationMetadata instantiates a new CustomDestinationMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomDestinationMetadata() *CustomDestinationMetadata {
	this := CustomDestinationMetadata{}
	return &this
}

// NewCustomDestinationMetadataWithDefaults instantiates a new CustomDestinationMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomDestinationMetadataWithDefaults() *CustomDestinationMetadata {
	this := CustomDestinationMetadata{}
	return &this
}

// GetCustomDestinationsLimit returns the CustomDestinationsLimit field value if set, zero value otherwise.
func (o *CustomDestinationMetadata) GetCustomDestinationsLimit() int64 {
	if o == nil || o.CustomDestinationsLimit == nil {
		var ret int64
		return ret
	}
	return *o.CustomDestinationsLimit
}

// GetCustomDestinationsLimitOk returns a tuple with the CustomDestinationsLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDestinationMetadata) GetCustomDestinationsLimitOk() (*int64, bool) {
	if o == nil || o.CustomDestinationsLimit == nil {
		return nil, false
	}
	return o.CustomDestinationsLimit, true
}

// HasCustomDestinationsLimit returns a boolean if a field has been set.
func (o *CustomDestinationMetadata) HasCustomDestinationsLimit() bool {
	return o != nil && o.CustomDestinationsLimit != nil
}

// SetCustomDestinationsLimit gets a reference to the given int64 and assigns it to the CustomDestinationsLimit field.
func (o *CustomDestinationMetadata) SetCustomDestinationsLimit(v int64) {
	o.CustomDestinationsLimit = &v
}

// GetDeletedEnclaveIds returns the DeletedEnclaveIds field value if set, zero value otherwise.
func (o *CustomDestinationMetadata) GetDeletedEnclaveIds() []string {
	if o == nil || o.DeletedEnclaveIds == nil {
		var ret []string
		return ret
	}
	return o.DeletedEnclaveIds
}

// GetDeletedEnclaveIdsOk returns a tuple with the DeletedEnclaveIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDestinationMetadata) GetDeletedEnclaveIdsOk() (*[]string, bool) {
	if o == nil || o.DeletedEnclaveIds == nil {
		return nil, false
	}
	return &o.DeletedEnclaveIds, true
}

// HasDeletedEnclaveIds returns a boolean if a field has been set.
func (o *CustomDestinationMetadata) HasDeletedEnclaveIds() bool {
	return o != nil && o.DeletedEnclaveIds != nil
}

// SetDeletedEnclaveIds gets a reference to the given []string and assigns it to the DeletedEnclaveIds field.
func (o *CustomDestinationMetadata) SetDeletedEnclaveIds(v []string) {
	o.DeletedEnclaveIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomDestinationMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.CustomDestinationsLimit != nil {
		toSerialize["customDestinationsLimit"] = o.CustomDestinationsLimit
	}
	if o.DeletedEnclaveIds != nil {
		toSerialize["deletedEnclaveIds"] = o.DeletedEnclaveIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomDestinationMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomDestinationsLimit *int64   `json:"customDestinationsLimit,omitempty"`
		DeletedEnclaveIds       []string `json:"deletedEnclaveIds,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"customDestinationsLimit", "deletedEnclaveIds"})
	} else {
		return err
	}
	o.CustomDestinationsLimit = all.CustomDestinationsLimit
	o.DeletedEnclaveIds = all.DeletedEnclaveIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
