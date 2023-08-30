#!/bin/bash
version=$(go run . -- --buildversion)
echo "$version"