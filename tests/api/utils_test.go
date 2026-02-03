package api

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"math/big"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/DataDog/datadog-api-client-go/v2/tests"
)

func TestContainsUnparsedObject(t *testing.T) {
	assert := tests.Assert(context.Background(), t)
	testCases := []struct {
		name                   string
		value                  interface{}
		expectedBool           bool
		expectedUnparsedObject interface{}
	}{
		{
			"top level unparsed struct",
			datadogV1.Dashboard{UnparsedObject: map[string]interface{}{"foo": "bar"}},
			true,
			map[string]interface{}{"foo": "bar"},
		},
		{
			"nested unparsed struct",
			datadogV1.SyntheticsAPITest{Name: "foo", Config: datadogV1.SyntheticsAPITestConfig{UnparsedObject: map[string]interface{}{"foo": "bar"}}},
			true,
			map[string]interface{}{"foo": "bar"},
		},
		{
			"unparsed struct in array",
			datadogV1.Dashboard{LayoutType: datadogV1.DASHBOARDLAYOUTTYPE_FREE, Widgets: []datadogV1.Widget{{Definition: datadogV1.WidgetDefinition{}}, {UnparsedObject: map[string]interface{}{"foo": "bar"}}}},
			true,
			map[string]interface{}{"foo": "bar"},
		},
		{
			"unparsed enum in array",
			datadogV2.DowntimeResponseAttributes{NotifyEndStates: []datadogV2.DowntimeNotifyEndStateTypes{"alert", "foobar"}},
			true,
			datadogV2.DowntimeNotifyEndStateTypes("foobar"),
		},
		{
			"unparsed enum in map",
			map[string]datadogV2.DowntimeNotifyEndStateTypes{"foo": "alert", "bar": "foobar"},
			true,
			datadogV2.DowntimeNotifyEndStateTypes("foobar"),
		},
		{
			"unparsed nullable",
			datadogV2.NewNullableLogsArchiveDestination(&datadogV2.LogsArchiveDestination{UnparsedObject: map[string]interface{}{"foo": "bar"}}),
			true,
			map[string]interface{}{"foo": "bar"},
		},
		{
			"unparsed nested in nullable",
			datadogV2.NewNullableLogsArchiveDestination(&datadogV2.LogsArchiveDestination{LogsArchiveDestinationAzure: &datadogV2.LogsArchiveDestinationAzure{UnparsedObject: map[string]interface{}{"foo": "bar"}}}),
			true,
			map[string]interface{}{"foo": "bar"},
		},
		{
			"valid nullable",
			datadogV2.NewNullableLogsArchiveDestination(&datadogV2.LogsArchiveDestination{LogsArchiveDestinationAzure: &datadogV2.LogsArchiveDestinationAzure{Type: datadogV2.LOGSARCHIVEDESTINATIONAZURETYPE_AZURE}}),
			false,
			nil,
		},
		{
			"valid struct",
			datadogV1.SyntheticsAPITest{Name: "foo", Type: datadogV1.SYNTHETICSAPITESTTYPE_API, Config: datadogV1.SyntheticsAPITestConfig{Assertions: []datadogV1.SyntheticsAssertion{{SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{Type: datadogV1.SYNTHETICSASSERTIONTYPE_BODY, Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_CONTAINS}}}}},
			false,
			nil,
		},
		{
			"valid simple type",
			"a simple string",
			false,
			nil,
		},
		{
			"valid simple pointer",
			datadog.PtrString("a simple pointer to string"),
			false,
			nil,
		},
	}

	for _, tc := range testCases {
		c := tc
		t.Run(tc.name, func(t *testing.T) {
			n, m := datadog.ContainsUnparsedObject(c.value)
			assert.Equal(c.expectedUnparsedObject, m)
			assert.Equal(c.expectedBool, n)
		})
	}
}

func TestSignString_RS256(t *testing.T) {
	assert := tests.Assert(context.Background(), t)

	// Generate test RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	assert.NoError(err)

	// Encode private key to PEM format (PKCS#1)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	testData := "test data to sign"

	// Sign using our function
	signature, err := datadog.SignString(testData, privateKeyPEM, datadog.SigningAlgorithmRS256)
	assert.NoError(err)
	assert.NotEmpty(signature)

	// Verify the signature manually
	signatureBytes, err := base64.RawURLEncoding.DecodeString(signature)
	assert.NoError(err)

	hash := sha256.Sum256([]byte(testData))
	err = rsa.VerifyPKCS1v15(&privateKey.PublicKey, crypto.SHA256, hash[:], signatureBytes)
	assert.NoError(err)
}

