### Golang Rest service - URL shortner
This service has two endpoints:
    1. one to create short url
    2. Another on to redirect shortURL to original URL
    
### Starting service
* for developement purpose use "make run"
* to run from docker use "make docker-run" (Make sure to download docker image from https://hub.docker.com/r/balrammani/url_shortner or create new image by using "make create-docker-image")

### Create ShortURL endpoint:

* Use Post call similar to localhost:8080/create_short_link?url=https://www.google.in/ to create short link

### Redirect ShortURL endpoint:
* Use Get call similar to localhost:8080/xnnFrDFW89
