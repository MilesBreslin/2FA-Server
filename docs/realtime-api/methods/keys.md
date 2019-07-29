# Key Methods

### Key Object
A fundamental part of any key method is the key object. It is declared in `pkg/keys/keys.go` internally, but can currently be summarized as the following:

```json
{
    "secret": SomeSecret,
    "id": KeyId
}
```

`SomeSecret` is a string which represents a supported key format.

`KeyId` is an integer which represents a reference to the key.

### GetKey Method
Retrieves a list of keys by keyids.

Client Message
```json
{
    "type":"method",
    "method": "GetKey",
    "obj": [
        ... KeyIds
    ],
    "id": n
}
```

`KeyIds` are integers which represents a reference to the key.

Server Message
```json
{
    "type":"result",
    "result": 200,
    "obj": [
        ... KeyObjects
    ],
    "id": n
}
```

`KeyObjects` are defined at the top of this document.

### GetKeyToken Method
Retrieves a list of keys by keyids.

Client Message
```json
{
    "type":"method",
    "method": "GetKeyToken",
    "obj": [
        ... KeyIds
    ],
    "id": n
}
```

`KeyIds` are integers which represents a reference to the key.

Server Message
```json
{
    "type":"result",
    "result": 200,
    "obj": [
        ... Tokens
    ],
    "id": n
}
```

`Tokens` are strings which represent a time based token derived from the given key.

### ListKeys Method
Lists all valid key ids.

Client Message
```json
{
    "type":"method",
    "method": "ListKeys",
    "id": n
}
```

Server Message
```json
{
    "type":"result",
    "result": 200,
    "obj": [
        ... KeyIds
    ],
    "id": n
}
```

`KeyIds` are integers which represents a reference to the key.

### AddKey Method
Adds a list of keys to the KeyChain. Note that this method may partially complete if an error occurred while adding one of the keys. If it encounters an error, it will abort, returning all keys that were added successfully and the result code of the failed addition.

Client Message
```json
{
    "type":"method",
    "method": "AddKey",
    "obj": [
        {
            "secret": SomeSecret
        }
        ... (Repeat for as many key additions as necessary)
    ],
    "id": n
}
```

`SomeSecret` is a string which represents a supported key format.

Server Message
```json
{
    "type":"result",
    "result": 200,
    "obj": [
        ... KeyIds
    ],
    "id": n
}
```

`KeyIds` are integers which represents a reference to the key.

### GetTOTPToken Method
Given an input of a TOTP Secret, it will return a token.

Client Message
```json
{
    "type":"method",
    "method": "GetTOTPToken",
    "obj": [
        ... SomeSecrets
    ],
    "id": n
}
```

`SomeSecrets` are strings which represents a supported key format.

Server Message
```json
{
    "type":"result",
    "result": 200,
    "obj": [
        ... Tokens
    ],
    "id": n
}
```

`Tokens` are strings which represent a time based token derived from the given key.