func TestSignString_RS256_PKCS8(t *testing.T) {
	assert := tests.Assert(context.Background(), t)

	// Generate test RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	assert.NoError(err)

	// Encode private key to PEM format (PKCS#8)
	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	assert.NoError(err)

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	testData := "test data to sign with PKCS#8"

	// Sign using our function
	signature, err := datadog.SignString(testData, privateKeyPEM, datadog.SigningAlgorithmRS256)
	assert.NoError(err)
	assert.NotEmpty(signature)

	// Verify the signature manually
	signatureBytes, err := base64.RawURLEncoding.DecodeString(signature)
	assert.NoError(err)

	hash := sha256.Sum256([]byte(testData))
	err = rsa.VerifyPKCS1v15(&privateKey.PublicKey, crypto.SHA256, hash[:], signatureBytes)
	assert.NoError(err)
}

func TestSignString_ES256(t *testing.T) {
	assert := tests.Assert(context.Background(), t)

	// Generate test ECDSA private key (P-256 curve for ES256)
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(err)

	// Encode private key to PEM format
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	assert.NoError(err)

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	testData := "test data to sign with ES256"

	// Sign using our function
	signature, err := datadog.SignString(testData, privateKeyPEM, datadog.SigningAlgorithmES256)
	assert.NoError(err)
	assert.NotEmpty(signature)

	// Verify the signature manually
	signatureBytes, err := base64.RawURLEncoding.DecodeString(signature)
	assert.NoError(err)
	assert.Equal(64, len(signatureBytes), "ES256 signature should be 64 bytes")

	// Split signature into r and s
	r := new(big.Int).SetBytes(signatureBytes[:32])
	s := new(big.Int).SetBytes(signatureBytes[32:])

	hash := sha256.Sum256([]byte(testData))
	verified := ecdsa.Verify(&privateKey.PublicKey, hash[:], r, s)
	assert.True(verified, "Signature verification should succeed")
}

func TestSignString_ES256_PKCS8(t *testing.T) {
	assert := tests.Assert(context.Background(), t)

	// Generate test ECDSA private key (P-256 curve for ES256)
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(err)

	// Encode private key to PEM format (PKCS#8)
	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	assert.NoError(err)

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	testData := "test data to sign with ES256 PKCS#8"

	// Sign using our function
	signature, err := datadog.SignString(testData, privateKeyPEM, datadog.SigningAlgorithmES256)
	assert.NoError(err)
	assert.NotEmpty(signature)

	// Verify the signature manually
	signatureBytes, err := base64.RawURLEncoding.DecodeString(signature)
	assert.NoError(err)
	assert.Equal(64, len(signatureBytes), "ES256 signature should be 64 bytes")

	// Split signature into r and s
	r := new(big.Int).SetBytes(signatureBytes[:32])
	s := new(big.Int).SetBytes(signatureBytes[32:])

	hash := sha256.Sum256([]byte(testData))
	verified := ecdsa.Verify(&privateKey.PublicKey, hash[:], r, s)
	assert.True(verified, "Signature verification should succeed")
}

func TestSignString_InvalidAlgorithm(t *testing.T) {
	assert := tests.Assert(context.Background(), t)

	// Generate a simple RSA key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	assert.NoError(err)

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	_, err = datadog.SignString("test data", privateKeyPEM, datadog.SigningAlgorithm("HS256"))
	assert.Error(err)
	assert.Contains(err.Error(), "unsupported algorithm")
}

func TestSignString_InvalidPEM(t *testing.T) {
	assert := tests.Assert(context.Background(), t)

	_, err := datadog.SignString("test data", []byte("not a valid PEM"), datadog.SigningAlgorithmRS256)
	assert.Error(err)
	assert.Contains(err.Error(), "failed to decode PEM block")
}

func TestSignString_WrongKeyTypeForAlgorithm(t *testing.T) {
	assert := tests.Assert(context.Background(), t)

	// Generate ECDSA key but try to use with RS256
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(err)

	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	assert.NoError(err)

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	_, err = datadog.SignString("test data", privateKeyPEM, datadog.SigningAlgorithmRS256)
	assert.Error(err)
	assert.Contains(err.Error(), "not an RSA private key")
}
