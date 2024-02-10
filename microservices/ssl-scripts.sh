#! /bin/bash
SERVER_CN=localhost
openssl genrsa -out ca.key 4096
openssl req -x509 -new -nodes -key ca.key -sha256 -days 365 -out ca.crt -subj "/CN=${SERVER_CN}"
openssl genrsa -out server.key 4096
openssl req -new -key server.key -out server.csr -subj "/CN=${SERVER_CN}"
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365 -sha256
openssl pkcs8 -topk8 -inform PEM -outform PEM -in server.key -out server.pem -nocrypt
