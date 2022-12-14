if [ $# -lt 1 ]; then
    echo Please provide the notes path
    exit 1
fi

echo Building docker container ...
docker build . -t go-kb-server

echo Running docker container ...
docker run -t -d --restart unless-stopped -v $1:/app/notes -p 8080:8080 go-kb-server /app/notes
