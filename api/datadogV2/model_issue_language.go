// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueLanguage Programming language associated with the issue.
type IssueLanguage string

// List of IssueLanguage.
const (
	ISSUELANGUAGE_JAVASCRIPT   IssueLanguage = "javascript"
	ISSUELANGUAGE_JVM          IssueLanguage = "jvm"
	ISSUELANGUAGE_RUBY         IssueLanguage = "ruby"
	ISSUELANGUAGE_TYPESCRIPT   IssueLanguage = "typescript"
	ISSUELANGUAGE_JAVA         IssueLanguage = "java"
	ISSUELANGUAGE_KOTLIN       IssueLanguage = "kotlin"
	ISSUELANGUAGE_SCALA        IssueLanguage = "scala"
	ISSUELANGUAGE_GROOVY       IssueLanguage = "groovy"
	ISSUELANGUAGE_CLOJURE      IssueLanguage = "clojure"
	ISSUELANGUAGE_GO           IssueLanguage = "go"
	ISSUELANGUAGE_PYTHON       IssueLanguage = "python"
	ISSUELANGUAGE_PHP          IssueLanguage = "php"
	ISSUELANGUAGE_DOT_NET      IssueLanguage = "dot_net"
	ISSUELANGUAGE_C_SHARP      IssueLanguage = "c_sharp"
	ISSUELANGUAGE_C_PLUS_PLUS  IssueLanguage = "c_plus_plus"
	ISSUELANGUAGE_OBJECTIVE_C  IssueLanguage = "objective_c"
	ISSUELANGUAGE_SWIFT        IssueLanguage = "swift"
	ISSUELANGUAGE_BRIGHTSCRIPT IssueLanguage = "brightscript"
	ISSUELANGUAGE_C            IssueLanguage = "c"
	ISSUELANGUAGE_ELIXIR       IssueLanguage = "elixir"
	ISSUELANGUAGE_ERLANG       IssueLanguage = "erlang"
	ISSUELANGUAGE_PERL         IssueLanguage = "perl"
	ISSUELANGUAGE_HASKELL      IssueLanguage = "haskell"
	ISSUELANGUAGE_RUST         IssueLanguage = "rust"
)

var allowedIssueLanguageEnumValues = []IssueLanguage{
	ISSUELANGUAGE_JAVASCRIPT,
	ISSUELANGUAGE_JVM,
	ISSUELANGUAGE_RUBY,
	ISSUELANGUAGE_TYPESCRIPT,
	ISSUELANGUAGE_JAVA,
	ISSUELANGUAGE_KOTLIN,
	ISSUELANGUAGE_SCALA,
	ISSUELANGUAGE_GROOVY,
	ISSUELANGUAGE_CLOJURE,
	ISSUELANGUAGE_GO,
	ISSUELANGUAGE_PYTHON,
	ISSUELANGUAGE_PHP,
	ISSUELANGUAGE_DOT_NET,
	ISSUELANGUAGE_C_SHARP,
	ISSUELANGUAGE_C_PLUS_PLUS,
	ISSUELANGUAGE_OBJECTIVE_C,
	ISSUELANGUAGE_SWIFT,
	ISSUELANGUAGE_BRIGHTSCRIPT,
	ISSUELANGUAGE_C,
	ISSUELANGUAGE_ELIXIR,
	ISSUELANGUAGE_ERLANG,
	ISSUELANGUAGE_PERL,
	ISSUELANGUAGE_HASKELL,
	ISSUELANGUAGE_RUST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueLanguage) GetAllowedValues() []IssueLanguage {
	return allowedIssueLanguageEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueLanguage) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueLanguage(value)
	return nil
}

// NewIssueLanguageFromValue returns a pointer to a valid IssueLanguage
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueLanguageFromValue(v string) (*IssueLanguage, error) {
	ev := IssueLanguage(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueLanguage: valid values are %v", v, allowedIssueLanguageEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueLanguage) IsValid() bool {
	for _, existing := range allowedIssueLanguageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueLanguage value.
func (v IssueLanguage) Ptr() *IssueLanguage {
	return &v
}
