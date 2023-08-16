package main

import (
	"fmt"
	"time"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/database"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/stream"
)

func main() {}

//export start
func _start(rid uint32) int32 {
	log.Log(fmt.Sprintf("start received: %d", rid))
	message, err := stream.GetDataByRID(rid)
	if err != nil {
		log.Log("error: " + err.Error())
		return -1
	}
	log.Log("wasm received: " + string(message))
	key := time.Now().Unix()
	if err := database.Set(string(key), []byte(message)); err != nil {
		return -1
	}
	log.Log("set key success")
	res, err := database.Get(string(key))
	if err != nil {
		log.Log("get key failed")
		return -1
	}
	log.Log("get key success")
	log.Log(fmt.Sprintf("get data %s by key %s", string(res), string(key)))
	return 0
}
