import os
import json
import re

go_code_dir = "./go/"

settings_file = "../.vscode/settings.json"


def get_build_tags(directory):
    tags = set()
    for filename in os.listdir(directory):
        if filename.endswith(".go"):
            with open(os.path.join(directory, filename), "r") as file:
                first_line = file.readline().strip()
                match = re.search(r"//go:build\s+(\d+)", first_line)
                if match:
                    tags.add(match.group(1))
    return sorted(list(tags))


def update_vscode_settings(settings_file, tags):
    try:
        with open(settings_file, "r") as file:
            settings = json.load(file)
    except FileNotFoundError:
        settings = {}

    if "gopls" not in settings:
        settings["gopls"] = {}

    settings["gopls"]["buildFlags"] = ["-tags=" + ",".join(tags)]

    with open(settings_file, "w") as file:
        json.dump(settings, file, indent=2)


if __name__ == "__main__":
    tags = get_build_tags(go_code_dir)
    update_vscode_settings(settings_file, tags)
    print(f"Updated VSCode settings with tags: {', '.join(tags)}")
