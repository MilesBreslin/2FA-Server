package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "../../pkg/realtime_api/client"
    // "../../pkg/totp"
    /*"github.com/opensource2fa/server/pkg/realtime_api/server"
    "net/http/httptest"
    "net/http"*/
    // "github.com/opensource2fa/server/pkg/keys/keychain"
)
// TODO: Close out the server connection
// TODO: Cycle through all the gotten keys?
// TODO: How do we want to deal with getting the name for the appropriate keys maybe add name field
// TODO: return count of the number of keys?
func main() {
    // pull the url
    url := os.Args[2] + "/api/v0/realtime"

    // get the command
    arg := os.Args[3]

    // add a key to the server
    if strings.Compare("add-key", arg) == 0 {
        // get the key
        key := os.Args[4]
        // example key "lhe4kfhfqapxipzmohswb6i5adg2gauh"
        c, err := client.NewClient(url)
        if err != nil {
            fmt.Println(err)
            return
        }
        id, err := c.AddKey(key)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("This key has been stored with the ID:", id)
        // get a key
    } else if strings.Compare("get-key", arg) == 0 {
        id_string := os.Args[4]
        c, err := client.NewClient(url)
        if err != nil {
            fmt.Println(err)
            return
        }
        // string to convert to uint64
        id, err := strconv.ParseUint(id_string, 10, 64)
        if err != nil {
            fmt.Println(err)
        }
        key, err := c.GetKey(id)
        if err != nil {
            fmt.Println(err)
            return 
        }
        fmt.Println("ID:", key.Id, "Key:", key.Secret)
        // get the token
    } else if strings.Compare("get-token", arg) == 0 {
        id_string := os.Args[4]
        c, err := client.NewClient(url)
        if err != nil {
            fmt.Println(err)
            return
        }
        // string to convert to uint64
        id, err := strconv.ParseUint(id_string, 10, 64)
        if err != nil {
            fmt.Println(err)
        }
        // get the token
        key, err := c.GetKeyToken(id)
        fmt.Println("Token:", key)

        // list all IDs
    } else if strings.Compare("list-keys", arg) == 0 || 
              strings.Compare("list-tokens", arg) == 0{
        c, err := client.NewClient(url)
        if err != nil {
            fmt.Println(err)
            return
        }
        key, err := c.ListKeys()
        if err != nil {
            fmt.Println(err)
            return 
        }
        fmt.Println("Stored IDs:", key)
        // error, user input command does not exist
    } else {
        fmt.Println("2FA Unkown command:", arg)
    }
}
