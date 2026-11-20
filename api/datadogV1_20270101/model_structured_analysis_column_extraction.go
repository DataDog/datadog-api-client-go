// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisColumnExtraction Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisColumnExtraction struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	GrokExpression string `json:"grok_expression"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	KeepNulls *bool `json:"keep_nulls,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SourceColumn string `json:"source_column"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type StructuredAnalysisColumnExtractionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewStructuredAnalysisColumnExtraction instantiates a new StructuredAnalysisColumnExtraction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStructuredAnalysisColumnExtraction(grokExpression string, sourceColumn string, typeVar StructuredAnalysisColumnExtractionType) *StructuredAnalysisColumnExtraction {
	this := StructuredAnalysisColumnExtraction{}
	this.GrokExpression = grokExpression
	this.SourceColumn = sourceColumn
	this.Type = typeVar
	return &this
}

// NewStructuredAnalysisColumnExtractionWithDefaults instantiates a new StructuredAnalysisColumnExtraction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStructuredAnalysisColumnExtractionWithDefaults() *StructuredAnalysisColumnExtraction {
	this := StructuredAnalysisColumnExtraction{}
	return &this
}

// GetGrokExpression returns the GrokExpression field value.
func (o *StructuredAnalysisColumnExtraction) GetGrokExpression() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GrokExpression
}

// GetGrokExpressionOk returns a tuple with the GrokExpression field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisColumnExtraction) GetGrokExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrokExpression, true
}

// SetGrokExpression sets field value.
func (o *StructuredAnalysisColumnExtraction) SetGrokExpression(v string) {
	o.GrokExpression = v
}

// GetKeepNulls returns the KeepNulls field value if set, zero value otherwise.
func (o *StructuredAnalysisColumnExtraction) GetKeepNulls() bool {
	if o == nil || o.KeepNulls == nil {
		var ret bool
		return ret
	}
	return *o.KeepNulls
}

// GetKeepNullsOk returns a tuple with the KeepNulls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisColumnExtraction) GetKeepNullsOk() (*bool, bool) {
	if o == nil || o.KeepNulls == nil {
		return nil, false
	}
	return o.KeepNulls, true
}

// HasKeepNulls returns a boolean if a field has been set.
func (o *StructuredAnalysisColumnExtraction) HasKeepNulls() bool {
	return o != nil && o.KeepNulls != nil
}

// SetKeepNulls gets a reference to the given bool and assigns it to the KeepNulls field.
func (o *StructuredAnalysisColumnExtraction) SetKeepNulls(v bool) {
	o.KeepNulls = &v
}

// GetSourceColumn returns the SourceColumn field value.
func (o *StructuredAnalysisColumnExtraction) GetSourceColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SourceColumn
}

// GetSourceColumnOk returns a tuple with the SourceColumn field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisColumnExtraction) GetSourceColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceColumn, true
}

// SetSourceColumn sets field value.
func (o *StructuredAnalysisColumnExtraction) SetSourceColumn(v string) {
	o.SourceColumn = v
}

// GetType returns the Type field value.
func (o *StructuredAnalysisColumnExtraction) GetType() StructuredAnalysisColumnExtractionType {
	if o == nil {
		var ret StructuredAnalysisColumnExtractionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisColumnExtraction) GetTypeOk() (*StructuredAnalysisColumnExtractionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StructuredAnalysisColumnExtraction) SetType(v StructuredAnalysisColumnExtractionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StructuredAnalysisColumnExtraction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["grok_expression"] = o.GrokExpression
	if o.KeepNulls != nil {
		toSerialize["keep_nulls"] = o.KeepNulls
	}
	toSerialize["source_column"] = o.SourceColumn
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StructuredAnalysisColumnExtraction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GrokExpression *string                                 `json:"grok_expression"`
		KeepNulls      *bool                                   `json:"keep_nulls,omitempty"`
		SourceColumn   *string                                 `json:"source_column"`
		Type           *StructuredAnalysisColumnExtractionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.GrokExpression == nil {
		return fmt.Errorf("required field grok_expression missing")
	}
	if all.SourceColumn == nil {
		return fmt.Errorf("required field source_column missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.GrokExpression = *all.GrokExpression
	o.KeepNulls = all.KeepNulls
	o.SourceColumn = *all.SourceColumn
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
