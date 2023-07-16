#!/bin/sh

set -e

host="$1"
port="$2"

until mysqladmin ping -h "$host" -P "$port" --silent; do
  >&2 echo "MySQL is unavailable - sleeping"
  sleep 1
done

>&2 echo "MySQL is up - executing command"
