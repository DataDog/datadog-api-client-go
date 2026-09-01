// Package awsauth provides AWS Workload Identity Federation authentication for
// the Datadog API client.
//
// AWS credentials are discovered and refreshed by the AWS SDK for Go v2
// default configuration chain. Importing this module is optional; the core
// Datadog API client does not depend on the AWS SDK.
package awsauth
