// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentOnCallPageDataAttributesRequest Attributes for linking a page to an incident.
type IncidentOnCallPageDataAttributesRequest struct {
	// The key of the on-call page.
	Key *string `json:"key,omitempty"`
	// The target of an on-call page.
	PageTarget *IncidentOnCallPageTarget `json:"page_target,omitempty"`
	// The team ID associated with the page (deprecated, use page_target instead).
	TeamId *string `json:"team_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentOnCallPageDataAttributesRequest instantiates a new IncidentOnCallPageDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentOnCallPageDataAttributesRequest() *IncidentOnCallPageDataAttributesRequest {
	this := IncidentOnCallPageDataAttributesRequest{}
	return &this
}

// NewIncidentOnCallPageDataAttributesRequestWithDefaults instantiates a new IncidentOnCallPageDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentOnCallPageDataAttributesRequestWithDefaults() *IncidentOnCallPageDataAttributesRequest {
	this := IncidentOnCallPageDataAttributesRequest{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *IncidentOnCallPageDataAttributesRequest) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentOnCallPageDataAttributesRequest) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *IncidentOnCallPageDataAttributesRequest) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *IncidentOnCallPageDataAttributesRequest) SetKey(v string) {
	o.Key = &v
}

// GetPageTarget returns the PageTarget field value if set, zero value otherwise.
func (o *IncidentOnCallPageDataAttributesRequest) GetPageTarget() IncidentOnCallPageTarget {
	if o == nil || o.PageTarget == nil {
		var ret IncidentOnCallPageTarget
		return ret
	}
	return *o.PageTarget
}

// GetPageTargetOk returns a tuple with the PageTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentOnCallPageDataAttributesRequest) GetPageTargetOk() (*IncidentOnCallPageTarget, bool) {
	if o == nil || o.PageTarget == nil {
		return nil, false
	}
	return o.PageTarget, true
}

// HasPageTarget returns a boolean if a field has been set.
func (o *IncidentOnCallPageDataAttributesRequest) HasPageTarget() bool {
	return o != nil && o.PageTarget != nil
}

// SetPageTarget gets a reference to the given IncidentOnCallPageTarget and assigns it to the PageTarget field.
func (o *IncidentOnCallPageDataAttributesRequest) SetPageTarget(v IncidentOnCallPageTarget) {
	o.PageTarget = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *IncidentOnCallPageDataAttributesRequest) GetTeamId() string {
	if o == nil || o.TeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentOnCallPageDataAttributesRequest) GetTeamIdOk() (*string, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *IncidentOnCallPageDataAttributesRequest) HasTeamId() bool {
	return o != nil && o.TeamId != nil
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *IncidentOnCallPageDataAttributesRequest) SetTeamId(v string) {
	o.TeamId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentOnCallPageDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.PageTarget != nil {
		toSerialize["page_target"] = o.PageTarget
	}
	if o.TeamId != nil {
		toSerialize["team_id"] = o.TeamId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentOnCallPageDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key        *string                   `json:"key,omitempty"`
		PageTarget *IncidentOnCallPageTarget `json:"page_target,omitempty"`
		TeamId     *string                   `json:"team_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"key", "page_target", "team_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Key = all.Key
	if all.PageTarget != nil && all.PageTarget.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PageTarget = all.PageTarget
	o.TeamId = all.TeamId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
