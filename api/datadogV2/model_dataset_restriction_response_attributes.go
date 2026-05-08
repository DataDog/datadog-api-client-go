// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetRestrictionResponseAttributes The current configuration of a dataset restriction, including restriction mode,
// ownership mode, and exempt principals.
type DatasetRestrictionResponseAttributes struct {
	// Controls how dataset ownership is determined. `disabled` turns off ownership-based access
	// entirely. `team_tag_based` assigns dataset ownership based on the team tags applied to the
	// data, allowing team members to see their own team's datasets.
	OwnershipMode *DatasetRestrictionOwnershipMode `json:"ownership_mode,omitempty"`
	// Internal key used by the restriction enforcement system to identify this restriction rule.
	RestrictionKey *string `json:"restriction_key,omitempty"`
	// Controls the default data visibility for the product type. `standard` makes data visible
	// to all users with appropriate product access. `default_hide` hides data by default and
	// requires explicit grants for each dataset.
	RestrictionMode DatasetRestrictionRestrictionMode `json:"restriction_mode"`
	// Principals (users or roles) that are exempt from this restriction and retain
	// full data access regardless of the restriction mode.
	UnrestrictedPrincipals []DatasetRestrictionPrincipal `json:"unrestricted_principals,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatasetRestrictionResponseAttributes instantiates a new DatasetRestrictionResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatasetRestrictionResponseAttributes(restrictionMode DatasetRestrictionRestrictionMode) *DatasetRestrictionResponseAttributes {
	this := DatasetRestrictionResponseAttributes{}
	this.RestrictionMode = restrictionMode
	return &this
}

// NewDatasetRestrictionResponseAttributesWithDefaults instantiates a new DatasetRestrictionResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatasetRestrictionResponseAttributesWithDefaults() *DatasetRestrictionResponseAttributes {
	this := DatasetRestrictionResponseAttributes{}
	return &this
}

// GetOwnershipMode returns the OwnershipMode field value if set, zero value otherwise.
func (o *DatasetRestrictionResponseAttributes) GetOwnershipMode() DatasetRestrictionOwnershipMode {
	if o == nil || o.OwnershipMode == nil {
		var ret DatasetRestrictionOwnershipMode
		return ret
	}
	return *o.OwnershipMode
}

// GetOwnershipModeOk returns a tuple with the OwnershipMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetRestrictionResponseAttributes) GetOwnershipModeOk() (*DatasetRestrictionOwnershipMode, bool) {
	if o == nil || o.OwnershipMode == nil {
		return nil, false
	}
	return o.OwnershipMode, true
}

// HasOwnershipMode returns a boolean if a field has been set.
func (o *DatasetRestrictionResponseAttributes) HasOwnershipMode() bool {
	return o != nil && o.OwnershipMode != nil
}

// SetOwnershipMode gets a reference to the given DatasetRestrictionOwnershipMode and assigns it to the OwnershipMode field.
func (o *DatasetRestrictionResponseAttributes) SetOwnershipMode(v DatasetRestrictionOwnershipMode) {
	o.OwnershipMode = &v
}

// GetRestrictionKey returns the RestrictionKey field value if set, zero value otherwise.
func (o *DatasetRestrictionResponseAttributes) GetRestrictionKey() string {
	if o == nil || o.RestrictionKey == nil {
		var ret string
		return ret
	}
	return *o.RestrictionKey
}

// GetRestrictionKeyOk returns a tuple with the RestrictionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetRestrictionResponseAttributes) GetRestrictionKeyOk() (*string, bool) {
	if o == nil || o.RestrictionKey == nil {
		return nil, false
	}
	return o.RestrictionKey, true
}

// HasRestrictionKey returns a boolean if a field has been set.
func (o *DatasetRestrictionResponseAttributes) HasRestrictionKey() bool {
	return o != nil && o.RestrictionKey != nil
}

// SetRestrictionKey gets a reference to the given string and assigns it to the RestrictionKey field.
func (o *DatasetRestrictionResponseAttributes) SetRestrictionKey(v string) {
	o.RestrictionKey = &v
}

// GetRestrictionMode returns the RestrictionMode field value.
func (o *DatasetRestrictionResponseAttributes) GetRestrictionMode() DatasetRestrictionRestrictionMode {
	if o == nil {
		var ret DatasetRestrictionRestrictionMode
		return ret
	}
	return o.RestrictionMode
}

// GetRestrictionModeOk returns a tuple with the RestrictionMode field value
// and a boolean to check if the value has been set.
func (o *DatasetRestrictionResponseAttributes) GetRestrictionModeOk() (*DatasetRestrictionRestrictionMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestrictionMode, true
}

// SetRestrictionMode sets field value.
func (o *DatasetRestrictionResponseAttributes) SetRestrictionMode(v DatasetRestrictionRestrictionMode) {
	o.RestrictionMode = v
}

// GetUnrestrictedPrincipals returns the UnrestrictedPrincipals field value if set, zero value otherwise.
func (o *DatasetRestrictionResponseAttributes) GetUnrestrictedPrincipals() []DatasetRestrictionPrincipal {
	if o == nil || o.UnrestrictedPrincipals == nil {
		var ret []DatasetRestrictionPrincipal
		return ret
	}
	return o.UnrestrictedPrincipals
}

// GetUnrestrictedPrincipalsOk returns a tuple with the UnrestrictedPrincipals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetRestrictionResponseAttributes) GetUnrestrictedPrincipalsOk() (*[]DatasetRestrictionPrincipal, bool) {
	if o == nil || o.UnrestrictedPrincipals == nil {
		return nil, false
	}
	return &o.UnrestrictedPrincipals, true
}

// HasUnrestrictedPrincipals returns a boolean if a field has been set.
func (o *DatasetRestrictionResponseAttributes) HasUnrestrictedPrincipals() bool {
	return o != nil && o.UnrestrictedPrincipals != nil
}

// SetUnrestrictedPrincipals gets a reference to the given []DatasetRestrictionPrincipal and assigns it to the UnrestrictedPrincipals field.
func (o *DatasetRestrictionResponseAttributes) SetUnrestrictedPrincipals(v []DatasetRestrictionPrincipal) {
	o.UnrestrictedPrincipals = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatasetRestrictionResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.OwnershipMode != nil {
		toSerialize["ownership_mode"] = o.OwnershipMode
	}
	if o.RestrictionKey != nil {
		toSerialize["restriction_key"] = o.RestrictionKey
	}
	toSerialize["restriction_mode"] = o.RestrictionMode
	if o.UnrestrictedPrincipals != nil {
		toSerialize["unrestricted_principals"] = o.UnrestrictedPrincipals
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatasetRestrictionResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OwnershipMode          *DatasetRestrictionOwnershipMode   `json:"ownership_mode,omitempty"`
		RestrictionKey         *string                            `json:"restriction_key,omitempty"`
		RestrictionMode        *DatasetRestrictionRestrictionMode `json:"restriction_mode"`
		UnrestrictedPrincipals []DatasetRestrictionPrincipal      `json:"unrestricted_principals,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RestrictionMode == nil {
		return fmt.Errorf("required field restriction_mode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ownership_mode", "restriction_key", "restriction_mode", "unrestricted_principals"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.OwnershipMode != nil && !all.OwnershipMode.IsValid() {
		hasInvalidField = true
	} else {
		o.OwnershipMode = all.OwnershipMode
	}
	o.RestrictionKey = all.RestrictionKey
	if !all.RestrictionMode.IsValid() {
		hasInvalidField = true
	} else {
		o.RestrictionMode = *all.RestrictionMode
	}
	o.UnrestrictedPrincipals = all.UnrestrictedPrincipals

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
