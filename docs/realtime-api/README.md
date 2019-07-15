# Realtime API

This documentation is for the v0 realtime api. During v0, changes to implementation are expected and all messages are subject to change.

The realtime uses the HTTP WebSocket Protocol, defined by [RFC 6455](https://tools.ietf.org/html/rfc6455), to send and recieve messages. Any client libraries that implement support for the defined protocol should be able to interface with this library. The server itself uses the Go implementation [Gorilla Websocket](https://github.com/gorilla/websocket).

### Accessing the Realtime API

The realtime api is accessible by using a websocket client to connect to the following address. No authentication or other initiation is required at this time, but is intended for the future.
```bash
ws://$HOSTNAME:$PORT/api/v0/realtime
```

### User Sent Messages
There are 3 kinds of acceptable sent messages over the realtime api. The method message, the lookup message, and the subscribe message.

Each message sent _must_ include an `id` field where `id` is a unique integer per connection to identify the message. This field will be sent back on the reply from the server to identify the resulting messages of the sent command.

#### Method
This message will essentially run a function on the server. That function could be to upload a key. It could be to delete a key. It could be to run some entirely special, server-implemented function and return a response. This exists to be a framework message to send for all kinds of methods.

```json
{
    "type":"method",
    "method": methodType,
    "obj": [
        {
            ... (any method parameters)
        }
        ... (more times to repeat function if applicable)
    ],
    "id": n
}
```
`n` in this example is a unique integer per connection to identify the message when the response is generated.

`methodType` is a string to identify the method to be run by name.

#### Lookup

This message will look up a key by any attribute known to it. It should return the first match it has.

```json
{
    "type":"lookup",
    "obj": [
        {
            "keyName": keyName,
            ... (other key attributes)
        },
        ... (more keys to lookup)
    ],
    "id": n
}
```

`n` in this example is a unique integer per connection to identify the message when the response is generated.

`keyName` is a partial keyname to match.

#### Subscribe

This message subcribes you to updates of a keytype. Everytime a new six-digit code gets generated, this will ensure that updates are delivered.

```json
{
    "type":"subscribe",
    "obj": [
        {
            "keyName": keyName,
            ... (other key attributes)
        },
        ... (more keys to lookup)
    ],
    "id": n
}
```

`n` in this example is a unique integer per connection to identify the message when the response is generated.

`keyName` is a partial keyname to match.

### Server Sent Messages

The server has 2 types of messages it can reply: result and update. Result is used for any synchronous results that need to be returned. Update is used for asynchronous changes.

#### Result

This message will be sent only after a client has sent a message. It must be sent in reply to all messages. It must contain a result code that matches with common HTTP status codes. It must contain an id field to match with the command that triggered it. It may contain an object field which is an array of object types, but may omit the field if unnecessary for the command that triggered it.

```json
{
    "type":"result",
    "result": httpStatusCode,
    "obj": [
        {
            ... (a response specific datatype)
        },
        ... (more responses)
    ],
    "id": n
}
```

`n` in this example is a unique integer per connection to identify which message this is a response to.

`httpStatusCode` is a integer of a status code defined [here](https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html).

#### Update

This message will be sent only after a client has previously recieved a `202` response code from a `subscribe` client-sent message. In the `obj` field, each data object should be interpreted without any context for any of the other data objects in the message. This message may come at any time.

```json
{
    "type":"update",
    "obj": [
        {
            ... (a response specific datatype)
        },
        ... (more responses)
    ],
    "id": n
}
```

`n` in this example is a unique integer per connection to identify which message this is a response to.
