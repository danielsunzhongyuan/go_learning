package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/astaxie/beego/logs"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
	CsvSeparator      string   `yaml:"csv_separator"`
	BaseKeyIndices    []uint32 `yaml:"base_key_index"`
	ChangeKeyIndices  []uint32 `yaml:"change_key_index"`
	TimestampKeyIndex uint32   `yaml:"time_key_index"`
	Timeout           uint32   `yaml:"timeout"`
}

var config Config

func IsFile(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func CheckArgs() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println("./deduplicate config_path input_path output_path")
		os.Exit(1)
	}

	if !IsFile(args[0]) {
		fmt.Println("config file not exists:", args[0])
		os.Exit(2)
	}

	if !IsFile(args[1]) && !IsDir(args[1]) {
		fmt.Println("input file or input folder not exists")
		os.Exit(3)
	}

	if !IsFile(args[2]) && !IsDir(args[2]) {
		fmt.Println("output file or output folder not exists")
		os.Exit(4)
	}
}

type Record struct {
	Base      string
	Change    string
	StartTime string
	Original  []string
	EndTime   string
	Timeout   uint32
}

type Records []*Record

func (r Records) Len() int {
	return len(r)
}

func (r Records) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Records) Less(i, j int) bool {
	if r[i].Change != r[j].Change {
		return r[i].Change < r[j].Change
	}
	return r[i].StartTime < r[j].StartTime
}

func (r *Record) GetKey() string {
	return r.Base
}

func (r *Record) SetEndTime(end string) {
	r.EndTime = end
}

func (r *Record) Equals(other Record) bool {
	if r == &other {
		return true
	}
	if r.Base != other.Base || r.Change != other.Change {
		return false
	}
	return true
}

func (r *Record) ShouldCompress(other Record, timeout uint32) bool {
	if r.Equals(other) {
		return true
	}
	return Recent(r.EndTime, other.StartTime, timeout)
}

func Recent(time1, time2 string, timeout uint32) bool {
	var timestamp1, timestamp2 uint64
	if time1 != "" {
		timestamp1, _ = parseTimestamp(time1)
	}
	if time2 != "" {
		timestamp2, _ = parseTimestamp(time2)
	}
	return timestamp1 > 0 && timestamp2 > 0 && math.Abs(float64(timestamp1-timestamp2)) <= float64(timeout)
}

func parseTime(t string) (uint64, error) {
	timeFormat := "2006-01-02"
	tt, err := time.Parse(timeFormat, t)
	return uint64(tt.Unix()), err
}

func parseTimestamp(t string) (uint64, error) {
	return strconv.ParseUint(t, 10, 64)
}

func main() {
	CheckArgs()

	// parse config
	configFile, err := ioutil.ReadFile(os.Args[1])
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}
	fmt.Println(config.BaseKeyIndices)

	// parse input file
	inputFile, err := os.Open(os.Args[2])
	defer inputFile.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(6)
	}

	reader := csv.NewReader(bufio.NewReader(inputFile))
	reader.Comma = []rune(config.CsvSeparator)[0]
	reader.TrimLeadingSpace = true
	reader.LazyQuotes = true
	reader.Comment = '#'
	recordMappings := make(map[string]Records)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("read csv line failed:", err)
			logs.Error(err)
		}

		tmpBaseKey, tmpChangeKey := "", ""
		for _, ind := range config.BaseKeyIndices {
			tmpBaseKey += line[ind]
		}
		for _, ind := range config.ChangeKeyIndices {
			tmpChangeKey += line[ind]
		}
		recordMappings[tmpBaseKey] = append(recordMappings[tmpBaseKey],
			&Record{
				Base:      tmpBaseKey,
				Change:    tmpChangeKey,
				Original:  line,
				StartTime: line[config.TimestampKeyIndex],
				EndTime:   line[config.TimestampKeyIndex],
				Timeout:   config.Timeout,
			})
	}
	for k, v := range recordMappings {
		fmt.Print(k, ": ")
		for _, r := range v {
			fmt.Print(r.EndTime, ", ")
		}
		fmt.Println()
	}

	inputDone := make(chan struct{})
	var outputDone atomic.Value

	originalData := make(chan Records, 1000)
	go func(c chan Records) {
		for _, v := range recordMappings {
			c <- v
		}
		close(c)
	}(originalData)
	wg := sync.WaitGroup{}
	compressedData := make(chan *Record, 4)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go merge(originalData, compressedData, inputDone, &wg, &outputDone)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for x := range compressedData {
			fmt.Println(x)
		}
	}()
	wg.Wait()
}

func merge(inputChannel chan Records, outputChannel chan *Record, done <-chan struct{}, wg *sync.WaitGroup, outputDone *atomic.Value) {
	defer wg.Done()
	for {
		select {
		case d, ok := <-inputChannel:
			if ok {
				length := len(d)
				sort.Sort(d)
				newRecord := d[0]
				for i := 1; i < length; i++ {
					if newRecord.ShouldCompress(*d[i], 2500) {
						newRecord.EndTime = d[i].StartTime
						continue
					} else {
						outputChannel <- newRecord
					}
					newRecord = d[i]
				}
				outputChannel <- newRecord
			} else {
				return
			}
		case _, ok := <-done:
			if !ok {
				return
			}
		}
	}
}
