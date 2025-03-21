## CA

```bash
docker cp k3d-kaixin-demo-server-0:./var/lib/rancher/k3s/agent/server-ca.crt .
docker cp k3d-kaixin-demo-server-0:./var/lib/rancher/k3s/agent/client-ca.crt .
docker cp k3d-kaixin-demo-server-0./var/lib/rancher/k3s/server/tls/client-ca.key .


curl --cert ./client-ca.crt --key ./client-ca.key --cacert ./server-ca.crt https://0.0.0.0:62639/version
```

```json
{
  "major": "1",
  "minor": "31",
  "gitVersion": "v1.31.5+k3s1",
  "gitCommit": "56ec5dd4d012c1d77dc7f50f21f65183044c92e7",
  "gitTreeState": "clean",
  "buildDate": "2025-01-28T16:50:04Z",
  "goVersion": "go1.22.10",
  "compiler": "gc",
  "platform": "linux/arm64"
}
```