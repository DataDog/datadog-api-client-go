"""Go data formatter."""
from functools import singledispatch
import warnings
import re

import dateutil.parser

from .utils import camel_case, schema_name


def go_name(name):
    """Convert key to Go name.

    Example:

    >>> go_name("DASHBOARD_ID")
    DashboardID
    """
    return "".join(
        part.capitalize() if part not in {"API", "ID", "HTTP", "URL", "DNS"} else part for part in name.split("_")
    )


def simple_type(schema):
    """Return the simple type of a schema.

    :param schema: The schema to extract the type from
    :return: The simple type name
    """
    if "enum" in schema:
        # enums have named types in Go client
        return None

    type_name = schema.get("type")
    type_format = schema.get("format")

    if type_name == "integer":
        return {
            "int32": "int",
            "int64": "int64",
            None: "int",
        }[type_format]

    if type_name == "number":
        return {
            "double": "float64",
            None: "float",
        }[type_format]

    if type_name == "string":
        return {
            "date": "time.Time",
            "date-time": "time.Time",
            "email": "string",
            None: "string",
        }[type_format]

    if type_name == "boolean":
        return "bool"

    return None


def reference_to_value(schema, value, print_nullable=True):
    """Return a reference to a value.

    :param schema: The schema to extract the type from
    :param value: The value to reference
    :return: The simple type name
    """
    type_name = schema.get("type")
    type_format = schema.get("format")
    nullable = schema.get("nullable", False)

    if nullable and print_nullable:
        if value == "nil":
            formatter = "*datadog.NewNullable{function_name}({value})"
        else:
            formatter = "*datadog.NewNullable{function_name}(datadog.Ptr{function_name}({value}))"
    else:
        formatter = "datadog.Ptr{function_name}({value})"

    if type_name == "integer":
        function_name = {
            "int": "Int",
            "int32": "Int32",
            "int64": "Int64",
            None: "Int",
        }[type_format]
        return formatter.format(function_name=function_name, value=value)

    if type_name == "number":
        function_name = {
            "float": "Float32",
            "double": "Float64",
            None: "Float32",
        }[type_format]
        return formatter.format(function_name=function_name, value=value)

    if type_name == "string":
        function_name = {
            "date": "Time",
            "date-time": "Time",
            "email": "String",
            None: "String",
        }[type_format]
        return formatter.format(function_name=function_name, value=value)

    if type_name == "boolean":
        return formatter.format(function_name="Bool", value=value)

    if nullable:
        function_name = schema_name(schema)
        if function_name is None:
            raise NotImplementedError(f"nullable {schema} is not supported")
        return formatter.format(function_name=function_name, value=value)
    return f"&{value}"


def format_parameters(data, spec, replace_values=None, has_body=False, **kwargs):
    parameters_spec = {p["name"]: p for p in spec.get("parameters", [])}
    if "requestBody" in spec and "multipart/form-data" in spec["requestBody"]["content"]:
        parent = spec["requestBody"]["content"]["multipart/form-data"]["schema"]
        for name, schema in parent["properties"].items():
            parameters_spec[name] = {
                "in": "form",
                "schema": schema,
                "name": name,
                "description": schema.get("description"),
                "required": name in parent.get("required", []),
            }

    parameters = ""
    has_optional = False
    for p in parameters_spec.values():
        required = p.get("required", False)
        if required:
            k = p["name"]
            v = data.pop(k)  # otherwise there is a missing required parameters
            value = format_data_with_schema(
                v["value"],
                p["schema"],
                name_prefix="datadog.",
                replace_values=replace_values,
                required=True,
                **kwargs,
            )
            parameters += f"{value}, "
        else:
            has_optional = True

    body_is_required = spec.get("requestBody", {"required": None}).get("required", False)

    if has_body and body_is_required:
        parameters += "body, "

    if has_optional or body_is_required is False:
        parameters += f"*datadog.New{spec['operationId'][0].upper()}{spec['operationId'][1:]}OptionalParameters()"
        if has_body and not body_is_required:
            parameters += ".WithBody(body)"

        for k, v in data.items():
            value = format_data_with_schema(
                v["value"],
                parameters_spec[k]["schema"],
                name_prefix="datadog.",
                replace_values=replace_values,
                required=True,
                **kwargs,
            )
            parameters += f".With{camel_case(k)}({value})"

        parameters += ", "

    return parameters


