package keychain

import (
    "github.com/opensource2fa/pkg/keys"
    "errors"
)

// See GetKey
type keyRequest struct {
    Result          chan keys.Key
    Id              uint64
}

// Declare Global Channels + KeyChain
var add chan keys.Key
var del chan uint64
var get chan keyRequest
var getId chan uint64
var getList chan chan []uint64
var keyChain map[uint64] *keys.Key

func init() {
    // Initialize Global Channels + KeyChain
    add = make(chan keys.Key, 0)
    del = make(chan uint64, 0)
    get = make(chan keyRequest, 0)
    getId = make(chan uint64, 4)
    getList = make(chan chan []uint64, 0)
    keyChain = make(map[uint64] *keys.Key)

    // Start Handling the Channel Requests
    go handleKeys()
}

// All Read/Writes to keyChain MUST go through this single thread
// Or potential difficult to reproduce runtime panics are created
func handleKeys() {
    // Assume first key is ID 1
    new_id := uint64(1)

    // Loop forever
    for {
        select {    // Block until any one of these threads has a message
        // See DelKey
        case id := <- del:
            delete(keyChain, id)

        // See AddKey
        case k := <- add:
            keyChain[k.Id] = &k

        // See GetKey
        case transport := <- get:
            if val, ok := keyChain[transport.Id]; ok {
                transport.Result <- *val
            } else {
                close(transport.Result)
            }

        // See GetList
        case transport := <- getList:
            // Declare a list with static size
            list := make([]uint64, len(keyChain))
            // Iterate through all 
            i:=0
            for id := range keyChain {
                list[i] = id
                i++
            }
            transport <- list

        // See getUID
        // Generate a new, unique ID
        // Maybe it should be more difficult to predict
        case getId <- new_id:
            new_id = new_id+1
        }
    }
}

func getUID() (uint64) {
    return <- getId
}

func GetKey(id uint64) (*keys.Key, error) {
    result := make(chan keys.Key, 0)
    get <- keyRequest {
        Result: result,
        Id: id,
    }
    if k, alive := <- result; alive {
        close(result)
        return &k, nil
    } else {
        return nil, errors.New("Key not found")
    }
}

func DelKey(id uint64) {
    del <- id
}

func AddKey(secret string) (*keys.Key) {
    var key = keys.Key{
        Secret: secret,
        Id: getUID(),
    }
    add <- key
    return &key
}

func GetList() ([]uint64) {
    result := make(chan []uint64, 0)
    defer close(result)
    getList <- result
    return <- result
}