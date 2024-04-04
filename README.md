# go-resolve

Dyndns without dyndns. A rewrite of [godaddy-dyndns](https://github.com/leonlatsch/godaddy-dyndns) in go.

### Setup with docker

```
docker volume create go-resolve-config
docker run -d \ 
-v go-resolve-config:/go-resolve/config \
--name go-resolve \
ghcr.io/leonlatsch/go-resolve:latest
```
