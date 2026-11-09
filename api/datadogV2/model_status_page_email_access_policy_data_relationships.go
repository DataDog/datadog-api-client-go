// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPageEmailAccessPolicyDataRelationships Relationships of a status page email access policy.
type StatusPageEmailAccessPolicyDataRelationships struct {
	// The Datadog user who created the email access policy.
	CreatedByUser StatusPageEmailAccessPolicyDataRelationshipsCreatedByUser `json:"created_by_user"`
	// The status page associated with the email access policy.
	StatusPage StatusPageEmailAccessPolicyDataRelationshipsStatusPage `json:"status_page"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStatusPageEmailAccessPolicyDataRelationships instantiates a new StatusPageEmailAccessPolicyDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStatusPageEmailAccessPolicyDataRelationships(createdByUser StatusPageEmailAccessPolicyDataRelationshipsCreatedByUser, statusPage StatusPageEmailAccessPolicyDataRelationshipsStatusPage) *StatusPageEmailAccessPolicyDataRelationships {
	this := StatusPageEmailAccessPolicyDataRelationships{}
	this.CreatedByUser = createdByUser
	this.StatusPage = statusPage
	return &this
}

// NewStatusPageEmailAccessPolicyDataRelationshipsWithDefaults instantiates a new StatusPageEmailAccessPolicyDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStatusPageEmailAccessPolicyDataRelationshipsWithDefaults() *StatusPageEmailAccessPolicyDataRelationships {
	this := StatusPageEmailAccessPolicyDataRelationships{}
	return &this
}

// GetCreatedByUser returns the CreatedByUser field value.
func (o *StatusPageEmailAccessPolicyDataRelationships) GetCreatedByUser() StatusPageEmailAccessPolicyDataRelationshipsCreatedByUser {
	if o == nil {
		var ret StatusPageEmailAccessPolicyDataRelationshipsCreatedByUser
		return ret
	}
	return o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value
// and a boolean to check if the value has been set.
func (o *StatusPageEmailAccessPolicyDataRelationships) GetCreatedByUserOk() (*StatusPageEmailAccessPolicyDataRelationshipsCreatedByUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByUser, true
}

// SetCreatedByUser sets field value.
func (o *StatusPageEmailAccessPolicyDataRelationships) SetCreatedByUser(v StatusPageEmailAccessPolicyDataRelationshipsCreatedByUser) {
	o.CreatedByUser = v
}

// GetStatusPage returns the StatusPage field value.
func (o *StatusPageEmailAccessPolicyDataRelationships) GetStatusPage() StatusPageEmailAccessPolicyDataRelationshipsStatusPage {
	if o == nil {
		var ret StatusPageEmailAccessPolicyDataRelationshipsStatusPage
		return ret
	}
	return o.StatusPage
}

// GetStatusPageOk returns a tuple with the StatusPage field value
// and a boolean to check if the value has been set.
func (o *StatusPageEmailAccessPolicyDataRelationships) GetStatusPageOk() (*StatusPageEmailAccessPolicyDataRelationshipsStatusPage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusPage, true
}

// SetStatusPage sets field value.
func (o *StatusPageEmailAccessPolicyDataRelationships) SetStatusPage(v StatusPageEmailAccessPolicyDataRelationshipsStatusPage) {
	o.StatusPage = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StatusPageEmailAccessPolicyDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["created_by_user"] = o.CreatedByUser
	toSerialize["status_page"] = o.StatusPage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StatusPageEmailAccessPolicyDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedByUser *StatusPageEmailAccessPolicyDataRelationshipsCreatedByUser `json:"created_by_user"`
		StatusPage    *StatusPageEmailAccessPolicyDataRelationshipsStatusPage    `json:"status_page"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedByUser == nil {
		return fmt.Errorf("required field created_by_user missing")
	}
	if all.StatusPage == nil {
		return fmt.Errorf("required field status_page missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by_user", "status_page"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CreatedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedByUser = *all.CreatedByUser
	if all.StatusPage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.StatusPage = *all.StatusPage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
