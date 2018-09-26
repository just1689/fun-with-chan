# fun-with-chan
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fjust1689%2Ffun-with-chan.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fjust1689%2Ffun-with-chan?ref=badge_shield)


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

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fjust1689%2Ffun-with-chan.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fjust1689%2Ffun-with-chan?ref=badge_large)