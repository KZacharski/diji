#!/bin/bash
echo "diji-build 1.0"
version=$(go run . -- --buildversion)
echo "Building diji $version"
rldir="release"

if [ -f "$rldir" ] ; then
    rm -rf "$rldir"
    echo "Removed files from the previous build"
fi
mkdir release
mkdir release/diji
cp -r diji-config release/diji
echo "diji-config prepared"
GOOS=linux GOARCH=amd64 go build -o release/diji/diji
(cd release; zip -r diji-$version-linux-amd64.zip diji; rm -rf diji/diji)
echo "linux/amd64 built"
GOOS=linux GOARCH=386 go build -o release/diji/diji
(cd release; zip -r diji-$version-linux-i386.zip diji; rm -rf diji/diji)
echo "linux/386 built"
GOOS=linux GOARCH=arm64 go build -o release/diji/diji
(cd release; zip -r diji-$version-linux-arm64.zip diji; rm -rf diji/diji)
echo "linux/arm64 built"
GOOS=linux GOARCH=arm go build -o release/diji/diji
(cd release; zip -r diji-$version-linux-arm.zip diji; rm -rf diji/diji)
echo "linux/arm built"
GOOS=darwin GOARCH=amd64 go build -o release/diji/diji
(cd release; zip -r diji-$version-macos-amd64.zip diji; rm -rf diji/diji)
echo "darwin/amd64 built"
GOOS=darwin GOARCH=arm64 go build -o release/diji/diji
(cd release; zip -r diji-$version-macos-arm64.zip diji; rm -rf diji/diji)
echo "darwin/arm64 built"
GOOS=windows GOARCH=amd64 go build -o release/diji/diji.exe
(cd release; zip -r diji-$version-windows-amd64.zip diji; rm -rf diji/diji.exe)
echo "windows/amd64 built"
GOOS=windows GOARCH=386 go build -o release/diji/diji.exe
(cd release; zip -r diji-$version-windows-i386.zip diji; rm -rf diji/diji.exe)
echo "windows/386 built"
GOOS=windows GOARCH=arm64 go build -o release/diji/diji.exe
(cd release; zip -r diji-$version-windows-arm64.zip diji; rm -rf diji/diji.exe)
echo "windows/arm64 built"
GOOS=windows GOARCH=arm go build -o release/diji/diji.exe
(cd release; zip -r diji-$version-windows-arm.zip diji; rm -rf diji/diji.exe)
echo "windows/arm built"
rm -rf release/diji
echo "Cleaned up"
echo "Done, the zips are ready in the release directory"