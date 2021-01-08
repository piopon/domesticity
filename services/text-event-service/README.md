# **Text Event Service**

### **Description**
This service is responsible for handling text events in Domesticity Smart Home Organizer project.

### **Getting started**
#### _Prerequisites_
To build this project locally one have to install:
* Golang - min. version: 1.15.3
* MongoDB - min. version: 4.2.0

When using Docker containers then one should install:
* Docker - min. version: 19.03.13

#### _Running service_
1. clone repository
2. navigate to `services/text-event-service` directory
3. invoke:
   - MongoDB (has to be installed first):
     ```
     mongod --bind_ip 0.0.0.0
     ```    
   - build service application:
     ```
     make build
     ```
   - run service application
     ```
     ./bin/text-event-service
     ```
4. alternatively one can use Docker
   - use standalone Docker image (MongoDB has to be manually installed)
     ```
     make docker-build && make docker-run
     ```
   - or invoke Docker compose
     ```
     make docker-up
     ```

#### _Stopping service_
Service can be stopped by using Ctrl + C shortkey.

When using docker one can use below command to stop and clean all downloaded data
```
make docker-stop && make docker-clean
```
Alternatively when using docker compose one can stop service with command
```
make docker-down
```

### **Endpoints**
All endpoint documentation in more detailed manner is present after hitting `/docs` endpoint in running service.
* GET
  * `/docs` - load detailed documentation
  * `/events` - read all events
  * `/events/{id}` - read an event with specified id
* POST
  * `/events` - add a new event
* PUT
  * `/events/{id}` - update an event with specified id
* DELETE
  * `/events/{id}` - delete an event with specified id

### **Roadmap**
- [x] add swagger documentation
- [x] host swagger documentation as an service endpoint
- [x] add filtering to MongoDB dataservice
- [x] add configuration script support
- [x] add missing unit tests
- [x] update home page with basic service information
- [x] add docker containers and docker compose features
- [ ] think about adding event types to this service (instead of creating new one)
- [ ] configure CI process

---
Developed with :heartpulse: by PNK @ 2020 :vulcan_salute:
