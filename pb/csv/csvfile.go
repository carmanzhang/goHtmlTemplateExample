package csv

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"
)

const FileBase = "./word_tag/"

type CsvTable struct {
	FileName string
	Records  []CsvRecord
}

type CsvRecord struct {
	Record map[string]string
}

func (c *CsvRecord) GetInt(field string) int {
	var r int
	var err error
	if r, err = strconv.Atoi(c.Record[field]); err != nil {
		log.Println(err)
		panic(err)
	}
	return r
}

func (c *CsvRecord) GetString(field string) string {
	data, ok := c.Record[field]
	if ok {
		return data
	} else {
		log.Println("Get fileld failed! fileld:", field)
		return ""
	}
}

func LoadCsvCfg(filename string, row int) *CsvTable {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if reader == nil {
		log.Println("NewReader return nil, file:", file)
		return nil
	}
	records, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return nil
	}
	if len(records) < row {
		log.Println(filename, " is empty")
		return nil
	}
	colNum := len(records[0])
	recordNum := len(records)
	var allRecords []CsvRecord
	for i := row; i < recordNum; i++ {
		record := &CsvRecord{make(map[string]string)}
		for k := 0; k < colNum; k++ {
			record.Record[records[0][k]] = records[i][k]
		}
		allRecords = append(allRecords, *record)
	}
	var result = &CsvTable{
		filename,
		allRecords,
	}
	return result
}

func ReadCsv(file string, row int) [][]string {
	cfg := LoadCsvCfg(file, row)
	records := cfg.Records
	arrs := [][]string{}
	for _, reccfg := range records {
		kv := reccfg.Record
		w1 := kv["0"]
		w2 := kv["1"]
		s1 := kv["2"]
		s2 := kv["4"]
		s1 = s1[:int(math.Min(5, float64(len(s1))))]
		s2 = s2[:int(math.Min(5, float64(len(s2))))]
		arr := []string{w1, w2, s1, s2}
		arrs = append(arrs, arr)

	}
	return arrs
}
