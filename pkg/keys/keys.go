package keys

import (
    "../totp"
    "errors"
)

type key struct {
    Secret          string              `json:"secret"`
    Id              uint64              `json:"id"`
}

// See GetKey
type keyRequest struct {
    Result          chan key
    Id              uint64
}

// Declare Global Channels + KeyChain
var add chan key
var del chan uint64
var get chan keyRequest
var getId chan uint64
var getList chan chan []uint64
var keyChain map[uint64] *key

func init() {
    // Initialize Global Channels + KeyChain
    add = make(chan key, 0)
    del = make(chan uint64, 0)
    get = make(chan keyRequest, 0)
    getId = make(chan uint64, 4)
    getList = make(chan chan []uint64, 0)
    keyChain = make(map[uint64] *key)

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

func (k *key) GetCode() (string, error) {
    return totp.GetTOTPToken(k.Secret) , nil
}

func GetKey(id uint64) (*key, error) {
    result := make(chan key, 0)
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

func AddKey(secret string) (*key) {
    var key = key{
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