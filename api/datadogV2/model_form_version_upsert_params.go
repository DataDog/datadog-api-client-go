// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormVersionUpsertParams Parameters for upserting a form version.
type FormVersionUpsertParams struct {
	// The entity tag for conflict detection.
	Etag *string `json:"etag,omitempty"`
	// The match policy for upserting.
	MatchPolicy *string `json:"match_policy,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormVersionUpsertParams instantiates a new FormVersionUpsertParams object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormVersionUpsertParams() *FormVersionUpsertParams {
	this := FormVersionUpsertParams{}
	return &this
}

// NewFormVersionUpsertParamsWithDefaults instantiates a new FormVersionUpsertParams object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormVersionUpsertParamsWithDefaults() *FormVersionUpsertParams {
	this := FormVersionUpsertParams{}
	return &this
}

// GetEtag returns the Etag field value if set, zero value otherwise.
func (o *FormVersionUpsertParams) GetEtag() string {
	if o == nil || o.Etag == nil {
		var ret string
		return ret
	}
	return *o.Etag
}

// GetEtagOk returns a tuple with the Etag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormVersionUpsertParams) GetEtagOk() (*string, bool) {
	if o == nil || o.Etag == nil {
		return nil, false
	}
	return o.Etag, true
}

// HasEtag returns a boolean if a field has been set.
func (o *FormVersionUpsertParams) HasEtag() bool {
	return o != nil && o.Etag != nil
}

// SetEtag gets a reference to the given string and assigns it to the Etag field.
func (o *FormVersionUpsertParams) SetEtag(v string) {
	o.Etag = &v
}

// GetMatchPolicy returns the MatchPolicy field value if set, zero value otherwise.
func (o *FormVersionUpsertParams) GetMatchPolicy() string {
	if o == nil || o.MatchPolicy == nil {
		var ret string
		return ret
	}
	return *o.MatchPolicy
}

// GetMatchPolicyOk returns a tuple with the MatchPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormVersionUpsertParams) GetMatchPolicyOk() (*string, bool) {
	if o == nil || o.MatchPolicy == nil {
		return nil, false
	}
	return o.MatchPolicy, true
}

// HasMatchPolicy returns a boolean if a field has been set.
func (o *FormVersionUpsertParams) HasMatchPolicy() bool {
	return o != nil && o.MatchPolicy != nil
}

// SetMatchPolicy gets a reference to the given string and assigns it to the MatchPolicy field.
func (o *FormVersionUpsertParams) SetMatchPolicy(v string) {
	o.MatchPolicy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormVersionUpsertParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Etag != nil {
		toSerialize["etag"] = o.Etag
	}
	if o.MatchPolicy != nil {
		toSerialize["match_policy"] = o.MatchPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormVersionUpsertParams) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Etag        *string `json:"etag,omitempty"`
		MatchPolicy *string `json:"match_policy,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"etag", "match_policy"})
	} else {
		return err
	}
	o.Etag = all.Etag
	o.MatchPolicy = all.MatchPolicy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
