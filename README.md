![Build Status](https://travis-ci.org/opensource2fa/Server.svg?branch=master)](https://travis-ci.org/opensource2fa/Server)
Copyright (c) 2019 Miles Breslin Robert Pelayo

# 2FA Server
 
This program provides web access to a database of One-time Passwords that are generated server side. The security of this project has not and will not be audited professionally, so use at your own     risk. The general idea is that this server contains the keys you have to be used in combination with the password you know and should not be trusted a single point of authentication.
  
### Build

To build, open a shell in the root of the repo and run the following:
 
```sh
 ./build.sh
 ```

It will generate a new folder called `build` and inside of that, there will be a server
executable file. Run that.

### Example Run
When the root of the repository, to start the server run:
```sh
 ./build/server
 ```

All client commands must begin with the following, if in the root of the repository:
```sh
 ./build/client --url ws://localhost:8000
 ```
 
#### Add key 
To add a key enter:
```sh
 ./build/client --url ws://localhost:8000 add-key lhe4kfhfqapxipzmohswb6i5adg2gauh
 ```
This code will add the key: `lhe4kfhfqapxipzmohswb6i5adg2gauh` to the server so it can now generate a one time token. The displayed ID will be where the token is stored, for access later.
#### Get key
To get a token:
```sh
 ./build/client --url ws://localhost:8000 get-key-token 1
 ```
This command will display token ID 1 to standard out. 
### License

This repository is licensed under the MIT License. See the `LICENSE` file for license conditions and details.