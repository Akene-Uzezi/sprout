set -e

go build -o sprout ./cmd/main.go

sudo mv sprout /usr/bin/sprout
