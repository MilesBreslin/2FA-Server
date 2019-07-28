package keys

import (
    "../totp"
    "errors"
)

type key struct {
    Secret          string              `json:"Secret"`
    Id              uint64              `json:"Id"`
}

type keyRequest struct {
    Result          chan *key
    Id              uint64
}

var add chan *key
var del chan uint64
var get chan keyRequest
var getId chan uint64
var keyChain map[uint64] *key

func init() {
    add = make(chan *key, 0)
    del = make(chan uint64, 0)
    get = make(chan keyRequest)
    getId = make(chan uint64)
    keyChain = make(map[uint64] *key)
    go handleKeys()
}

func handleKeys() {
    new_id := uint64(1)
    for {
        select {
        case id := <- del:
            delete(keyChain, id)
        case k := <- add:
            keyChain[k.Id] = k
        case transport := <- get:
            if val, ok := keyChain[transport.Id]; ok {
                transport.Result <- val
            } else {
                close(transport.Result)
            }
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
    result := make(chan *key, 0)
    get <- keyRequest {
        Result: result,
        Id: id,
    }
    if k, alive := <- result; alive {
        return k, nil
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
    add <- &key
    return &key
}