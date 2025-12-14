package main

import (
    "encoding/hex"
    "encoding/json"
    "fmt"
)

type Slice []byte

// MarshalJSON реализует интерфейс json.Marshaler.
func (s Slice) MarshalJSON() ([]byte, error) {
    // используйте hex.EncodeToString для преобразования в hex-строку
    // и затем json.Marshal
    // ...
	res := hex.EncodeToString(s)
	return json.Marshal(res)
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler.
func (s *Slice) UnmarshalJSON(data []byte) error {
    var tmp string
    if err := json.Unmarshal(data, &tmp); err != nil {
        return err
    }
    v, err := hex.DecodeString(tmp)
    if err == nil {
        *s = v
    }
    return err
}

type MySlice struct {
    ID    int
    Slice Slice
}

func main() {
    ret, err := json.Marshal(MySlice{ID: 7, Slice: []byte{1, 2, 3, 10, 11, 255}})
    if err != nil {
        panic(err)
    }
    fmt.Println(string(ret))
    var result MySlice
    if err = json.Unmarshal(ret, &result); err != nil {
        panic(err)
    }
    fmt.Println(result)
}