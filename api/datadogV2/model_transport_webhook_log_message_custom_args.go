// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TransportWebhookLogMessageCustomArgs Custom arguments passed through the email transport provider for tracking.
type TransportWebhookLogMessageCustomArgs struct {
	// The unique email identifier.
	EmailId *string `json:"email_id,omitempty"`
	// The human-readable email type name.
	EmailTypeDisplayName *string `json:"email_type_display_name,omitempty"`
	// The organization UUID.
	OrgUuid *string `json:"org_uuid,omitempty"`
	// The timestamp when the email was queued.
	QueueTime *string `json:"queue_time,omitempty"`
	// The email subject line.
	Subject *string `json:"subject,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTransportWebhookLogMessageCustomArgs instantiates a new TransportWebhookLogMessageCustomArgs object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTransportWebhookLogMessageCustomArgs() *TransportWebhookLogMessageCustomArgs {
	this := TransportWebhookLogMessageCustomArgs{}
	return &this
}

// NewTransportWebhookLogMessageCustomArgsWithDefaults instantiates a new TransportWebhookLogMessageCustomArgs object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTransportWebhookLogMessageCustomArgsWithDefaults() *TransportWebhookLogMessageCustomArgs {
	this := TransportWebhookLogMessageCustomArgs{}
	return &this
}

// GetEmailId returns the EmailId field value if set, zero value otherwise.
func (o *TransportWebhookLogMessageCustomArgs) GetEmailId() string {
	if o == nil || o.EmailId == nil {
		var ret string
		return ret
	}
	return *o.EmailId
}

// GetEmailIdOk returns a tuple with the EmailId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessageCustomArgs) GetEmailIdOk() (*string, bool) {
	if o == nil || o.EmailId == nil {
		return nil, false
	}
	return o.EmailId, true
}

// HasEmailId returns a boolean if a field has been set.
func (o *TransportWebhookLogMessageCustomArgs) HasEmailId() bool {
	return o != nil && o.EmailId != nil
}

// SetEmailId gets a reference to the given string and assigns it to the EmailId field.
func (o *TransportWebhookLogMessageCustomArgs) SetEmailId(v string) {
	o.EmailId = &v
}

// GetEmailTypeDisplayName returns the EmailTypeDisplayName field value if set, zero value otherwise.
func (o *TransportWebhookLogMessageCustomArgs) GetEmailTypeDisplayName() string {
	if o == nil || o.EmailTypeDisplayName == nil {
		var ret string
		return ret
	}
	return *o.EmailTypeDisplayName
}

// GetEmailTypeDisplayNameOk returns a tuple with the EmailTypeDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessageCustomArgs) GetEmailTypeDisplayNameOk() (*string, bool) {
	if o == nil || o.EmailTypeDisplayName == nil {
		return nil, false
	}
	return o.EmailTypeDisplayName, true
}

// HasEmailTypeDisplayName returns a boolean if a field has been set.
func (o *TransportWebhookLogMessageCustomArgs) HasEmailTypeDisplayName() bool {
	return o != nil && o.EmailTypeDisplayName != nil
}

// SetEmailTypeDisplayName gets a reference to the given string and assigns it to the EmailTypeDisplayName field.
func (o *TransportWebhookLogMessageCustomArgs) SetEmailTypeDisplayName(v string) {
	o.EmailTypeDisplayName = &v
}

// GetOrgUuid returns the OrgUuid field value if set, zero value otherwise.
func (o *TransportWebhookLogMessageCustomArgs) GetOrgUuid() string {
	if o == nil || o.OrgUuid == nil {
		var ret string
		return ret
	}
	return *o.OrgUuid
}

// GetOrgUuidOk returns a tuple with the OrgUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessageCustomArgs) GetOrgUuidOk() (*string, bool) {
	if o == nil || o.OrgUuid == nil {
		return nil, false
	}
	return o.OrgUuid, true
}

// HasOrgUuid returns a boolean if a field has been set.
func (o *TransportWebhookLogMessageCustomArgs) HasOrgUuid() bool {
	return o != nil && o.OrgUuid != nil
}

// SetOrgUuid gets a reference to the given string and assigns it to the OrgUuid field.
func (o *TransportWebhookLogMessageCustomArgs) SetOrgUuid(v string) {
	o.OrgUuid = &v
}

// GetQueueTime returns the QueueTime field value if set, zero value otherwise.
func (o *TransportWebhookLogMessageCustomArgs) GetQueueTime() string {
	if o == nil || o.QueueTime == nil {
		var ret string
		return ret
	}
	return *o.QueueTime
}

// GetQueueTimeOk returns a tuple with the QueueTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessageCustomArgs) GetQueueTimeOk() (*string, bool) {
	if o == nil || o.QueueTime == nil {
		return nil, false
	}
	return o.QueueTime, true
}

// HasQueueTime returns a boolean if a field has been set.
func (o *TransportWebhookLogMessageCustomArgs) HasQueueTime() bool {
	return o != nil && o.QueueTime != nil
}

// SetQueueTime gets a reference to the given string and assigns it to the QueueTime field.
func (o *TransportWebhookLogMessageCustomArgs) SetQueueTime(v string) {
	o.QueueTime = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *TransportWebhookLogMessageCustomArgs) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessageCustomArgs) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *TransportWebhookLogMessageCustomArgs) HasSubject() bool {
	return o != nil && o.Subject != nil
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *TransportWebhookLogMessageCustomArgs) SetSubject(v string) {
	o.Subject = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TransportWebhookLogMessageCustomArgs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EmailId != nil {
		toSerialize["email_id"] = o.EmailId
	}
	if o.EmailTypeDisplayName != nil {
		toSerialize["email_type_display_name"] = o.EmailTypeDisplayName
	}
	if o.OrgUuid != nil {
		toSerialize["org_uuid"] = o.OrgUuid
	}
	if o.QueueTime != nil {
		toSerialize["queue_time"] = o.QueueTime
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TransportWebhookLogMessageCustomArgs) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EmailId              *string `json:"email_id,omitempty"`
		EmailTypeDisplayName *string `json:"email_type_display_name,omitempty"`
		OrgUuid              *string `json:"org_uuid,omitempty"`
		QueueTime            *string `json:"queue_time,omitempty"`
		Subject              *string `json:"subject,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"email_id", "email_type_display_name", "org_uuid", "queue_time", "subject"})
	} else {
		return err
	}
	o.EmailId = all.EmailId
	o.EmailTypeDisplayName = all.EmailTypeDisplayName
	o.OrgUuid = all.OrgUuid
	o.QueueTime = all.QueueTime
	o.Subject = all.Subject

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
