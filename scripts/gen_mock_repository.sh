if [[ ! "$1" =~ ^[a-z_]+$ ]]; then
  echo "Error: Repository name is not in snake_case format"
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

go_project=$(go list -m)
repository_name=$(snake_to_pascal "$1")
directory="mock/domain/repository"
if [ ! -d "$directory" ]; then
  mkdir -p "$directory"
fi
mockgen -package=repository ${go_project}/domain/repository ${repository_name}Repository > mock/domain/repository/$1.go
