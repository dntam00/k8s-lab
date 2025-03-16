## configMaps

```bash
kubectl create configmap colors \
--from-literal=text=black \
--from-file=./favorite \
--from-file=./primary/

kubectl get configmap colors
```

Get env variables inside a container.
```
kubectl exec -it <pod_name> -c simpleapp -- env

PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=try1-68484d7ff9-q8g95
yellow=y

ilike=blue

black=k
known as key

cyan=c

favorite=blue

magenta=m

text=black
KUBERNETES_PORT_443_TCP_PORT=443
KUBERNETES_PORT_443_TCP_ADDR=10.43.0.1
KUBERNETES_SERVICE_HOST=10.43.0.1
KUBERNETES_SERVICE_PORT=443
KUBERNETES_SERVICE_PORT_HTTPS=443
KUBERNETES_PORT=tcp://10.43.0.1:443
KUBERNETES_PORT_443_TCP=tcp://10.43.0.1:443
KUBERNETES_PORT_443_TCP_PROTO=tcp
TERM=xterm
HOME=/root
```

![env](../images/env.png)

## pv/pvc

```bash
kubectl apply -f pv.yaml
kubectl apply -f pvc.yaml
kubectl apply -f nginx.yaml
kubectl patch pv nfs-pv -p '{"spec":{"claimRef": null}}'
```