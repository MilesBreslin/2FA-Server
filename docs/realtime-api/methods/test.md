# Test Method

The test method is a simple echo method for the realtime API v0. It will likely be removed in future versions, but while major developmental changes are taking place, it will stay. 

Client Message
```json
{
    "type":"method",
    "method": "test",
    "obj": [
        ... (any valid JSON; it will echo it back)
    ],
    "id": n
}
```

Server Message
```json
{
    "type":"result",
    "result": 200,
    "obj": [
        ... (any valid JSON; based entirely upon sent request)
    ],
    "id": n
}
```