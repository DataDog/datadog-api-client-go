// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentStatuspageIncidentDataAttributesRequest Attributes for creating or updating a Statuspage incident.
type IncidentStatuspageIncidentDataAttributesRequest struct {
	// The body text of the Statuspage incident.
	Body datadog.NullableString `json:"body,omitempty"`
	// Map of component identifiers to their status.
	Components map[string]string `json:"components,omitempty"`
	// Whether to deliver notifications.
	DeliverNotifications datadog.NullableBool `json:"deliver_notifications,omitempty"`
	// The impact level of the incident.
	Impact datadog.NullableString `json:"impact,omitempty"`
	// The name of the Statuspage incident.
	Name datadog.NullableString `json:"name,omitempty"`
	// The Statuspage page identifier.
	PageId *string `json:"page_id,omitempty"`
	// The status of the Statuspage incident.
	Status datadog.NullableString `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentStatuspageIncidentDataAttributesRequest instantiates a new IncidentStatuspageIncidentDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentStatuspageIncidentDataAttributesRequest() *IncidentStatuspageIncidentDataAttributesRequest {
	this := IncidentStatuspageIncidentDataAttributesRequest{}
	return &this
}

// NewIncidentStatuspageIncidentDataAttributesRequestWithDefaults instantiates a new IncidentStatuspageIncidentDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentStatuspageIncidentDataAttributesRequestWithDefaults() *IncidentStatuspageIncidentDataAttributesRequest {
	this := IncidentStatuspageIncidentDataAttributesRequest{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentStatuspageIncidentDataAttributesRequest) GetBody() string {
	if o == nil || o.Body.Get() == nil {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentStatuspageIncidentDataAttributesRequest) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *IncidentStatuspageIncidentDataAttributesRequest) HasBody() bool {
	return o != nil && o.Body.IsSet()
}

// SetBody gets a reference to the given datadog.NullableString and assigns it to the Body field.
func (o *IncidentStatuspageIncidentDataAttributesRequest) SetBody(v string) {
	o.Body.Set(&v)
}

// SetBodyNil sets the value for Body to be an explicit nil.
func (o *IncidentStatuspageIncidentDataAttributesRequest) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil.
func (o *IncidentStatuspageIncidentDataAttributesRequest) UnsetBody() {
	o.Body.Unset()
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *IncidentStatuspageIncidentDataAttributesRequest) GetComponents() map[string]string {
	if o == nil || o.Components == nil {
		var ret map[string]string
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentStatuspageIncidentDataAttributesRequest) GetComponentsOk() (*map[string]string, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *IncidentStatuspageIncidentDataAttributesRequest) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given map[string]string and assigns it to the Components field.
func (o *IncidentStatuspageIncidentDataAttributesRequest) SetComponents(v map[string]string) {
	o.Components = v
}

// GetDeliverNotifications returns the DeliverNotifications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentStatuspageIncidentDataAttributesRequest) GetDeliverNotifications() bool {
	if o == nil || o.DeliverNotifications.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DeliverNotifications.Get()
}

// GetDeliverNotificationsOk returns a tuple with the DeliverNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentStatuspageIncidentDataAttributesRequest) GetDeliverNotificationsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliverNotifications.Get(), o.DeliverNotifications.IsSet()
}

// HasDeliverNotifications returns a boolean if a field has been set.
func (o *IncidentStatuspageIncidentDataAttributesRequest) HasDeliverNotifications() bool {
	return o != nil && o.DeliverNotifications.IsSet()
}

// SetDeliverNotifications gets a reference to the given datadog.NullableBool and assigns it to the DeliverNotifications field.
func (o *IncidentStatuspageIncidentDataAttributesRequest) SetDeliverNotifications(v bool) {
	o.DeliverNotifications.Set(&v)
}

// SetDeliverNotificationsNil sets the value for DeliverNotifications to be an explicit nil.
func (o *IncidentStatuspageIncidentDataAttributesRequest) SetDeliverNotificationsNil() {
	o.DeliverNotifications.Set(nil)
}

// UnsetDeliverNotifications ensures that no value is present for DeliverNotifications, not even an explicit nil.
func (o *IncidentStatuspageIncidentDataAttributesRequest) UnsetDeliverNotifications() {
	o.DeliverNotifications.Unset()
}

// GetImpact returns the Impact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentStatuspageIncidentDataAttributesRequest) GetImpact() string {
	if o == nil || o.Impact.Get() == nil {
		var ret string
		return ret
	}
	return *o.Impact.Get()
}

// GetImpactOk returns a tuple with the Impact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentStatuspageIncidentDataAttributesRequest) GetImpactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Impact.Get(), o.Impact.IsSet()
}

// HasImpact returns a boolean if a field has been set.
func (o *IncidentStatuspageIncidentDataAttributesRequest) HasImpact() bool {
	return o != nil && o.Impact.IsSet()
}

// SetImpact gets a reference to the given datadog.NullableString and assigns it to the Impact field.
func (o *IncidentStatuspageIncidentDataAttributesRequest) SetImpact(v string) {
	o.Impact.Set(&v)
}

// SetImpactNil sets the value for Impact to be an explicit nil.
func (o *IncidentStatuspageIncidentDataAttributesRequest) SetImpactNil() {
	o.Impact.Set(nil)
}

// UnsetImpact ensures that no value is present for Impact, not even an explicit nil.
func (o *IncidentStatuspageIncidentDataAttributesRequest) UnsetImpact() {
	o.Impact.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentStatuspageIncidentDataAttributesRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentStatuspageIncidentDataAttributesRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IncidentStatuspageIncidentDataAttributesRequest) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given datadog.NullableString and assigns it to the Name field.
func (o *IncidentStatuspageIncidentDataAttributesRequest) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *IncidentStatuspageIncidentDataAttributesRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *IncidentStatuspageIncidentDataAttributesRequest) UnsetName() {
	o.Name.Unset()
}

// GetPageId returns the PageId field value if set, zero value otherwise.
func (o *IncidentStatuspageIncidentDataAttributesRequest) GetPageId() string {
	if o == nil || o.PageId == nil {
		var ret string
		return ret
	}
	return *o.PageId
}

// GetPageIdOk returns a tuple with the PageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentStatuspageIncidentDataAttributesRequest) GetPageIdOk() (*string, bool) {
	if o == nil || o.PageId == nil {
		return nil, false
	}
	return o.PageId, true
}

// HasPageId returns a boolean if a field has been set.
func (o *IncidentStatuspageIncidentDataAttributesRequest) HasPageId() bool {
	return o != nil && o.PageId != nil
}

// SetPageId gets a reference to the given string and assigns it to the PageId field.
func (o *IncidentStatuspageIncidentDataAttributesRequest) SetPageId(v string) {
	o.PageId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentStatuspageIncidentDataAttributesRequest) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentStatuspageIncidentDataAttributesRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *IncidentStatuspageIncidentDataAttributesRequest) HasStatus() bool {
	return o != nil && o.Status.IsSet()
}

// SetStatus gets a reference to the given datadog.NullableString and assigns it to the Status field.
func (o *IncidentStatuspageIncidentDataAttributesRequest) SetStatus(v string) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil.
func (o *IncidentStatuspageIncidentDataAttributesRequest) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil.
func (o *IncidentStatuspageIncidentDataAttributesRequest) UnsetStatus() {
	o.Status.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentStatuspageIncidentDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Body.IsSet() {
		toSerialize["body"] = o.Body.Get()
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.DeliverNotifications.IsSet() {
		toSerialize["deliver_notifications"] = o.DeliverNotifications.Get()
	}
	if o.Impact.IsSet() {
		toSerialize["impact"] = o.Impact.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.PageId != nil {
		toSerialize["page_id"] = o.PageId
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentStatuspageIncidentDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Body                 datadog.NullableString `json:"body,omitempty"`
		Components           map[string]string      `json:"components,omitempty"`
		DeliverNotifications datadog.NullableBool   `json:"deliver_notifications,omitempty"`
		Impact               datadog.NullableString `json:"impact,omitempty"`
		Name                 datadog.NullableString `json:"name,omitempty"`
		PageId               *string                `json:"page_id,omitempty"`
		Status               datadog.NullableString `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"body", "components", "deliver_notifications", "impact", "name", "page_id", "status"})
	} else {
		return err
	}
	o.Body = all.Body
	o.Components = all.Components
	o.DeliverNotifications = all.DeliverNotifications
	o.Impact = all.Impact
	o.Name = all.Name
	o.PageId = all.PageId
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
