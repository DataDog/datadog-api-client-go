/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package scenarios

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/tests"
	"github.com/go-bdd/gobdd"
	"github.com/mcuadros/go-lookup"
)

var templateFunctions = map[string]func(map[string]interface{}, string) string{
	"timeISO":   relativeTime(true),
	"timestamp": relativeTime(false),
}

func relativeTime(iso bool) func(map[string]interface{}, string) string {
	timeRE := regexp.MustCompile(`now( *([+-]) *(\d+)([smhdMy]))?`)
	unitToDur := map[string]time.Duration{
		"s": time.Second,
		"m": time.Minute,
		"h": time.Hour,
	}
	return func(data map[string]interface{}, arg string) string {
		ret := data["now"].(time.Time)
		if res := timeRE.FindSubmatch([]byte(arg)); res != nil {
			if string(res[1]) != "" {
				// Add or substract a duration from now
				num, _ := strconv.ParseInt(string(res[2])+string(res[3]), 10, 64)
				unit := string(res[4])
				switch unit {
				case "s", "m", "h":
					ret = ret.Add(unitToDur[unit] * time.Duration(num))
				case "d":
					ret = ret.AddDate(0, 0, int(num))
				case "M":
					ret = ret.AddDate(0, int(num), 0)
				case "y":
					ret = ret.AddDate(int(num), 0, 0)
				}
			}
			if iso {
				return ret.Format(time.RFC3339)
			} else {
				return strconv.FormatInt(ret.Unix(), 10)
			}
		}
		return ""
	}
}


type versionKey struct{}
type ctxKey struct{}
type clientKey struct{}
type apiKey struct{}
type requestKey struct{}
type requestNameKey struct{}
type requestArgsKey struct{}
type requestParamsKey struct{}
type responseKey struct{}
type jsonResponseKey struct{}
type dataKey struct{}
type bodyKey struct{}
type cleanupKey struct{}
type pathParamCountKey struct{}

// GetIgnoredTags returns list of ignored tags.
func GetIgnoredTags() []string {
	tags := make([]string, 1)
	tags = append(tags, "@skip", "@skip-go")
	if tests.GetRecording() != tests.ModeIgnore {
		tags = append(tags, "@integration-only")
	}
	if tests.GetRecording() != tests.ModeReplaying {
		tags = append(tags, "@replay-only")
	}
	return tags
}

// Templated replaces {{ path }} in source with value from data[path]. If path contains `(`, `)`, then it's a function to call, with the args between brackets.
func Templated(t gobdd.StepTest, data map[string]interface{}, source string) string {
	re := regexp.MustCompile(`{{ *([^{}]+|'[^']+'|"[^"]+") *}}`)
	functionRE := regexp.MustCompile(`^(.+)\((.*)\)$`)
	replace := func(source string) string {
		path := strings.Trim(source, "{ }")
		if res := functionRE.FindSubmatch([]byte(path)); res != nil {
			functionName := string(res[1])
			arg := string(res[2])
			res := templateFunctions[functionName](data, arg)
			return res
		}
		if path[0] == '\'' || path[0] == '"' {
			return path[1 : len(path)-1]
		}
		v, err := lookup.LookupStringI(data, path)
		if err != nil {
			t.Fatalf("problem with replacement of %s (%s): %v", source, path, err)
		}
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return strconv.FormatInt(v.Int(), 10)
		case reflect.Float64, reflect.Float32:
			return strconv.FormatFloat(v.Float(), 'f', -1, 10)
		case reflect.Bool:
			return strconv.FormatBool(v.Bool())
		case reflect.Interface:
			return v.Elem().String()
		}
		return v.String()
	}
	return re.ReplaceAllStringFunc(source, replace)
}

var matchFirstUpper = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllUpper = regexp.MustCompile("([a-z0-9])([A-Z])")

// CamelToSnakeCase converts CamelCase to snake_case.
func CamelToSnakeCase(camel string) string {
    snake := matchFirstUpper.ReplaceAllString(camel, "${1}_${2}")
    snake = matchAllUpper.ReplaceAllString(snake, "${1}_${2}")
    return strings.ToLower(snake)
}

// ToVarName converts a string to a valid Go variable name.
func ToVarName(param string) (varName string) {
	isToUpper := true

	for _, v := range param {
		if isToUpper {
			varName += strings.ToUpper(string(v))
			isToUpper = false
		} else {
			if v == '_' {
				isToUpper = true
			} else if m, _ := regexp.Match("[()\\[\\].]", []byte{byte(v)}); m {
				isToUpper = true
			} else {
				varName += string(v)
			}
		}
	}
	return
}

