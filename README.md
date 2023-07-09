![Docker Build](https://github.com/pehlicd/unnecessarywebserver/actions/workflows/main.yml/badge.svg?branch=main)

# Unnecessary Web Server

## How to run?

### Using Source Code 
```bash
git clone git@github.com:pehlicd/unnecessarywebserver.git
export API_PORT=8080
go run .
```

### Using Docker
```bash
docker run -p 8080:8080 ghcr.io/pehlicd/unnecessarywebserver
```

### Using Helm Chart
```bash
git clone git@github.com:pehlicd/unnecessarywebserver.git
cd helm-chart
helm install unnecessarywebserver . -f values.yaml
```