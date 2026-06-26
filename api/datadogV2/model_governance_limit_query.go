// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceLimitQuery A metric query used to compute usage against a limit.
type GovernanceLimitQuery struct {
	// The metric query expression used to compute the limit value.
	Query string `json:"query"`
	// How the query results are aggregated into a single value (for example, sum, max, or avg).
	Reducer string `json:"reducer"`
	// The data source used to evaluate the metric query (for example, the metrics or events backend).
	Source string `json:"source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceLimitQuery instantiates a new GovernanceLimitQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceLimitQuery(query string, reducer string, source string) *GovernanceLimitQuery {
	this := GovernanceLimitQuery{}
	this.Query = query
	this.Reducer = reducer
	this.Source = source
	return &this
}

// NewGovernanceLimitQueryWithDefaults instantiates a new GovernanceLimitQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceLimitQueryWithDefaults() *GovernanceLimitQuery {
	this := GovernanceLimitQuery{}
	return &this
}

// GetQuery returns the Query field value.
func (o *GovernanceLimitQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *GovernanceLimitQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *GovernanceLimitQuery) SetQuery(v string) {
	o.Query = v
}

// GetReducer returns the Reducer field value.
func (o *GovernanceLimitQuery) GetReducer() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reducer
}

// GetReducerOk returns a tuple with the Reducer field value
// and a boolean to check if the value has been set.
func (o *GovernanceLimitQuery) GetReducerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reducer, true
}

// SetReducer sets field value.
func (o *GovernanceLimitQuery) SetReducer(v string) {
	o.Reducer = v
}

// GetSource returns the Source field value.
func (o *GovernanceLimitQuery) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *GovernanceLimitQuery) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *GovernanceLimitQuery) SetSource(v string) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceLimitQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["query"] = o.Query
	toSerialize["reducer"] = o.Reducer
	toSerialize["source"] = o.Source

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceLimitQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query   *string `json:"query"`
		Reducer *string `json:"reducer"`
		Source  *string `json:"source"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.Reducer == nil {
		return fmt.Errorf("required field reducer missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"query", "reducer", "source"})
	} else {
		return err
	}
	o.Query = *all.Query
	o.Reducer = *all.Reducer
	o.Source = *all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
