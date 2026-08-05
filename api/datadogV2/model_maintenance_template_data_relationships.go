// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MaintenanceTemplateDataRelationships The relationships of a maintenance template.
type MaintenanceTemplateDataRelationships struct {
	// The Datadog user who created the maintenance template.
	CreatedByUser *MaintenanceTemplateDataRelationshipsCreatedByUser `json:"created_by_user,omitempty"`
	// The Datadog user who last modified the maintenance template.
	LastModifiedByUser *MaintenanceTemplateDataRelationshipsLastModifiedByUser `json:"last_modified_by_user,omitempty"`
	// The status page the maintenance template belongs to.
	StatusPage *MaintenanceTemplateDataRelationshipsStatusPage `json:"status_page,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMaintenanceTemplateDataRelationships instantiates a new MaintenanceTemplateDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMaintenanceTemplateDataRelationships() *MaintenanceTemplateDataRelationships {
	this := MaintenanceTemplateDataRelationships{}
	return &this
}

// NewMaintenanceTemplateDataRelationshipsWithDefaults instantiates a new MaintenanceTemplateDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMaintenanceTemplateDataRelationshipsWithDefaults() *MaintenanceTemplateDataRelationships {
	this := MaintenanceTemplateDataRelationships{}
	return &this
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise.
func (o *MaintenanceTemplateDataRelationships) GetCreatedByUser() MaintenanceTemplateDataRelationshipsCreatedByUser {
	if o == nil || o.CreatedByUser == nil {
		var ret MaintenanceTemplateDataRelationshipsCreatedByUser
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceTemplateDataRelationships) GetCreatedByUserOk() (*MaintenanceTemplateDataRelationshipsCreatedByUser, bool) {
	if o == nil || o.CreatedByUser == nil {
		return nil, false
	}
	return o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *MaintenanceTemplateDataRelationships) HasCreatedByUser() bool {
	return o != nil && o.CreatedByUser != nil
}

// SetCreatedByUser gets a reference to the given MaintenanceTemplateDataRelationshipsCreatedByUser and assigns it to the CreatedByUser field.
func (o *MaintenanceTemplateDataRelationships) SetCreatedByUser(v MaintenanceTemplateDataRelationshipsCreatedByUser) {
	o.CreatedByUser = &v
}

// GetLastModifiedByUser returns the LastModifiedByUser field value if set, zero value otherwise.
func (o *MaintenanceTemplateDataRelationships) GetLastModifiedByUser() MaintenanceTemplateDataRelationshipsLastModifiedByUser {
	if o == nil || o.LastModifiedByUser == nil {
		var ret MaintenanceTemplateDataRelationshipsLastModifiedByUser
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceTemplateDataRelationships) GetLastModifiedByUserOk() (*MaintenanceTemplateDataRelationshipsLastModifiedByUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		return nil, false
	}
	return o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *MaintenanceTemplateDataRelationships) HasLastModifiedByUser() bool {
	return o != nil && o.LastModifiedByUser != nil
}

// SetLastModifiedByUser gets a reference to the given MaintenanceTemplateDataRelationshipsLastModifiedByUser and assigns it to the LastModifiedByUser field.
func (o *MaintenanceTemplateDataRelationships) SetLastModifiedByUser(v MaintenanceTemplateDataRelationshipsLastModifiedByUser) {
	o.LastModifiedByUser = &v
}

// GetStatusPage returns the StatusPage field value if set, zero value otherwise.
func (o *MaintenanceTemplateDataRelationships) GetStatusPage() MaintenanceTemplateDataRelationshipsStatusPage {
	if o == nil || o.StatusPage == nil {
		var ret MaintenanceTemplateDataRelationshipsStatusPage
		return ret
	}
	return *o.StatusPage
}

// GetStatusPageOk returns a tuple with the StatusPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceTemplateDataRelationships) GetStatusPageOk() (*MaintenanceTemplateDataRelationshipsStatusPage, bool) {
	if o == nil || o.StatusPage == nil {
		return nil, false
	}
	return o.StatusPage, true
}

// HasStatusPage returns a boolean if a field has been set.
func (o *MaintenanceTemplateDataRelationships) HasStatusPage() bool {
	return o != nil && o.StatusPage != nil
}

// SetStatusPage gets a reference to the given MaintenanceTemplateDataRelationshipsStatusPage and assigns it to the StatusPage field.
func (o *MaintenanceTemplateDataRelationships) SetStatusPage(v MaintenanceTemplateDataRelationshipsStatusPage) {
	o.StatusPage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MaintenanceTemplateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedByUser != nil {
		toSerialize["created_by_user"] = o.CreatedByUser
	}
	if o.LastModifiedByUser != nil {
		toSerialize["last_modified_by_user"] = o.LastModifiedByUser
	}
	if o.StatusPage != nil {
		toSerialize["status_page"] = o.StatusPage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MaintenanceTemplateDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedByUser      *MaintenanceTemplateDataRelationshipsCreatedByUser      `json:"created_by_user,omitempty"`
		LastModifiedByUser *MaintenanceTemplateDataRelationshipsLastModifiedByUser `json:"last_modified_by_user,omitempty"`
		StatusPage         *MaintenanceTemplateDataRelationshipsStatusPage         `json:"status_page,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by_user", "last_modified_by_user", "status_page"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CreatedByUser != nil && all.CreatedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedByUser = all.CreatedByUser
	if all.LastModifiedByUser != nil && all.LastModifiedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LastModifiedByUser = all.LastModifiedByUser
	if all.StatusPage != nil && all.StatusPage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.StatusPage = all.StatusPage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
