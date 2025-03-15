## selector

Filter resources based on labels.

```bash
kubectl create deployment design2 --image=nginx
kubectl get deployment design2 -o yaml > design2.yaml
kubectl get deployments.apps design2 -o wide

# -l: use label label
kubectl get -l app=design2 pod

kubectl edit pod <pod_name>
```

1. `-l`: only supports equality-based filtering (= and !=).

2. `--selector`: allows `in`, `notin`, and `exists` operators.

## labels

When using deployment to control pod, if I edit the label `app` of the existing pod, then deployment will create another pod to match with current definition.

If I delete the deployment, the existing pod will still exist.

## logs

```
kubectl logs <pod_name> -c <container_name>
```

## resources limit

## custom resource definition

```
kubectl get crd
```


