package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/blockchain"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/database"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/stream"
)

func main() {}

//export start
func _start(rid uint32) int32 {
	log.Log(fmt.Sprintf("start received: %d", rid))
	res, err := stream.GetDataByRID(rid)
	if err != nil {
		log.Log("error: " + err.Error())
		return -1
	}
	log.Log(fmt.Sprintf("receive message: %s", res))

	s, e := blockchain.SendTx(
		4689,
		"0x1ED83F5AD999262eC06Ed8f3B801e108024b3e9c",
		big.NewInt(1),
		fmt.Sprintf("40c10f19000000000000000000000000%s0000000000000000000000000000000000000000000000000de0b6b3a7640000",
			"97186a21fa8e7955c0f154f960d588c3aca44f14"))
	log.Log("tx:" + s)
	log.Log("err:" + e.Error())

	return 0
}

func getDB(key string) int32 {
	val, err := database.Get(key)
	if err != nil {
		return 0
	}
	var ret int32
	buf := bytes.NewBuffer(val)
	binary.Read(buf, binary.LittleEndian, &ret)
	return ret
}

func setDB(key string, val int32) {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, val)
	_ = database.Set(key, buf.Bytes())
}
