#!/bin/bash

echo $(pwd)

./scripts/setup-postfix.sh

useradd -m info
mkdir -p /home/info/Maildir
chown -R info:info /home/info/Maildir

sleep infinity
