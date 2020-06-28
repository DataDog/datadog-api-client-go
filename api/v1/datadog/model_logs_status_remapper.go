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
// LogsStatusRemapper Use this Processor if you want to assign some attributes as the official status.  Each incoming status value is mapped as follows.    - Integers from 0 to 7 map to the Syslog severity standards   - Strings beginning with `emerg` or f (case-insensitive) map to `emerg` (0)   - Strings beginning with `a` (case-insensitive) map to `alert` (1)   - Strings beginning with `c` (case-insensitive) map to `critical` (2)   - Strings beginning with `err` (case-insensitive) map to `error` (3)   - Strings beginning with `w` (case-insensitive) map to `warning` (4)   - Strings beginning with `n` (case-insensitive) map to `notice` (5)   - Strings beginning with `i` (case-insensitive) map to `info` (6)   - Strings beginning with `d`, `trace` or `verbose` (case-insensitive) map to `debug` (7)   - Strings beginning with `o` or matching `OK` or `Success` (case-insensitive) map to OK   - All others map to `info` (6)    **Note:** If multiple log status remapper processors can be applied to a given log,   only the first one (according to the pipelines order) is taken into account.
type LogsStatusRemapper struct {
	// Whether or not the processor is enabled.
	IsEnabled bool `json:"is_enabled,omitempty"`
	// Name of the processor.
	Name string `json:"name,omitempty"`
	// Array of source attributes.
	Sources []string `json:"sources"`
	Type LogsStatusRemapperType `json:"type"`
}
