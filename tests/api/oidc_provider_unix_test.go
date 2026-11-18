//go:build unix

package api

import (
	"path/filepath"
	"syscall"
	"testing"
	"time"

	oidcauth "github.com/DataDog/datadog-api-client-go/v2/auth/oidc"
)

func TestOIDCProofRefusesFIFOWithoutBlocking(t *testing.T) {
	path := filepath.Join(t.TempDir(), "token")
	if err := syscall.Mkfifo(path, 0o600); err != nil {
		t.Fatal(err)
	}
	provider, err := oidcauth.New(oidcauth.WithTokenFile(path))
	if err != nil {
		t.Fatal(err)
	}
	result := make(chan error, 1)
	go func() {
		_, err := provider.Proof(testOrgUUID)
		result <- err
	}()
	select {
	case err := <-result:
		if err == nil {
			t.Fatal("Proof read a FIFO, want an error")
		}
	case <-time.After(5 * time.Second):
		t.Fatal("Proof blocked opening a FIFO")
	}
}
