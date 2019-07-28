package keys

import (
    "testing"
    "fmt"
)

func TestAddKey(t *testing.T) {
    fmt.Println("add")
    key_1 := AddKey("lhe4kfhfqapxipzmohswb6i5adg2gauh")
    fmt.Printf("Key %d\n\t%s",key_1.Id, key_1.Secret)
    key_2, _ := GetKey(key_1.Id)
    if key_1.Secret != key_2.Secret {
        t.Errorf("Add key; Get Key has failed")
    }
    if _, err := GetKey(key_1.Id+1); err == nil {
        t.Errorf("Get key on a key that does not exist has failed")
    }
    DelKey(key_1.Id)
    if _, err := GetKey(key_1.Id); err == nil {
        t.Errorf("Key delete has failed")
    }
}