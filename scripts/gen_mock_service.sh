go_project=$(go list -m)
directory="mock/pkg/$1"
if [ ! -d "$directory" ]; then
  mkdir -p "$directory"
fi
mockgen -package=$1 ${go_project}/pkg/$1 Service > mock/pkg/$1/$1.go


