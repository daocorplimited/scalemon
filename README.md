# Scale monitoring Kubernetes operator

## Repository structure

1. scalemon: golang application with monitoring business logic (scale monitor, scalemon).
1. scalemon-ops: kubernetes operator which manages monitoring deployments (scalemons)

## Installing single instance with Helm

From the machine, connected to a Kubernetes cluster could be installed with Helm 3 by running the following command:

                $ helm upgrade --install scalemon ./scalemon/.helm/scalemon

The command will install the scalemon app instance together with standalone mongodb. To delete the application instance from a Kubernetes cluster run the following:

                $ helm delete scalemon

## Deploy Kubernetes operator

Kubernetes operator is implemented with Operator Framework (https://sdk.operatorframework.io/docs/building-operators/helm/quickstart/) to run it the following actions should be done:

1. Install the operator to the destination cluster.
1. Install 1+ instances of monitoring app (scalemon)

It could be achieved with the following commands:

        $ cd scalemon-ops
        $ make deploy IMG="aolishchuk/scalemon-operator:v04" //path could be yours
        $ kubectl apply -f config/resources/monitoring_v1alpha1_scalemon.yaml