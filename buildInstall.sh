#! /usr/bin/env bash

set -aeEux
set -o pipefail


echo "@@@ lets start"
rm -rf tmp
uname -a
yum install -y wget
mkdir tmp
wget https://dl.google.com/go/go1.13.10.linux-amd64.tar.gz
rm -rf tmp
echo "@@@ finish"