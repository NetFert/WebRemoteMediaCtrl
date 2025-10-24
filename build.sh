#!/bin/bash

PROJECT_NAME="WebRemoteMediaCtrl"
OUTPUT_DIR="./dist"
mkdir -p $OUTPUT_DIR

SRC_DIR=$(pwd)

echo "$PROJECT_NAME ..."

# 1. macOS x86_64 (Intel)
echo "macOS x86_64..."
GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o $OUTPUT_DIR/${PROJECT_NAME}_mac_x86_64 $SRC_DIR

# 2. macOS arm64 (M1/M2)
echo "macOS arm64..."
GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build -o $OUTPUT_DIR/${PROJECT_NAME}_mac_arm64 $SRC_DIR

# 3. Windows 32-bit x86_32
echo "Windows 32-bit (i386)..."
GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ CGO_ENABLED=1 go build -o $OUTPUT_DIR/${PROJECT_NAME}_win_i386.exe $SRC_DIR

# 4. Windows 64-bit x86_64
echo "Windows 64-bit (x86_64)..."
GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ CGO_ENABLED=1 go build -o $OUTPUT_DIR/${PROJECT_NAME}_win_x86_64.exe $SRC_DIR

echo "Done $OUTPUT_DIR"

cd 'dist' || exit

for file in *; do
    [[ -f "$file" ]] || continue
    [[ "$file" == *.zip ]] && continue
    zip_name="${file%.*}.zip"
    echo "📦 zip: $file -> $zip_name"
    if zip -q "$zip_name" "$file"; then
        echo "✅ delete file: $file"
        rm "$file"
    else
        echo "❌: $file"
    fi
done