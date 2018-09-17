package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	data := "x < y"

	raw, _ := json.Marshal(data)
	fmt.Println(string(raw))

	var b1 bytes.Buffer
	json.NewEncoder(&b1).Encode(data)
	fmt.Println(b1.String())

	var b2 bytes.Buffer
	enc := json.NewEncoder(&b2)
	enc.SetEscapeHTML(false)
	enc.Encode(data)
	fmt.Println(b2.String())

	// unmarshal number. Default type is float64, but you can:
	// 1. transfer it to what you need
	var data1 = []byte(`{"status": 200}`)
	var result1 map[string]interface{}
	if err := json.Unmarshal(data1, &result1); err != nil {
		fmt.Println("error:", err)
		return
	}
	var status1 = uint64(result1["status"].(float64))
	fmt.Println("status1 value:", status1)

	// 2. use Decoder type
	var data2 = []byte(`{"status": 200}`)
	var result2 map[string]interface{}
	var decoder = json.NewDecoder(bytes.NewReader(data2))
	decoder.UseNumber()
	if err := decoder.Decode(&result2); err != nil {
		fmt.Println("error:", err)
		return
	}
	var status2, _ = result2["status"].(json.Number).Int64()
	fmt.Println("status2 value:", status2)

	// 3. 定义一个结构体，专门承接这个数字
	var data3 = []byte(`{"status": 200}`)
	var result3 struct {
		Status uint64 `json:"status"`
	}
	if err := json.NewDecoder(bytes.NewReader(data3)).Decode(&result3); err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("result => %+v\n", result3)

	// 4. 定义一个结构体，承接整个数据（json.RawMessage）
	records := [][]byte{
		[]byte(`{"status": 200, "tag":"one"}`),
		[]byte(`{"status":"ok", "tag":"two"}`),
	}

	for idx, record := range records {
		var result struct {
			StatusCode uint64
			StatusName string
			Status     json.RawMessage `json:"status"`
			Tag        string          `json:"tag"`
		}

		if err := json.NewDecoder(bytes.NewReader(record)).Decode(&result); err != nil {
			fmt.Println("error:", err)
			return
		}

		var sstatus string
		if err := json.Unmarshal(result.Status, &sstatus); err == nil {
			result.StatusName = sstatus
		}

		var nstatus uint64
		if err := json.Unmarshal(result.Status, &nstatus); err == nil {
			result.StatusCode = nstatus
		}

		fmt.Printf("[%v] result => %+v\n", idx, result)
	}
}
