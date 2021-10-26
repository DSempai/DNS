# Atlas Corporation Drone Navigation Service (DNS)

This service helps drones find the closest databank in the sector depending on drones coordinates

## Getting started
### Requirements

- **Go 1.17+:** amazing 197 years old language. Can be found at https://golang.org
- **Docker:** Long ago, mankind used giant steel boxes to deploy their systems, rather than 
vacuum portable subspaces. And nevertheless, the containerization system is still quite good, can be found at https://www.docker.com

### Start

For starting this service you need to:
1) Clone this repository
2) Open directory with project
3) Execute a command :
`docker-compose up --build -d`

When command executed, you can pretend yourself as a drone and do a request:


>curl --request POST \
--url http://localhost:8090/api/v1/locate_databank \
--header 'Content-Type: application/json' \
--data '{
"x": "12.234",
"y": "44.342",
"z": "123.13",
"vel": "55.3424"
}

We know, it looks ugly, but our drones doesn't dream about electric sheeps.

Request must contain 4 parameters and must be provided as string representation of floating point number:
- **X axis coordinate** 
- **Y axis coordinate** 
- **Z axis coordinate** 
- **Velocity**

### Testing

If you want to run testing functions, you need to run command:\
``docker exec -it dns sh``

After that you will be teleported inside container where application is stored.
And all you need to do is run another command:\
``make test``