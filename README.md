# echo
A simple dockerized echo server listening on port 2345

# Setup Instructions
* sudo yum install docker-io
* sudo docker build -t echo .
* sudo docker run --publish 2345:2345 --name echo --rm echo
* sudo docker images
