## configMaps

```bash
kubectl create configmap colors \
--from-literal=text=black \
--from-file=./favorite \
--from-file=./primary/

kubectl get configmap colors
```

![env](../images/env.png)