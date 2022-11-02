# Monitoring application for an application autoscaler

## General structure

The monitoring system consists of two major components:

1. Monitoring application itself, which contains all the business logic.
1. Kubernetes operator, which contains all the service logic (deployment, scaling, health management etc)

Kubernetes operator is implemented with Kubernetes Operator Framework https://sdk.operatorframework.io/docs/building-operators/helm/quickstart/ with Helm plugin.



