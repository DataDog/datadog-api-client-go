// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionDataJobsQueryDefinition A formula and functions data jobs query.
type MonitorFormulaAndFunctionDataJobsQueryDefinition struct {
	// The type of job being monitored. Valid values include:
	// `databricks.job`, `spark.application`, `airflow.dag`, `dbt.job`, `glue.job`, `adf.pipeline`.
	// The `aastra` dialect also accepts `airflow.job`, `dbt.model`, and `dbt.test`.
	// Custom job types are supported with the `custom.ol.` prefix.
	JobType string `json:"job_type"`
	// Filter expression used to select the jobs to monitor. Its syntax depends on `query_dialect`:
	// a metric tag filter such as `job_name:smoke*` for `metric`, or an Aastra entity search
	// such as ``search for databricks_job where `@team:data-eng` `` for `aastra`.
	JobsQuery string `json:"jobs_query"`
	// Monitor configuration options for data jobs queries.
	MonitorOptions *MonitorFormulaAndFunctionDataJobsMonitorOptions `json:"monitor_options,omitempty"`
	// Name of the query for use in formulas. Must be `run_query`.
	Name string `json:"name"`
	// Query dialect for data jobs queries. Valid values are `metric` and `aastra`.
	QueryDialect string `json:"query_dialect"`
	// Optional filter on individual runs, using metrics filter syntax.
	// Only supported with the `aastra` query dialect.
	RunFilter *string `json:"run_filter,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorFormulaAndFunctionDataJobsQueryDefinition instantiates a new MonitorFormulaAndFunctionDataJobsQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorFormulaAndFunctionDataJobsQueryDefinition(jobType string, jobsQuery string, name string, queryDialect string) *MonitorFormulaAndFunctionDataJobsQueryDefinition {
	this := MonitorFormulaAndFunctionDataJobsQueryDefinition{}
	this.JobType = jobType
	this.JobsQuery = jobsQuery
	this.Name = name
	this.QueryDialect = queryDialect
	return &this
}

// NewMonitorFormulaAndFunctionDataJobsQueryDefinitionWithDefaults instantiates a new MonitorFormulaAndFunctionDataJobsQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorFormulaAndFunctionDataJobsQueryDefinitionWithDefaults() *MonitorFormulaAndFunctionDataJobsQueryDefinition {
	this := MonitorFormulaAndFunctionDataJobsQueryDefinition{}
	return &this
}

// GetJobType returns the JobType field value.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) GetJobType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) GetJobTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobType, true
}

// SetJobType sets field value.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) SetJobType(v string) {
	o.JobType = v
}

// GetJobsQuery returns the JobsQuery field value.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) GetJobsQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.JobsQuery
}

// GetJobsQueryOk returns a tuple with the JobsQuery field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) GetJobsQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobsQuery, true
}

// SetJobsQuery sets field value.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) SetJobsQuery(v string) {
	o.JobsQuery = v
}

// GetMonitorOptions returns the MonitorOptions field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) GetMonitorOptions() MonitorFormulaAndFunctionDataJobsMonitorOptions {
	if o == nil || o.MonitorOptions == nil {
		var ret MonitorFormulaAndFunctionDataJobsMonitorOptions
		return ret
	}
	return *o.MonitorOptions
}

// GetMonitorOptionsOk returns a tuple with the MonitorOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) GetMonitorOptionsOk() (*MonitorFormulaAndFunctionDataJobsMonitorOptions, bool) {
	if o == nil || o.MonitorOptions == nil {
		return nil, false
	}
	return o.MonitorOptions, true
}

// HasMonitorOptions returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) HasMonitorOptions() bool {
	return o != nil && o.MonitorOptions != nil
}

// SetMonitorOptions gets a reference to the given MonitorFormulaAndFunctionDataJobsMonitorOptions and assigns it to the MonitorOptions field.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) SetMonitorOptions(v MonitorFormulaAndFunctionDataJobsMonitorOptions) {
	o.MonitorOptions = &v
}

// GetName returns the Name field value.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) SetName(v string) {
	o.Name = v
}

// GetQueryDialect returns the QueryDialect field value.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) GetQueryDialect() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryDialect
}

// GetQueryDialectOk returns a tuple with the QueryDialect field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) GetQueryDialectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryDialect, true
}

// SetQueryDialect sets field value.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) SetQueryDialect(v string) {
	o.QueryDialect = v
}

// GetRunFilter returns the RunFilter field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) GetRunFilter() string {
	if o == nil || o.RunFilter == nil {
		var ret string
		return ret
	}
	return *o.RunFilter
}

// GetRunFilterOk returns a tuple with the RunFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) GetRunFilterOk() (*string, bool) {
	if o == nil || o.RunFilter == nil {
		return nil, false
	}
	return o.RunFilter, true
}

// HasRunFilter returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) HasRunFilter() bool {
	return o != nil && o.RunFilter != nil
}

// SetRunFilter gets a reference to the given string and assigns it to the RunFilter field.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) SetRunFilter(v string) {
	o.RunFilter = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorFormulaAndFunctionDataJobsQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["job_type"] = o.JobType
	toSerialize["jobs_query"] = o.JobsQuery
	if o.MonitorOptions != nil {
		toSerialize["monitor_options"] = o.MonitorOptions
	}
	toSerialize["name"] = o.Name
	toSerialize["query_dialect"] = o.QueryDialect
	if o.RunFilter != nil {
		toSerialize["run_filter"] = o.RunFilter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorFormulaAndFunctionDataJobsQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		JobType        *string                                          `json:"job_type"`
		JobsQuery      *string                                          `json:"jobs_query"`
		MonitorOptions *MonitorFormulaAndFunctionDataJobsMonitorOptions `json:"monitor_options,omitempty"`
		Name           *string                                          `json:"name"`
		QueryDialect   *string                                          `json:"query_dialect"`
		RunFilter      *string                                          `json:"run_filter,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.JobType == nil {
		return fmt.Errorf("required field job_type missing")
	}
	if all.JobsQuery == nil {
		return fmt.Errorf("required field jobs_query missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.QueryDialect == nil {
		return fmt.Errorf("required field query_dialect missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"job_type", "jobs_query", "monitor_options", "name", "query_dialect", "run_filter"})
	} else {
		return err
	}

	hasInvalidField := false
	o.JobType = *all.JobType
	o.JobsQuery = *all.JobsQuery
	if all.MonitorOptions != nil && all.MonitorOptions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MonitorOptions = all.MonitorOptions
	o.Name = *all.Name
	o.QueryDialect = *all.QueryDialect
	o.RunFilter = all.RunFilter

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
