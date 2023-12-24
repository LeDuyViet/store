#!/bin/sh

grep -v "replace\s.*=>.*" go.mod > tmpfile && mv tmpfile go.mod

go mod tidy

set -e

BIN_DIR=${BIN_DIR:-/app/bin}
mkdir -p "$BIN_DIR"

go install github.com/go-delve/delve/cmd/dlv@latest

echo "BIN_DIR: $BIN_DIR"

echo "****************************************"
echo "******** building applications **********"
echo "****************************************"

for file in *.go; do
    if [[ -f "$file" ]]; then
        echo building $file
        go build -o "$BIN_DIR"/${file%.go} $file
    fi
done

echo "****************************************"
echo "******** building workers **********"
echo "****************************************"

# Hàm đệ quy để duyệt qua các thư mục con
build_recursive() {
    local dir="$1"
    if [[ -d "$dir" ]]; then
        cd "$dir"
        for file in *.go; do
            if [[ -f "$file" ]]; then
                echo building "$file"
                go build -o "$BIN_DIR"/"${file%.go}" "$file"
            fi
        done
        for subdir in */; do
            if [[ -d "$subdir" ]]; then
                build_recursive "$subdir"
            fi
        done
        cd ..
    fi

    
}
build_recursive "cmd"

# ls -A | grep -Ev '/app/bin|/app/scripts' | xargs -d $"\n" rm -rf

# find . -name "*.go" -type f -delete