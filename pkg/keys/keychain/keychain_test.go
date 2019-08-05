package keychain

import (
    "testing"
)

func TestAddKey(t *testing.T) {
    key_1 := AddKey("lhe4kfhfqapxipzmohswb6i5adg2gauh")
    key_2, _ := GetKey(key_1.Id)
    if key_1.Secret != key_2.Secret {
        t.Errorf("Add key; Get Key has failed")
    }
    if _, err := GetKey(key_1.Id+1); err == nil {
        t.Errorf("Get key on a key that does not exist has failed")
    }

    list := GetList()
    if len(list) != 1 || list[0] != key_1.Id {
        t.Errorf("Unexpected error when getting list of keys\nExpecting [%d]\nGot %v", key_1.Id, list)
    }

    DelKey(key_1.Id)
    if _, err := GetKey(key_1.Id); err == nil || len(GetList()) != 0 {
        t.Errorf("Key delete has failed")
    }
}