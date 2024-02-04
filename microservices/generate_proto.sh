#!/bin/bash

protoc greet/greetpb/greet.proto --go-grpc_out=. --go_out=.
protoc calculator/calculatorpb/calculator.proto --go-grpc_out=. --go_out=.
