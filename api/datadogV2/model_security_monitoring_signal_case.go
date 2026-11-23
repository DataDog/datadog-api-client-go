// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalCase A case linked to a security signal.
type SecurityMonitoringSignalCase struct {
	// The UUID of the user assigned to the case.
	AssigneeId *string `json:"assigneeId,omitempty"`
	// The name of the user assigned to the case.
	AssigneeName *string `json:"assigneeName,omitempty"`
	// The unique ID of the case.
	Id *string `json:"id,omitempty"`
	// The key of the case.
	Key *string `json:"key,omitempty"`
	// The status of the case.
	Status *string `json:"status,omitempty"`
	// The title of the case.
	Title *string `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSignalCase instantiates a new SecurityMonitoringSignalCase object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSignalCase() *SecurityMonitoringSignalCase {
	this := SecurityMonitoringSignalCase{}
	return &this
}

// NewSecurityMonitoringSignalCaseWithDefaults instantiates a new SecurityMonitoringSignalCase object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSignalCaseWithDefaults() *SecurityMonitoringSignalCase {
	this := SecurityMonitoringSignalCase{}
	return &this
}

// GetAssigneeId returns the AssigneeId field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalCase) GetAssigneeId() string {
	if o == nil || o.AssigneeId == nil {
		var ret string
		return ret
	}
	return *o.AssigneeId
}

// GetAssigneeIdOk returns a tuple with the AssigneeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalCase) GetAssigneeIdOk() (*string, bool) {
	if o == nil || o.AssigneeId == nil {
		return nil, false
	}
	return o.AssigneeId, true
}

// HasAssigneeId returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalCase) HasAssigneeId() bool {
	return o != nil && o.AssigneeId != nil
}

// SetAssigneeId gets a reference to the given string and assigns it to the AssigneeId field.
func (o *SecurityMonitoringSignalCase) SetAssigneeId(v string) {
	o.AssigneeId = &v
}

// GetAssigneeName returns the AssigneeName field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalCase) GetAssigneeName() string {
	if o == nil || o.AssigneeName == nil {
		var ret string
		return ret
	}
	return *o.AssigneeName
}

// GetAssigneeNameOk returns a tuple with the AssigneeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalCase) GetAssigneeNameOk() (*string, bool) {
	if o == nil || o.AssigneeName == nil {
		return nil, false
	}
	return o.AssigneeName, true
}

// HasAssigneeName returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalCase) HasAssigneeName() bool {
	return o != nil && o.AssigneeName != nil
}

// SetAssigneeName gets a reference to the given string and assigns it to the AssigneeName field.
func (o *SecurityMonitoringSignalCase) SetAssigneeName(v string) {
	o.AssigneeName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalCase) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalCase) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalCase) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SecurityMonitoringSignalCase) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalCase) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalCase) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalCase) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SecurityMonitoringSignalCase) SetKey(v string) {
	o.Key = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalCase) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalCase) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalCase) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SecurityMonitoringSignalCase) SetStatus(v string) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalCase) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalCase) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalCase) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SecurityMonitoringSignalCase) SetTitle(v string) {
	o.Title = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSignalCase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssigneeId != nil {
		toSerialize["assigneeId"] = o.AssigneeId
	}
	if o.AssigneeName != nil {
		toSerialize["assigneeName"] = o.AssigneeName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSignalCase) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssigneeId   *string `json:"assigneeId,omitempty"`
		AssigneeName *string `json:"assigneeName,omitempty"`
		Id           *string `json:"id,omitempty"`
		Key          *string `json:"key,omitempty"`
		Status       *string `json:"status,omitempty"`
		Title        *string `json:"title,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assigneeId", "assigneeName", "id", "key", "status", "title"})
	} else {
		return err
	}
	o.AssigneeId = all.AssigneeId
	o.AssigneeName = all.AssigneeName
	o.Id = all.Id
	o.Key = all.Key
	o.Status = all.Status
	o.Title = all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
