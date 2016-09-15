# appcelerator/pinger

Simple service for testing purposes. Responds with `pong` to a `GET` request at: `http://<host>:<port>/ping`, and logs to `stdout`.

## Running pinger

    $ docker run -it --rm -p 3000:3000 --name pinger appcelerator/pinger

or

    $ docker service create -p 3000:3000 --name pinger appcelerator/pinger

## Development

`make.sh` - runs the go compiler in a container to create an alpine binary

`build.sh` - builds an alpine-based image: `appcelerator/pinger`

`run.sh` - convenience script for `docker run ...`

`service-create.sh` - convenience script for `docker service create ...`
