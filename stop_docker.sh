docker stop $(docker ps | grep 'go-kb-server' | awk '{ print $1 }')
