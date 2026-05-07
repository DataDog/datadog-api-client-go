// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TransportWebhookLog A single email transport webhook log event.
type TransportWebhookLog struct {
	// Top-level attributes for the webhook log event, including delivery status, recipient details, and provider metadata.
	Attributes TransportWebhookLogAttributes `json:"attributes"`
	// The ISO 8601 timestamp of the event.
	Date time.Time `json:"date"`
	// The unique log event identifier.
	LogId string `json:"log_id"`
	// The email transport provider.
	Source string `json:"source"`
	// The log status level.
	Status string `json:"status"`
	// A list of tags associated with the event.
	Tags []string `json:"tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTransportWebhookLog instantiates a new TransportWebhookLog object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTransportWebhookLog(attributes TransportWebhookLogAttributes, date time.Time, logId string, source string, status string, tags []string) *TransportWebhookLog {
	this := TransportWebhookLog{}
	this.Attributes = attributes
	this.Date = date
	this.LogId = logId
	this.Source = source
	this.Status = status
	this.Tags = tags
	return &this
}

// NewTransportWebhookLogWithDefaults instantiates a new TransportWebhookLog object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTransportWebhookLogWithDefaults() *TransportWebhookLog {
	this := TransportWebhookLog{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *TransportWebhookLog) GetAttributes() TransportWebhookLogAttributes {
	if o == nil {
		var ret TransportWebhookLogAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TransportWebhookLog) GetAttributesOk() (*TransportWebhookLogAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *TransportWebhookLog) SetAttributes(v TransportWebhookLogAttributes) {
	o.Attributes = v
}

// GetDate returns the Date field value.
func (o *TransportWebhookLog) GetDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *TransportWebhookLog) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value.
func (o *TransportWebhookLog) SetDate(v time.Time) {
	o.Date = v
}

// GetLogId returns the LogId field value.
func (o *TransportWebhookLog) GetLogId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LogId
}

// GetLogIdOk returns a tuple with the LogId field value
// and a boolean to check if the value has been set.
func (o *TransportWebhookLog) GetLogIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogId, true
}

// SetLogId sets field value.
func (o *TransportWebhookLog) SetLogId(v string) {
	o.LogId = v
}

// GetSource returns the Source field value.
func (o *TransportWebhookLog) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *TransportWebhookLog) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *TransportWebhookLog) SetSource(v string) {
	o.Source = v
}

// GetStatus returns the Status field value.
func (o *TransportWebhookLog) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TransportWebhookLog) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *TransportWebhookLog) SetStatus(v string) {
	o.Status = v
}

// GetTags returns the Tags field value.
func (o *TransportWebhookLog) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *TransportWebhookLog) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *TransportWebhookLog) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TransportWebhookLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	if o.Date.Nanosecond() == 0 {
		toSerialize["date"] = o.Date.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["date"] = o.Date.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["log_id"] = o.LogId
	toSerialize["source"] = o.Source
	toSerialize["status"] = o.Status
	toSerialize["tags"] = o.Tags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TransportWebhookLog) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *TransportWebhookLogAttributes `json:"attributes"`
		Date       *time.Time                     `json:"date"`
		LogId      *string                        `json:"log_id"`
		Source     *string                        `json:"source"`
		Status     *string                        `json:"status"`
		Tags       *[]string                      `json:"tags"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Date == nil {
		return fmt.Errorf("required field date missing")
	}
	if all.LogId == nil {
		return fmt.Errorf("required field log_id missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "date", "log_id", "source", "status", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Date = *all.Date
	o.LogId = *all.LogId
	o.Source = *all.Source
	o.Status = *all.Status
	o.Tags = *all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
