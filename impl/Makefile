swagger:
    doc/index.html
    doc/index.html:
        @mkdir -p doc
        curl https://codeload.github.com/swagger-api/swagger-ui/tar.gz/master | tar xzv -C doc swagger-ui-master/dist
        mv -f doc/swagger-ui-master/dist/* doc/
        rm -rf doc/swagger-ui-master doc/swagger-ui.js
        sed -i_ "s/swagger-ui\.js/swagger-ui\.min\.js/" doc/index.html
        sed -i_ "s/http:\/\/petstore\.swagger\.io\/v2\///" doc/index.html
        rm -f doc/*_
