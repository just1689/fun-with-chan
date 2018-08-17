# fun-with-chan

The goal of this project is to have fun with ring.Ring, chan, grpc and go routines


## Completed

### Basic operations
- Subscribe
- Put
- Push
- Mark done
- Multiple consumers
- Timeouts

## To do

### GRPC
- Messages for newTopic(), put(), subscribe(), push(), disconnect
- GRPC server
- Client for put()
- Client for subscribe()

### Logging
- Each put
- Each push
- Number of consumers

### Prettiness
- Web server
- Health check rest endpoint
- Fancy healthy rest endpoint (puts, pushes, count at time, up-time, processing time, ingestion time)
- HTML page for querying the endpoints

### Maturity
- Config yaml / json read on startup
- Config ports
- Config cache sizes