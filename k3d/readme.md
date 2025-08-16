## Command


```bash
 k3d registry create k3d-kaixin-registry --port 12345
cluster create --registry-use k3d-kaixin-registry:12345 -p "8083:32005@loadbalancer" another-kaixin-demo --agents 2
```

## Update hostfile

`/etc/hosts`

```
127.0.0.1 kaixin-registry
```