# Kubernetes Infrastructure

Helm charts for deploying VAP Platform to Kubernetes.

## Structure

```
k8s/
├── charts/
│   ├── gateway/
│   ├── auth-service/
│   ├── billing-service/
│   └── ...
├── values/
│   ├── development.yaml
│   ├── staging.yaml
│   └── production.yaml
└── README.md
```

## Deployment

```bash
# Deploy to development
helm upgrade --install vap-platform ./charts -f values/development.yaml

# Deploy to production
helm upgrade --install vap-platform ./charts -f values/production.yaml
```
