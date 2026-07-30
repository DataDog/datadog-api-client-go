// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// STIXIndicatorObject A STIX 2.1 indicator object.
type STIXIndicatorObject struct {
	// The confidence in the correctness of the indicator, from 0 through 100.
	Confidence *int32 `json:"confidence,omitempty"`
	// The time when the indicator was created.
	Created time.Time `json:"created"`
	// Optional external reference metadata preserved with the indicator but not interpreted during ingestion.
	ExternalReferences []map[string]interface{} `json:"external_references,omitempty"`
	// The STIX indicator identifier.
	Id string `json:"id"`
	// The open vocabulary terms that categorize the indicator.
	IndicatorTypes []string `json:"indicator_types,omitempty"`
	// Optional kill chain metadata preserved with the indicator but not interpreted during ingestion.
	KillChainPhases []map[string]interface{} `json:"kill_chain_phases,omitempty"`
	// Labels associated with the indicator.
	Labels []string `json:"labels,omitempty"`
	// The time when the indicator was last modified.
	Modified time.Time `json:"modified"`
	// References to marking definition objects that apply to the indicator.
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	// The STIX pattern that identifies the observable.
	Pattern string `json:"pattern"`
	// The supported STIX pattern language.
	PatternType STIXPatternType `json:"pattern_type"`
	// Whether the indicator has been revoked.
	Revoked *bool `json:"revoked,omitempty"`
	// The supported STIX specification version.
	SpecVersion STIXSpecVersion `json:"spec_version"`
	// The STIX object type for an indicator.
	Type STIXIndicatorType `json:"type"`
	// The time from which the indicator is considered valid.
	ValidFrom time.Time `json:"valid_from"`
	// The time until which the indicator is considered valid.
	ValidUntil *time.Time `json:"valid_until,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSTIXIndicatorObject instantiates a new STIXIndicatorObject object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSTIXIndicatorObject(created time.Time, id string, modified time.Time, pattern string, patternType STIXPatternType, specVersion STIXSpecVersion, typeVar STIXIndicatorType, validFrom time.Time) *STIXIndicatorObject {
	this := STIXIndicatorObject{}
	this.Created = created
	this.Id = id
	this.Modified = modified
	this.Pattern = pattern
	this.PatternType = patternType
	this.SpecVersion = specVersion
	this.Type = typeVar
	this.ValidFrom = validFrom
	return &this
}

// NewSTIXIndicatorObjectWithDefaults instantiates a new STIXIndicatorObject object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSTIXIndicatorObjectWithDefaults() *STIXIndicatorObject {
	this := STIXIndicatorObject{}
	return &this
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *STIXIndicatorObject) GetConfidence() int32 {
	if o == nil || o.Confidence == nil {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXIndicatorObject) GetConfidenceOk() (*int32, bool) {
	if o == nil || o.Confidence == nil {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *STIXIndicatorObject) HasConfidence() bool {
	return o != nil && o.Confidence != nil
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *STIXIndicatorObject) SetConfidence(v int32) {
	o.Confidence = &v
}

// GetCreated returns the Created field value.
func (o *STIXIndicatorObject) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *STIXIndicatorObject) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value.
func (o *STIXIndicatorObject) SetCreated(v time.Time) {
	o.Created = v
}

// GetExternalReferences returns the ExternalReferences field value if set, zero value otherwise.
func (o *STIXIndicatorObject) GetExternalReferences() []map[string]interface{} {
	if o == nil || o.ExternalReferences == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.ExternalReferences
}

// GetExternalReferencesOk returns a tuple with the ExternalReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXIndicatorObject) GetExternalReferencesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.ExternalReferences == nil {
		return nil, false
	}
	return &o.ExternalReferences, true
}

// HasExternalReferences returns a boolean if a field has been set.
func (o *STIXIndicatorObject) HasExternalReferences() bool {
	return o != nil && o.ExternalReferences != nil
}

// SetExternalReferences gets a reference to the given []map[string]interface{} and assigns it to the ExternalReferences field.
func (o *STIXIndicatorObject) SetExternalReferences(v []map[string]interface{}) {
	o.ExternalReferences = v
}

// GetId returns the Id field value.
func (o *STIXIndicatorObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *STIXIndicatorObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *STIXIndicatorObject) SetId(v string) {
	o.Id = v
}

// GetIndicatorTypes returns the IndicatorTypes field value if set, zero value otherwise.
func (o *STIXIndicatorObject) GetIndicatorTypes() []string {
	if o == nil || o.IndicatorTypes == nil {
		var ret []string
		return ret
	}
	return o.IndicatorTypes
}

// GetIndicatorTypesOk returns a tuple with the IndicatorTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXIndicatorObject) GetIndicatorTypesOk() (*[]string, bool) {
	if o == nil || o.IndicatorTypes == nil {
		return nil, false
	}
	return &o.IndicatorTypes, true
}

// HasIndicatorTypes returns a boolean if a field has been set.
func (o *STIXIndicatorObject) HasIndicatorTypes() bool {
	return o != nil && o.IndicatorTypes != nil
}

// SetIndicatorTypes gets a reference to the given []string and assigns it to the IndicatorTypes field.
func (o *STIXIndicatorObject) SetIndicatorTypes(v []string) {
	o.IndicatorTypes = v
}

// GetKillChainPhases returns the KillChainPhases field value if set, zero value otherwise.
func (o *STIXIndicatorObject) GetKillChainPhases() []map[string]interface{} {
	if o == nil || o.KillChainPhases == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.KillChainPhases
}

// GetKillChainPhasesOk returns a tuple with the KillChainPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXIndicatorObject) GetKillChainPhasesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.KillChainPhases == nil {
		return nil, false
	}
	return &o.KillChainPhases, true
}

// HasKillChainPhases returns a boolean if a field has been set.
func (o *STIXIndicatorObject) HasKillChainPhases() bool {
	return o != nil && o.KillChainPhases != nil
}

// SetKillChainPhases gets a reference to the given []map[string]interface{} and assigns it to the KillChainPhases field.
func (o *STIXIndicatorObject) SetKillChainPhases(v []map[string]interface{}) {
	o.KillChainPhases = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *STIXIndicatorObject) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXIndicatorObject) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *STIXIndicatorObject) HasLabels() bool {
	return o != nil && o.Labels != nil
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *STIXIndicatorObject) SetLabels(v []string) {
	o.Labels = v
}

// GetModified returns the Modified field value.
func (o *STIXIndicatorObject) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *STIXIndicatorObject) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value.
func (o *STIXIndicatorObject) SetModified(v time.Time) {
	o.Modified = v
}

// GetObjectMarkingRefs returns the ObjectMarkingRefs field value if set, zero value otherwise.
func (o *STIXIndicatorObject) GetObjectMarkingRefs() []string {
	if o == nil || o.ObjectMarkingRefs == nil {
		var ret []string
		return ret
	}
	return o.ObjectMarkingRefs
}

// GetObjectMarkingRefsOk returns a tuple with the ObjectMarkingRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXIndicatorObject) GetObjectMarkingRefsOk() (*[]string, bool) {
	if o == nil || o.ObjectMarkingRefs == nil {
		return nil, false
	}
	return &o.ObjectMarkingRefs, true
}

// HasObjectMarkingRefs returns a boolean if a field has been set.
func (o *STIXIndicatorObject) HasObjectMarkingRefs() bool {
	return o != nil && o.ObjectMarkingRefs != nil
}

// SetObjectMarkingRefs gets a reference to the given []string and assigns it to the ObjectMarkingRefs field.
func (o *STIXIndicatorObject) SetObjectMarkingRefs(v []string) {
	o.ObjectMarkingRefs = v
}

// GetPattern returns the Pattern field value.
func (o *STIXIndicatorObject) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *STIXIndicatorObject) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value.
func (o *STIXIndicatorObject) SetPattern(v string) {
	o.Pattern = v
}

// GetPatternType returns the PatternType field value.
func (o *STIXIndicatorObject) GetPatternType() STIXPatternType {
	if o == nil {
		var ret STIXPatternType
		return ret
	}
	return o.PatternType
}

// GetPatternTypeOk returns a tuple with the PatternType field value
// and a boolean to check if the value has been set.
func (o *STIXIndicatorObject) GetPatternTypeOk() (*STIXPatternType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PatternType, true
}

// SetPatternType sets field value.
func (o *STIXIndicatorObject) SetPatternType(v STIXPatternType) {
	o.PatternType = v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise.
func (o *STIXIndicatorObject) GetRevoked() bool {
	if o == nil || o.Revoked == nil {
		var ret bool
		return ret
	}
	return *o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXIndicatorObject) GetRevokedOk() (*bool, bool) {
	if o == nil || o.Revoked == nil {
		return nil, false
	}
	return o.Revoked, true
}

// HasRevoked returns a boolean if a field has been set.
func (o *STIXIndicatorObject) HasRevoked() bool {
	return o != nil && o.Revoked != nil
}

// SetRevoked gets a reference to the given bool and assigns it to the Revoked field.
func (o *STIXIndicatorObject) SetRevoked(v bool) {
	o.Revoked = &v
}

// GetSpecVersion returns the SpecVersion field value.
func (o *STIXIndicatorObject) GetSpecVersion() STIXSpecVersion {
	if o == nil {
		var ret STIXSpecVersion
		return ret
	}
	return o.SpecVersion
}

// GetSpecVersionOk returns a tuple with the SpecVersion field value
// and a boolean to check if the value has been set.
func (o *STIXIndicatorObject) GetSpecVersionOk() (*STIXSpecVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpecVersion, true
}

// SetSpecVersion sets field value.
func (o *STIXIndicatorObject) SetSpecVersion(v STIXSpecVersion) {
	o.SpecVersion = v
}

// GetType returns the Type field value.
func (o *STIXIndicatorObject) GetType() STIXIndicatorType {
	if o == nil {
		var ret STIXIndicatorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *STIXIndicatorObject) GetTypeOk() (*STIXIndicatorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *STIXIndicatorObject) SetType(v STIXIndicatorType) {
	o.Type = v
}

// GetValidFrom returns the ValidFrom field value.
func (o *STIXIndicatorObject) GetValidFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value
// and a boolean to check if the value has been set.
func (o *STIXIndicatorObject) GetValidFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidFrom, true
}

// SetValidFrom sets field value.
func (o *STIXIndicatorObject) SetValidFrom(v time.Time) {
	o.ValidFrom = v
}

// GetValidUntil returns the ValidUntil field value if set, zero value otherwise.
func (o *STIXIndicatorObject) GetValidUntil() time.Time {
	if o == nil || o.ValidUntil == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXIndicatorObject) GetValidUntilOk() (*time.Time, bool) {
	if o == nil || o.ValidUntil == nil {
		return nil, false
	}
	return o.ValidUntil, true
}

// HasValidUntil returns a boolean if a field has been set.
func (o *STIXIndicatorObject) HasValidUntil() bool {
	return o != nil && o.ValidUntil != nil
}

// SetValidUntil gets a reference to the given time.Time and assigns it to the ValidUntil field.
func (o *STIXIndicatorObject) SetValidUntil(v time.Time) {
	o.ValidUntil = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o STIXIndicatorObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Confidence != nil {
		toSerialize["confidence"] = o.Confidence
	}
	if o.Created.Nanosecond() == 0 {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ExternalReferences != nil {
		toSerialize["external_references"] = o.ExternalReferences
	}
	toSerialize["id"] = o.Id
	if o.IndicatorTypes != nil {
		toSerialize["indicator_types"] = o.IndicatorTypes
	}
	if o.KillChainPhases != nil {
		toSerialize["kill_chain_phases"] = o.KillChainPhases
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Modified.Nanosecond() == 0 {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ObjectMarkingRefs != nil {
		toSerialize["object_marking_refs"] = o.ObjectMarkingRefs
	}
	toSerialize["pattern"] = o.Pattern
	toSerialize["pattern_type"] = o.PatternType
	if o.Revoked != nil {
		toSerialize["revoked"] = o.Revoked
	}
	toSerialize["spec_version"] = o.SpecVersion
	toSerialize["type"] = o.Type
	if o.ValidFrom.Nanosecond() == 0 {
		toSerialize["valid_from"] = o.ValidFrom.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["valid_from"] = o.ValidFrom.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ValidUntil != nil {
		if o.ValidUntil.Nanosecond() == 0 {
			toSerialize["valid_until"] = o.ValidUntil.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["valid_until"] = o.ValidUntil.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *STIXIndicatorObject) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Confidence         *int32                   `json:"confidence,omitempty"`
		Created            *time.Time               `json:"created"`
		ExternalReferences []map[string]interface{} `json:"external_references,omitempty"`
		Id                 *string                  `json:"id"`
		IndicatorTypes     []string                 `json:"indicator_types,omitempty"`
		KillChainPhases    []map[string]interface{} `json:"kill_chain_phases,omitempty"`
		Labels             []string                 `json:"labels,omitempty"`
		Modified           *time.Time               `json:"modified"`
		ObjectMarkingRefs  []string                 `json:"object_marking_refs,omitempty"`
		Pattern            *string                  `json:"pattern"`
		PatternType        *STIXPatternType         `json:"pattern_type"`
		Revoked            *bool                    `json:"revoked,omitempty"`
		SpecVersion        *STIXSpecVersion         `json:"spec_version"`
		Type               *STIXIndicatorType       `json:"type"`
		ValidFrom          *time.Time               `json:"valid_from"`
		ValidUntil         *time.Time               `json:"valid_until,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Created == nil {
		return fmt.Errorf("required field created missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Modified == nil {
		return fmt.Errorf("required field modified missing")
	}
	if all.Pattern == nil {
		return fmt.Errorf("required field pattern missing")
	}
	if all.PatternType == nil {
		return fmt.Errorf("required field pattern_type missing")
	}
	if all.SpecVersion == nil {
		return fmt.Errorf("required field spec_version missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.ValidFrom == nil {
		return fmt.Errorf("required field valid_from missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"confidence", "created", "external_references", "id", "indicator_types", "kill_chain_phases", "labels", "modified", "object_marking_refs", "pattern", "pattern_type", "revoked", "spec_version", "type", "valid_from", "valid_until"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Confidence = all.Confidence
	o.Created = *all.Created
	o.ExternalReferences = all.ExternalReferences
	o.Id = *all.Id
	o.IndicatorTypes = all.IndicatorTypes
	o.KillChainPhases = all.KillChainPhases
	o.Labels = all.Labels
	o.Modified = *all.Modified
	o.ObjectMarkingRefs = all.ObjectMarkingRefs
	o.Pattern = *all.Pattern
	if !all.PatternType.IsValid() {
		hasInvalidField = true
	} else {
		o.PatternType = *all.PatternType
	}
	o.Revoked = all.Revoked
	if !all.SpecVersion.IsValid() {
		hasInvalidField = true
	} else {
		o.SpecVersion = *all.SpecVersion
	}
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.ValidFrom = *all.ValidFrom
	o.ValidUntil = all.ValidUntil

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
