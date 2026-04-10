# coding=utf-8
"""Unit tests for the formatter module."""

from generator.formatter import format_data_with_schema


class SchemaWithRef(dict):
    """A schema dict that carries a $ref, as the generator produces after resolving references."""

    def __init__(self, *args, ref=None, **kwargs):
        super().__init__(*args, **kwargs)
        if ref:
            self.__reference__ = {"$ref": ref}


class TestFormatDataWithSchemaAdditionalProperties:
    """Tests for format_data_with_schema with additionalProperties schemas."""

    def test_empty_additional_properties_with_nullable_generates_map_literal(self):
        """additionalProperties: {} + nullable: true must not generate NewNullableXxx.

        Before the fix, the Python truthiness of {} caused the additionalProperties
        block to be skipped, falling through to the nullable path which generated
        *NewNullableXxx(&Xxx{}) — a function that is never emitted by the model generator
        for map schemas, causing a Go compile error.
        """
        schema = SchemaWithRef(
            {"nullable": True, "additionalProperties": {}},
            ref="#/components/schemas/CustomFields",
        )
        result = format_data_with_schema({"key": "value"}, schema)
        assert result == 'map[string]interface{}{\n"key": "value",\n}'
        assert "NewNullable" not in result

    def test_empty_additional_properties_with_type_object_generates_map_literal(self):
        """type: object + additionalProperties: {} + nullable: true must not generate NewNullableXxx."""
        schema = SchemaWithRef(
            {"type": "object", "nullable": True, "additionalProperties": {}},
            ref="#/components/schemas/IncidentImpactFieldsObject",
        )
        result = format_data_with_schema({"field": "val"}, schema)
        assert result == 'map[string]interface{}{\n"field": "val",\n}'
        assert "NewNullable" not in result

    def test_typed_additional_properties_generates_typed_map(self):
        """additionalProperties: {type: string} must generate map[string]string{} as before."""
        schema = SchemaWithRef(
            {"type": "object", "nullable": True, "additionalProperties": {"type": "string"}},
            ref="#/components/schemas/StringMap",
        )
        result = format_data_with_schema({"k": "v"}, schema)
        assert result == 'map[string]string{\n"k": "v",\n}'
