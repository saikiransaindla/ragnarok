# Ragnarok
Trying to prevent the inevitable Doooooom!!!

Install go and use the below command to run the application

`go run main.go`

Application would be available at `localhost:8080`

App setup
- Install go
- Clone this app inside $GOPATH/src (Not sure about this, will update once I confirm)
- go mod init
- download gin `go get -u github.com/gin-gonic/gin`

DB Migration
- brew install golang-migrate
- `make migrateup`

TODO:
- Add env specific configs