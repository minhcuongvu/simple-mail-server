#!/bin/bash
docker build -t simple-mail-server .
docker stop simple-mail-server
docker rm simple-mail-server
docker run -it -d -p 25:25 --env-file .env --name simple-mail-server simple-mail-server

