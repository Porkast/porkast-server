./build.sh

echo "remove guoshaofm-web image"
docker rmi beegedelow/guoshaofm-web

./build_docker.sh

docker push beegedelow/guoshaofm-web
