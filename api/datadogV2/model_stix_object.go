// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// STIXObject A STIX 2.1 object. Indicator objects are processed and ingested; other STIX object types are accepted and counted in the `unsupported` response counter.
type STIXObject struct {
	// The confidence in the correctness of the indicator, from 0 through 100.
	Confidence *int32 `json:"confidence,omitempty"`
	// The time when the object was created.
	Created *time.Time `json:"created,omitempty"`
	// Optional external reference metadata preserved with the indicator but not interpreted during ingestion.
	ExternalReferences []map[string]interface{} `json:"external_references,omitempty"`
	// The STIX object identifier.
	Id string `json:"id"`
	// The open vocabulary terms that categorize the indicator.
	IndicatorTypes []string `json:"indicator_types,omitempty"`
	// Optional kill chain metadata preserved with the indicator but not interpreted during ingestion.
	KillChainPhases []map[string]interface{} `json:"kill_chain_phases,omitempty"`
	// Labels associated with the indicator.
	Labels []string `json:"labels,omitempty"`
	// The time when the object was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// References to marking definition objects that apply to the indicator.
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	// The STIX pattern that identifies the observable. Present on indicator objects.
	Pattern *string `json:"pattern,omitempty"`
	// The supported STIX pattern language.
	PatternType *STIXPatternType `json:"pattern_type,omitempty"`
	// Whether the indicator has been revoked.
	Revoked *bool `json:"revoked,omitempty"`
	// The supported STIX specification version.
	SpecVersion *STIXSpecVersion `json:"spec_version,omitempty"`
	// The STIX object type.
	Type string `json:"type"`
	// The time from which the indicator is considered valid. Present on indicator objects.
	ValidFrom *time.Time `json:"valid_from,omitempty"`
	// The time until which the indicator is considered valid.
	ValidUntil *time.Time `json:"valid_until,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSTIXObject instantiates a new STIXObject object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSTIXObject(id string, typeVar string) *STIXObject {
	this := STIXObject{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewSTIXObjectWithDefaults instantiates a new STIXObject object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSTIXObjectWithDefaults() *STIXObject {
	this := STIXObject{}
	return &this
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *STIXObject) GetConfidence() int32 {
	if o == nil || o.Confidence == nil {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXObject) GetConfidenceOk() (*int32, bool) {
	if o == nil || o.Confidence == nil {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *STIXObject) HasConfidence() bool {
	return o != nil && o.Confidence != nil
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *STIXObject) SetConfidence(v int32) {
	o.Confidence = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *STIXObject) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXObject) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *STIXObject) HasCreated() bool {
	return o != nil && o.Created != nil
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *STIXObject) SetCreated(v time.Time) {
	o.Created = &v
}

// GetExternalReferences returns the ExternalReferences field value if set, zero value otherwise.
func (o *STIXObject) GetExternalReferences() []map[string]interface{} {
	if o == nil || o.ExternalReferences == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.ExternalReferences
}

// GetExternalReferencesOk returns a tuple with the ExternalReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXObject) GetExternalReferencesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.ExternalReferences == nil {
		return nil, false
	}
	return &o.ExternalReferences, true
}

// HasExternalReferences returns a boolean if a field has been set.
func (o *STIXObject) HasExternalReferences() bool {
	return o != nil && o.ExternalReferences != nil
}

// SetExternalReferences gets a reference to the given []map[string]interface{} and assigns it to the ExternalReferences field.
func (o *STIXObject) SetExternalReferences(v []map[string]interface{}) {
	o.ExternalReferences = v
}

// GetId returns the Id field value.
func (o *STIXObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *STIXObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *STIXObject) SetId(v string) {
	o.Id = v
}

// GetIndicatorTypes returns the IndicatorTypes field value if set, zero value otherwise.
func (o *STIXObject) GetIndicatorTypes() []string {
	if o == nil || o.IndicatorTypes == nil {
		var ret []string
		return ret
	}
	return o.IndicatorTypes
}

// GetIndicatorTypesOk returns a tuple with the IndicatorTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXObject) GetIndicatorTypesOk() (*[]string, bool) {
	if o == nil || o.IndicatorTypes == nil {
		return nil, false
	}
	return &o.IndicatorTypes, true
}

// HasIndicatorTypes returns a boolean if a field has been set.
func (o *STIXObject) HasIndicatorTypes() bool {
	return o != nil && o.IndicatorTypes != nil
}

// SetIndicatorTypes gets a reference to the given []string and assigns it to the IndicatorTypes field.
func (o *STIXObject) SetIndicatorTypes(v []string) {
	o.IndicatorTypes = v
}

// GetKillChainPhases returns the KillChainPhases field value if set, zero value otherwise.
func (o *STIXObject) GetKillChainPhases() []map[string]interface{} {
	if o == nil || o.KillChainPhases == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.KillChainPhases
}

// GetKillChainPhasesOk returns a tuple with the KillChainPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXObject) GetKillChainPhasesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.KillChainPhases == nil {
		return nil, false
	}
	return &o.KillChainPhases, true
}

// HasKillChainPhases returns a boolean if a field has been set.
func (o *STIXObject) HasKillChainPhases() bool {
	return o != nil && o.KillChainPhases != nil
}

// SetKillChainPhases gets a reference to the given []map[string]interface{} and assigns it to the KillChainPhases field.
func (o *STIXObject) SetKillChainPhases(v []map[string]interface{}) {
	o.KillChainPhases = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *STIXObject) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXObject) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *STIXObject) HasLabels() bool {
	return o != nil && o.Labels != nil
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *STIXObject) SetLabels(v []string) {
	o.Labels = v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *STIXObject) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXObject) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *STIXObject) HasModified() bool {
	return o != nil && o.Modified != nil
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *STIXObject) SetModified(v time.Time) {
	o.Modified = &v
}

// GetObjectMarkingRefs returns the ObjectMarkingRefs field value if set, zero value otherwise.
func (o *STIXObject) GetObjectMarkingRefs() []string {
	if o == nil || o.ObjectMarkingRefs == nil {
		var ret []string
		return ret
	}
	return o.ObjectMarkingRefs
}

// GetObjectMarkingRefsOk returns a tuple with the ObjectMarkingRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXObject) GetObjectMarkingRefsOk() (*[]string, bool) {
	if o == nil || o.ObjectMarkingRefs == nil {
		return nil, false
	}
	return &o.ObjectMarkingRefs, true
}

// HasObjectMarkingRefs returns a boolean if a field has been set.
func (o *STIXObject) HasObjectMarkingRefs() bool {
	return o != nil && o.ObjectMarkingRefs != nil
}

// SetObjectMarkingRefs gets a reference to the given []string and assigns it to the ObjectMarkingRefs field.
func (o *STIXObject) SetObjectMarkingRefs(v []string) {
	o.ObjectMarkingRefs = v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *STIXObject) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXObject) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *STIXObject) HasPattern() bool {
	return o != nil && o.Pattern != nil
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *STIXObject) SetPattern(v string) {
	o.Pattern = &v
}

// GetPatternType returns the PatternType field value if set, zero value otherwise.
func (o *STIXObject) GetPatternType() STIXPatternType {
	if o == nil || o.PatternType == nil {
		var ret STIXPatternType
		return ret
	}
	return *o.PatternType
}

// GetPatternTypeOk returns a tuple with the PatternType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXObject) GetPatternTypeOk() (*STIXPatternType, bool) {
	if o == nil || o.PatternType == nil {
		return nil, false
	}
	return o.PatternType, true
}

// HasPatternType returns a boolean if a field has been set.
func (o *STIXObject) HasPatternType() bool {
	return o != nil && o.PatternType != nil
}

// SetPatternType gets a reference to the given STIXPatternType and assigns it to the PatternType field.
func (o *STIXObject) SetPatternType(v STIXPatternType) {
	o.PatternType = &v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise.
func (o *STIXObject) GetRevoked() bool {
	if o == nil || o.Revoked == nil {
		var ret bool
		return ret
	}
	return *o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXObject) GetRevokedOk() (*bool, bool) {
	if o == nil || o.Revoked == nil {
		return nil, false
	}
	return o.Revoked, true
}

// HasRevoked returns a boolean if a field has been set.
func (o *STIXObject) HasRevoked() bool {
	return o != nil && o.Revoked != nil
}

// SetRevoked gets a reference to the given bool and assigns it to the Revoked field.
func (o *STIXObject) SetRevoked(v bool) {
	o.Revoked = &v
}

// GetSpecVersion returns the SpecVersion field value if set, zero value otherwise.
func (o *STIXObject) GetSpecVersion() STIXSpecVersion {
	if o == nil || o.SpecVersion == nil {
		var ret STIXSpecVersion
		return ret
	}
	return *o.SpecVersion
}

// GetSpecVersionOk returns a tuple with the SpecVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXObject) GetSpecVersionOk() (*STIXSpecVersion, bool) {
	if o == nil || o.SpecVersion == nil {
		return nil, false
	}
	return o.SpecVersion, true
}

// HasSpecVersion returns a boolean if a field has been set.
func (o *STIXObject) HasSpecVersion() bool {
	return o != nil && o.SpecVersion != nil
}

// SetSpecVersion gets a reference to the given STIXSpecVersion and assigns it to the SpecVersion field.
func (o *STIXObject) SetSpecVersion(v STIXSpecVersion) {
	o.SpecVersion = &v
}

// GetType returns the Type field value.
func (o *STIXObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *STIXObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *STIXObject) SetType(v string) {
	o.Type = v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *STIXObject) GetValidFrom() time.Time {
	if o == nil || o.ValidFrom == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXObject) GetValidFromOk() (*time.Time, bool) {
	if o == nil || o.ValidFrom == nil {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *STIXObject) HasValidFrom() bool {
	return o != nil && o.ValidFrom != nil
}

// SetValidFrom gets a reference to the given time.Time and assigns it to the ValidFrom field.
func (o *STIXObject) SetValidFrom(v time.Time) {
	o.ValidFrom = &v
}

// GetValidUntil returns the ValidUntil field value if set, zero value otherwise.
func (o *STIXObject) GetValidUntil() time.Time {
	if o == nil || o.ValidUntil == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STIXObject) GetValidUntilOk() (*time.Time, bool) {
	if o == nil || o.ValidUntil == nil {
		return nil, false
	}
	return o.ValidUntil, true
}

// HasValidUntil returns a boolean if a field has been set.
func (o *STIXObject) HasValidUntil() bool {
	return o != nil && o.ValidUntil != nil
}

// SetValidUntil gets a reference to the given time.Time and assigns it to the ValidUntil field.
func (o *STIXObject) SetValidUntil(v time.Time) {
	o.ValidUntil = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o STIXObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Confidence != nil {
		toSerialize["confidence"] = o.Confidence
	}
	if o.Created != nil {
		if o.Created.Nanosecond() == 0 {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
		}
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
	if o.Modified != nil {
		if o.Modified.Nanosecond() == 0 {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ObjectMarkingRefs != nil {
		toSerialize["object_marking_refs"] = o.ObjectMarkingRefs
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.PatternType != nil {
		toSerialize["pattern_type"] = o.PatternType
	}
	if o.Revoked != nil {
		toSerialize["revoked"] = o.Revoked
	}
	if o.SpecVersion != nil {
		toSerialize["spec_version"] = o.SpecVersion
	}
	toSerialize["type"] = o.Type
	if o.ValidFrom != nil {
		if o.ValidFrom.Nanosecond() == 0 {
			toSerialize["valid_from"] = o.ValidFrom.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["valid_from"] = o.ValidFrom.Format("2006-01-02T15:04:05.000Z07:00")
		}
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
func (o *STIXObject) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Confidence         *int32                   `json:"confidence,omitempty"`
		Created            *time.Time               `json:"created,omitempty"`
		ExternalReferences []map[string]interface{} `json:"external_references,omitempty"`
		Id                 *string                  `json:"id"`
		IndicatorTypes     []string                 `json:"indicator_types,omitempty"`
		KillChainPhases    []map[string]interface{} `json:"kill_chain_phases,omitempty"`
		Labels             []string                 `json:"labels,omitempty"`
		Modified           *time.Time               `json:"modified,omitempty"`
		ObjectMarkingRefs  []string                 `json:"object_marking_refs,omitempty"`
		Pattern            *string                  `json:"pattern,omitempty"`
		PatternType        *STIXPatternType         `json:"pattern_type,omitempty"`
		Revoked            *bool                    `json:"revoked,omitempty"`
		SpecVersion        *STIXSpecVersion         `json:"spec_version,omitempty"`
		Type               *string                  `json:"type"`
		ValidFrom          *time.Time               `json:"valid_from,omitempty"`
		ValidUntil         *time.Time               `json:"valid_until,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"confidence", "created", "external_references", "id", "indicator_types", "kill_chain_phases", "labels", "modified", "object_marking_refs", "pattern", "pattern_type", "revoked", "spec_version", "type", "valid_from", "valid_until"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Confidence = all.Confidence
	o.Created = all.Created
	o.ExternalReferences = all.ExternalReferences
	o.Id = *all.Id
	o.IndicatorTypes = all.IndicatorTypes
	o.KillChainPhases = all.KillChainPhases
	o.Labels = all.Labels
	o.Modified = all.Modified
	o.ObjectMarkingRefs = all.ObjectMarkingRefs
	o.Pattern = all.Pattern
	if all.PatternType != nil && !all.PatternType.IsValid() {
		hasInvalidField = true
	} else {
		o.PatternType = all.PatternType
	}
	o.Revoked = all.Revoked
	if all.SpecVersion != nil && !all.SpecVersion.IsValid() {
		hasInvalidField = true
	} else {
		o.SpecVersion = all.SpecVersion
	}
	o.Type = *all.Type
	o.ValidFrom = all.ValidFrom
	o.ValidUntil = all.ValidUntil

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
