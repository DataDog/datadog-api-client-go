// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OIDCCloudAuthClaimMatcherCreate Claim condition for creating a provider or persona mapping. A token claim must match at least one configured value.
type OIDCCloudAuthClaimMatcherCreate struct {
	// Name of the token claim.
	Claim string `json:"claim"`
	// Case-sensitive matching mode. Exact compares strings. Glob matches the entire value: * matches any number of characters, including /, and ? matches one character. Other characters are literal; escaping is not supported.
	Mode *OIDCCloudAuthClaimMatchMode `json:"mode,omitempty"`
	// Accepted values for this claim. Each value is limited to 500 UTF-8 bytes when creating the resource.
	Values []string `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOIDCCloudAuthClaimMatcherCreate instantiates a new OIDCCloudAuthClaimMatcherCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOIDCCloudAuthClaimMatcherCreate(claim string, values []string) *OIDCCloudAuthClaimMatcherCreate {
	this := OIDCCloudAuthClaimMatcherCreate{}
	this.Claim = claim
	var mode OIDCCloudAuthClaimMatchMode = OIDCCLOUDAUTHCLAIMMATCHMODE_EXACT
	this.Mode = &mode
	this.Values = values
	return &this
}

// NewOIDCCloudAuthClaimMatcherCreateWithDefaults instantiates a new OIDCCloudAuthClaimMatcherCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOIDCCloudAuthClaimMatcherCreateWithDefaults() *OIDCCloudAuthClaimMatcherCreate {
	this := OIDCCloudAuthClaimMatcherCreate{}
	var mode OIDCCloudAuthClaimMatchMode = OIDCCLOUDAUTHCLAIMMATCHMODE_EXACT
	this.Mode = &mode
	return &this
}

// GetClaim returns the Claim field value.
func (o *OIDCCloudAuthClaimMatcherCreate) GetClaim() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Claim
}

// GetClaimOk returns a tuple with the Claim field value
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthClaimMatcherCreate) GetClaimOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Claim, true
}

// SetClaim sets field value.
func (o *OIDCCloudAuthClaimMatcherCreate) SetClaim(v string) {
	o.Claim = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *OIDCCloudAuthClaimMatcherCreate) GetMode() OIDCCloudAuthClaimMatchMode {
	if o == nil || o.Mode == nil {
		var ret OIDCCloudAuthClaimMatchMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthClaimMatcherCreate) GetModeOk() (*OIDCCloudAuthClaimMatchMode, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *OIDCCloudAuthClaimMatcherCreate) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given OIDCCloudAuthClaimMatchMode and assigns it to the Mode field.
func (o *OIDCCloudAuthClaimMatcherCreate) SetMode(v OIDCCloudAuthClaimMatchMode) {
	o.Mode = &v
}

// GetValues returns the Values field value.
func (o *OIDCCloudAuthClaimMatcherCreate) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthClaimMatcherCreate) GetValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *OIDCCloudAuthClaimMatcherCreate) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OIDCCloudAuthClaimMatcherCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["claim"] = o.Claim
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OIDCCloudAuthClaimMatcherCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Claim  *string                      `json:"claim"`
		Mode   *OIDCCloudAuthClaimMatchMode `json:"mode,omitempty"`
		Values *[]string                    `json:"values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Claim == nil {
		return fmt.Errorf("required field claim missing")
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"claim", "mode", "values"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Claim = *all.Claim
	if all.Mode != nil && !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = all.Mode
	}
	o.Values = *all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
