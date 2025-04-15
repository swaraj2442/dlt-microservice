DLT-microservice Monitoring, Observability & API Gateway

This document includes setup instructions for integrating:
- Prometheus (monitoring & alerting)
- Jaeger (distributed tracing)
- KrakenD API Gateway with custom authentication

--------------------------------------------------------------------------------

Prometheus Setup (Kubernetes)

Prometheus is used to collect and query metrics from your services and infrastructure.

1. Add Prometheus Helm Repo:

  helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
  helm repo update

2. Install Prometheus Stack:

  kubectl create namespace monitoring
  helm install monitoring prometheus-community/kube-prometheus-stack -n monitoring

This deploys Prometheus, Alertmanager, PushGateway, and supporting components.

--------------------------------------------------------------------------------

Prometheus Access

In-cluster DNS:
  http://prometheus-server.monitoring.svc.cluster.local:80

Local Port Forwarding:

  export POD_NAME=$(kubectl get pods --namespace monitoring \
    -l "app.kubernetes.io/name=prometheus,app.kubernetes.io/instance=prometheus" \
    -o jsonpath="{.items[0].metadata.name}")

  kubectl --namespace monitoring port-forward $POD_NAME 9090

Access Prometheus locally at: http://localhost:9090

--------------------------------------------------------------------------------

Alertmanager Access

In-cluster DNS:
  http://prometheus-alertmanager.monitoring.svc.cluster.local:9093

Local Port Forwarding:

  export POD_NAME=$(kubectl get pods --namespace monitoring \
    -l "app.kubernetes.io/name=alertmanager,app.kubernetes.io/instance=prometheus" \
    -o jsonpath="{.items[0].metadata.name}")

  kubectl --namespace monitoring port-forward $POD_NAME 9093

Access Alertmanager locally at: http://localhost:9093

--------------------------------------------------------------------------------

PushGateway Access

In-cluster DNS:
  http://prometheus-prometheus-pushgateway.monitoring.svc.cluster.local:9091

Local Port Forwarding:

  export POD_NAME=$(kubectl get pods --namespace monitoring \
    -l "app=prometheus-pushgateway,component=pushgateway" \
    -o jsonpath="{.items[0].metadata.name}")

  kubectl --namespace monitoring port-forward $POD_NAME 9091

Access PushGateway locally at: http://localhost:9091

--------------------------------------------------------------------------------

Jaeger Setup (Docker)

Jaeger provides distributed tracing for understanding request flows and performance bottlenecks.

Run Jaeger All-in-One via Docker:

  docker run -d --name jaeger \
    -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
    -p 5775:5775/udp \
    -p 6831:6831/udp \
    -p 6832:6832/udp \
    -p 5778:5778 \
    -p 16686:16686 \
    -p 14268:14268 \
    -p 14250:14250 \
    -p 9411:9411 \
    jaegertracing/all-in-one:1.46

Jaeger UI available at: http://localhost:16686

--------------------------------------------------------------------------------

KrakenD API Gateway Setup

KrakenD acts as a stateless, high-performance API Gateway to manage and aggregate microservice APIs.

1. Create krakend.json Configuration File (example):

  {
    "version": 3,
    "name": "dlt-gateway",
    "port": 8080,
    "endpoints": [
      {
        "endpoint": "/secure-data",
        "method": "GET",
        "extra_config": {
          "github_com/devopsfaith/krakend-jose/validator": {
            "alg": "HS256",
            "jwk-url": "http://your-auth-server/.well-known/jwks.json",
            "roles_key": "roles",
            "roles": ["admin"]
          }
        },
        "backend": [
          {
            "url_pattern": "/api/data",
            "host": ["http://your-internal-service"]
          }
        ]
      }
    ]
  }

2. Run KrakenD (Docker):

  docker run -d --name krakend \
    -p 8080:8080 \
    -v $(pwd)/krakend.json:/etc/krakend/krakend.json \
    devopsfaith/krakend run -c /etc/krakend/krakend.json

Access the API Gateway at: http://localhost:8080

--------------------------------------------------------------------------------

Summary

Tool          | Purpose                        | Access URL
--------------|--------------------------------|----------------------------
Prometheus    | Metrics & Monitoring           | http://localhost:9090
Alertmanager  | Alerts & Notification Routing  | http://localhost:9093
PushGateway   | Push-based Metrics             | http://localhost:9091
Jaeger        | Distributed Tracing            | http://localhost:16686
KrakenD       | API Gateway with Auth          | http://localhost:8080

--------------------------------------------------------------------------------
