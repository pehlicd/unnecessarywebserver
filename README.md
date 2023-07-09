![Docker Build](https://github.com/pehlicd/unnecessarywebserver/actions/workflows/main.yml/badge.svg?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/pehlicd/unnecessarywebserver)](https://goreportcard.com/report/github.com/pehlicd/unnecessarywebserver)

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
docker run -p 8080:8080 ghcr.io/pehlicd/unnecessarywebserver:latest -e API_PORT=8080
```

### Using Helm Chart
```bash
git clone git@github.com:pehlicd/unnecessarywebserver.git
cd helm-chart
helm install unnecessarywebserver . -f values.yaml
```
