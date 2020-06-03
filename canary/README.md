# Canary
You are going to create a Rollout with a canary strategy and play around with the different functionalities available.

## Mimicking Rolling Update

1. Create the [namespace](https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/canary/example/namespace.yaml), [service](https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/canary/example/service.yaml), [rollout](https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/canary/example/rollout.yaml), and [ingress](https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/canary/example/ingress.yaml) by running:
    ```bash
    kubectl apply -f https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/canary/example/namespace.yaml
    kubectl apply -f https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/canary/example/rollout.yaml
    kubectl apply -f https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/canary/example/service.yaml
    kubectl apply -f https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/canary/example/ingress.yaml
    ```
1. Set the max surge to 0 and max unavailable to 1 in the Rollout
   * See [Canary docs](https://argoproj.github.io/argo-rollouts/features/canary/#other-configurable-features) for more info.
   * You can either have to run `kubectl edit rollout canary-demo` or change the rollout.yaml and kubectl apply it to edit the rollout.
1. In a separate tab, run `kubectl get rs -o wide -w` and then change the rollout image by running:
   ```bash
   kubectl argo rollouts set image canary-demo "*=argoproj/rollouts-demo:green"
   ```
   * How does the Rollout scale the new and old ReplicaSets to switch to the new version?
   * Repeat this step with a max unavailable to 0 and max surge to “20%”

## Add Steps
Note: all these steps assume that the rollout is in a fully progressed out state before you start making changes.

1. Wait until your Canary rollout is progressed fully into the desired ReplicaSet
2. Take a look at the following [Rollout](https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/canary/example/rollout.yaml)
   * What do you think will happen when you apply this rollout?
3. Run `kubectl argo rollouts get rollout canary-demo -w` in a seperate tab and apply the yaml by running:
   ```
   kubectl apply -f https://github.intuit.com/argoproj-internal/argo-rollouts-onboarding/blob/master/canary/example/rollout-with-steps.yaml
   ```
   
4. Observe the output in the `kubectl argo rollouts get rollout canary-demo -w` tab until the new replicaset becomes the stable RS
   * If nothing happens, try changing the image by running `kubectl argo rollouts set image canary-demo "*=argoproj/rollouts-demo:green"` with a different color.
5. Edit the rollout and change the first setWeight step from 20 to 23.
   * What do you think will happen if you change the image?
6. Change the image and write what happens.
   ```






   ```
7.  Afterward, you can run `kubectl argo rollouts promote --skip-all-steps canary-demo` so you don’t have to wait for all the durations to pass
## Manual Pause
For all these steps, have `kubectl argo rollouts get rollout canary-demo -w` running and in view. We will be referencing these as we walk through the steps.

1. Edit the rollout and remove the `duration:30` from the second step and fourth step.
   ```yaml
   ...
   canary:
     steps:
     - setWeight: 20
     - pause: {}
     - setWeight: 20
     - pause: {}
     ...
   ```
2. Change the image and watch the output of `kubectl argo rollouts get rollout canary-demo -w` until it doesn't change
3. Look at the output of `kubectl get rollouts canary-demo -o yaml`
   * How can you tell it's at the second step and the rollout is paused
     <details>
     <summary>Hint:</summary>
     Look at Status.CurrentStepIndex and Status.PauseConditions fields
     </details>
4. Manually edit the rollout to unpause the rollout.
   * How did you do it?
   <details>
   <summary>Answer:</summary>
   Remove the Status.PauseConditions fields
   </details>
   * Note: Increasing the `currentStepIndex` is not the correct way
5. When the rollout pauses, run `kubectl argo rollouts get rollout canary-demo -w`and then run `kubectl-argo-rollouts abort canary-demo`
   * What happens in the watch?
   * Run `kubectl get rollouts canary-demo -o yaml`. What is different about the yaml?
6. Run `kubectl argo rollouts retry rollout canary-demo`
   * What happens in the watch?

## Clean up
Run the following command to delete the namespaces:
```bash
kubectl delete namespace canary-demo
```

This will delete all the resources with the cluster.