# Docker instructions (NOT WORKING FOR NOW)

## Build

- Run `build_docker.sh <notes folder path>` to build and run the docker container

## Rebuild

- Stop the docker container with `docker stop <CONTAINER ID>`
- Remove the docker container with `docker rm <CONTAINER ID>`
- Run `clean_docker.sh` to clean ALL images
- Go to `Build` section in this document

## Interact 

- Run `docker exec -it <CONTAINER ID> /bin/bash` to get inside the docker container
