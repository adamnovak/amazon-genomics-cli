## Toil AWS Mirror

A Toil mono-container WES server for use with Amazon AGC.

### Building the Container Manually

Go to this directory and run:

```bash
docker build . -f Dockerfile -t adamnovak/toil-agc
```

### Running for Testing

Having built the container, run:

```bash
docker run -ti --rm -p "127.0.0.1:8000:8001" adamnovak/toil-agc
```

This will start the containerized server and make it available on port 8000 on the loopback interface. You can inspect the port mapping with:

```bash
docker port "$(docker ps | grep adamnovak/toil-agc | rev | cut -f1 -d' ' | rev)"
```

Then you can talk to it with e.g.:

```bash
curl -vvv "http://localhost:8000/ga4gh/wes/v1/service-info"
```

For debugging, you can get inside the container with:

```bash
docker exec -ti "$(docker ps | grep adamnovak/toil-agc | rev | cut -f1 -d' ' | rev)" /bin/bash
```
