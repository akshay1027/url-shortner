
## Learning

[Should we check for context done/cancel on every handler](https://github.com/gofiber/fiber/issues/805#issuecomment-699614997)

## Docker commands:

### Delete all docker containers
docker container rm $(docker container ls -aq)
docker rmi - f $(docker images -aq)

### Stop a conatiner
docker stop 796d2ed8c94e

### Show all running containers
docker container ls

### Create docker container 
docker-compose up -d