COVERAGE_FILE="profile.cov"

rm -f "$COVERAGE_FILE"

MODE_SET=false

for s in $(go list "./..."); do
  go test -coverprofile=profile-child.cov -failfast -p 1 "$s"
  if [ $? -eq 0 ]; then
    if [ "$MODE_SET" = false ]; then
      cat profile-child.cov >> "$COVERAGE_FILE"
      MODE_SET=true
    else
      sed '1d' profile-child.cov >> "$COVERAGE_FILE"
    fi
  else
    echo "Test failed for package: $s"
    exit 1
  fi
done

rm -f profile-child.cov

echo "Done"

