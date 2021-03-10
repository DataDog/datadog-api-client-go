#!/usr/bin/env python3
import csv
import pathlib
import re
import yaml

RE_COPYRIGHT = re.compile("Copyright [^ ]+ (.*)$", re.MULTILINE)


def copyrights(info):
    for license in info["licenses"]:
        for match in re.findall(RE_COPYRIGHT, license["text"]):
            yield match.replace(",", "").replace("  ", " ")


fieldnames = ["Origin", "License", "Authors"]

with (pathlib.Path(__file__).parent.parent / "LICENSE-3rdparty.csv").open(
    "w", newline=""
) as output:
    writer = csv.DictWriter(output, fieldnames=fieldnames)
    writer.writeheader()

    for dep in sorted(pathlib.Path(__file__).parent.rglob("*.dep.yml")):
        with dep.open() as fp:
            info = yaml.safe_load(fp)
            writer.writerow(
                {
                    "Origin": info["name"],
                    "License": info["license"],
                    "Authors": " ".join(copyrights(info)),
                }
            )
