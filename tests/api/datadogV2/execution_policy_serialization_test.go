package test

import (
	"encoding/json"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func TestExecutionPolicyRshellAllowedCommandsSerialization(t *testing.T) {
	rule := datadogV2.ExecutionPolicyRemoteActionRshellScopeRule{
		Access:      datadogV2.EXECUTIONPOLICYREMOTEACTIONRSHELLACCESS_READ_ONLY,
		TargetPaths: []string{"/var/log"},
	}
	tests := []struct {
		name         string
		allowed      *datadogV2.ExecutionPolicyRemoteActionRshellAllowedCommands
		wantOuter    bool
		wantCommands []string
	}{
		{
			name:      "missing outer field",
			wantOuter: false,
		},
		{
			name:         "present empty list",
			allowed:      &datadogV2.ExecutionPolicyRemoteActionRshellAllowedCommands{Commands: []datadogV2.ExecutionPolicyRemoteActionRshellAllowedCommand{}},
			wantOuter:    true,
			wantCommands: []string{},
		},
		{
			name: "exact wildcard",
			allowed: &datadogV2.ExecutionPolicyRemoteActionRshellAllowedCommands{
				Commands: []datadogV2.ExecutionPolicyRemoteActionRshellAllowedCommand{{Command: "rshell:*"}},
			},
			wantOuter:    true,
			wantCommands: []string{"rshell:*"},
		},
		{
			name: "literal list",
			allowed: &datadogV2.ExecutionPolicyRemoteActionRshellAllowedCommands{
				Commands: []datadogV2.ExecutionPolicyRemoteActionRshellAllowedCommand{
					{Command: "rshell:cat"},
					{Command: "rshell:ls"},
				},
			},
			wantOuter:    true,
			wantCommands: []string{"rshell:cat", "rshell:ls"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			scope := datadogV2.ExecutionPolicyRemoteActionRshellScope{
				AllowedCommands: tc.allowed,
				Rules:           []datadogV2.ExecutionPolicyRemoteActionRshellScopeRule{rule},
			}
			encoded, err := json.Marshal(scope)
			if err != nil {
				t.Fatal(err)
			}

			var payload map[string]json.RawMessage
			if err := json.Unmarshal(encoded, &payload); err != nil {
				t.Fatal(err)
			}
			_, hasOuter := payload["allowed_commands"]
			if hasOuter != tc.wantOuter {
				t.Fatalf("allowed_commands presence = %t, want %t: %s", hasOuter, tc.wantOuter, encoded)
			}
			if scope.HasAllowedCommands() != tc.wantOuter {
				t.Fatalf("HasAllowedCommands() = %t, want %t", scope.HasAllowedCommands(), tc.wantOuter)
			}
			var decoded datadogV2.ExecutionPolicyRemoteActionRshellScope
			if err := json.Unmarshal(encoded, &decoded); err != nil {
				t.Fatal(err)
			}
			if decoded.HasAllowedCommands() != tc.wantOuter {
				t.Fatalf("decoded HasAllowedCommands() = %t, want %t", decoded.HasAllowedCommands(), tc.wantOuter)
			}
			if !hasOuter {
				return
			}

			var allowed struct {
				Commands []datadogV2.ExecutionPolicyRemoteActionRshellAllowedCommand `json:"commands"`
			}
			if err := json.Unmarshal(payload["allowed_commands"], &allowed); err != nil {
				t.Fatal(err)
			}
			if allowed.Commands == nil {
				t.Fatal("commands was serialized as null or omitted")
			}
			if len(allowed.Commands) != len(tc.wantCommands) {
				t.Fatalf("command count = %d, want %d", len(allowed.Commands), len(tc.wantCommands))
			}
			for i, want := range tc.wantCommands {
				if got := allowed.Commands[i].Command; got != want {
					t.Fatalf("command %d = %q, want %q", i, got, want)
				}
			}

			if len(tc.wantCommands) == 0 && decoded.GetAllowedCommands().Commands == nil {
				t.Fatal("decoded empty commands was nil")
			}
		})
	}
}
