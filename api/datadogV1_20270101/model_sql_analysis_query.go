// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SqlAnalysisQuery Test-only generated schema for the dashboard versioning proof of concept.
type SqlAnalysisQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	SqlQuery string `json:"sql_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type SqlAnalysisQueryType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSqlAnalysisQuery instantiates a new SqlAnalysisQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSqlAnalysisQuery(sqlQuery string, typeVar SqlAnalysisQueryType) *SqlAnalysisQuery {
	this := SqlAnalysisQuery{}
	this.SqlQuery = sqlQuery
	this.Type = typeVar
	return &this
}

// NewSqlAnalysisQueryWithDefaults instantiates a new SqlAnalysisQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSqlAnalysisQueryWithDefaults() *SqlAnalysisQuery {
	this := SqlAnalysisQuery{}
	return &this
}

// GetSqlQuery returns the SqlQuery field value.
func (o *SqlAnalysisQuery) GetSqlQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SqlQuery
}

// GetSqlQueryOk returns a tuple with the SqlQuery field value
// and a boolean to check if the value has been set.
func (o *SqlAnalysisQuery) GetSqlQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SqlQuery, true
}

// SetSqlQuery sets field value.
func (o *SqlAnalysisQuery) SetSqlQuery(v string) {
	o.SqlQuery = v
}

// GetType returns the Type field value.
func (o *SqlAnalysisQuery) GetType() SqlAnalysisQueryType {
	if o == nil {
		var ret SqlAnalysisQueryType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SqlAnalysisQuery) GetTypeOk() (*SqlAnalysisQueryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SqlAnalysisQuery) SetType(v SqlAnalysisQueryType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SqlAnalysisQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["sql_query"] = o.SqlQuery
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SqlAnalysisQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SqlQuery *string               `json:"sql_query"`
		Type     *SqlAnalysisQueryType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SqlQuery == nil {
		return fmt.Errorf("required field sql_query missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.SqlQuery = *all.SqlQuery
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
