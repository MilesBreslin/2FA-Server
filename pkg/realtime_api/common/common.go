package common

type IncommingMessage struct {
    Type        string          `json:"type"`
    Method      string          `json:"method,omitempty"`
    Obj         []interface{}   `json:"obj,omitempty"`
    Id          uint64          `json:"id"`
}

type OutgoingMessage struct {
    Type        string          `json:"type"`
    Result      uint16          `json:"result,omitempty"`
    Obj         []interface{}   `json:"obj,omitempty"`
    Id          uint64          `json:"id"`
}