// SnakeToCamelCase converts snake_case to SnakeCase.
func SnakeToCamelCase(snake string) (camel string) {
	isToUpper := false

	for k, v := range snake {
		if k == 0 {
			camel = strings.ToUpper(string(v))
		} else {
			if isToUpper {
				camel += strings.ToUpper(string(v))
				isToUpper = false
			} else {
				if v == '_' {
					isToUpper = true
				} else if v == '.' { // support for lookup paths
					isToUpper = true
					camel += string(v)
				} else {
					camel += string(v)
				}
			}
		}
	}
	return
}



// GetT returns stored reference to the testing object.
func GetT(ctx gobdd.Context) *testing.T {
	t, err := ctx.Get(gobdd.TestingTKey{})
	if err != nil {
		return nil
	}
	tt, ok := t.(*testing.T)
	if !ok {
		return nil
	}
	return tt
}

// GetCtx returns Go context.
func GetCtx(ctx gobdd.Context) context.Context {
	c, err := ctx.Get(ctxKey{})
	if err != nil {
		GetT(ctx).Logf("could not get a context: %v", err)
	}
	return c.(context.Context)
}

// GetResponse returns request response.
func GetResponse(ctx gobdd.Context) []reflect.Value {
	r, err := ctx.Get(responseKey{})
	if err != nil {
		GetT(ctx).Logf("could not get a response: %v", err)
	}
	return r.([]reflect.Value)
}

// GetJSONResponse returns request response.
func GetJSONResponse(ctx gobdd.Context) interface{} {
	r, err := ctx.Get(jsonResponseKey{})
	if err != nil {
		GetT(ctx).Logf("could not get a json response: %v", err)
	}
	return r
}

// GetResponseStatusCode returns request response status code.
func GetResponseStatusCode(resp []reflect.Value) (int, error) {
	if len(resp) < 2 {
		return -1, fmt.Errorf("needs at least 2 values, got %d", len(resp))
	}
	// Execute() returns triples -> 2nd value is *http.Response -> get StatusCode]
	response := resp[len(resp)-2].Interface().(*http.Response)
	if response == nil {
		return -1, errors.New("response is nil")
	}
	return response.StatusCode, nil
}

// SetCtx sets Go context in BDD context.
func SetCtx(ctx gobdd.Context, value context.Context) {
	ctx.Set(ctxKey{}, value)
}

// SetFixtureData sets the fixture data in BDD context
func SetFixtureData(ctx gobdd.Context) {
	ct, _ := ctx.Get(gobdd.TestingTKey{})
	cctx := GetCtx(ctx)
	nameParts := strings.Split(ct.(*testing.T).Name(), "/")
	testName := strings.SplitN(nameParts[3], "_", 2)[1]
	unique := tests.WithUniqueSurrounding(cctx, testName)
	alnum := regexp.MustCompile(`[^A-Za-z0-9]+`)
	data := GetData(ctx)
	data["unique"] = unique
	data["unique_lower"] = strings.ToLower(unique)
	data["unique_upper"] = strings.ToUpper(unique)
	data["unique_alnum"] = string(alnum.ReplaceAll([]byte(unique), []byte("")))
	data["unique_lower_alnum"] = strings.ToLower(data["unique_alnum"].(string))
	data["unique_upper_alnum"] = strings.ToUpper(data["unique_alnum"].(string))
	data["now"] = tests.ClockFromContext(cctx).Now()
}

// SetClient sets client reflection.
func SetClient(ctx gobdd.Context, value interface{}) {
	ctx.Set(clientKey{}, value)
}

// GetClient get reflected value of client instance.
func GetClient(ctx gobdd.Context) reflect.Value {
	client, _ := GetClientVersion(ctx, GetVersion(ctx))

	return client
}

// SetVersion sets package version.
func SetVersion(ctx gobdd.Context, version string) {
	ctx.Set(versionKey{}, version)
}

// GetVersion gets package version.
func GetVersion(ctx gobdd.Context) string {
	c, err := ctx.Get(versionKey{})
	if err != nil {
		GetT(ctx).Fatalf("could not get version: %v", err)
	}
	return c.(string)
}

// SetAPI sets client API.
func SetAPI(ctx gobdd.Context, value interface{}) {
	ctx.Set(apiKey{}, value)
}

// GetData returns feature context.
func GetData(ctx gobdd.Context) map[string]interface{} {
	c, err := ctx.Get(dataKey{})
	if err != nil {
		GetT(ctx).Fatalf("could not get data: %v", err)
	}
	return c.(map[string]interface{})
}

// SetData sets feature context.
func SetData(ctx gobdd.Context, value map[string]interface{}) {
	ctx.Set(dataKey{}, value)
}

