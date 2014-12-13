#!/bin/bash

openssl genrsa -out private-key.pem -rand /var/log/messages 4096
openssl req -new -key private-key.pem -out csr.pem
openssl x509 -req -in csr.pem -signkey private-key.pem -out public-cert.pem
