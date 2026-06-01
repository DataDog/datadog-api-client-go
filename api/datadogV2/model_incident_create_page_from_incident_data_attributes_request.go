// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentCreatePageFromIncidentDataAttributesRequest Attributes for creating a page from an incident.
type IncidentCreatePageFromIncidentDataAttributesRequest struct {
	// The description of the page.
	Description *string `json:"description,omitempty"`
	// The public ID of the incident.
	IncidentPublicId *string `json:"incident_public_id,omitempty"`
	// A reference to an incident role for a page.
	Role *IncidentPageRoleReference `json:"role,omitempty"`
	// List of affected services.
	Services []string `json:"services,omitempty"`
	// List of tags for the page.
	Tags []string `json:"tags,omitempty"`
	// The target recipient for a page.
	Target *IncidentPageTarget `json:"target,omitempty"`
	// The title of the page.
	Title *string `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentCreatePageFromIncidentDataAttributesRequest instantiates a new IncidentCreatePageFromIncidentDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentCreatePageFromIncidentDataAttributesRequest() *IncidentCreatePageFromIncidentDataAttributesRequest {
	this := IncidentCreatePageFromIncidentDataAttributesRequest{}
	return &this
}

// NewIncidentCreatePageFromIncidentDataAttributesRequestWithDefaults instantiates a new IncidentCreatePageFromIncidentDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentCreatePageFromIncidentDataAttributesRequestWithDefaults() *IncidentCreatePageFromIncidentDataAttributesRequest {
	this := IncidentCreatePageFromIncidentDataAttributesRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIncidentPublicId returns the IncidentPublicId field value if set, zero value otherwise.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) GetIncidentPublicId() string {
	if o == nil || o.IncidentPublicId == nil {
		var ret string
		return ret
	}
	return *o.IncidentPublicId
}

// GetIncidentPublicIdOk returns a tuple with the IncidentPublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) GetIncidentPublicIdOk() (*string, bool) {
	if o == nil || o.IncidentPublicId == nil {
		return nil, false
	}
	return o.IncidentPublicId, true
}

// HasIncidentPublicId returns a boolean if a field has been set.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) HasIncidentPublicId() bool {
	return o != nil && o.IncidentPublicId != nil
}

// SetIncidentPublicId gets a reference to the given string and assigns it to the IncidentPublicId field.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) SetIncidentPublicId(v string) {
	o.IncidentPublicId = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) GetRole() IncidentPageRoleReference {
	if o == nil || o.Role == nil {
		var ret IncidentPageRoleReference
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) GetRoleOk() (*IncidentPageRoleReference, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given IncidentPageRoleReference and assigns it to the Role field.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) SetRole(v IncidentPageRoleReference) {
	o.Role = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) GetServicesOk() (*[]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return &o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) HasServices() bool {
	return o != nil && o.Services != nil
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) SetServices(v []string) {
	o.Services = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) SetTags(v []string) {
	o.Tags = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) GetTarget() IncidentPageTarget {
	if o == nil || o.Target == nil {
		var ret IncidentPageTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) GetTargetOk() (*IncidentPageTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given IncidentPageTarget and assigns it to the Target field.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) SetTarget(v IncidentPageTarget) {
	o.Target = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) SetTitle(v string) {
	o.Title = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentCreatePageFromIncidentDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.IncidentPublicId != nil {
		toSerialize["incident_public_id"] = o.IncidentPublicId
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
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
func (o *IncidentCreatePageFromIncidentDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description      *string                    `json:"description,omitempty"`
		IncidentPublicId *string                    `json:"incident_public_id,omitempty"`
		Role             *IncidentPageRoleReference `json:"role,omitempty"`
		Services         []string                   `json:"services,omitempty"`
		Tags             []string                   `json:"tags,omitempty"`
		Target           *IncidentPageTarget        `json:"target,omitempty"`
		Title            *string                    `json:"title,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "incident_public_id", "role", "services", "tags", "target", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.IncidentPublicId = all.IncidentPublicId
	if all.Role != nil && all.Role.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Role = all.Role
	o.Services = all.Services
	o.Tags = all.Tags
	if all.Target != nil && all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = all.Target
	o.Title = all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
