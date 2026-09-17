// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamsOwnershipMappingBatchError An error encountered while validating or applying an operation.
type TeamsOwnershipMappingBatchError struct {
	// A human-readable explanation specific to this error.
	Detail *string `json:"detail,omitempty"`
	// The HTTP status code applicable to this error.
	Status string `json:"status"`
	// A short, human-readable summary of the error.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamsOwnershipMappingBatchError instantiates a new TeamsOwnershipMappingBatchError object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamsOwnershipMappingBatchError(status string, title string) *TeamsOwnershipMappingBatchError {
	this := TeamsOwnershipMappingBatchError{}
	this.Status = status
	this.Title = title
	return &this
}

// NewTeamsOwnershipMappingBatchErrorWithDefaults instantiates a new TeamsOwnershipMappingBatchError object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamsOwnershipMappingBatchErrorWithDefaults() *TeamsOwnershipMappingBatchError {
	this := TeamsOwnershipMappingBatchError{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *TeamsOwnershipMappingBatchError) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingBatchError) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *TeamsOwnershipMappingBatchError) HasDetail() bool {
	return o != nil && o.Detail != nil
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *TeamsOwnershipMappingBatchError) SetDetail(v string) {
	o.Detail = &v
}

// GetStatus returns the Status field value.
func (o *TeamsOwnershipMappingBatchError) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingBatchError) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *TeamsOwnershipMappingBatchError) SetStatus(v string) {
	o.Status = v
}

// GetTitle returns the Title field value.
func (o *TeamsOwnershipMappingBatchError) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingBatchError) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *TeamsOwnershipMappingBatchError) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamsOwnershipMappingBatchError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	toSerialize["status"] = o.Status
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamsOwnershipMappingBatchError) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Detail *string `json:"detail,omitempty"`
		Status *string `json:"status"`
		Title  *string `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"detail", "status", "title"})
	} else {
		return err
	}
	o.Detail = all.Detail
	o.Status = *all.Status
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
