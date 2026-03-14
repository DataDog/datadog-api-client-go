// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableRequest Request definition for the guided table widget — a structured, guided format for defining a table of data. Used when the `graphing_new_table_widget_editor` feature flag is enabled; otherwise use the classic `TableWidgetRequest`. Defines base queries, how to produce rows, and how to compute columns.
type GuidedTableRequest struct {
	// Column definitions that describe how to compute and display table values.
	Columns []GuidedTableColumn `json:"columns"`
	// Base queries that provide the source data for the table.
	Queries []GuidedTableQuery `json:"queries"`
	// Identifies this as a guided table request.
	RequestType GuidedTableRequestRequestType `json:"request_type"`
	// Defines how to compute the rows of a guided table.
	Rows GuidedTableRows `json:"rows"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGuidedTableRequest instantiates a new GuidedTableRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGuidedTableRequest(columns []GuidedTableColumn, queries []GuidedTableQuery, requestType GuidedTableRequestRequestType, rows GuidedTableRows) *GuidedTableRequest {
	this := GuidedTableRequest{}
	this.Columns = columns
	this.Queries = queries
	this.RequestType = requestType
	this.Rows = rows
	return &this
}

// NewGuidedTableRequestWithDefaults instantiates a new GuidedTableRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGuidedTableRequestWithDefaults() *GuidedTableRequest {
	this := GuidedTableRequest{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *GuidedTableRequest) GetColumns() []GuidedTableColumn {
	if o == nil {
		var ret []GuidedTableColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *GuidedTableRequest) GetColumnsOk() (*[]GuidedTableColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *GuidedTableRequest) SetColumns(v []GuidedTableColumn) {
	o.Columns = v
}

// GetQueries returns the Queries field value.
func (o *GuidedTableRequest) GetQueries() []GuidedTableQuery {
	if o == nil {
		var ret []GuidedTableQuery
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *GuidedTableRequest) GetQueriesOk() (*[]GuidedTableQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *GuidedTableRequest) SetQueries(v []GuidedTableQuery) {
	o.Queries = v
}

// GetRequestType returns the RequestType field value.
func (o *GuidedTableRequest) GetRequestType() GuidedTableRequestRequestType {
	if o == nil {
		var ret GuidedTableRequestRequestType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *GuidedTableRequest) GetRequestTypeOk() (*GuidedTableRequestRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *GuidedTableRequest) SetRequestType(v GuidedTableRequestRequestType) {
	o.RequestType = v
}

// GetRows returns the Rows field value.
func (o *GuidedTableRequest) GetRows() GuidedTableRows {
	if o == nil {
		var ret GuidedTableRows
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value
// and a boolean to check if the value has been set.
func (o *GuidedTableRequest) GetRowsOk() (*GuidedTableRows, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rows, true
}

// SetRows sets field value.
func (o *GuidedTableRequest) SetRows(v GuidedTableRows) {
	o.Rows = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GuidedTableRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["columns"] = o.Columns
	toSerialize["queries"] = o.Queries
	toSerialize["request_type"] = o.RequestType
	toSerialize["rows"] = o.Rows

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GuidedTableRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns     *[]GuidedTableColumn           `json:"columns"`
		Queries     *[]GuidedTableQuery            `json:"queries"`
		RequestType *GuidedTableRequestRequestType `json:"request_type"`
		Rows        *GuidedTableRows               `json:"rows"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Columns == nil {
		return fmt.Errorf("required field columns missing")
	}
	if all.Queries == nil {
		return fmt.Errorf("required field queries missing")
	}
	if all.RequestType == nil {
		return fmt.Errorf("required field request_type missing")
	}
	if all.Rows == nil {
		return fmt.Errorf("required field rows missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"columns", "queries", "request_type", "rows"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Columns = *all.Columns
	o.Queries = *all.Queries
	if !all.RequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.RequestType = *all.RequestType
	}
	if all.Rows.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rows = *all.Rows

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
