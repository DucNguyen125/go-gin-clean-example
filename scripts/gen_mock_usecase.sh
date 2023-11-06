go_project=$(go list -m)
directory="mock/usecase/$1"
if [ ! -d "$directory" ]; then
  mkdir -p "$directory"
fi
mockgen -package=$1 ${go_project}/usecase/$1 UseCase > mock/usecase/$1/main.go


