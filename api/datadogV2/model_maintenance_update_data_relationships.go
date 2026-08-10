// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MaintenanceUpdateDataRelationships Relationships of a maintenance update resource.
type MaintenanceUpdateDataRelationships struct {
	// A user relationship of a maintenance update.
	CreatedByUser *MaintenanceUpdateDataRelationshipsUser `json:"created_by_user,omitempty"`
	// A user relationship of a maintenance update.
	LastModifiedByUser *MaintenanceUpdateDataRelationshipsUser `json:"last_modified_by_user,omitempty"`
	// The parent maintenance of the update.
	Maintenance *MaintenanceUpdateDataRelationshipsMaintenance `json:"maintenance,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMaintenanceUpdateDataRelationships instantiates a new MaintenanceUpdateDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMaintenanceUpdateDataRelationships() *MaintenanceUpdateDataRelationships {
	this := MaintenanceUpdateDataRelationships{}
	return &this
}

// NewMaintenanceUpdateDataRelationshipsWithDefaults instantiates a new MaintenanceUpdateDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMaintenanceUpdateDataRelationshipsWithDefaults() *MaintenanceUpdateDataRelationships {
	this := MaintenanceUpdateDataRelationships{}
	return &this
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise.
func (o *MaintenanceUpdateDataRelationships) GetCreatedByUser() MaintenanceUpdateDataRelationshipsUser {
	if o == nil || o.CreatedByUser == nil {
		var ret MaintenanceUpdateDataRelationshipsUser
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceUpdateDataRelationships) GetCreatedByUserOk() (*MaintenanceUpdateDataRelationshipsUser, bool) {
	if o == nil || o.CreatedByUser == nil {
		return nil, false
	}
	return o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *MaintenanceUpdateDataRelationships) HasCreatedByUser() bool {
	return o != nil && o.CreatedByUser != nil
}

// SetCreatedByUser gets a reference to the given MaintenanceUpdateDataRelationshipsUser and assigns it to the CreatedByUser field.
func (o *MaintenanceUpdateDataRelationships) SetCreatedByUser(v MaintenanceUpdateDataRelationshipsUser) {
	o.CreatedByUser = &v
}

// GetLastModifiedByUser returns the LastModifiedByUser field value if set, zero value otherwise.
func (o *MaintenanceUpdateDataRelationships) GetLastModifiedByUser() MaintenanceUpdateDataRelationshipsUser {
	if o == nil || o.LastModifiedByUser == nil {
		var ret MaintenanceUpdateDataRelationshipsUser
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceUpdateDataRelationships) GetLastModifiedByUserOk() (*MaintenanceUpdateDataRelationshipsUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		return nil, false
	}
	return o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *MaintenanceUpdateDataRelationships) HasLastModifiedByUser() bool {
	return o != nil && o.LastModifiedByUser != nil
}

// SetLastModifiedByUser gets a reference to the given MaintenanceUpdateDataRelationshipsUser and assigns it to the LastModifiedByUser field.
func (o *MaintenanceUpdateDataRelationships) SetLastModifiedByUser(v MaintenanceUpdateDataRelationshipsUser) {
	o.LastModifiedByUser = &v
}

// GetMaintenance returns the Maintenance field value if set, zero value otherwise.
func (o *MaintenanceUpdateDataRelationships) GetMaintenance() MaintenanceUpdateDataRelationshipsMaintenance {
	if o == nil || o.Maintenance == nil {
		var ret MaintenanceUpdateDataRelationshipsMaintenance
		return ret
	}
	return *o.Maintenance
}

// GetMaintenanceOk returns a tuple with the Maintenance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceUpdateDataRelationships) GetMaintenanceOk() (*MaintenanceUpdateDataRelationshipsMaintenance, bool) {
	if o == nil || o.Maintenance == nil {
		return nil, false
	}
	return o.Maintenance, true
}

// HasMaintenance returns a boolean if a field has been set.
func (o *MaintenanceUpdateDataRelationships) HasMaintenance() bool {
	return o != nil && o.Maintenance != nil
}

// SetMaintenance gets a reference to the given MaintenanceUpdateDataRelationshipsMaintenance and assigns it to the Maintenance field.
func (o *MaintenanceUpdateDataRelationships) SetMaintenance(v MaintenanceUpdateDataRelationshipsMaintenance) {
	o.Maintenance = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MaintenanceUpdateDataRelationships) MarshalJSON() ([]byte, error) {
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
	if o.Maintenance != nil {
		toSerialize["maintenance"] = o.Maintenance
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MaintenanceUpdateDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedByUser      *MaintenanceUpdateDataRelationshipsUser        `json:"created_by_user,omitempty"`
		LastModifiedByUser *MaintenanceUpdateDataRelationshipsUser        `json:"last_modified_by_user,omitempty"`
		Maintenance        *MaintenanceUpdateDataRelationshipsMaintenance `json:"maintenance,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by_user", "last_modified_by_user", "maintenance"})
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
	if all.Maintenance != nil && all.Maintenance.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Maintenance = all.Maintenance

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
