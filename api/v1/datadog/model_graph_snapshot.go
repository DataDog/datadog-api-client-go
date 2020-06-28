/*
 * Datadog API Collection
 *
 * Collection of all Datadog Public endpoints.
 *
 * API version: 1.0
 * Contact: support@datadoghq.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog
// GraphSnapshot Object representing a graph snapshot.
type GraphSnapshot struct {
	// A JSON document defining the graph. `graph_def` can be used instead of `metric_query`. The JSON document uses the [grammar defined here](https://docs.datadoghq.com/graphing/graphing_json/#grammar) and should be formatted to a single line then URL encoded.
	GraphDef string `json:"graph_def,omitempty"`
	// The metric query. One of `metric_query` or `graph_def` is required.
	MetricQuery string `json:"metric_query,omitempty"`
	// URL of your [graph snapshot](https://docs.datadoghq.com/metrics/explorer/#snapshot).
	SnapshotUrl string `json:"snapshot_url,omitempty"`
}
