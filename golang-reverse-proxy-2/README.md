About
==

See : https://ben-appsec.github.io/golang-reverse-proxy-2-4/

Build and launch
==
`docker-compose up .`

Requesting 
==
Using the docker client :
`docker -H=127.0.0.1:8888 container ls`

Using curl :
`curl -Lv -k  http://localhost:8888/containers/json`