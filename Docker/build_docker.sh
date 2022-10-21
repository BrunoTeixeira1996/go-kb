#echo y | docker system prune -a # comment this when everything is working

if [ $# -lt 1 ]; then
    echo Please provide the notes path
    exit 1
fi

echo Building docker container ...
docker build . -t go-kb-server

echo Running docker container ...
docker run -t -d -p 8080:8080 go-kb-server $(basename $1)
