// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureDataExplorerDestinationAuth - Authentication configuration for Azure Data Explorer. The `azure_credential_kind` field selects the credential type.
type ObservabilityPipelineAzureDataExplorerDestinationAuth struct {
	ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli                       *ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli
	ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret                   *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret
	ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate              *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate
	ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity                *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity
	ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion
	ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity               *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliAsObservabilityPipelineAzureDataExplorerDestinationAuth is a convenience function that returns ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli wrapped in ObservabilityPipelineAzureDataExplorerDestinationAuth.
func ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliAsObservabilityPipelineAzureDataExplorerDestinationAuth(v *ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli) ObservabilityPipelineAzureDataExplorerDestinationAuth {
	return ObservabilityPipelineAzureDataExplorerDestinationAuth{ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli: v}
}

// ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretAsObservabilityPipelineAzureDataExplorerDestinationAuth is a convenience function that returns ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret wrapped in ObservabilityPipelineAzureDataExplorerDestinationAuth.
func ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretAsObservabilityPipelineAzureDataExplorerDestinationAuth(v *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret) ObservabilityPipelineAzureDataExplorerDestinationAuth {
	return ObservabilityPipelineAzureDataExplorerDestinationAuth{ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret: v}
}

// ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateAsObservabilityPipelineAzureDataExplorerDestinationAuth is a convenience function that returns ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate wrapped in ObservabilityPipelineAzureDataExplorerDestinationAuth.
func ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateAsObservabilityPipelineAzureDataExplorerDestinationAuth(v *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) ObservabilityPipelineAzureDataExplorerDestinationAuth {
	return ObservabilityPipelineAzureDataExplorerDestinationAuth{ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate: v}
}

// ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityAsObservabilityPipelineAzureDataExplorerDestinationAuth is a convenience function that returns ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity wrapped in ObservabilityPipelineAzureDataExplorerDestinationAuth.
func ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityAsObservabilityPipelineAzureDataExplorerDestinationAuth(v *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity) ObservabilityPipelineAzureDataExplorerDestinationAuth {
	return ObservabilityPipelineAzureDataExplorerDestinationAuth{ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity: v}
}

// ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionAsObservabilityPipelineAzureDataExplorerDestinationAuth is a convenience function that returns ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion wrapped in ObservabilityPipelineAzureDataExplorerDestinationAuth.
func ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertionAsObservabilityPipelineAzureDataExplorerDestinationAuth(v *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion) ObservabilityPipelineAzureDataExplorerDestinationAuth {
	return ObservabilityPipelineAzureDataExplorerDestinationAuth{ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion: v}
}

// ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityAsObservabilityPipelineAzureDataExplorerDestinationAuth is a convenience function that returns ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity wrapped in ObservabilityPipelineAzureDataExplorerDestinationAuth.
func ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityAsObservabilityPipelineAzureDataExplorerDestinationAuth(v *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) ObservabilityPipelineAzureDataExplorerDestinationAuth {
	return ObservabilityPipelineAzureDataExplorerDestinationAuth{ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineAzureDataExplorerDestinationAuth) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli)
	if err == nil {
		if obj.ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli != nil && obj.ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli.UnparsedObject == nil {
			jsonObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli, _ := datadog.Marshal(obj.ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli)
			if string(jsonObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli) == "{}" { // empty struct
				obj.ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli = nil
		}
	} else {
		obj.ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli = nil
	}

	// try to unmarshal data into ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret)
	if err == nil {
		if obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret != nil && obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret.UnparsedObject == nil {
			jsonObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret, _ := datadog.Marshal(obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret)
			if string(jsonObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret) == "{}" { // empty struct
				obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret = nil
		}
	} else {
		obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret = nil
	}

	// try to unmarshal data into ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate)
	if err == nil {
		if obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate != nil && obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate.UnparsedObject == nil {
			jsonObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate, _ := datadog.Marshal(obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate)
			if string(jsonObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) == "{}" { // empty struct
				obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate = nil
		}
	} else {
		obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate = nil
	}

	// try to unmarshal data into ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity)
	if err == nil {
		if obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity != nil && obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity.UnparsedObject == nil {
			jsonObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity, _ := datadog.Marshal(obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity)
			if string(jsonObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity) == "{}" { // empty struct
				obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity = nil
		}
	} else {
		obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity = nil
	}

	// try to unmarshal data into ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion)
	if err == nil {
		if obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion != nil && obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion.UnparsedObject == nil {
			jsonObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion, _ := datadog.Marshal(obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion)
			if string(jsonObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion) == "{}" { // empty struct
				obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion = nil
		}
	} else {
		obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion = nil
	}

	// try to unmarshal data into ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity)
	if err == nil {
		if obj.ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity != nil && obj.ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity.UnparsedObject == nil {
			jsonObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity, _ := datadog.Marshal(obj.ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity)
			if string(jsonObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) == "{}" { // empty struct
				obj.ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity = nil
		}
	} else {
		obj.ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli = nil
		obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret = nil
		obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate = nil
		obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity = nil
		obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion = nil
		obj.ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineAzureDataExplorerDestinationAuth) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli)
	}

	if obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret)
	}

	if obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate)
	}

	if obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity)
	}

	if obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion)
	}

	if obj.ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineAzureDataExplorerDestinationAuth) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli != nil {
		return obj.ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli
	}

	if obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret != nil {
		return obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret
	}

	if obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate != nil {
		return obj.ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate
	}

	if obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity != nil {
		return obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity
	}

	if obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion != nil {
		return obj.ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityClientAssertion
	}

	if obj.ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity != nil {
		return obj.ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity
	}

	// all schemas are nil
	return nil
}
