# appcelerator/pinger

Simple service for testing purposes. Responds with `pong` to a `GET` request at: `http://<host>:<port>/ping`.

`make.sh` - runs the go compiler in a container to create an alpine binary

`build.sh` - builds an alpine-based image: `appcelerator/pinger`

`run.sh` - convenience script for `docker run ...`

`service-create.sh` - convenience script for `docker service create ...`
