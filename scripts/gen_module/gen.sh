if [[ ! "$1" =~ ^[a-z_]+$ ]]; then
  echo "Error: Module name is not in snake_case format"
  exit 1
fi

snake_to_pascal() {
  local input="$1"
  local result=""
  local make_upper=1

  for ((i = 0; i < ${#input}; i++)); do
    local char="${input:i:1}"

    if [[ "$char" == "_" ]]; then
      make_upper=1
    else
      if [ "$make_upper" -eq 1 ]; then
        char="$(tr '[:lower:]' '[:upper:]' <<< "$char")"
        make_upper=0
      else
        char="$(tr '[:upper:]' '[:lower:]' <<< "$char")"
      fi
      result="${result}${char}"
    fi
  done

  echo "$result"
}

snake_to_camel() {
  local input="$1"
  local result=""
  local make_upper=0

  for ((i = 0; i < ${#input}; i++)); do
    local char="${input:i:1}"

    if [[ "$char" == "_" ]]; then
      make_upper=1
    else
      if [ "$make_upper" -eq 1 ]; then
        char="$(tr '[:lower:]' '[:upper:]' <<< "$char")"
        make_upper=0
      fi
      result="${result}${char}"
    fi
  done

  echo "$result"
}


module_pascal_name=$(snake_to_pascal "$1")
module_camel_name=$(snake_to_camel "$1")
file_list=$(find ./scripts/gen_module -type f -name "*.gotmpl")
for file in $file_list; do
  source_file="$file"
  destination_file="${source_file#./scripts/gen_module/}"
  destination_file="${destination_file//\{\{module\}\}/$1}"
  destination_file="${destination_file%.gotmpl}.go"
  destination_directory=$(dirname "$destination_file")
  if [ ! -d "$destination_directory" ]; then
    mkdir -p "$destination_directory"
  fi
  cp "$source_file" "$destination_file"
  sed -i "s/{{module_snake}}/$1/g" "$destination_file"
  sed -i "s/{{module}}/${module_camel_name}/g" "$destination_file"
  sed -i "s/{{Module}}/${module_pascal_name}/g" "$destination_file"
done

  model="
type ${module_pascal_name} struct {
	internal.${module_pascal_name}
	BaseModel
}"
  echo -e "$model" >> "infra/postgresql/model/model.go"
