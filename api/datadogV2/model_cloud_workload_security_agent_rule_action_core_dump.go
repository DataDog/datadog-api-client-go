// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentRuleActionCoreDump The core dump action applied on the process matching the rule.
type CloudWorkloadSecurityAgentRuleActionCoreDump struct {
	// Whether the directory entry information is included in the core dump.
	Dentry *bool `json:"dentry,omitempty"`
	// Whether the mount information is included in the core dump.
	Mount *bool `json:"mount,omitempty"`
	// Whether the core dump is left uncompressed.
	NoCompression *bool `json:"no_compression,omitempty"`
	// Whether the process memory is included in the core dump.
	Process *bool `json:"process,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudWorkloadSecurityAgentRuleActionCoreDump instantiates a new CloudWorkloadSecurityAgentRuleActionCoreDump object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudWorkloadSecurityAgentRuleActionCoreDump() *CloudWorkloadSecurityAgentRuleActionCoreDump {
	this := CloudWorkloadSecurityAgentRuleActionCoreDump{}
	return &this
}

// NewCloudWorkloadSecurityAgentRuleActionCoreDumpWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleActionCoreDump object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudWorkloadSecurityAgentRuleActionCoreDumpWithDefaults() *CloudWorkloadSecurityAgentRuleActionCoreDump {
	this := CloudWorkloadSecurityAgentRuleActionCoreDump{}
	return &this
}

// GetDentry returns the Dentry field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionCoreDump) GetDentry() bool {
	if o == nil || o.Dentry == nil {
		var ret bool
		return ret
	}
	return *o.Dentry
}

// GetDentryOk returns a tuple with the Dentry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionCoreDump) GetDentryOk() (*bool, bool) {
	if o == nil || o.Dentry == nil {
		return nil, false
	}
	return o.Dentry, true
}

// HasDentry returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionCoreDump) HasDentry() bool {
	return o != nil && o.Dentry != nil
}

// SetDentry gets a reference to the given bool and assigns it to the Dentry field.
func (o *CloudWorkloadSecurityAgentRuleActionCoreDump) SetDentry(v bool) {
	o.Dentry = &v
}

// GetMount returns the Mount field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionCoreDump) GetMount() bool {
	if o == nil || o.Mount == nil {
		var ret bool
		return ret
	}
	return *o.Mount
}

// GetMountOk returns a tuple with the Mount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionCoreDump) GetMountOk() (*bool, bool) {
	if o == nil || o.Mount == nil {
		return nil, false
	}
	return o.Mount, true
}

// HasMount returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionCoreDump) HasMount() bool {
	return o != nil && o.Mount != nil
}

// SetMount gets a reference to the given bool and assigns it to the Mount field.
func (o *CloudWorkloadSecurityAgentRuleActionCoreDump) SetMount(v bool) {
	o.Mount = &v
}

// GetNoCompression returns the NoCompression field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionCoreDump) GetNoCompression() bool {
	if o == nil || o.NoCompression == nil {
		var ret bool
		return ret
	}
	return *o.NoCompression
}

// GetNoCompressionOk returns a tuple with the NoCompression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionCoreDump) GetNoCompressionOk() (*bool, bool) {
	if o == nil || o.NoCompression == nil {
		return nil, false
	}
	return o.NoCompression, true
}

// HasNoCompression returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionCoreDump) HasNoCompression() bool {
	return o != nil && o.NoCompression != nil
}

// SetNoCompression gets a reference to the given bool and assigns it to the NoCompression field.
func (o *CloudWorkloadSecurityAgentRuleActionCoreDump) SetNoCompression(v bool) {
	o.NoCompression = &v
}

// GetProcess returns the Process field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionCoreDump) GetProcess() bool {
	if o == nil || o.Process == nil {
		var ret bool
		return ret
	}
	return *o.Process
}

// GetProcessOk returns a tuple with the Process field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionCoreDump) GetProcessOk() (*bool, bool) {
	if o == nil || o.Process == nil {
		return nil, false
	}
	return o.Process, true
}

// HasProcess returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionCoreDump) HasProcess() bool {
	return o != nil && o.Process != nil
}

// SetProcess gets a reference to the given bool and assigns it to the Process field.
func (o *CloudWorkloadSecurityAgentRuleActionCoreDump) SetProcess(v bool) {
	o.Process = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudWorkloadSecurityAgentRuleActionCoreDump) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Dentry != nil {
		toSerialize["dentry"] = o.Dentry
	}
	if o.Mount != nil {
		toSerialize["mount"] = o.Mount
	}
	if o.NoCompression != nil {
		toSerialize["no_compression"] = o.NoCompression
	}
	if o.Process != nil {
		toSerialize["process"] = o.Process
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudWorkloadSecurityAgentRuleActionCoreDump) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Dentry        *bool `json:"dentry,omitempty"`
		Mount         *bool `json:"mount,omitempty"`
		NoCompression *bool `json:"no_compression,omitempty"`
		Process       *bool `json:"process,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dentry", "mount", "no_compression", "process"})
	} else {
		return err
	}
	o.Dentry = all.Dentry
	o.Mount = all.Mount
	o.NoCompression = all.NoCompression
	o.Process = all.Process

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
