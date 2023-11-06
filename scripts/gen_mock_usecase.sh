if [[ ! "$1" =~ ^[a-z_]+$ ]]; then
  echo "Error: Usecase name is not in snake_case format"
  exit 1
fi

go_project=$(go list -m)
directory="mock/usecase/$1"
if [ ! -d "$directory" ]; then
  mkdir -p "$directory"
fi
mockgen -package=$1 ${go_project}/usecase/$1 UseCase > mock/usecase/$1/main.go
