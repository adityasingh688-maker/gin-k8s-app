# Gin app on Kubernetes (Docker + CI/CD + TLS)

## 1. Run locally
    go mod tidy        # generates go.sum (commit it!)
    go run .           # http://localhost:8080

## 2. Docker
    docker build -t gin-app .
    docker run -p 8080:8080 gin-app

## 3. Create TLS secret (self-signed, for local clusters)
    openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
      -keyout tls.key -out tls.crt -subj "/CN=gin-app.local" \
      -addext "subjectAltName=DNS:gin-app.local"
    kubectl create secret tls gin-app-tls --cert=tls.crt --key=tls.key

## 4. Deploy manually (minikube example)
    minikube addons enable ingress
    # edit image in k8s/deployment.yaml: ghcr.io/<owner>/<repo>:latest
    kubectl apply -f k8s/
    echo "$(minikube ip) gin-app.local" | sudo tee -a /etc/hosts
    curl -k https://gin-app.local/health

## 5. CI/CD secrets
Repo > Settings > Secrets > Actions > add `KUBE_CONFIG`:
    base64 -w0 ~/.kube/config
