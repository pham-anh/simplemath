# simplemath

```bash
set -x
TAG=$(git rev-parse --short HEAD)
docker build --push -t asia-east1-docker.pkg.dev/inspired-micron-198514/pqa/simplemath:$TAG .
set +x
```