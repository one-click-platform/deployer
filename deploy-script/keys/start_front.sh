#!/bin/bash
su

docker login docker.pkg.github.com --username xxxvik-xakerxxx --password 510bb6a0997c9b621ccaa5a1a7dd6f1e0955ce58

docker pull docker.pkg.github.com/one-click-platform/web-client/web_eth:latest

docker run -v /home/ubuntu/env.js:/usr/share/nginx/html/static/env.js -p 81:80 -d docker.pkg.github.com/one-click-platform/web-client/web_eth

systemctl restart nginx.service