# [k8s-grpc-gateway-example](https://github.com/vivekmarakana/k8s-grpc-gateway-example)

This is a sample gRPC Server with HTTP gateway using [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) library.

It implements following proto file:

```
syntax = "proto3";
option go_package = "proto;proto";

package proto;

import "google/api/annotations.proto";

message RequestMessage {
	string message = 1;
}

message ResponseMessage {
	string host = 1;
	string message = 2;
}

service EchoService {
	rpc Echo(RequestMessage) returns (ResponseMessage) {
		option (google.api.http) = {
			get: "/v1/echo/{message}"
			additional_bindings {
				get: "/v1/echo"
			}
		};
	}
}
``` 

## How to run?

Run the image on docker using following command:
```
docker run -p 8080:8080 -p 9090:9090 vivekmarakana/k8s-grpc-gateway:latest
```

### Make gRPC Calls
gRPC server is run on port 9090 with reflections enabled so you won't need to provide the proto files to client like [evans](https://github.com/ktr0731/evans). You can directly call the service by running:
```
bash-3.2$ evans repl --host localhost --port 9090 --reflection

  ______
 |  ____|
 | |__    __   __   __ _   _ __    ___
 |  __|   \ \ / /  / _. | | '_ \  / __|
 | |____   \ V /  | (_| | | | | | \__ \
 |______|   \_/    \__,_| |_| |_| |___/

 more expressive universal gRPC client


proto.EchoService@127.0.0.1:9090> call Echo
message (TYPE_STRING) => hello
{
  "host": "127.0.0.1",
  "message": "hello"
}

```

### Make REST Calls

HTTP gateway is running on port 8080. You can call it using browser or curl:
```
bash-3.2$ curl -s localhost:8080/v1/echo | jq
{
  "host": "127.0.0.1",
  "message": ":("
}
bash-3.2$ curl -s localhost:8080/v1/echo/hello | jq
{
  "host": "127.0.0.1",
  "message": "hello"
}
```

## Deployment on Kubernetes

The repo contains the kubernetes [kustomization](https://kubernetes.io/docs/tasks/manage-kubernetes-objects/kustomization) files in `k8s` directory. If you have kubernetes cluster setup already, you can deploy this using the following command:
```
kubectl apply -k ./k8s 
```

This will run the deployment with 1 replica and a NodePort service. To get the port to connect to use the command:
```
bash-3.2$ kubectl get service grpc-gateway-service
NAME                   TYPE       CLUSTER-IP       EXTERNAL-IP   PORT(S)                         AGE
grpc-gateway-service   NodePort   10.107.229.130   <none>        9090:32013/TCP,8080:31520/TCP   17s
```

Now you can make grpc calls on port `32013`(container port 9090) and http calls on `31520`(container port 8080).

## How to make changes & build locally?

- If you want to change the `proto/echo_service.proto`, you can run `make gp` after changing the file. 
- To run locally, use the command `go run src/main.go`
- To build a binary, run the command `make build`. This will create a binary `bin/server`
- To build docker image you can run `DOCKER_HUB_USERNAME=<YOUR_USERNAME> make dockerbuild`. This will make the local image and tag it with `<YOUR_USERNAME>/k8s-grpc-gateway:latest` which you can publish with `docker push`
