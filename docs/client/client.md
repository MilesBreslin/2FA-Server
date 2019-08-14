# 2FA Client

This documentation is for the v0 2FA client. During v0, changes to implementation are expected and all messages are subject to change.

The client is the main way to communicate with the server. This allows for the user to add, view, and manage the storage of keys (to be implemented, hopefully!). This is mainly used to get command line arguments and use them to call the corresponding method and return the information or store the information with the user/server, respectively. 

### Accessing the Client

To interact with the client, all command line arguments must be as follows:
```bash
./build/client --url ws://localhost:8000
```
To interact with the client, the user must include `--url ws://localhost:8000 command` Anything less than this many arguments, will cause the program to throw an error, and let the user know what must be entered.

### List-keys

To list all keys stored on the server:
```bash
./build/client --url ws://localhost:8000 list-keys
```
The list-keys command will call the `ListKeys` method which uses the `runMethod` method to query the server and run its specific method to return the data stored inside an interface. This method returns the array of keys to the ListKeys method and returns again to the list-keys command in the client.

#### Add-key

To add a key to the database:
```bash
./build/client --url ws://localhost:8000 add-key AKSUEJNHDUO0918Y
```
In this case, `AKSUEJNHDUO0918Y` is the secret key that will be hashed. The add-key command will run similarly to the list-keys command, but will call the `AddKey` method passing in the string of secret characters to the server (through the runMethod) and return the unique ID in which the secret key has been stored. 

#### Get-key

To retrieve a key from the server:
```bash
./build/client --url ws://localhost:8000 get-key 2
```
With this command, the client will call the `GetKey` method with the argument `2`. This will return the key that is associated with the UID `2`. This command was meant for admin, and would not normally be available to regular users, but has not yet been implemented.

### Get-token
To retrieve a token:
```bash
./build/client --url ws://localhost:8000 get-token 1
```
With this command, the client will call the `GetKeyToken` method and pass in the UID 1 to return the OTP token. This number can be changed to return the desired token within the index. If the entered key was not long enough, the client will return a message indicating failure and recommendation to remove the invalid key (removal to be implemented).