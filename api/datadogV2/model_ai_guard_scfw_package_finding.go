// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardScfwPackageFinding One package the Supply Chain Firewall evaluated, with the outcome and every matched rule.
type AIGuardScfwPackageFinding struct {
	// The package ecosystem the firewall evaluated, such as `PyPI`, `npm`, or `crates.io`.
	Ecosystem string `json:"ecosystem"`
	// Every rule the firewall matched on the package.
	MatchedRules []AIGuardScfwMatchedRule `json:"matched_rules,omitempty"`
	// The name of the package.
	Name string `json:"name"`
	// The Supply Chain Firewall outcome for a package, a matched rule, or the whole request.
	Outcome AIGuardScfwOutcome `json:"outcome"`
	// The version of the package.
	Version string `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAIGuardScfwPackageFinding instantiates a new AIGuardScfwPackageFinding object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAIGuardScfwPackageFinding(ecosystem string, name string, outcome AIGuardScfwOutcome, version string) *AIGuardScfwPackageFinding {
	this := AIGuardScfwPackageFinding{}
	this.Ecosystem = ecosystem
	this.Name = name
	this.Outcome = outcome
	this.Version = version
	return &this
}

// NewAIGuardScfwPackageFindingWithDefaults instantiates a new AIGuardScfwPackageFinding object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAIGuardScfwPackageFindingWithDefaults() *AIGuardScfwPackageFinding {
	this := AIGuardScfwPackageFinding{}
	return &this
}

// GetEcosystem returns the Ecosystem field value.
func (o *AIGuardScfwPackageFinding) GetEcosystem() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Ecosystem
}

// GetEcosystemOk returns a tuple with the Ecosystem field value
// and a boolean to check if the value has been set.
func (o *AIGuardScfwPackageFinding) GetEcosystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ecosystem, true
}

// SetEcosystem sets field value.
func (o *AIGuardScfwPackageFinding) SetEcosystem(v string) {
	o.Ecosystem = v
}

// GetMatchedRules returns the MatchedRules field value if set, zero value otherwise.
func (o *AIGuardScfwPackageFinding) GetMatchedRules() []AIGuardScfwMatchedRule {
	if o == nil || o.MatchedRules == nil {
		var ret []AIGuardScfwMatchedRule
		return ret
	}
	return o.MatchedRules
}

// GetMatchedRulesOk returns a tuple with the MatchedRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardScfwPackageFinding) GetMatchedRulesOk() (*[]AIGuardScfwMatchedRule, bool) {
	if o == nil || o.MatchedRules == nil {
		return nil, false
	}
	return &o.MatchedRules, true
}

// HasMatchedRules returns a boolean if a field has been set.
func (o *AIGuardScfwPackageFinding) HasMatchedRules() bool {
	return o != nil && o.MatchedRules != nil
}

// SetMatchedRules gets a reference to the given []AIGuardScfwMatchedRule and assigns it to the MatchedRules field.
func (o *AIGuardScfwPackageFinding) SetMatchedRules(v []AIGuardScfwMatchedRule) {
	o.MatchedRules = v
}

// GetName returns the Name field value.
func (o *AIGuardScfwPackageFinding) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AIGuardScfwPackageFinding) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AIGuardScfwPackageFinding) SetName(v string) {
	o.Name = v
}

// GetOutcome returns the Outcome field value.
func (o *AIGuardScfwPackageFinding) GetOutcome() AIGuardScfwOutcome {
	if o == nil {
		var ret AIGuardScfwOutcome
		return ret
	}
	return o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value
// and a boolean to check if the value has been set.
func (o *AIGuardScfwPackageFinding) GetOutcomeOk() (*AIGuardScfwOutcome, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outcome, true
}

// SetOutcome sets field value.
func (o *AIGuardScfwPackageFinding) SetOutcome(v AIGuardScfwOutcome) {
	o.Outcome = v
}

// GetVersion returns the Version field value.
func (o *AIGuardScfwPackageFinding) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AIGuardScfwPackageFinding) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *AIGuardScfwPackageFinding) SetVersion(v string) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AIGuardScfwPackageFinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["ecosystem"] = o.Ecosystem
	if o.MatchedRules != nil {
		toSerialize["matched_rules"] = o.MatchedRules
	}
	toSerialize["name"] = o.Name
	toSerialize["outcome"] = o.Outcome
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AIGuardScfwPackageFinding) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ecosystem    *string                  `json:"ecosystem"`
		MatchedRules []AIGuardScfwMatchedRule `json:"matched_rules,omitempty"`
		Name         *string                  `json:"name"`
		Outcome      *AIGuardScfwOutcome      `json:"outcome"`
		Version      *string                  `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Ecosystem == nil {
		return fmt.Errorf("required field ecosystem missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Outcome == nil {
		return fmt.Errorf("required field outcome missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ecosystem", "matched_rules", "name", "outcome", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Ecosystem = *all.Ecosystem
	o.MatchedRules = all.MatchedRules
	o.Name = *all.Name
	if !all.Outcome.IsValid() {
		hasInvalidField = true
	} else {
		o.Outcome = *all.Outcome
	}
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
