// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalInvestigationFeedbackSection A feedback section containing metrics.
type SecurityMonitoringSignalInvestigationFeedbackSection struct {
	// The section identifier.
	Id string `json:"id"`
	// Array of feedback metrics.
	Metrics []SecurityMonitoringSignalInvestigationFeedbackMetric `json:"metrics"`
	// The section title.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSignalInvestigationFeedbackSection instantiates a new SecurityMonitoringSignalInvestigationFeedbackSection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSignalInvestigationFeedbackSection(id string, metrics []SecurityMonitoringSignalInvestigationFeedbackMetric, title string) *SecurityMonitoringSignalInvestigationFeedbackSection {
	this := SecurityMonitoringSignalInvestigationFeedbackSection{}
	this.Id = id
	this.Metrics = metrics
	this.Title = title
	return &this
}

// NewSecurityMonitoringSignalInvestigationFeedbackSectionWithDefaults instantiates a new SecurityMonitoringSignalInvestigationFeedbackSection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSignalInvestigationFeedbackSectionWithDefaults() *SecurityMonitoringSignalInvestigationFeedbackSection {
	this := SecurityMonitoringSignalInvestigationFeedbackSection{}
	return &this
}

// GetId returns the Id field value.
func (o *SecurityMonitoringSignalInvestigationFeedbackSection) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackSection) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *SecurityMonitoringSignalInvestigationFeedbackSection) SetId(v string) {
	o.Id = v
}

// GetMetrics returns the Metrics field value.
func (o *SecurityMonitoringSignalInvestigationFeedbackSection) GetMetrics() []SecurityMonitoringSignalInvestigationFeedbackMetric {
	if o == nil {
		var ret []SecurityMonitoringSignalInvestigationFeedbackMetric
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackSection) GetMetricsOk() (*[]SecurityMonitoringSignalInvestigationFeedbackMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value.
func (o *SecurityMonitoringSignalInvestigationFeedbackSection) SetMetrics(v []SecurityMonitoringSignalInvestigationFeedbackMetric) {
	o.Metrics = v
}

// GetTitle returns the Title field value.
func (o *SecurityMonitoringSignalInvestigationFeedbackSection) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationFeedbackSection) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *SecurityMonitoringSignalInvestigationFeedbackSection) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSignalInvestigationFeedbackSection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["metrics"] = o.Metrics
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSignalInvestigationFeedbackSection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id      *string                                                `json:"id"`
		Metrics *[]SecurityMonitoringSignalInvestigationFeedbackMetric `json:"metrics"`
		Title   *string                                                `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Metrics == nil {
		return fmt.Errorf("required field metrics missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "metrics", "title"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.Metrics = *all.Metrics
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
