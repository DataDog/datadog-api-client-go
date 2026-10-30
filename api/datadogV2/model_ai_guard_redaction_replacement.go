// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardRedactionReplacement The fully redacted content for one path in the request.
type AIGuardRedactionReplacement struct {
	// The path of the redacted content relative to the request attributes.
	Path string `json:"path"`
	// The redacted replacement content.
	Replacement string `json:"replacement"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAIGuardRedactionReplacement instantiates a new AIGuardRedactionReplacement object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAIGuardRedactionReplacement(path string, replacement string) *AIGuardRedactionReplacement {
	this := AIGuardRedactionReplacement{}
	this.Path = path
	this.Replacement = replacement
	return &this
}

// NewAIGuardRedactionReplacementWithDefaults instantiates a new AIGuardRedactionReplacement object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAIGuardRedactionReplacementWithDefaults() *AIGuardRedactionReplacement {
	this := AIGuardRedactionReplacement{}
	return &this
}

// GetPath returns the Path field value.
func (o *AIGuardRedactionReplacement) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *AIGuardRedactionReplacement) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value.
func (o *AIGuardRedactionReplacement) SetPath(v string) {
	o.Path = v
}

// GetReplacement returns the Replacement field value.
func (o *AIGuardRedactionReplacement) GetReplacement() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Replacement
}

// GetReplacementOk returns a tuple with the Replacement field value
// and a boolean to check if the value has been set.
func (o *AIGuardRedactionReplacement) GetReplacementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replacement, true
}

// SetReplacement sets field value.
func (o *AIGuardRedactionReplacement) SetReplacement(v string) {
	o.Replacement = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AIGuardRedactionReplacement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["path"] = o.Path
	toSerialize["replacement"] = o.Replacement

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AIGuardRedactionReplacement) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Path        *string `json:"path"`
		Replacement *string `json:"replacement"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Path == nil {
		return fmt.Errorf("required field path missing")
	}
	if all.Replacement == nil {
		return fmt.Errorf("required field replacement missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"path", "replacement"})
	} else {
		return err
	}
	o.Path = *all.Path
	o.Replacement = *all.Replacement

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
