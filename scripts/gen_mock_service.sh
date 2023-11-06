if [[ ! "$1" =~ ^[a-z_]+$ ]]; then
  echo "Error: Service name is not in snake_case format"
  exit 1
fi

go_project=$(go list -m)
directory="mock/pkg/$1"
if [ ! -d "$directory" ]; then
  mkdir -p "$directory"
fi
mockgen -package=$1 ${go_project}/pkg/$1 Service > mock/pkg/$1/$1.go