@singledispatch
def format_data_with_schema(
    data,
    schema,
    name_prefix="",
    replace_values=None,
    default_name=None,
    required=False,
    in_list=False,
    **kwargs,
):
    if not schema:
        return ""

    nullable = schema.get("nullable", False)
    variables = kwargs.get("variables", set())

    name = schema_name(schema)

    if "enum" in schema and data not in schema["enum"]:
        raise ValueError(f"{data} is not valid enum value {schema['enum']}")

    if replace_values and data in replace_values:
        parameters = replace_values[data]

        # Make sure that variables used in given statements are camelCase for Go linter
        if parameters in variables:
            parameters = go_name(parameters)

        simple_type_value = simple_type(schema)
        if isinstance(data, int) and simple_type_value in {
            "float",
            "float32",
            "float64",
        }:
            parameters = f"{simple_type_value}({parameters})"
    else:
        if nullable and data is None:
            parameters = "nil"
        else:

            def format_string(x):
                if "`" in x:
                    x = re.sub(r"(`+)", r'` + "\1" + `', x)
                if x and ('"' in x or "\n" in x):
                    x = f"`{x}`"
                    x = re.sub(r" \+ ``$", "", x)
                    return x
                return f'"{x}"' if x else '""'

            def format_datetime(x):
                d = dateutil.parser.isoparse(x)
                return f"time.Date({d.year}, {d.month}, {d.day}, {d.hour}, {d.minute}, {d.second}, {d.microsecond}, time.UTC)"

            schema = schema.copy()

            def format_interface(x):
                if isinstance(x, int):
                    return str(x)
                if isinstance(x, float):
                    return str(x)
                if isinstance(x, str):
                    return format_string(x)
                raise TypeError(f"{x} is not supported type {schema}")

            def format_bool(x):
                if not isinstance(x, bool):
                    raise TypeError(f"{x} is not supported type {schema}")
                return "true" if x else "false"

            def open_file(x):
                return f"func() *os.File {{ fp, _ := os.Open({format_string(x)}); return fp }}()"

            formatter = {
                "int32": str,
                "int64": str,
                "double": str,
                "date-time": format_datetime,
                "number": str,
                "integer": str,
                "boolean": format_bool,
                "string": format_string,
                "email": format_string,
                "binary": open_file,
                None: format_interface,
            }[schema.get("format", schema.get("type"))]

            # TODO format date and datetime
            parameters = formatter(data)

    if "enum" in schema and name:
        # find schema index and get name from x-enum-varnames
        index = schema["enum"].index(data)
        enum_varnames = schema["x-enum-varnames"][index]
        parameters = f"{name_prefix}{name.upper()}_{enum_varnames}"
        if not required:
            parameters = f"{parameters}.Ptr()"
        return parameters
        # TODO handle nullable enums if necessary
        # return f"{name_prefix}{name}({parameters}){'.Ptr()' if not required else ''}"

    if in_list and nullable:
        schema = schema.copy()
        schema["nullable"] = False

    if (not required or schema.get("nullable")) and schema.get("type") is not None:
        return reference_to_value(schema, parameters, print_nullable=not in_list)

    if "oneOf" in schema and name:
        matched = 0
        one_of_schema = None
        for sub_schema in schema["oneOf"]:
            try:
                if sub_schema.get("nullable") and data is None:
                    # only one schema can be nullable
                    formatted = "nil"
                else:
                    sub_schema["nullable"] = False
                    formatted = format_data_with_schema(
                        data,
                        sub_schema,
                        name_prefix=name_prefix,
                        replace_values=replace_values,
                        **kwargs,
                    )
                if matched == 0:
                    one_of_schema = sub_schema
                    # NOTE we do not support mixed schemas with oneOf
                    # parameters += formatted
                    parameters = formatted
                matched += 1
            except (KeyError, ValueError, TypeError) as e:
                print(f"{e}")

        if matched == 0:
            raise ValueError(f"[{matched}] {data} is not valid for schema {name}")
        elif matched > 1:
            warnings.warn(f"[{matched}] {data} is not valid for schema {name}")

        one_of_schema_name = schema_name(one_of_schema)
        if not one_of_schema_name:
            one_of_schema_name = simple_type(one_of_schema).title()
        reference = "" if required or nullable else "&"
        return f"{reference}{name_prefix}{name}{{\n{one_of_schema_name}: {parameters}}}"

    return parameters


