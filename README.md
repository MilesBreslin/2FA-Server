Copyright (c) 2019 Miles Breslin Robert Pelayo

# 2FA Server

This program provides web access to a database of One-time Passwords that are generated server side. The security of this project has not and will not be audited professionally, so use at your own risk. The general idea is that this server contains the keys you have to be used in combination with the password you know and should not be trusted a single point of authentication.

### Build and Run
Must be in the directory: 
```sh
/Server
```
Run commmand:
```sh
go get github.com/gorilla/websocket
```

After the package has been installed, run:
```sh
go run cmd/server/main.go
```

### License

This repository is licensed under the MIT License. See the `LICENSE` file for license conditions and details.

