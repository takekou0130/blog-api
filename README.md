## use docker

build image

```bash
docker image build -t [tag]:[version] .
ex) docker image build -t gin:latest .
```

run container

```bash
docker container run -it --rm -p [port]:[port] --name [name] [image tag]:[version tag]
ex) docker container run -it --rm -p 3030:3030 --name gin gin:latest
```
