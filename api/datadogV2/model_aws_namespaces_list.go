// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSNamespacesList AWS Metrics namespace filters
type AWSNamespacesList struct {
	// Exclude all namespaces
	ExcludeAll *bool `json:"exclude_all,omitempty"`
	// Exclude only these namespaces
	ExcludeOnly []string `json:"exclude_only,omitempty"`
	// Include all namespaces
	IncludeAll *bool `json:"include_all,omitempty"`
	// Include only these namespaces
	IncludeOnly []string `json:"include_only,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewAWSNamespacesList instantiates a new AWSNamespacesList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSNamespacesList() *AWSNamespacesList {
	this := AWSNamespacesList{}
	return &this
}

// NewAWSNamespacesListWithDefaults instantiates a new AWSNamespacesList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSNamespacesListWithDefaults() *AWSNamespacesList {
	this := AWSNamespacesList{}
	return &this
}

// GetExcludeAll returns the ExcludeAll field value if set, zero value otherwise.
func (o *AWSNamespacesList) GetExcludeAll() bool {
	if o == nil || o.ExcludeAll == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeAll
}

// GetExcludeAllOk returns a tuple with the ExcludeAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSNamespacesList) GetExcludeAllOk() (*bool, bool) {
	if o == nil || o.ExcludeAll == nil {
		return nil, false
	}
	return o.ExcludeAll, true
}

// HasExcludeAll returns a boolean if a field has been set.
func (o *AWSNamespacesList) HasExcludeAll() bool {
	return o != nil && o.ExcludeAll != nil
}

// SetExcludeAll gets a reference to the given bool and assigns it to the ExcludeAll field.
func (o *AWSNamespacesList) SetExcludeAll(v bool) {
	o.ExcludeAll = &v
}

// GetExcludeOnly returns the ExcludeOnly field value if set, zero value otherwise.
func (o *AWSNamespacesList) GetExcludeOnly() []string {
	if o == nil || o.ExcludeOnly == nil {
		var ret []string
		return ret
	}
	return o.ExcludeOnly
}

// GetExcludeOnlyOk returns a tuple with the ExcludeOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSNamespacesList) GetExcludeOnlyOk() (*[]string, bool) {
	if o == nil || o.ExcludeOnly == nil {
		return nil, false
	}
	return &o.ExcludeOnly, true
}

// HasExcludeOnly returns a boolean if a field has been set.
func (o *AWSNamespacesList) HasExcludeOnly() bool {
	return o != nil && o.ExcludeOnly != nil
}

// SetExcludeOnly gets a reference to the given []string and assigns it to the ExcludeOnly field.
func (o *AWSNamespacesList) SetExcludeOnly(v []string) {
	o.ExcludeOnly = v
}

// GetIncludeAll returns the IncludeAll field value if set, zero value otherwise.
func (o *AWSNamespacesList) GetIncludeAll() bool {
	if o == nil || o.IncludeAll == nil {
		var ret bool
		return ret
	}
	return *o.IncludeAll
}

// GetIncludeAllOk returns a tuple with the IncludeAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSNamespacesList) GetIncludeAllOk() (*bool, bool) {
	if o == nil || o.IncludeAll == nil {
		return nil, false
	}
	return o.IncludeAll, true
}

// HasIncludeAll returns a boolean if a field has been set.
func (o *AWSNamespacesList) HasIncludeAll() bool {
	return o != nil && o.IncludeAll != nil
}

// SetIncludeAll gets a reference to the given bool and assigns it to the IncludeAll field.
func (o *AWSNamespacesList) SetIncludeAll(v bool) {
	o.IncludeAll = &v
}

// GetIncludeOnly returns the IncludeOnly field value if set, zero value otherwise.
func (o *AWSNamespacesList) GetIncludeOnly() []string {
	if o == nil || o.IncludeOnly == nil {
		var ret []string
		return ret
	}
	return o.IncludeOnly
}

// GetIncludeOnlyOk returns a tuple with the IncludeOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSNamespacesList) GetIncludeOnlyOk() (*[]string, bool) {
	if o == nil || o.IncludeOnly == nil {
		return nil, false
	}
	return &o.IncludeOnly, true
}

// HasIncludeOnly returns a boolean if a field has been set.
func (o *AWSNamespacesList) HasIncludeOnly() bool {
	return o != nil && o.IncludeOnly != nil
}

// SetIncludeOnly gets a reference to the given []string and assigns it to the IncludeOnly field.
func (o *AWSNamespacesList) SetIncludeOnly(v []string) {
	o.IncludeOnly = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSNamespacesList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExcludeAll != nil {
		toSerialize["exclude_all"] = o.ExcludeAll
	}
	if o.ExcludeOnly != nil {
		toSerialize["exclude_only"] = o.ExcludeOnly
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
func (o *AWSNamespacesList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExcludeAll  *bool    `json:"exclude_all,omitempty"`
		ExcludeOnly []string `json:"exclude_only,omitempty"`
		IncludeAll  *bool    `json:"include_all,omitempty"`
		IncludeOnly []string `json:"include_only,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"exclude_all", "exclude_only", "include_all", "include_only"})
	} else {
		return err
	}
	o.ExcludeAll = all.ExcludeAll
	o.ExcludeOnly = all.ExcludeOnly
	o.IncludeAll = all.IncludeAll
	o.IncludeOnly = all.IncludeOnly

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
