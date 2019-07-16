package realtime_api

type incommingMessage struct {
    Type        string          `json:"type"`
    Method      string          `json:"method,omitempty"`
    Obj         []interface{}   `json:"obj,omitempty"`
    Id          uint64          `json:"id"`
}

type outgoingMessage struct {
    Type        string          `json:"type"`
    Result      uint8           `json:"result,omitempty"`
    Obj         []interface{}   `json:"obj,omitempty"`
    Id          uint64          `json:"id"`
}