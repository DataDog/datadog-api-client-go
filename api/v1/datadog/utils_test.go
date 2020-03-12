/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"testing"
	"runtime"
)

func TestReplaceRuntimePlaceholders(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{{
		"runtime",
		"runtime={runtime}",
		"runtime=" + runtime.Version(),
	}, {
		"os",
		"os={os}",
		"os=" + runtime.GOOS,
	}, {
		"arch",
		"arch={arch}",
		"arch=" + runtime.GOARCH,
	}, {
		"User-Agent",
		"datadog-api-client-go/0.1.0 ({runtime}; {os}/{arch})",
		"datadog-api-client-go/0.1.0 (" + runtime.Version() + "; " + runtime.GOOS + "/" + runtime.GOARCH + ")",
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceRuntimePlaceholders(tt.args); got != tt.want {
				t.Errorf("ReplaceRuntimePlaceholders(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
