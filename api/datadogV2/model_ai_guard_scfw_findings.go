// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardScfwFindings The Supply Chain Firewall verdict for the version-pinned install packages
// found in the request. Evaluation requires Supply Chain Firewall to be enabled
// for the organization and the caller to have the `appsec_vm_read` permission.
// Without that permission, Supply Chain Firewall evaluation is skipped; the
// AI Guard evaluation request itself is not rejected. Findings are omitted
// when Supply Chain Firewall does not produce a result.
// The recommended action combines the package outcomes.
// Confirmation support is inferred from the evaluator profile, not from a
// separate confirmation-capability flag. By default, a non-empty
// `meta.coding_agent` selects the coding-agent profile, which assumes the caller
// can ask for confirmation and retains WARN in `recommended_action`.
// Omitting `meta.coding_agent` or setting it to an empty string selects the
// default profile, which assumes the caller cannot ask for confirmation and
// returns ALLOW instead of WARN in `recommended_action`. Package outcomes
// still retain WARN, and BLOCK is never changed to ALLOW by this rule.
type AIGuardScfwFindings struct {
	// The packages the firewall evaluated, `ALLOW` ones included.
	Packages []AIGuardScfwPackageFinding `json:"packages,omitempty"`
	// The Supply Chain Firewall outcome for a package, a matched rule, or the whole request.
	RecommendedAction AIGuardScfwOutcome `json:"recommended_action"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAIGuardScfwFindings instantiates a new AIGuardScfwFindings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAIGuardScfwFindings(recommendedAction AIGuardScfwOutcome) *AIGuardScfwFindings {
	this := AIGuardScfwFindings{}
	this.RecommendedAction = recommendedAction
	return &this
}

// NewAIGuardScfwFindingsWithDefaults instantiates a new AIGuardScfwFindings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAIGuardScfwFindingsWithDefaults() *AIGuardScfwFindings {
	this := AIGuardScfwFindings{}
	return &this
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *AIGuardScfwFindings) GetPackages() []AIGuardScfwPackageFinding {
	if o == nil || o.Packages == nil {
		var ret []AIGuardScfwPackageFinding
		return ret
	}
	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardScfwFindings) GetPackagesOk() (*[]AIGuardScfwPackageFinding, bool) {
	if o == nil || o.Packages == nil {
		return nil, false
	}
	return &o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *AIGuardScfwFindings) HasPackages() bool {
	return o != nil && o.Packages != nil
}

// SetPackages gets a reference to the given []AIGuardScfwPackageFinding and assigns it to the Packages field.
func (o *AIGuardScfwFindings) SetPackages(v []AIGuardScfwPackageFinding) {
	o.Packages = v
}

// GetRecommendedAction returns the RecommendedAction field value.
func (o *AIGuardScfwFindings) GetRecommendedAction() AIGuardScfwOutcome {
	if o == nil {
		var ret AIGuardScfwOutcome
		return ret
	}
	return o.RecommendedAction
}

// GetRecommendedActionOk returns a tuple with the RecommendedAction field value
// and a boolean to check if the value has been set.
func (o *AIGuardScfwFindings) GetRecommendedActionOk() (*AIGuardScfwOutcome, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendedAction, true
}

// SetRecommendedAction sets field value.
func (o *AIGuardScfwFindings) SetRecommendedAction(v AIGuardScfwOutcome) {
	o.RecommendedAction = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AIGuardScfwFindings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Packages != nil {
		toSerialize["packages"] = o.Packages
	}
	toSerialize["recommended_action"] = o.RecommendedAction

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AIGuardScfwFindings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Packages          []AIGuardScfwPackageFinding `json:"packages,omitempty"`
		RecommendedAction *AIGuardScfwOutcome         `json:"recommended_action"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RecommendedAction == nil {
		return fmt.Errorf("required field recommended_action missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"packages", "recommended_action"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Packages = all.Packages
	if !all.RecommendedAction.IsValid() {
		hasInvalidField = true
	} else {
		o.RecommendedAction = *all.RecommendedAction
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
