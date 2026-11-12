// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AgentHealthRemediationScript A script that can be used to remediate the issue.
type AgentHealthRemediationScript struct {
	// The content of the script.
	Content *string `json:"content,omitempty"`
	// The filename of the script.
	Filename *string `json:"filename,omitempty"`
	// The language of the script.
	Language *string `json:"language,omitempty"`
	// The language version used by the script.
	LanguageVersion *string `json:"language_version,omitempty"`
	// Whether the script requires root privileges.
	RequiresRoot *bool `json:"requires_root,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAgentHealthRemediationScript instantiates a new AgentHealthRemediationScript object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAgentHealthRemediationScript() *AgentHealthRemediationScript {
	this := AgentHealthRemediationScript{}
	return &this
}

// NewAgentHealthRemediationScriptWithDefaults instantiates a new AgentHealthRemediationScript object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAgentHealthRemediationScriptWithDefaults() *AgentHealthRemediationScript {
	this := AgentHealthRemediationScript{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *AgentHealthRemediationScript) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHealthRemediationScript) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *AgentHealthRemediationScript) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *AgentHealthRemediationScript) SetContent(v string) {
	o.Content = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *AgentHealthRemediationScript) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHealthRemediationScript) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *AgentHealthRemediationScript) HasFilename() bool {
	return o != nil && o.Filename != nil
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *AgentHealthRemediationScript) SetFilename(v string) {
	o.Filename = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *AgentHealthRemediationScript) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHealthRemediationScript) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *AgentHealthRemediationScript) HasLanguage() bool {
	return o != nil && o.Language != nil
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *AgentHealthRemediationScript) SetLanguage(v string) {
	o.Language = &v
}

// GetLanguageVersion returns the LanguageVersion field value if set, zero value otherwise.
func (o *AgentHealthRemediationScript) GetLanguageVersion() string {
	if o == nil || o.LanguageVersion == nil {
		var ret string
		return ret
	}
	return *o.LanguageVersion
}

// GetLanguageVersionOk returns a tuple with the LanguageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHealthRemediationScript) GetLanguageVersionOk() (*string, bool) {
	if o == nil || o.LanguageVersion == nil {
		return nil, false
	}
	return o.LanguageVersion, true
}

// HasLanguageVersion returns a boolean if a field has been set.
func (o *AgentHealthRemediationScript) HasLanguageVersion() bool {
	return o != nil && o.LanguageVersion != nil
}

// SetLanguageVersion gets a reference to the given string and assigns it to the LanguageVersion field.
func (o *AgentHealthRemediationScript) SetLanguageVersion(v string) {
	o.LanguageVersion = &v
}

// GetRequiresRoot returns the RequiresRoot field value if set, zero value otherwise.
func (o *AgentHealthRemediationScript) GetRequiresRoot() bool {
	if o == nil || o.RequiresRoot == nil {
		var ret bool
		return ret
	}
	return *o.RequiresRoot
}

// GetRequiresRootOk returns a tuple with the RequiresRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHealthRemediationScript) GetRequiresRootOk() (*bool, bool) {
	if o == nil || o.RequiresRoot == nil {
		return nil, false
	}
	return o.RequiresRoot, true
}

// HasRequiresRoot returns a boolean if a field has been set.
func (o *AgentHealthRemediationScript) HasRequiresRoot() bool {
	return o != nil && o.RequiresRoot != nil
}

// SetRequiresRoot gets a reference to the given bool and assigns it to the RequiresRoot field.
func (o *AgentHealthRemediationScript) SetRequiresRoot(v bool) {
	o.RequiresRoot = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AgentHealthRemediationScript) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.LanguageVersion != nil {
		toSerialize["language_version"] = o.LanguageVersion
	}
	if o.RequiresRoot != nil {
		toSerialize["requires_root"] = o.RequiresRoot
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AgentHealthRemediationScript) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Content         *string `json:"content,omitempty"`
		Filename        *string `json:"filename,omitempty"`
		Language        *string `json:"language,omitempty"`
		LanguageVersion *string `json:"language_version,omitempty"`
		RequiresRoot    *bool   `json:"requires_root,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content", "filename", "language", "language_version", "requires_root"})
	} else {
		return err
	}
	o.Content = all.Content
	o.Filename = all.Filename
	o.Language = all.Language
	o.LanguageVersion = all.LanguageVersion
	o.RequiresRoot = all.RequiresRoot

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
