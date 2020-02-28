#!/bin/bash
set -eu

protoc -I proto proto/*.proto --go_out=plugins=grpc:proto
protoc -I proto proto/*.proto --swift_out=grpcvoiceover/grpcvoiceover/Proto --grpc-swift_out=grpcvoiceover/grpcvoiceover/Proto