#!/bin/bash

mkdir fis-receiver
cd fis-receiver

wget --no-check-certificate https://github.com/lrenc/fis-receiver/archive/master.zip

unzip master

cd fis-receiver-master

cp main ../

cd ../
rm master
rm -rf fis-receiver-master