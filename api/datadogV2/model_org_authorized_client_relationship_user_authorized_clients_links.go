// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgAuthorizedClientRelationshipUserAuthorizedClientsLinks Links for the user authorized clients relationship.
type OrgAuthorizedClientRelationshipUserAuthorizedClientsLinks struct {
	// Link to the user authorized clients for this org authorized client.
	Related string `json:"related"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgAuthorizedClientRelationshipUserAuthorizedClientsLinks instantiates a new OrgAuthorizedClientRelationshipUserAuthorizedClientsLinks object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgAuthorizedClientRelationshipUserAuthorizedClientsLinks(related string) *OrgAuthorizedClientRelationshipUserAuthorizedClientsLinks {
	this := OrgAuthorizedClientRelationshipUserAuthorizedClientsLinks{}
	this.Related = related
	return &this
}

// NewOrgAuthorizedClientRelationshipUserAuthorizedClientsLinksWithDefaults instantiates a new OrgAuthorizedClientRelationshipUserAuthorizedClientsLinks object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgAuthorizedClientRelationshipUserAuthorizedClientsLinksWithDefaults() *OrgAuthorizedClientRelationshipUserAuthorizedClientsLinks {
	this := OrgAuthorizedClientRelationshipUserAuthorizedClientsLinks{}
	return &this
}

// GetRelated returns the Related field value.
func (o *OrgAuthorizedClientRelationshipUserAuthorizedClientsLinks) GetRelated() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Related
}

// GetRelatedOk returns a tuple with the Related field value
// and a boolean to check if the value has been set.
func (o *OrgAuthorizedClientRelationshipUserAuthorizedClientsLinks) GetRelatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Related, true
}

// SetRelated sets field value.
func (o *OrgAuthorizedClientRelationshipUserAuthorizedClientsLinks) SetRelated(v string) {
	o.Related = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgAuthorizedClientRelationshipUserAuthorizedClientsLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["related"] = o.Related

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgAuthorizedClientRelationshipUserAuthorizedClientsLinks) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Related *string `json:"related"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Related == nil {
		return fmt.Errorf("required field related missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"related"})
	} else {
		return err
	}
	o.Related = *all.Related

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
