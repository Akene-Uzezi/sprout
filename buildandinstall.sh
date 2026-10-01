set -e

go build -o sprout main.go

sudo mv sprout /usr/bin/sprout
