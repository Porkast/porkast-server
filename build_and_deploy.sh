go test ./...

if [[ $? -ne 0 ]]; then
    echo "The go test failed"
    exit 1
fi

./build.sh

echo "remove guoshaofm-web image"
docker rmi beegedelow/guoshaofm-web

./build_docker.sh

docker push beegedelow/guoshaofm-web
