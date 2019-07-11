# Websocket Protocol

### User Sent Messages
There are 3 kinds of acceptable sent messages over the websocket protocol. The method message, the lookup message, and the subscribe message.

Each message sent _must_ include an `id` field where `id` is a unique integer per websocket session to identify the message. This field will be sent back on the reply from the server to identify the resulting messages of the sent command.

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
    ],
    "id": n
}
```
`n` in this example is a unique integer per websocket session to identify the message when the response is generated.

`methodType` is a string to identify the method to be run by name.

#### Lookup

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

`n` in this example is a unique integer per websocket session to identify the message when the response is generated.

`keyName` is a partial keyname to match.

#### Subscribe

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

`n` in this example is a unique integer per websocket session to identify the message when the response is generated.

`keyName` is a partial keyname to match.