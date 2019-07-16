# REST API

This documentation is for the v0 REST api. During v0, changes to implementation are expected and all messages are subject to change.

This api is for general usage of client apps. Unless a client application needs to maintain state information, this api should be used.

This api is intended to be built using a [RESTful style](https://restfulapi.net/). Any straying from this style is to be considered a bug unless otherwise noted. This is also intended to be written in a manner similar to other REST api's.

### Accessing the REST API

The rest api is accessible by using a HTTP client to connect to the following address and its children. No authentication or other initiation is required at this time, but is intended for the future.
```bash
http://$HOSTNAME:$PORT/api/v0/
```