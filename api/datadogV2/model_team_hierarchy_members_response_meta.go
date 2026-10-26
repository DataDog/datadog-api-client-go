// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyMembersResponseMeta Pagination metadata for the filtered members.
type TeamHierarchyMembersResponseMeta struct {
	// Pagination metadata for the selected pagination mode.
	Page TeamHierarchyMembersPage `json:"page"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamHierarchyMembersResponseMeta instantiates a new TeamHierarchyMembersResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamHierarchyMembersResponseMeta(page TeamHierarchyMembersPage) *TeamHierarchyMembersResponseMeta {
	this := TeamHierarchyMembersResponseMeta{}
	this.Page = page
	return &this
}

// NewTeamHierarchyMembersResponseMetaWithDefaults instantiates a new TeamHierarchyMembersResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamHierarchyMembersResponseMetaWithDefaults() *TeamHierarchyMembersResponseMeta {
	this := TeamHierarchyMembersResponseMeta{}
	return &this
}

// GetPage returns the Page field value.
func (o *TeamHierarchyMembersResponseMeta) GetPage() TeamHierarchyMembersPage {
	if o == nil {
		var ret TeamHierarchyMembersPage
		return ret
	}
	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMembersResponseMeta) GetPageOk() (*TeamHierarchyMembersPage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value.
func (o *TeamHierarchyMembersResponseMeta) SetPage(v TeamHierarchyMembersPage) {
	o.Page = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamHierarchyMembersResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["page"] = o.Page

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamHierarchyMembersResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Page *TeamHierarchyMembersPage `json:"page"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Page == nil {
		return fmt.Errorf("required field page missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"page"})
	} else {
		return err
	}
	o.Page = *all.Page

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
