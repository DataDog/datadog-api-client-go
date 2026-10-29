// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSResourceTypeTagFilter AWS Metrics Collection resource type tag filter.
// Each filter targets either a specific AWS service's resource types (`service` and
// `resource_types`) or custom metrics (`custom`). These two modes are mutually exclusive.
// The `expression` defines a tag filter that Datadog uses when collecting metrics from the
// resource types of a specified service. Wildcards, such as `?` (match a single character)
// and `*` (match multiple characters), and exclusion using `!` before the tag are supported.
// For example, `env:production AND instance-type:c?.* AND !region:us-east-1`.
type AWSResourceTypeTagFilter struct {
	// Whether the tag filter applies to custom metrics. When `true`, `service` and `resource_types` must be omitted.
	Custom *bool `json:"custom,omitempty"`
	// The tag filter expression to apply when collecting metrics for the matching resource types.
	Expression string `json:"expression"`
	// The AWS resource types of the service to which the tag filter applies. Required when `service` is set. Use `*` as the only entry to match all resource types of the service; `*` cannot be combined with specific resource types.
	ResourceTypes []string `json:"resource_types,omitempty"`
	// The AWS service for which the tag filter defined in `expression` will be applied. Required unless `custom` is `true`, and must be set together with `resource_types`.
	Service *string `json:"service,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSResourceTypeTagFilter instantiates a new AWSResourceTypeTagFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSResourceTypeTagFilter(expression string) *AWSResourceTypeTagFilter {
	this := AWSResourceTypeTagFilter{}
	this.Expression = expression
	return &this
}

// NewAWSResourceTypeTagFilterWithDefaults instantiates a new AWSResourceTypeTagFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSResourceTypeTagFilterWithDefaults() *AWSResourceTypeTagFilter {
	this := AWSResourceTypeTagFilter{}
	return &this
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *AWSResourceTypeTagFilter) GetCustom() bool {
	if o == nil || o.Custom == nil {
		var ret bool
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSResourceTypeTagFilter) GetCustomOk() (*bool, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *AWSResourceTypeTagFilter) HasCustom() bool {
	return o != nil && o.Custom != nil
}

// SetCustom gets a reference to the given bool and assigns it to the Custom field.
func (o *AWSResourceTypeTagFilter) SetCustom(v bool) {
	o.Custom = &v
}

// GetExpression returns the Expression field value.
func (o *AWSResourceTypeTagFilter) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *AWSResourceTypeTagFilter) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value.
func (o *AWSResourceTypeTagFilter) SetExpression(v string) {
	o.Expression = v
}

// GetResourceTypes returns the ResourceTypes field value if set, zero value otherwise.
func (o *AWSResourceTypeTagFilter) GetResourceTypes() []string {
	if o == nil || o.ResourceTypes == nil {
		var ret []string
		return ret
	}
	return o.ResourceTypes
}

// GetResourceTypesOk returns a tuple with the ResourceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSResourceTypeTagFilter) GetResourceTypesOk() (*[]string, bool) {
	if o == nil || o.ResourceTypes == nil {
		return nil, false
	}
	return &o.ResourceTypes, true
}

// HasResourceTypes returns a boolean if a field has been set.
func (o *AWSResourceTypeTagFilter) HasResourceTypes() bool {
	return o != nil && o.ResourceTypes != nil
}

// SetResourceTypes gets a reference to the given []string and assigns it to the ResourceTypes field.
func (o *AWSResourceTypeTagFilter) SetResourceTypes(v []string) {
	o.ResourceTypes = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *AWSResourceTypeTagFilter) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSResourceTypeTagFilter) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *AWSResourceTypeTagFilter) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *AWSResourceTypeTagFilter) SetService(v string) {
	o.Service = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSResourceTypeTagFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Custom != nil {
		toSerialize["custom"] = o.Custom
	}
	toSerialize["expression"] = o.Expression
	if o.ResourceTypes != nil {
		toSerialize["resource_types"] = o.ResourceTypes
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSResourceTypeTagFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Custom        *bool    `json:"custom,omitempty"`
		Expression    *string  `json:"expression"`
		ResourceTypes []string `json:"resource_types,omitempty"`
		Service       *string  `json:"service,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Expression == nil {
		return fmt.Errorf("required field expression missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"custom", "expression", "resource_types", "service"})
	} else {
		return err
	}
	o.Custom = all.Custom
	o.Expression = *all.Expression
	o.ResourceTypes = all.ResourceTypes
	o.Service = all.Service

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
