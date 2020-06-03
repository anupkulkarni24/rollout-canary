# Blue Green
You are going to create a Rollout with a blue-green strategy and play around with the different functionalities available.

## The Basics

1. Create the [namespace](https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/blue-green/example/namespace.yaml) for the resources by running:
    ```bash
    kubectl apply -f https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/blue-green/example/namespace.yaml
    ```
1. Create the following [active service](https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/blue-green/example/active-service.yaml) and [preview service](https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/blue-green/example/preview-service.yaml) by running:
    ```bash
    kubectl apply -f https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/blue-green/example/active-service.yaml
    kubectl apply -f https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/blue-green/example/preview-service.yaml
    ```
1. Run `kubectl get service bluegreen-service -o yaml -w` in a separate tab and do not close it.
1. Please post the yaml result here
    ```yaml
    




    
    ```
1. Create the [Rollout](https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/blue-green/example/rollout.yaml) by running: 
   ```bash
   kubectl apply -f https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/blue-green/example/rollout.yaml
   ```
1. How does the bluegreen-service Service change?
     ```
         




    
     ```
     * If you do not see the difference, try deleting the Rollout and looking at the service again. What was removed?
     * Once you see the difference, apply the rollout if you deleted
1. Run `kubectl get pods` and make sure that all the pods are running. Afterward, run `kubectl get rollouts rollout-bluegreen -o yaml` and post the results here.
     ```yaml
         




    
     ```
1. Create the [Ingress](https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/blue-green/example/ingress.yaml) objects by running the follow: 
    ```bash
    kubectl apply -f https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/blue-green/example/ingress.yaml
   ```
   * Run `kubectl get ingress -o wide` to get an IP address you can hit for the active and preview Ingresses.
## Changing the image
1. Open the active service IP, take a screenshot of the UI and post it here
    ```






   ```
1. Change the image from blue to green by running:
    ```bash
    kubectl argo rollouts set image bluegreen-demo "*=argoproj.io/rollouts-demo:green"
    ```
1. Open the preview and active URLs (from the ingress), take a screenshot of the UI each one and post it here
   ```






   ```
1. Run `kubectl get service bluegreen-demo` and `kubectl get service bluegreen-demo-preview` describe the differences below:
    ```






    ```
2.  Run `kubectl get rollouts rollout-bluegreen -o yaml` and post the results here. How is this rollout different than the previous rollouts?
    <details>
    <summary>Hint:</summary>

    Look at Status.BlueGreen fields
    </details>

    ```






    ```
3.  Promote the preview version to the active by running:
    ```
    kubectl argo rollouts promote bluegreen-demo
    ```
    * Look at the bluegreen-service yaml again. How has it changed?
4.  Run `kubectl get rollouts rollout-bluegreen -o yaml` again and compare it to the previous. How has it changed?

## Scale Down Delay

1. Read the important notice in https://argoproj.github.io/argo-rollouts/features/bluegreen/
   *  Each Kubernetes Service with any label selector has a corresponding Kubernetes endpoint, and that Kubernetes Endpoint holds all the IPs of the pods that the Service selects. When a pod is added/removed or the service selector changes, the Endpoint controller will update the corresponding endpoint to match the new set of Pods selected by the Service’s selector. This endpoint is used by kube-proxy to understand how to route traffic in the cluster by using a tool called IP tables. Since Kube-Proxy runs on every node, every node needs to process the change in Endpoints. If the traffic from the application is high enough or the cluster is large enough, there’s a high chance traffic will go to the previous version. 
1. Run `kubectl get endpoint bluegreen-service -o yaml` and then `kubectl get pods -o wide`. Each IP listed in the endpoint will be one of the pod IPs
1. Take a look at the following [rollout](https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/blue-green/example/rollout-with-increased-delay.yaml)
   * What’s different about this rollout compared to the previous?
      <details>
      <summary>Answer:</summary>
      The Rollout now has a scaleDownDelaySeconds field in the blueGreen section with a value of 600 and the autoPromotionEnabled is set to true.
      </details>
1. Run the following command in a seperate tab to create a watch on a Rollout:
   ```bash
   kubectl argo rollouts get rollout bluegreen-demo
   ```
1. Run the following command to apply the new version from above:
   ```bash
   kubectl apply -f https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/blue-green/example/rollout-with-increased-delay.yaml
   ```
   * What happened in the terminal with the watch? How long do you think the previous version will stick around?

## Other Features

1. Take a look at [Blue Green docs](https://argoproj.github.io/argo-rollouts/features/bluegreen/) and try modifying the rollout to add some of these fields.
What happens when you change the image to a different color with these fields?

## Clean up
Run the following command to delete the namespaces:
```bash
kubectl delete namespace bluegreen-demo
```

This will delete all the resources with the cluster.
