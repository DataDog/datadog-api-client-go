// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyMembersResponse Direct and inherited members of the requested team.
type TeamHierarchyMembersResponse struct {
	// Membership relationships for users who belong to the requested team or inherit membership through its descendant teams. Each user appears once, even if they belong to multiple teams in the hierarchy. Each relationship identifies the user and requested team, and indicates whether the user is a direct member of the requested team.
	Data []TeamHierarchyMember `json:"data"`
	// Requested users and the team. May be absent when the result is empty, even if requested.
	Included []TeamHierarchyMemberIncluded `json:"included,omitempty"`
	// Page navigation links. Unavailable links are omitted.
	Links TeamHierarchyMembersResponseLinks `json:"links"`
	// Pagination metadata for the filtered members.
	Meta TeamHierarchyMembersResponseMeta `json:"meta"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamHierarchyMembersResponse instantiates a new TeamHierarchyMembersResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamHierarchyMembersResponse(data []TeamHierarchyMember, links TeamHierarchyMembersResponseLinks, meta TeamHierarchyMembersResponseMeta) *TeamHierarchyMembersResponse {
	this := TeamHierarchyMembersResponse{}
	this.Data = data
	this.Links = links
	this.Meta = meta
	return &this
}

// NewTeamHierarchyMembersResponseWithDefaults instantiates a new TeamHierarchyMembersResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamHierarchyMembersResponseWithDefaults() *TeamHierarchyMembersResponse {
	this := TeamHierarchyMembersResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *TeamHierarchyMembersResponse) GetData() []TeamHierarchyMember {
	if o == nil {
		var ret []TeamHierarchyMember
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMembersResponse) GetDataOk() (*[]TeamHierarchyMember, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *TeamHierarchyMembersResponse) SetData(v []TeamHierarchyMember) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *TeamHierarchyMembersResponse) GetIncluded() []TeamHierarchyMemberIncluded {
	if o == nil || o.Included == nil {
		var ret []TeamHierarchyMemberIncluded
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMembersResponse) GetIncludedOk() (*[]TeamHierarchyMemberIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *TeamHierarchyMembersResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []TeamHierarchyMemberIncluded and assigns it to the Included field.
func (o *TeamHierarchyMembersResponse) SetIncluded(v []TeamHierarchyMemberIncluded) {
	o.Included = v
}

// GetLinks returns the Links field value.
func (o *TeamHierarchyMembersResponse) GetLinks() TeamHierarchyMembersResponseLinks {
	if o == nil {
		var ret TeamHierarchyMembersResponseLinks
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMembersResponse) GetLinksOk() (*TeamHierarchyMembersResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value.
func (o *TeamHierarchyMembersResponse) SetLinks(v TeamHierarchyMembersResponseLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value.
func (o *TeamHierarchyMembersResponse) GetMeta() TeamHierarchyMembersResponseMeta {
	if o == nil {
		var ret TeamHierarchyMembersResponseMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMembersResponse) GetMetaOk() (*TeamHierarchyMembersResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *TeamHierarchyMembersResponse) SetMeta(v TeamHierarchyMembersResponseMeta) {
	o.Meta = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamHierarchyMembersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamHierarchyMembersResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *[]TeamHierarchyMember             `json:"data"`
		Included []TeamHierarchyMemberIncluded      `json:"included,omitempty"`
		Links    *TeamHierarchyMembersResponseLinks `json:"links"`
		Meta     *TeamHierarchyMembersResponseMeta  `json:"meta"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	if all.Links == nil {
		return fmt.Errorf("required field links missing")
	}
	if all.Meta == nil {
		return fmt.Errorf("required field meta missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included", "links", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	o.Included = all.Included
	if all.Links.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Links = *all.Links
	if all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = *all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
