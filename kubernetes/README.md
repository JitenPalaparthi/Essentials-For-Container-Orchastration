

- Only for development purpose 

 ```
 minikube tunnel
 ```


- To check and get ip and port of a service in minikube

```
minikube service nginx-service -n default
```

- To scale deployment

 ```
 kubectl scale deployment --replicas=1 nginx-deployment

 ```

 - To work with different context

 ```
 kubectl config current-context
 kubectl config use-context aws-dev-context
 ```
 
 - Fetch all information from minikube context (aka local server)

 ```
 kubectl --context=minikube get all -A 
 ```
