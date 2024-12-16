# otel-samples


A client-server example for otel implementation

`client` side: hello service
`server` side: world service
`otel-collector`: to receive metrics/traces from client and server, and export to jaeger

## how to run 

the `docker-compose.yml` in **otel-collector** dir only runs two services: `otel-collector` and `jaeger`

client and server services are executed by go cmd

open a new terminal and run

```bash
$ cd hello # client service
$ go run main.go
```

open another new terminal and run

```bash
$ cd world # server service
$ go run main.go
```

you can also create a new docker-compose file and put all these services inside