echo "stop guoshaofm-web"
docker stop guoshaofm-web
echo "remove guoshaofm-web container"
docker container rm guoshaofm-web
echo "remove guoshaofm-web image"
docker rmi guoshaofm-web
echo "pull guoshaofm-web image"
docker pull beegedelow/guoshaofm-web

LOGS_DIR=/home/guoshaofm-web/logs
if [[ ! -e $LOGS_DIR ]]; then
    mkdir -p $LOGS_DIR
elif [[ ! -d $LOGS_DIR ]]; then
    echo "$LOGS_DIR already exists but is not a directory" 1>&2
fi
echo "run guoshaofm-web container"
docker run --name guoshaofm-web --network host --log-opt max-size=500m --log-opt max-file=3 -v $LOGS_DIR:/app/logs -d beegedelow/guoshaofm-web
