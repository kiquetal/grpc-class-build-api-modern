#!/bin/bash

protoc blog/blogpb/blog.proto --go-grpc_out=. --go_out=.
