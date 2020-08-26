docker ps --filter "name=golang-server"
docker pull $IMAGE
docker rm -f golang-server || true
docker run -d \
           --name golang-server \
           -p 8080:8080 \
           $IMAGE
docker ps --filter "name=golang-server"
