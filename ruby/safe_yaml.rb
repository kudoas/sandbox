file = Rails.root.join("config/database.yml").read

# Ruby3.0だとYAML.loadがセキュリティで警告が出る
# Ruby3.1だとYAML.loadがYAML.safe_loadと同じなので、警告はでない
# Prefer using `YAML.safe_load` over `YAML.load`. (convention:Security/YAMLLoad)
# https://github.com/rubocop/rubocop/blob/6b96cbd045f6f4dd9d33ead29219aa588e065129/config/default.yml#L2970
YAML.load(file)
