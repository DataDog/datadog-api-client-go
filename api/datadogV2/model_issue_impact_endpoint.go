// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueImpactEndpoint An endpoint whose traces have the issue's errors.
type IssueImpactEndpoint struct {
	// Number of errors reported by each source. Several sources can report the same error, such as an APM span and
	// a log, so the counts overlap: don't add them together.
	ErrorCounts IssueImpactErrorCounts `json:"error_counts"`
	// Resource name of the trace's root span.
	ResourceName string `json:"resource_name"`
	// Service of the trace's root span.
	Service string `json:"service"`
	// Whether the entry stands out as a likely driver of the issue, because it accounts for most of the issue's
	// errors or for far more errors than the other listed entries.
	Significant bool `json:"significant"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueImpactEndpoint instantiates a new IssueImpactEndpoint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueImpactEndpoint(errorCounts IssueImpactErrorCounts, resourceName string, service string, significant bool) *IssueImpactEndpoint {
	this := IssueImpactEndpoint{}
	this.ErrorCounts = errorCounts
	this.ResourceName = resourceName
	this.Service = service
	this.Significant = significant
	return &this
}

// NewIssueImpactEndpointWithDefaults instantiates a new IssueImpactEndpoint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueImpactEndpointWithDefaults() *IssueImpactEndpoint {
	this := IssueImpactEndpoint{}
	return &this
}

// GetErrorCounts returns the ErrorCounts field value.
func (o *IssueImpactEndpoint) GetErrorCounts() IssueImpactErrorCounts {
	if o == nil {
		var ret IssueImpactErrorCounts
		return ret
	}
	return o.ErrorCounts
}

// GetErrorCountsOk returns a tuple with the ErrorCounts field value
// and a boolean to check if the value has been set.
func (o *IssueImpactEndpoint) GetErrorCountsOk() (*IssueImpactErrorCounts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCounts, true
}

// SetErrorCounts sets field value.
func (o *IssueImpactEndpoint) SetErrorCounts(v IssueImpactErrorCounts) {
	o.ErrorCounts = v
}

// GetResourceName returns the ResourceName field value.
func (o *IssueImpactEndpoint) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *IssueImpactEndpoint) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value.
func (o *IssueImpactEndpoint) SetResourceName(v string) {
	o.ResourceName = v
}

// GetService returns the Service field value.
func (o *IssueImpactEndpoint) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *IssueImpactEndpoint) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *IssueImpactEndpoint) SetService(v string) {
	o.Service = v
}

// GetSignificant returns the Significant field value.
func (o *IssueImpactEndpoint) GetSignificant() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Significant
}

// GetSignificantOk returns a tuple with the Significant field value
// and a boolean to check if the value has been set.
func (o *IssueImpactEndpoint) GetSignificantOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Significant, true
}

// SetSignificant sets field value.
func (o *IssueImpactEndpoint) SetSignificant(v bool) {
	o.Significant = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueImpactEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["error_counts"] = o.ErrorCounts
	toSerialize["resource_name"] = o.ResourceName
	toSerialize["service"] = o.Service
	toSerialize["significant"] = o.Significant

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueImpactEndpoint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ErrorCounts  *IssueImpactErrorCounts `json:"error_counts"`
		ResourceName *string                 `json:"resource_name"`
		Service      *string                 `json:"service"`
		Significant  *bool                   `json:"significant"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ErrorCounts == nil {
		return fmt.Errorf("required field error_counts missing")
	}
	if all.ResourceName == nil {
		return fmt.Errorf("required field resource_name missing")
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}
	if all.Significant == nil {
		return fmt.Errorf("required field significant missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error_counts", "resource_name", "service", "significant"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ErrorCounts.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ErrorCounts = *all.ErrorCounts
	o.ResourceName = *all.ResourceName
	o.Service = *all.Service
	o.Significant = *all.Significant

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
