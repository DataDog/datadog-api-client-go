import pathlib

import click
from jinja2 import Environment, FileSystemLoader

from . import openapi
from . import formatter
from . import utils

PACKAGE_NAME = "datadog"


@click.command()
@click.option(
    "-i",
    "--input",
    type=click.Path(exists=True, file_okay=True, dir_okay=False, path_type=pathlib.Path),
)
@click.option(
    "-o",
    "--output",
    type=click.Path(path_type=pathlib.Path),
)
def cli(input, output):
    """
    Generate a Ruby code snippet from OpenAPI specification.
    """
    spec = openapi.load(input)

    version = input.parent.name

    env = Environment(loader=FileSystemLoader(str(pathlib.Path(__file__).parent / "templates")))

    env.filters["accept_headers"] = openapi.accept_headers
    env.filters["attribute_name"] = formatter.attribute_name
    env.filters["block_comment"] = formatter.block_comment
    env.filters["camel_case"] = formatter.camel_case
    env.filters["collection_format"] = openapi.collection_format
    env.filters["format_server"] = openapi.format_server
    env.filters["format_value"] = formatter.format_value
    env.filters["is_reference"] = formatter.is_reference
    env.filters["parameter_schema"] = openapi.parameter_schema
    env.filters["parameters"] = openapi.parameters
    env.filters["form_parameter"] = openapi.form_parameter
    env.filters["response_type"] = openapi.get_type_for_response
    env.filters["return_type"] = openapi.return_type
    env.filters["simple_type"] = formatter.simple_type
    env.filters["snake_case"] = formatter.snake_case
    env.filters["untitle_case"] = formatter.untitle_case
    env.filters["upperfirst"] = utils.upperfirst
    env.filters["variable_name"] = formatter.variable_name

    env.globals["enumerate"] = enumerate
    env.globals["get_name"] = openapi.get_name
    env.globals["get_type_for_attribute"] = openapi.get_type_for_attribute
    env.globals["get_type_for_parameter"] = openapi.get_type_for_parameter
    env.globals["get_type"] = openapi.type_to_go
    env.globals["openapi"] = spec
    env.globals["package_name"] = PACKAGE_NAME
    env.globals["version"] = version
    env.globals["get_default"] = openapi.get_default
    env.globals["get_container"] = openapi.get_container
    env.globals["get_container_type"] = openapi.get_container_type
    env.globals["get_type_at_path"] = openapi.get_type_at_path

    api_j2 = env.get_template("api.j2")
    model_j2 = env.get_template("model.j2")

    extra_files = {
        "client.go": env.get_template("client.j2"),
        "configuration.go": env.get_template("configuration.j2"),
        "utils.go": env.get_template("utils.j2"),
        "zstd.go": env.get_template("zstd.j2"),
        "no_zstd.go": env.get_template("no_zstd.j2"),
    }

    apis = openapi.apis(spec)
    models = openapi.models(spec)

    output.mkdir(parents=True, exist_ok=True)

    for name, template in extra_files.items():
        filename = output / name
        with filename.open("w") as fp:
            fp.write(template.render(apis=apis, models=models))

    for name, model in models.items():
        filename = "model_" + formatter.model_filename(name) + ".go"
        model_path = output / filename
        model_path.parent.mkdir(parents=True, exist_ok=True)
        with model_path.open("w") as fp:
            fp.write(model_j2.render(name=name, model=model, models=models))

    for name, operations in apis.items():
        filename = "api_" + formatter.snake_case(name) + ".go"
        api_path = output / filename
        api_path.parent.mkdir(parents=True, exist_ok=True)
        with api_path.open("w") as fp:
            fp.write(api_j2.render(name=name, operations=operations))
