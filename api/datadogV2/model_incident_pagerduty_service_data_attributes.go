// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentPagerdutyServiceDataAttributes Attributes of a PagerDuty service.
type IncidentPagerdutyServiceDataAttributes struct {
	// The handle for the PagerDuty service.
	Handle string `json:"handle"`
	// The name of the PagerDuty service.
	Name string `json:"name"`
	// The PagerDuty service identifier.
	ServiceId string `json:"service_id"`
	// Whether webhooks are enabled for this service.
	WebhooksEnabled bool `json:"webhooks_enabled"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentPagerdutyServiceDataAttributes instantiates a new IncidentPagerdutyServiceDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentPagerdutyServiceDataAttributes(handle string, name string, serviceId string, webhooksEnabled bool) *IncidentPagerdutyServiceDataAttributes {
	this := IncidentPagerdutyServiceDataAttributes{}
	this.Handle = handle
	this.Name = name
	this.ServiceId = serviceId
	this.WebhooksEnabled = webhooksEnabled
	return &this
}

// NewIncidentPagerdutyServiceDataAttributesWithDefaults instantiates a new IncidentPagerdutyServiceDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentPagerdutyServiceDataAttributesWithDefaults() *IncidentPagerdutyServiceDataAttributes {
	this := IncidentPagerdutyServiceDataAttributes{}
	return &this
}

// GetHandle returns the Handle field value.
func (o *IncidentPagerdutyServiceDataAttributes) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *IncidentPagerdutyServiceDataAttributes) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *IncidentPagerdutyServiceDataAttributes) SetHandle(v string) {
	o.Handle = v
}

// GetName returns the Name field value.
func (o *IncidentPagerdutyServiceDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentPagerdutyServiceDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *IncidentPagerdutyServiceDataAttributes) SetName(v string) {
	o.Name = v
}

// GetServiceId returns the ServiceId field value.
func (o *IncidentPagerdutyServiceDataAttributes) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *IncidentPagerdutyServiceDataAttributes) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value.
func (o *IncidentPagerdutyServiceDataAttributes) SetServiceId(v string) {
	o.ServiceId = v
}

// GetWebhooksEnabled returns the WebhooksEnabled field value.
func (o *IncidentPagerdutyServiceDataAttributes) GetWebhooksEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.WebhooksEnabled
}

// GetWebhooksEnabledOk returns a tuple with the WebhooksEnabled field value
// and a boolean to check if the value has been set.
func (o *IncidentPagerdutyServiceDataAttributes) GetWebhooksEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebhooksEnabled, true
}

// SetWebhooksEnabled sets field value.
func (o *IncidentPagerdutyServiceDataAttributes) SetWebhooksEnabled(v bool) {
	o.WebhooksEnabled = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentPagerdutyServiceDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["handle"] = o.Handle
	toSerialize["name"] = o.Name
	toSerialize["service_id"] = o.ServiceId
	toSerialize["webhooks_enabled"] = o.WebhooksEnabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentPagerdutyServiceDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Handle          *string `json:"handle"`
		Name            *string `json:"name"`
		ServiceId       *string `json:"service_id"`
		WebhooksEnabled *bool   `json:"webhooks_enabled"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Handle == nil {
		return fmt.Errorf("required field handle missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ServiceId == nil {
		return fmt.Errorf("required field service_id missing")
	}
	if all.WebhooksEnabled == nil {
		return fmt.Errorf("required field webhooks_enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"handle", "name", "service_id", "webhooks_enabled"})
	} else {
		return err
	}
	o.Handle = *all.Handle
	o.Name = *all.Name
	o.ServiceId = *all.ServiceId
	o.WebhooksEnabled = *all.WebhooksEnabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
