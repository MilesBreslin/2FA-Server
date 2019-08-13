
[![Build Status](https://travis-ci.org/opensource2fa/Server.svg?branch=master)](https://travis-ci.org/opensource2fa/Server)
Copyright (c) 2019 Miles Breslin Robert Pelayo

# 2FA Server
 
This program provides web access to a database of One-time Passwords that are generated server side. The security of this project has not and will not be audited professionally, so use at your own risk. The general idea is that this server contains the keys you have to be used in combination with the password you know and should not be trusted a single point of authentication.
  
### Build

To build, open a shell in the root of the repository and run the following:
 
```sh
 ./build.sh
```

It will generate a new folder called `build` containing server and client executable files.

### Example Usage

#### Start the server

When the root of the repository, to run the server enter:
```sh
./build/server
```

#### Run the client

Leave the terminal open to keep the connection alive. Open a new terminal window to use the client. **All** client commands must begin with the following, if in the root of the repository:
```sh
./build/client --url ws://localhost:8000
```

You can add this as an alias to your bashrc file to configure this for your shell.

#### Add key 

To add a key enter:
```sh
./build/client --url ws://localhost:8000 add-key lhe4kfhfqapxipzmohswb6i5adg2gauh
```
 
This code will add the key: `lhe4kfhfqapxipzmohswb6i5adg2gauh` to the server so it can now generate a one time token. The displayed ID will be where the token is stored, for access later.

#### Get token

To get the first token:
```sh
./build/client --url ws://localhost:8000 get-token 1
```

#### Get key

To get the first key from the server:
```sh
./build/client --url ws://localhost:8000 get-key 1
```
 
#### List keys

To list all IDs stored:
```sh
./build/client --url ws://localhost:8000 list-keys
```
 
### License

This repository is licensed under the MIT License. See the `LICENSE` file for license conditions and details.
