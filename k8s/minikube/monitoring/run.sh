#!/usr/bin/env bash
kubectl config set-context minikube

kubectl get componentstatus

kubectl cluster-info

# deploy an application and a service
# describe a namespace, deployment, and a service

kubectl create -f monitoring-namespace.yaml
kubectl create -f prometheus-config.yaml
kubectl create -f prometheus-deployment.yaml
kubectl create -f prometheus-service.yaml

# view via CLI:
kubectl get services --namespace=monitoring
kubectl get deployments --namespace=monitoring

# show the dashboard
minikube dashboard

# show node ports under services for prometheus. exposes

# show prometheus service

minikube service --namespace=monitoring prometheus

# click around
# show /targets

# show graph and query of container_memory_usage_bytes{kubernetes_namespace="monitoring"}

# deploy grafana
kubectl create -f grafana-deployment.yaml
kubectl create -f grafana-service.yaml

# show grafana
minikube service --namespace=monitoring grafana

# add datasource. make sure type is prometheus http://prometheus:8080
# describe kubernetes DNS

# create a graph:
#  container_memory_usage_bytes{kubernetes_namespace="monitoring"}
# {{kubernetes_pod_name}

# lets add node metrics
# deploy node exporter. explain daemonser
kubectl create -f node-exporter-daemonset.yml

# show new target in prometheus. explain it autodiscovering the pods

# now create a graph

# node_load1
# {{kubernetes_pod_node_name}}

