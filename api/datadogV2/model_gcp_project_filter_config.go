// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPProjectFilterConfig Configuration for filtering the GCP projects that Datadog monitors by project ID or by project label.
type GCPProjectFilterConfig struct {
	// List of filters matched against the project field set by `type`. Filters support the `*` and `?`
	// wildcards, and a leading `!` makes a filter an exclusion. A project is skipped if it matches any
	// exclusion filter. When inclusion filters are set, a project must also match at least one of them.
	// For `project_id`, each filter is a project ID pattern, such as `prod-*`. You can set up to 15,000
	// filters. For `project_label`, each filter is a label in the form `key:value`, such as `env:prod*`.
	// You can set up to 100 filters. Keys are matched exactly, and wildcards apply only to values, so
	// `key:*` matches any value.
	Filters []string `json:"filters,omitempty"`
	// The GCP project field that the filters are matched against.
	Type *GCPProjectFilterConfigType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGCPProjectFilterConfig instantiates a new GCPProjectFilterConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGCPProjectFilterConfig() *GCPProjectFilterConfig {
	this := GCPProjectFilterConfig{}
	return &this
}

// NewGCPProjectFilterConfigWithDefaults instantiates a new GCPProjectFilterConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGCPProjectFilterConfigWithDefaults() *GCPProjectFilterConfig {
	this := GCPProjectFilterConfig{}
	return &this
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *GCPProjectFilterConfig) GetFilters() []string {
	if o == nil || o.Filters == nil {
		var ret []string
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPProjectFilterConfig) GetFiltersOk() (*[]string, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return &o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *GCPProjectFilterConfig) HasFilters() bool {
	return o != nil && o.Filters != nil
}

// SetFilters gets a reference to the given []string and assigns it to the Filters field.
func (o *GCPProjectFilterConfig) SetFilters(v []string) {
	o.Filters = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GCPProjectFilterConfig) GetType() GCPProjectFilterConfigType {
	if o == nil || o.Type == nil {
		var ret GCPProjectFilterConfigType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPProjectFilterConfig) GetTypeOk() (*GCPProjectFilterConfigType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GCPProjectFilterConfig) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given GCPProjectFilterConfigType and assigns it to the Type field.
func (o *GCPProjectFilterConfig) SetType(v GCPProjectFilterConfigType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GCPProjectFilterConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GCPProjectFilterConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filters []string                    `json:"filters,omitempty"`
		Type    *GCPProjectFilterConfigType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filters", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Filters = all.Filters
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
