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
    "provider": "hetzner", // Available: updateUrl, goDaddy, hetzner, hetznerCloud
    "interval": "1h",
    "domain": "yourdomain.dom",
    "ipProvider": {
        "enableUpnp": true, // Enable upnp ip resolving
        "upnpGateway": "", // Manually specify a upnp gateway (your router). This is needed if the container is not running in netowrk mode host
        "enableExternal": true // Enable external ip services
    },
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
    },
    "hetznerCloudConfig": {
        "cloudApiToken": "HETZNER_CLOUD_API_TOKEN"
    }
}
```
