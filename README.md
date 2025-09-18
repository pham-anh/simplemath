# simplemath

```bash
set -x
TAG=$(git rev-parse --short HEAD)
docker build --push -t asia-east1-docker.pkg.dev/inspired-micron-198514/pqa/simplemath:$TAG .
set +x
```

Rate Limit

```
simplemath git:(main) ✗ for i in {1..301}; do curl -s -o /dev/null -w "%{http_code}\n" http://localhost:8080; done
```