// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSXRayServicesList AWS X-Ray services to collect traces from
type AWSXRayServicesList struct {
	// Include all services
	IncludeAll *bool `json:"include_all,omitempty"`
	// Include only these services
	IncludeOnly []string `json:"include_only,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewAWSXRayServicesList instantiates a new AWSXRayServicesList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSXRayServicesList() *AWSXRayServicesList {
	this := AWSXRayServicesList{}
	return &this
}

// NewAWSXRayServicesListWithDefaults instantiates a new AWSXRayServicesList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSXRayServicesListWithDefaults() *AWSXRayServicesList {
	this := AWSXRayServicesList{}
	return &this
}

// GetIncludeAll returns the IncludeAll field value if set, zero value otherwise.
func (o *AWSXRayServicesList) GetIncludeAll() bool {
	if o == nil || o.IncludeAll == nil {
		var ret bool
		return ret
	}
	return *o.IncludeAll
}

// GetIncludeAllOk returns a tuple with the IncludeAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSXRayServicesList) GetIncludeAllOk() (*bool, bool) {
	if o == nil || o.IncludeAll == nil {
		return nil, false
	}
	return o.IncludeAll, true
}

// HasIncludeAll returns a boolean if a field has been set.
func (o *AWSXRayServicesList) HasIncludeAll() bool {
	return o != nil && o.IncludeAll != nil
}

// SetIncludeAll gets a reference to the given bool and assigns it to the IncludeAll field.
func (o *AWSXRayServicesList) SetIncludeAll(v bool) {
	o.IncludeAll = &v
}

// GetIncludeOnly returns the IncludeOnly field value if set, zero value otherwise.
func (o *AWSXRayServicesList) GetIncludeOnly() []string {
	if o == nil || o.IncludeOnly == nil {
		var ret []string
		return ret
	}
	return o.IncludeOnly
}

// GetIncludeOnlyOk returns a tuple with the IncludeOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSXRayServicesList) GetIncludeOnlyOk() (*[]string, bool) {
	if o == nil || o.IncludeOnly == nil {
		return nil, false
	}
	return &o.IncludeOnly, true
}

// HasIncludeOnly returns a boolean if a field has been set.
func (o *AWSXRayServicesList) HasIncludeOnly() bool {
	return o != nil && o.IncludeOnly != nil
}

// SetIncludeOnly gets a reference to the given []string and assigns it to the IncludeOnly field.
func (o *AWSXRayServicesList) SetIncludeOnly(v []string) {
	o.IncludeOnly = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSXRayServicesList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IncludeAll != nil {
		toSerialize["include_all"] = o.IncludeAll
	}
	if o.IncludeOnly != nil {
		toSerialize["include_only"] = o.IncludeOnly
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSXRayServicesList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncludeAll  *bool    `json:"include_all,omitempty"`
		IncludeOnly []string `json:"include_only,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"include_all", "include_only"})
	} else {
		return err
	}
	o.IncludeAll = all.IncludeAll
	o.IncludeOnly = all.IncludeOnly

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
