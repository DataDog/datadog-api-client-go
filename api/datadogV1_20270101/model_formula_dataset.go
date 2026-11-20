// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaDataset Test-only generated schema for the dashboard versioning proof of concept.
type FormulaDataset struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaDatasetDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Formulas []FormulaDatasetFormulasItems `json:"formulas"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Queries []DashboardInlineSchema40 `json:"queries"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat FormulaDatasetResponseFormat `json:"response_format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TimeWindow *TimeWindow `json:"time_window,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaDataset instantiates a new FormulaDataset object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaDataset(dataSource FormulaDatasetDataSource, formulas []FormulaDatasetFormulasItems, name string, queries []DashboardInlineSchema40, responseFormat FormulaDatasetResponseFormat) *FormulaDataset {
	this := FormulaDataset{}
	this.DataSource = dataSource
	this.Formulas = formulas
	this.Name = name
	this.Queries = queries
	this.ResponseFormat = responseFormat
	return &this
}

// NewFormulaDatasetWithDefaults instantiates a new FormulaDataset object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaDatasetWithDefaults() *FormulaDataset {
	this := FormulaDataset{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *FormulaDataset) GetDataSource() FormulaDatasetDataSource {
	if o == nil {
		var ret FormulaDatasetDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaDataset) GetDataSourceOk() (*FormulaDatasetDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaDataset) SetDataSource(v FormulaDatasetDataSource) {
	o.DataSource = v
}

// GetFormulas returns the Formulas field value.
func (o *FormulaDataset) GetFormulas() []FormulaDatasetFormulasItems {
	if o == nil {
		var ret []FormulaDatasetFormulasItems
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value
// and a boolean to check if the value has been set.
func (o *FormulaDataset) GetFormulasOk() (*[]FormulaDatasetFormulasItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// SetFormulas sets field value.
func (o *FormulaDataset) SetFormulas(v []FormulaDatasetFormulasItems) {
	o.Formulas = v
}

// GetName returns the Name field value.
func (o *FormulaDataset) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaDataset) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaDataset) SetName(v string) {
	o.Name = v
}

// GetQueries returns the Queries field value.
func (o *FormulaDataset) GetQueries() []DashboardInlineSchema40 {
	if o == nil {
		var ret []DashboardInlineSchema40
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *FormulaDataset) GetQueriesOk() (*[]DashboardInlineSchema40, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *FormulaDataset) SetQueries(v []DashboardInlineSchema40) {
	o.Queries = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *FormulaDataset) GetResponseFormat() FormulaDatasetResponseFormat {
	if o == nil {
		var ret FormulaDatasetResponseFormat
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *FormulaDataset) GetResponseFormatOk() (*FormulaDatasetResponseFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *FormulaDataset) SetResponseFormat(v FormulaDatasetResponseFormat) {
	o.ResponseFormat = v
}

// GetTimeWindow returns the TimeWindow field value if set, zero value otherwise.
func (o *FormulaDataset) GetTimeWindow() TimeWindow {
	if o == nil || o.TimeWindow == nil {
		var ret TimeWindow
		return ret
	}
	return *o.TimeWindow
}

// GetTimeWindowOk returns a tuple with the TimeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaDataset) GetTimeWindowOk() (*TimeWindow, bool) {
	if o == nil || o.TimeWindow == nil {
		return nil, false
	}
	return o.TimeWindow, true
}

// HasTimeWindow returns a boolean if a field has been set.
func (o *FormulaDataset) HasTimeWindow() bool {
	return o != nil && o.TimeWindow != nil
}

// SetTimeWindow gets a reference to the given TimeWindow and assigns it to the TimeWindow field.
func (o *FormulaDataset) SetTimeWindow(v TimeWindow) {
	o.TimeWindow = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaDataset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["formulas"] = o.Formulas
	toSerialize["name"] = o.Name
	toSerialize["queries"] = o.Queries
	toSerialize["response_format"] = o.ResponseFormat
	if o.TimeWindow != nil {
		toSerialize["time_window"] = o.TimeWindow
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaDataset) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource     *FormulaDatasetDataSource      `json:"data_source"`
		Formulas       *[]FormulaDatasetFormulasItems `json:"formulas"`
		Name           *string                        `json:"name"`
		Queries        *[]DashboardInlineSchema40     `json:"queries"`
		ResponseFormat *FormulaDatasetResponseFormat  `json:"response_format"`
		TimeWindow     *TimeWindow                    `json:"time_window,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Formulas == nil {
		return fmt.Errorf("required field formulas missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Queries == nil {
		return fmt.Errorf("required field queries missing")
	}
	if all.ResponseFormat == nil {
		return fmt.Errorf("required field response_format missing")
	}

	hasInvalidField := false
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Formulas = *all.Formulas
	o.Name = *all.Name
	o.Queries = *all.Queries
	if !all.ResponseFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.ResponseFormat = *all.ResponseFormat
	}
	if all.TimeWindow != nil && all.TimeWindow.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeWindow = all.TimeWindow

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
