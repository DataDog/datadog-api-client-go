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
import (
	"time"
)
// UsageLambdaHour Number of lambda functions and sum of the invocations of all lambda functions for each hour for a given organization.
type UsageLambdaHour struct {
	// Contains the number of different functions for each region and AWS account.
	FuncCount int64 `json:"func_count,omitempty"`
	// The hour for the usage.
	Hour time.Time `json:"hour,omitempty"`
	// Contains the sum of invocations of all functions.
	InvocationsSum int64 `json:"invocations_sum,omitempty"`
}
