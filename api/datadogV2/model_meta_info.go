// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetaInfo Additional information related to your service account.
type MetaInfo struct {
	// The current list of projects accessible from your service account.
	AccessibleProjects []string `json:"accessible_projects,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewMetaInfo instantiates a new MetaInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetaInfo() *MetaInfo {
	this := MetaInfo{}
	return &this
}

// NewMetaInfoWithDefaults instantiates a new MetaInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetaInfoWithDefaults() *MetaInfo {
	this := MetaInfo{}
	return &this
}

// GetAccessibleProjects returns the AccessibleProjects field value if set, zero value otherwise.
func (o *MetaInfo) GetAccessibleProjects() []string {
	if o == nil || o.AccessibleProjects == nil {
		var ret []string
		return ret
	}
	return o.AccessibleProjects
}

// GetAccessibleProjectsOk returns a tuple with the AccessibleProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaInfo) GetAccessibleProjectsOk() (*[]string, bool) {
	if o == nil || o.AccessibleProjects == nil {
		return nil, false
	}
	return &o.AccessibleProjects, true
}

// HasAccessibleProjects returns a boolean if a field has been set.
func (o *MetaInfo) HasAccessibleProjects() bool {
	return o != nil && o.AccessibleProjects != nil
}

// SetAccessibleProjects gets a reference to the given []string and assigns it to the AccessibleProjects field.
func (o *MetaInfo) SetAccessibleProjects(v []string) {
	o.AccessibleProjects = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetaInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.AccessibleProjects != nil {
		toSerialize["accessible_projects"] = o.AccessibleProjects
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetaInfo) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		AccessibleProjects []string `json:"accessible_projects,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"accessible_projects"})
	} else {
		return err
	}
	o.AccessibleProjects = all.AccessibleProjects
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
