#!/bin/bash
rm -rf doc
mkdir doc
curl -o swagger.tar.gz https://codeload.github.com/swagger-api/swagger-ui/tar.gz/master
tar -xvf swagger.tar.gz
rm -f swagger.tar.gz
mv -f swagger-ui-master/dist/* doc/
rm -rf swagger-ui-master doc/swagger-ui.js
sed -i_ "s/swagger-ui\.js/swagger-ui\.min\.js/" doc/index.html
sed -i_ "s/http:\/\/petstore\.swagger\.io\/v2\///" doc/index.html
rm -f doc/*_
swagger generate spec -b ../impl/ -o swagger.json
cd ../ipml
go build -o ../server
cd ..
