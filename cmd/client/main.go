package main

import (
    "fmt"
    "os"
    "strings"
    //"strconv"
    "github.com/opensource2fa/server/pkg/realtime_api/client"
    "github.com/opensource2fa/server/pkg/totp"
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

    fmt.Println("Running:", arg, "on URL:", url)
    // add a key to the server
    if strings.Compare("add-key", arg) == 0 {
        // get the key
        key := os.Args[4]
        //key := "lhe4kfhfqapxipzmohswb6i5adg2gauh"
        fmt.Println(key)
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
        fmt.Println("This key has been stored with the ID: %i", id)
    } else if strings.Compare("get-key", arg) == 0 {
        id := os.Args[4]
        c, err := client.NewClient(url)
        if err != nil {
            fmt.Println(err)
            return
        }
        key, err := c.GetKey(id)
        if err != nil {
            fmt.Println(err)
            return 
        }
        token := totp.GetTOTPToken(key.Secret)
        fmt.Println("2FA Token:", token)
        } else if strings.Compare("list-keys", arg) == 0 {
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
        fmt.Println(key)
        }
    }
