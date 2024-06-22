#!/bin/bash

# Generate Go files
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./classification.proto

# Generate Python files
python -m grpc_tools.protoc -I./ --python_out=. --pyi_out=. --grpc_python_out=. ./classification.proto
