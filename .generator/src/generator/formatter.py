"""Data formatter."""

from .utils import snake_case, camel_case, untitle_case, schema_name

KEYWORDS = {
    "break",
    "case",
    "chan",
    "const",
    "continue",
    "default",
    "defer",
    "else",
    "fallthrough",
    "for",
    "func",
    "go",
    "goto",
    "if",
    "import",
    "interface",
    "map",
    "package",
    "range",
    "return",
    "select",
    "struct",
    "switch",
    "type",
    "var",
}

SUFFIXES = {
    # Test
    "test",
    # $GOOS
    "aix",
    "android",
    "darwin",
    "dragonfly",
    "freebsd",
    "illumos",
    "js",
    "linux",
    "netbsd",
    "openbsd",
    "plan9",
    "solaris",
    "windows",
    # $GOARCH
    "386",
    "amd64",
    "arm",
    "arm64",
    "mips",
    "mips64",
    "mips64le",
    "mipsle",
    "ppc64",
    "ppc64le",
    "s390x",
    "wasm",
}


def block_comment(comment, prefix="#", first_line=True):
    lines = comment.split("\n")
    start = "" if first_line else lines[0] + "\n"
    return (start + "\n".join(f"{prefix} {line}".rstrip() for line in lines[(0 if first_line else 1) :])).rstrip()


def model_filename(name):
    filename = snake_case(name)
    last = filename.split("_")[-1]
    if last in SUFFIXES:
        filename += "_"
    return filename


def escape_reserved_keyword(word):
    """
    Escape reserved language keywords like openapi generator does it
    :param word: Word to escape
    :return: The escaped word if it was a reserved keyword, the word unchanged otherwise
    """
    if word in KEYWORDS:
        return f"{word}_"
    return word


def attribute_name(attribute):
    return escape_reserved_keyword(camel_case(attribute))


def variable_name(attribute):
    return escape_reserved_keyword(untitle_case(camel_case(attribute)))


def format_value(value, quotes='"', schema=None):
    if schema and "enum" in schema:
        index = schema["enum"].index(value)
        enum_varnames = schema["x-enum-varnames"][index]
        name = schema_name(schema)
        return f"{name.upper()}_{enum_varnames}"

    if isinstance(value, str):
        return f"{quotes}{value}{quotes}"
    elif isinstance(value, bool):
        return "true" if value else "false"
    elif value is None:
        return "nil"
    return value


def simple_type(schema, render_nullable=False):
    """Return the simple type of a schema.

    :param schema: The schema to extract the type from
    :return: The simple type name
    """
    type_name = schema.get("type")
    type_format = schema.get("format")
    nullable = render_nullable and schema.get("nullable", False)

    if type_name == "integer":
        return {
            "int32": "int32" if not nullable else "NullableInt32",
            "int64": "int64" if not nullable else "NullableInt64",
            None: "int32" if not nullable else "NullableInt32",
        }[type_format]

    if type_name == "number":
        return {
            "double": "float64" if not nullable else "NullableFloat64",
            None: "float" if not nullable else "NullableFloat",
        }[type_format]

    if type_name == "string":
        return {
            "date": "time.Time" if not nullable else "NullableTime",
            "date-time": "time.Time" if not nullable else "NullableTime",
            "email": "string" if not nullable else "NullableString",
            "binary": "*os.File",
            None: "string" if not nullable else "NullableString",
        }[type_format]
    if type_name == "boolean":
        return "bool" if not nullable else "NullableBool"

    return None


def is_reference(schema, attribute):
    """Check if an attribute is a reference."""
    is_required = attribute in schema.get("required", [])
    if is_required:
        return False

    attribute_schema = schema.get("properties", {}).get(attribute, {})

    is_nullable = attribute_schema.get("nullable", False)
    if is_nullable:
        return False

    is_anytype = attribute_schema.get("type", "object") == "object" and not (
        "properties" in attribute_schema
        or "oneOf" in attribute_schema
        or "anyOf" in attribute_schema
        or "allOf" in attribute_schema
    )
    if is_anytype:
        return False

    # no reference to arrays
    if attribute_schema.get("type", "object") == "array" or "items" in attribute_schema:
        return False

    return True


def attribute_path(attribute):
    return ".".join(attribute_name(a) for a in attribute.split("."))