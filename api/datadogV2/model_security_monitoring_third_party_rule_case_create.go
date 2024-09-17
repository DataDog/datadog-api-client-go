// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV2

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"

)


// SecurityMonitoringThirdPartyRuleCaseCreate Case when a signal is generated by a third party rule.
type SecurityMonitoringThirdPartyRuleCaseCreate struct {
	// Name of the case.
	Name *string `json:"name,omitempty"`
	// Notification targets for each rule case.
	Notifications []string `json:"notifications,omitempty"`
	// A query to map a third party event to this case.
	Query *string `json:"query,omitempty"`
	// Severity of the Security Signal.
	Status SecurityMonitoringRuleSeverity `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewSecurityMonitoringThirdPartyRuleCaseCreate instantiates a new SecurityMonitoringThirdPartyRuleCaseCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringThirdPartyRuleCaseCreate(status SecurityMonitoringRuleSeverity) *SecurityMonitoringThirdPartyRuleCaseCreate {
	this := SecurityMonitoringThirdPartyRuleCaseCreate{}
	this.Status = status
	return &this
}

// NewSecurityMonitoringThirdPartyRuleCaseCreateWithDefaults instantiates a new SecurityMonitoringThirdPartyRuleCaseCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringThirdPartyRuleCaseCreateWithDefaults() *SecurityMonitoringThirdPartyRuleCaseCreate {
	this := SecurityMonitoringThirdPartyRuleCaseCreate{}
	return &this
}
// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringThirdPartyRuleCaseCreate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringThirdPartyRuleCaseCreate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringThirdPartyRuleCaseCreate) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringThirdPartyRuleCaseCreate) SetName(v string) {
	o.Name = &v
}


// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *SecurityMonitoringThirdPartyRuleCaseCreate) GetNotifications() []string {
	if o == nil || o.Notifications == nil {
		var ret []string
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringThirdPartyRuleCaseCreate) GetNotificationsOk() (*[]string, bool) {
	if o == nil || o.Notifications == nil {
		return nil, false
	}
	return &o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *SecurityMonitoringThirdPartyRuleCaseCreate) HasNotifications() bool {
	return o != nil && o.Notifications != nil
}

// SetNotifications gets a reference to the given []string and assigns it to the Notifications field.
func (o *SecurityMonitoringThirdPartyRuleCaseCreate) SetNotifications(v []string) {
	o.Notifications = v
}


// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SecurityMonitoringThirdPartyRuleCaseCreate) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringThirdPartyRuleCaseCreate) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SecurityMonitoringThirdPartyRuleCaseCreate) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *SecurityMonitoringThirdPartyRuleCaseCreate) SetQuery(v string) {
	o.Query = &v
}


// GetStatus returns the Status field value.
func (o *SecurityMonitoringThirdPartyRuleCaseCreate) GetStatus() SecurityMonitoringRuleSeverity {
	if o == nil {
		var ret SecurityMonitoringRuleSeverity
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringThirdPartyRuleCaseCreate) GetStatusOk() (*SecurityMonitoringRuleSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *SecurityMonitoringThirdPartyRuleCaseCreate) SetStatus(v SecurityMonitoringRuleSeverity) {
	o.Status = v
}



// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringThirdPartyRuleCaseCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Notifications != nil {
		toSerialize["notifications"] = o.Notifications
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringThirdPartyRuleCaseCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string `json:"name,omitempty"`
		Notifications []string `json:"notifications,omitempty"`
		Query *string `json:"query,omitempty"`
		Status *SecurityMonitoringRuleSeverity `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{ "name", "notifications", "query", "status",  })
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	o.Notifications = all.Notifications
	o.Query = all.Query
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
