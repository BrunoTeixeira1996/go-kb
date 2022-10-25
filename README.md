# go-kb

I tend to create notes in markdown files and those notes are everywhere.
This projects aims to help me visualize notes, like a knowledge base.
This was built to improve my Go experience.


# Build instructions

The app will be listening on port `8080`

## Build localy

```shell
git clone https://github.com/BrunoTeixeira1996/go-kb
cd go-kb
go build
./go-kb <full path of the notes folder>
```

## Docker

### Build

- Run `build_docker.sh <notes folder full path>` to build and run the docker container

### Rebuild

- TODO

### Interact 

- Run `docker exec -it <CONTAINER ID> /bin/bash` to get inside the docker container

# Images

If you want to use images, create an `images` folder inside your `notes` folder and place the images in there.
Then, if you want to call those images on the markdown file, you simply need to do `![image1](/images/image1.png)`

# Todo

Check `TODO.md`
