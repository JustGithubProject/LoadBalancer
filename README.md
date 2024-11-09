# Load Balancer

### Launching three servers

```
    cd server
```

```
    go run start_server.go target_server.go 8001 
```

```
    go run start_server.go target_server.go 8002
```

```
    go run start_server.go target_server.go 8003
```


### Sending requests using load balancer (6 requests)

```
    cd client
```

```
    go run client.go load_balancer.go
```

