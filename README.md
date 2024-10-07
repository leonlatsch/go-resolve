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

### Config

```json
{
    "provider": "hetzner", // Available: updateUrl, goDaddy, hetzner
    "interval": "1h",
    "domain": "yourdomain.dom",
    "hosts": [
        "subdomain1",
        "subdomain2"
    ],
    "updateUrlConfig": {
        "url": "UPDATE_URL"
    },
    "goDaddyConfig": {
        "apiKey": "GODADDY_API_KEY",
        "apiSecret": "GODADDY_API_SECRET"
    },
    "hetznerConfig": {
        "zoneId": "HETZNER_ZONE_ID",
        "apiToken": "HETZNER_API_TOKEN"
    }
}
```
