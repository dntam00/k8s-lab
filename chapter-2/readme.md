
## pod

```
kubectl create -f basic.yaml
kubectl get pod
kubectl describe pod basicpod
kubectl delete pod basicpod
```

## service

```
kubectl create -f basicservice.yaml
kubectl get svc -o wide
kubectl delete svc basicservice
```

## deployment

```
kubectl create deployment firstpod --image=nginx
kubectl get deploy,rs,po,svc,ep
```

## namespace

```
kubectl get namespaces
kubectl get pod --all-namespaces
```