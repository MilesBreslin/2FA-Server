package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "../../pkg/realtime_api/client"
)
func main() {
    // check correct number of args
    // lengths are all -1
    // arg 1 = --url
    // arg 2 = ws://localhost:8000
    // arg 3 = command
    // arg 4 = id
    if os.Args == nil || len(os.Args) < 3 {
        fmt.Println("2FA: Error, please follow format: client --url ws://localhost:8000 command")
        return
    } else if len(os.Args) == 3{
        fmt.Println("2FA: Please enter a command:")
        fmt.Println("add-key")
        fmt.Println("get-key")
        fmt.Println("get-token")
        fmt.Println("list-keys")
        return
    }

    option:= os.Args[1]
    if strings.Compare("--url", option) != 0{
        fmt.Println("Error argument:", option)
        fmt.Println("Be sure to include '--url' as an option followed by the URL of the websocket")
        return
    }
    
    // we will add in the version 
    url := os.Args[2] + "/api/v0/realtime"

    // get the command
    arg := os.Args[3]

    // add a key to the server
    if strings.Compare("add-key", arg) == 0 {
        // check the user added the key
        if len(os.Args) == 4{
            fmt.Println("2FA: Please enter the key after add-key.")
            return
        }
        // get the key
        key := os.Args[4]
        // example key "lhe4kfhfqapxipzmohswb6i5adg2gauh"
        c, err := client.NewClient(url)
        if err != nil {
            fmt.Println(err)
            fmt.Println("Please check you entered the correct url.")
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
        // check the user entered an ID 
        if len(os.Args) == 4{
            fmt.Println("2FA: Please enter the ID after get-key.")
            return
        }
        id_string := os.Args[4]
        c, err := client.NewClient(url)
        if err != nil {
            fmt.Println(err)
            fmt.Println("Please check you entered the correct url.")
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
        if len(os.Args) == 4{
            fmt.Println("2FA: Please enter the ID after get-token.")
            return
        } 
        id_string := os.Args[4]
        c, err := client.NewClient(url)
        if err != nil {
            fmt.Println(err)
            fmt.Println("Please check you entered the correct url.")
            return
        }
        // string to convert to uint64
        id, err := strconv.ParseUint(id_string, 10, 64)
        if err != nil {
            fmt.Println(err)
        }
        // get the token
        key, err := c.GetKeyToken(id)
        if (key == "first fail"){
            fmt.Println("The key was rejected, please remove this key. ID:", id)
            return
        }
        fmt.Println("Token:", key)

        // list all IDs
    } else if strings.Compare("list-keys", arg) == 0 || 
    strings.Compare("list-tokens", arg) == 0{
        c, err := client.NewClient(url)
        if err != nil {
            fmt.Println(err)
            fmt.Println("Please check you entered the correct url.")
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
