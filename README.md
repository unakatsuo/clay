# Summary

Clay is an abstract system model store to automate something.
It provides some APIs to access the system model store, and sample UI.

### Sample UI - network design
![Network design](./images/sample1.png)

### Sample UI - physial diagram from the system model store
![Physical diagram](./images/sample2.png)

### Sample UI - logical diagram from the system model store
![Logical diagram](./images/sample3.png)

# How to use

```
$ go build
$ AUTOMIGRATE=1 ./clay
```
server runs at http://localhost:8080

# Example model

```
$ curl -X PUT localhost:8080/v1/designs/present -H "Content-Type: application/json" -d @examples/design.json
```

# Example template

```
$ # register template and external parameters
$ curl -X POST "localhost:8080/v1/templates" -H "Content-Type: multipart/form-data" -F name=terraform -F template_content=@examples/terraform.template
$ curl -X POST "localhost:8080/v1/template_external_parameters" -H "Content-Type: application/json" -d '{"template_id": 1, "name": "dpid", "value": "dp-pica8"}'
$ # show generated template
$ curl -X GET "localhost:8080/v1/templates/1"
```

# Example requirement

```
$ # register protocols
$ curl -X POST "localhost:8080/v1/protocols" -H "Content-Type: application/json" -d '{"name": "icmp"}'
$ curl -X POST "localhost:8080/v1/protocols" -H "Content-Type: application/json" -d '{"name": "tcp"}'
$ curl -X POST "localhost:8080/v1/protocols" -H "Content-Type: application/json" -d '{"name": "udp"}'
$ # register services
$ curl -X POST "localhost:8080/v1/services" -H "Content-Type: application/json" -d '{"name": "ping"}'
$ curl -X POST "localhost:8080/v1/services" -H "Content-Type: application/json" -d '{"name": "ssh"}'
$ curl -X POST "localhost:8080/v1/services" -H "Content-Type: application/json" -d '{"name": "http"}'
$ curl -X POST "localhost:8080/v1/services" -H "Content-Type: application/json" -d '{"name": "https"}'
$ curl -X POST "localhost:8080/v1/services" -H "Content-Type: application/json" -d '{"name": "dns"}'
$ # register connections
$ curl -X POST "localhost:8080/v1/connections" -H "Content-Type: application/json" -d '{"service_id": 1, "protocol_id": "1", "port_number": -1}'
$ curl -X POST "localhost:8080/v1/connections" -H "Content-Type: application/json" -d '{"service_id": 2, "protocol_id": "2", "port_number": 22}'
$ curl -X POST "localhost:8080/v1/connections" -H "Content-Type: application/json" -d '{"service_id": 3 "protocol_id": "2", "port_number": 80}'
$ curl -X POST "localhost:8080/v1/connections" -H "Content-Type: application/json" -d '{"service_id": 4, "protocol_id": "2", "port_number": 443}'
$ curl -X POST "localhost:8080/v1/connections" -H "Content-Type: application/json" -d '{"service_id": 5, "protocol_id": "2", "port_number": 53}'
$ curl -X POST "localhost:8080/v1/connections" -H "Content-Type: application/json" -d '{"service_id": 5, "protocol_id": "3", "port_number": 53}'
$ # register communication requirements
$ curl -X POST "localhost:8080/v1/requirements" -H "Content-Type: application/json" -d '{"source_port_id": 1, "destination_port_id": 2, "service_id": 1, "accessibility": true}'
```

# Example testcase

```
$ # register test case
$ curl -X POST "localhost:8080/v1/test_cases"
$ # register test scripts
$ curl -X POST "localhost:8080/v1/test_commands" -H "Content-Type: multipart/form-data" -F "service_name=ping" -F "server_script_template=@examples/testscripts/ping_server.sh" -F "client_script_template=@examples/testscripts/ping_client.sh"
$ curl -X POST "localhost:8080/v1/test_commands" -H "Content-Type: multipart/form-data" -F "service_name=ssh" -F "server_script_template=@examples/testscripts/ssh_server.sh" -F "client_script_template=@examples/testscripts/ssh_client.sh"
$ curl -X POST "localhost:8080/v1/test_commands" -H "Content-Type: multipart/form-data" -F "service_name=http" -F "server_script_template=@examples/testscripts/http_server.sh" -F "client_script_template=@examples/testscripts/http_client.sh"
$ curl -X POST "localhost:8080/v1/test_commands" -H "Content-Type: multipart/form-data" -F "service_name=https" -F "server_script_template=@examples/testscripts/https_server.sh" -F "client_script_template=@examples/testscripts/https_client.sh"
$ # register test patterns
$ curl -X POST "localhost:8080/v1/test_patterns" -H "Content-Type: application/json" -d '{"test_case_id": 1, "test_command_id": 1}'
$ curl -X POST "localhost:8080/v1/test_patterns" -H "Content-Type: application/json" -d '{"test_case_id": 1, "test_command_id": 2}'
$ curl -X POST "localhost:8080/v1/test_patterns" -H "Content-Type: application/json" -d '{"test_case_id": 1, "test_command_id": 3}'
$ curl -X POST "localhost:8080/v1/test_patterns" -H "Content-Type: application/json" -d '{"test_case_id": 1, "test_command_id": 4}'
$ # check test case
$ curl -X GET "localhost:8080/v1/test_cases?preloads=TestPatterns,TestPatterns.TestCommand"
```

# API Server

Simple Rest API using gin(framework) & gorm(orm)

## Endpoint list

### Nodes Resource

```
GET    /<version>/nodes
GET    /<version>/nodes/:id
POST   /<version>/nodes
PUT    /<version>/nodes/:id
DELETE /<version>/nodes/:id
```

### NodeGroups Resource

```
GET    /<version>/nodegroups
GET    /<version>/nodegroups/:id
POST   /<version>/nodegroups
PUT    /<version>/nodegroups/:id
DELETE /<version>/nodegroups/:id
```

### NodePvs Resource

```
GET    /<version>/nodepvs
GET    /<version>/nodepvs/:id
POST   /<version>/nodepvs
PUT    /<version>/nodepvs/:id
DELETE /<version>/nodepvs/:id
```

### NodeTypes Resource

```
GET    /<version>/nodetypes
GET    /<version>/nodetypes/:id
POST   /<version>/nodetypes
PUT    /<version>/nodetypes/:id
DELETE /<version>/nodetypes/:id
```

### Ports Resource

```
GET    /<version>/ports
GET    /<version>/ports/:id
POST   /<version>/ports
PUT    /<version>/ports/:id
DELETE /<version>/ports/:id
```

# UI

Access http://localhost:8080/ui/

# Thanks

* Clay was partially generated by https://github.com/wantedly/apig
* Clay uses https://github.com/codeout/inet-henge
