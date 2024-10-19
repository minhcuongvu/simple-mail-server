#!/bin/bash
cd "$(dirname "$0")/.."
cp ./postfix/main.cf /etc/postfix/main.cf

if [[ -z "$SES_SMTP_SERVER" || -z "$SES_SMTP_USERNAME" || -z "$SES_SMTP_PASSWORD" ]]; then
  echo "Error: SES_SMTP_SERVER, SES_SMTP_USERNAME, or SES_SMTP_PASSWORD are not set in the environment."
  exit 1
fi
echo "$SES_SMTP_SERVER    $SES_SMTP_USERNAME:$SES_SMTP_PASSWORD" > /etc/postfix/sasl_passwd
chmod 600 /etc/postfix/sasl_passwd
postmap /etc/postfix/sasl_passwd

sed -i "s/^relayhost = .*/relayhost = [$SES_SMTP_SERVER]:$SES_SMTP_SERVER_PORT/" /etc/postfix/main.cf

service postfix start