// GetRequestParameters helps to build a request.
func GetRequestParameters(ctx gobdd.Context) map[string]interface{} {
	c, err := ctx.Get(requestParamsKey{})
	if err != nil {
		GetT(ctx).Fatalf("could not get request parameters: %v", err)
	}
	return c.(map[string]interface{})
}

// GetRequestArguments helps to build a request.
func GetRequestArguments(ctx gobdd.Context) []interface{} {
	c, err := ctx.Get(requestArgsKey{})
	if err != nil {
		GetT(ctx).Fatalf("could not get request arguments: %v", err)
	}
	return c.([]interface{})
}

// SetCleanup functions.
func SetCleanup(ctx gobdd.Context, value map[string]func()) {
	ctx.Set(cleanupKey{}, value)
}

// GetCleanup functions.
func GetCleanup(ctx gobdd.Context) map[string]func() {
	c, err := ctx.Get(cleanupKey{})
	if err != nil {
		GetT(ctx).Fatalf("could not get cleanup: %v", err)
	}
	return c.(map[string]func())
}

// RunCleanup executes cleanup functions in controlled ordered based on registered keys.
func RunCleanup(ctx gobdd.Context) {
	c, _ := ctx.Get(cleanupKey{})
	cc := c.(map[string]func())
	keys := make([]string, 0, len(cc))

	for k := range cc {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, name := range keys {
		cc[name]()
	}
}


// getRequestBuilder returns the reflect value of the current request in ctx
func getRequestBuilder(ctx gobdd.Context) (reflect.Value, []reflect.Value, error) {
	c, err := ctx.Get(requestKey{})
	if err != nil {
		return reflect.Value{}, []reflect.Value{}, fmt.Errorf("Missing requestKey{}")
	}
	f := c.(reflect.Value)

	in := make([]reflect.Value, f.Type().NumIn())

	// first argument is always context.Context
	in[0] = reflect.ValueOf(GetCtx(ctx))
	requestArgs := GetRequestArguments(ctx)
	requestParams := GetRequestParameters(ctx)
	numArgs := f.Type().NumIn()

	if f.Type().IsVariadic() {
		for i := 1; i < numArgs-1 && i <= len(requestArgs); i++ {
			object := requestArgs[i-1]
			in[i] = object.(reflect.Value)
		}

		optionalParams := reflect.New(f.Type().In(numArgs - 1).Elem())
		for k, v := range requestParams {
			paramName := ToVarName(k)
			if method := optionalParams.MethodByName("With" + paramName); method.IsValid() {
				at := reflect.New(method.Type().In(0))
				switch v.(type) {
				case reflect.Value:
					at.Elem().Set(v.(reflect.Value))
				case interface{}:
					json.Unmarshal([]byte(v.(string)), at.Interface())
				}
				optionalParams = method.Call([]reflect.Value{at.Elem()})[0]
			} else if k == "body" {
				a := reflect.New(f.Type().In(numArgs - 2))
				json.Unmarshal([]byte(v.(string)), a.Interface())
				in[numArgs-2] = a.Elem()
			}
		}

		in[len(in)-1] = optionalParams.Elem()
	} else {
		// Set request args
		for i := 1; i <= len(requestArgs); i++ {
			object := requestArgs[i-1].(reflect.Value)
			in[i] = object.Convert(f.Type().In(i))
		}
		// Append body to args if it exists
		if val, ok := requestParams["body"]; ok {
			a := reflect.New(f.Type().In(numArgs - 1))
			json.Unmarshal([]byte(val.(string)), a.Interface())
			in[numArgs-1] = a.Elem()
		}
	}

	return f, in, nil
}

func toJSON(o reflect.Value) (interface{}, error) {
	var jsonResult interface{}
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(o.Interface()); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(buf.Bytes(), &jsonResult); err != nil {
		return nil, err
	}
	return jsonResult, nil
}


func stringToType(s string, t interface{}) (interface{}, error) {
	switch t.(type) {
	case int:
		return strconv.Atoi(s)
	case int64:
		return strconv.ParseInt(s, 10, 64)
	case float64:
		return strconv.ParseFloat(s, 64)
	case string:
		return strconv.Unquote(s)
	case bool:
		return strconv.ParseBool(s)
	case map[string]interface{}, []interface{}:
		var res map[string]interface{}
		if err := json.Unmarshal([]byte(s), &res); err != nil {
			return nil, fmt.Errorf("error converting %s to %T: %v", s, t, err)
		}
		return res, nil
	default:
		return nil, fmt.Errorf("unknown type '%T' to convert", t)
	}
}
