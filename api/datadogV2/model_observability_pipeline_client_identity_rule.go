// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineClientIdentityRule A named rule. It matches when every constraint it defines matches at least one identity in the client's certificate. It must define at least one of `common_names`, `dns_sans`, or `uri_sans`.
type ObservabilityPipelineClientIdentityRule struct {
	// Patterns for the certificate subject common name. Matching is case-sensitive, and `*` matches any characters, including dots.
	CommonNames []string `json:"common_names,omitempty"`
	// Patterns for DNS subject alternative names. Matching is case-insensitive, and `*` matches any characters within one DNS label.
	DnsSans []string `json:"dns_sans,omitempty"`
	// Unique rule name, reported in metrics.
	Name string `json:"name"`
	// Patterns for URI subject alternative names. Matching is case-sensitive. `*` matches any characters within one path segment, and a trailing `/**` matches the path and everything below it. Wildcards are only supported in the path.
	UriSans []string `json:"uri_sans,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineClientIdentityRule instantiates a new ObservabilityPipelineClientIdentityRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineClientIdentityRule(name string) *ObservabilityPipelineClientIdentityRule {
	this := ObservabilityPipelineClientIdentityRule{}
	this.Name = name
	return &this
}

// NewObservabilityPipelineClientIdentityRuleWithDefaults instantiates a new ObservabilityPipelineClientIdentityRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineClientIdentityRuleWithDefaults() *ObservabilityPipelineClientIdentityRule {
	this := ObservabilityPipelineClientIdentityRule{}
	return &this
}

// GetCommonNames returns the CommonNames field value if set, zero value otherwise.
func (o *ObservabilityPipelineClientIdentityRule) GetCommonNames() []string {
	if o == nil || o.CommonNames == nil {
		var ret []string
		return ret
	}
	return o.CommonNames
}

// GetCommonNamesOk returns a tuple with the CommonNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClientIdentityRule) GetCommonNamesOk() (*[]string, bool) {
	if o == nil || o.CommonNames == nil {
		return nil, false
	}
	return &o.CommonNames, true
}

// HasCommonNames returns a boolean if a field has been set.
func (o *ObservabilityPipelineClientIdentityRule) HasCommonNames() bool {
	return o != nil && o.CommonNames != nil
}

// SetCommonNames gets a reference to the given []string and assigns it to the CommonNames field.
func (o *ObservabilityPipelineClientIdentityRule) SetCommonNames(v []string) {
	o.CommonNames = v
}

// GetDnsSans returns the DnsSans field value if set, zero value otherwise.
func (o *ObservabilityPipelineClientIdentityRule) GetDnsSans() []string {
	if o == nil || o.DnsSans == nil {
		var ret []string
		return ret
	}
	return o.DnsSans
}

// GetDnsSansOk returns a tuple with the DnsSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClientIdentityRule) GetDnsSansOk() (*[]string, bool) {
	if o == nil || o.DnsSans == nil {
		return nil, false
	}
	return &o.DnsSans, true
}

// HasDnsSans returns a boolean if a field has been set.
func (o *ObservabilityPipelineClientIdentityRule) HasDnsSans() bool {
	return o != nil && o.DnsSans != nil
}

// SetDnsSans gets a reference to the given []string and assigns it to the DnsSans field.
func (o *ObservabilityPipelineClientIdentityRule) SetDnsSans(v []string) {
	o.DnsSans = v
}

// GetName returns the Name field value.
func (o *ObservabilityPipelineClientIdentityRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClientIdentityRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ObservabilityPipelineClientIdentityRule) SetName(v string) {
	o.Name = v
}

// GetUriSans returns the UriSans field value if set, zero value otherwise.
func (o *ObservabilityPipelineClientIdentityRule) GetUriSans() []string {
	if o == nil || o.UriSans == nil {
		var ret []string
		return ret
	}
	return o.UriSans
}

// GetUriSansOk returns a tuple with the UriSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClientIdentityRule) GetUriSansOk() (*[]string, bool) {
	if o == nil || o.UriSans == nil {
		return nil, false
	}
	return &o.UriSans, true
}

// HasUriSans returns a boolean if a field has been set.
func (o *ObservabilityPipelineClientIdentityRule) HasUriSans() bool {
	return o != nil && o.UriSans != nil
}

// SetUriSans gets a reference to the given []string and assigns it to the UriSans field.
func (o *ObservabilityPipelineClientIdentityRule) SetUriSans(v []string) {
	o.UriSans = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineClientIdentityRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CommonNames != nil {
		toSerialize["common_names"] = o.CommonNames
	}
	if o.DnsSans != nil {
		toSerialize["dns_sans"] = o.DnsSans
	}
	toSerialize["name"] = o.Name
	if o.UriSans != nil {
		toSerialize["uri_sans"] = o.UriSans
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineClientIdentityRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CommonNames []string `json:"common_names,omitempty"`
		DnsSans     []string `json:"dns_sans,omitempty"`
		Name        *string  `json:"name"`
		UriSans     []string `json:"uri_sans,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"common_names", "dns_sans", "name", "uri_sans"})
	} else {
		return err
	}
	o.CommonNames = all.CommonNames
	o.DnsSans = all.DnsSans
	o.Name = *all.Name
	o.UriSans = all.UriSans

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