@format_data_with_schema.register(list)
def format_data_with_schema_list(
    data,
    schema,
    name_prefix="",
    replace_values=None,
    default_name=None,
    required=False,
    in_list=False,
    **kwargs,
):
    if not schema:
        return ""

    parameters = ""
    # collect nested array types until you find a non-array type
    schema_parts = [(required, "[]")]
    list_schema = schema["items"]
    while list_schema.get("type") == "array":
        schema_parts.append((not list_schema.get("nullable", False), "[]"))
        list_schema = list_schema["items"]

    nested_prefix = list_schema.get("nullable", False) and "*" or ""
    nested_schema_name = schema_name(list_schema)
    nested_schema_name = f"{name_prefix}{nested_schema_name}" if nested_schema_name else "interface{}"
    schema_parts.append(
        (
            not list_schema.get("nullable", False),
            nested_prefix + (simple_type(list_schema) or nested_schema_name),
        )
    )
    nested_simple_type_name = "".join(p[1] for p in schema_parts)

    parameters = ""
    for d in data:
        value = format_data_with_schema(
            d,
            schema["items"],
            name_prefix=name_prefix,
            replace_values=replace_values,
            required=not schema["items"].get("nullable", False),
            in_list=True,
            **kwargs,
        )
        parameters += f"{value},\n"

    if in_list:
        return f"{{\n{parameters}}}"

    return f"{nested_simple_type_name}{{\n{parameters}}}"


@format_data_with_schema.register(dict)
def format_data_with_schema_dict(
    data,
    schema,
    name_prefix="",
    replace_values=None,
    default_name=None,
    required=False,
    in_list=False,
    **kwargs,
):
    if not schema:
        return ""

    reference = "" if required else "&"
    nullable = schema.get("nullable", False)

    name = schema_name(schema) or default_name

    parameters = ""
    if "properties" in schema:
        required_properties = set(schema.get("required", []))

        for k, v in data.items():
            if k not in schema["properties"]:
                continue
            value = format_data_with_schema(
                v,
                schema["properties"][k],
                name_prefix=name_prefix,
                replace_values=replace_values,
                default_name=name + camel_case(k) if name else None,
                required=k in required_properties,
                **kwargs,
            )
            parameters += f"{camel_case(k)}: {value},\n"

    if schema.get("additionalProperties"):
        saved_parameters = ""
        if schema.get("properties"):
            saved_parameters = parameters
            parameters = ""
        nested_schema = schema["additionalProperties"]
        nested_schema_name = simple_type(nested_schema)
        if not nested_schema_name:
            nested_schema_name = schema_name(nested_schema)
            if nested_schema_name:
                nested_schema_name = name_prefix + nested_schema_name

        has_properties = schema.get("properties")

        for k, v in data.items():
            if has_properties and k in schema["properties"]:
                continue
            value = format_data_with_schema(
                v,
                schema["additionalProperties"],
                name_prefix=name_prefix,
                replace_values=replace_values,
                required=True,
                **kwargs,
            )
            parameters += f'"{k}": {value},\n'

            # IMPROVE: find a better way to get nested schema name
            if not nested_schema_name:
                nested_schema_name = value.split("{")[0]

        if has_properties:
            if parameters:
                parameters = f"{saved_parameters}AdditionalProperties: map[string]{nested_schema_name}{{\n{parameters}}},\n"
            else:
                parameters = saved_parameters
        else:
            return f"map[string]{nested_schema_name}{{\n{parameters}}}"

    if "oneOf" in schema:
        matched = 0
        one_of_schema = None
        for sub_schema in schema["oneOf"]:
            try:
                if sub_schema.get("nullable") and data is None:
                    # only one schema can be nullable
                    formatted = "nil"
                else:
                    sub_schema["nullable"] = False
                    formatted = format_data_with_schema(
                        data,
                        sub_schema,
                        name_prefix=name_prefix,
                        replace_values=replace_values,
                        **kwargs,
                    )
                if matched == 0:
                    one_of_schema = sub_schema
                    # NOTE we do not support mixed schemas with oneOf
                    # parameters += formatted
                    parameters = formatted
                matched += 1
            except (KeyError, ValueError) as e:
                print(f"{e}")

        if matched == 0:
            raise ValueError(f"[{matched}] {data} is not valid for schema {name}")
        elif matched > 1:
            warnings.warn(f"[{matched}] {data} is not valid for schema {name}")

        one_of_schema_name = simple_type(one_of_schema) or f"{schema_name(one_of_schema)}"
        return f"{reference}{name_prefix}{name}{{\n{one_of_schema_name}: {parameters}}}"

    if schema.get("type") == "object" and "properties" not in schema:
        return "new(interface{})"

    if not name:
        warnings.warn(f"Unnamed schema {schema} for {data}")
        name_prefix = ""
        name = "map[string]interface"

    if parameters == "":
        # TODO select oneOf based on data
        warnings.warn(f"No schema matched for {data}")

    if nullable:
        return f"*{name_prefix}NewNullable{name}(&{name_prefix}{name}{{\n{parameters}}})"

    if in_list:
        return f"{{\n{parameters}}}"
    return f"{reference}{name_prefix}{name}{{\n{parameters}}}"
