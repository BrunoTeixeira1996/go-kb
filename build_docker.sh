echo y | docker system prune -a # comment this when everything is working
docker build . -t go-kb-server
docker run -t -d -v /notes:/notes -p 8080:8080 go-kb-server notes
