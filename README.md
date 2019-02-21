## Installation via Docker
Build the docker image

```docker build -t fast-text-server .```

Running the docker image

```docker run --name fast-text-server -v /models:/models -p 8080:8080 -e MODEL=/models/model_name.bin -d fast-text-server```

## How to Use
```curl -d '{"data":"Test Data"}' -H "Content-Type: application/json" -X POST http://localhost:8080```
