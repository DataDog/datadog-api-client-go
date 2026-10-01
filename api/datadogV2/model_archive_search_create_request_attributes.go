// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArchiveSearchCreateRequestAttributes Attributes accepted when creating an Archive Search.
type ArchiveSearchCreateRequestAttributes struct {
	// ID of the archive to search. Use the Logs Archives API to list the archives of the organization.
	ArchiveId string `json:"archive_id"`
	// Free-text description of the Archive Search.
	Description *string `json:"description,omitempty"`
	// Start of the time range to search, as an ISO 8601 timestamp.
	From time.Time `json:"from"`
	// Name of the Archive Search.
	Name string `json:"name"`
	// Log search query used to filter the archived logs.
	Query string `json:"query"`
	// Rehydration settings. Include this object to index the matched logs into a retained historical view.
	// Omit it to run an Archive Search that only scans the archive.
	Rehydration *ArchiveSearchCreateRehydration `json:"rehydration,omitempty"`
	// End of the time range to search, as an ISO 8601 timestamp. Must be after `from`.
	To time.Time `json:"to"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewArchiveSearchCreateRequestAttributes instantiates a new ArchiveSearchCreateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewArchiveSearchCreateRequestAttributes(archiveId string, from time.Time, name string, query string, to time.Time) *ArchiveSearchCreateRequestAttributes {
	this := ArchiveSearchCreateRequestAttributes{}
	this.ArchiveId = archiveId
	this.From = from
	this.Name = name
	this.Query = query
	this.To = to
	return &this
}

// NewArchiveSearchCreateRequestAttributesWithDefaults instantiates a new ArchiveSearchCreateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewArchiveSearchCreateRequestAttributesWithDefaults() *ArchiveSearchCreateRequestAttributes {
	this := ArchiveSearchCreateRequestAttributes{}
	return &this
}

// GetArchiveId returns the ArchiveId field value.
func (o *ArchiveSearchCreateRequestAttributes) GetArchiveId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ArchiveId
}

// GetArchiveIdOk returns a tuple with the ArchiveId field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchCreateRequestAttributes) GetArchiveIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArchiveId, true
}

// SetArchiveId sets field value.
func (o *ArchiveSearchCreateRequestAttributes) SetArchiveId(v string) {
	o.ArchiveId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ArchiveSearchCreateRequestAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveSearchCreateRequestAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ArchiveSearchCreateRequestAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ArchiveSearchCreateRequestAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetFrom returns the From field value.
func (o *ArchiveSearchCreateRequestAttributes) GetFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchCreateRequestAttributes) GetFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *ArchiveSearchCreateRequestAttributes) SetFrom(v time.Time) {
	o.From = v
}

// GetName returns the Name field value.
func (o *ArchiveSearchCreateRequestAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchCreateRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ArchiveSearchCreateRequestAttributes) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value.
func (o *ArchiveSearchCreateRequestAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchCreateRequestAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ArchiveSearchCreateRequestAttributes) SetQuery(v string) {
	o.Query = v
}

// GetRehydration returns the Rehydration field value if set, zero value otherwise.
func (o *ArchiveSearchCreateRequestAttributes) GetRehydration() ArchiveSearchCreateRehydration {
	if o == nil || o.Rehydration == nil {
		var ret ArchiveSearchCreateRehydration
		return ret
	}
	return *o.Rehydration
}

// GetRehydrationOk returns a tuple with the Rehydration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveSearchCreateRequestAttributes) GetRehydrationOk() (*ArchiveSearchCreateRehydration, bool) {
	if o == nil || o.Rehydration == nil {
		return nil, false
	}
	return o.Rehydration, true
}

// HasRehydration returns a boolean if a field has been set.
func (o *ArchiveSearchCreateRequestAttributes) HasRehydration() bool {
	return o != nil && o.Rehydration != nil
}

// SetRehydration gets a reference to the given ArchiveSearchCreateRehydration and assigns it to the Rehydration field.
func (o *ArchiveSearchCreateRequestAttributes) SetRehydration(v ArchiveSearchCreateRehydration) {
	o.Rehydration = &v
}

// GetTo returns the To field value.
func (o *ArchiveSearchCreateRequestAttributes) GetTo() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchCreateRequestAttributes) GetToOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *ArchiveSearchCreateRequestAttributes) SetTo(v time.Time) {
	o.To = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ArchiveSearchCreateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["archive_id"] = o.ArchiveId
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.From.Nanosecond() == 0 {
		toSerialize["from"] = o.From.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["from"] = o.From.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query
	if o.Rehydration != nil {
		toSerialize["rehydration"] = o.Rehydration
	}
	if o.To.Nanosecond() == 0 {
		toSerialize["to"] = o.To.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["to"] = o.To.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ArchiveSearchCreateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ArchiveId   *string                         `json:"archive_id"`
		Description *string                         `json:"description,omitempty"`
		From        *time.Time                      `json:"from"`
		Name        *string                         `json:"name"`
		Query       *string                         `json:"query"`
		Rehydration *ArchiveSearchCreateRehydration `json:"rehydration,omitempty"`
		To          *time.Time                      `json:"to"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ArchiveId == nil {
		return fmt.Errorf("required field archive_id missing")
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"archive_id", "description", "from", "name", "query", "rehydration", "to"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ArchiveId = *all.ArchiveId
	o.Description = all.Description
	o.From = *all.From
	o.Name = *all.Name
	o.Query = *all.Query
	if all.Rehydration != nil && all.Rehydration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rehydration = all.Rehydration
	o.To = *all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
