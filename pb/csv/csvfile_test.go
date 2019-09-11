package csv

import (
	"fmt"
	"testing"
)

func TestLoadCsvCfg(t *testing.T) {
	cfg := LoadCsvCfg("C:/Users/Zhang_Li/idea-workspace/near_synonym_extraction/src/test/0_62_word_tag.csv", 1)
	fmt.Println(cfg.FileName)
	fmt.Println(len(cfg.Records))
	kv := cfg.Records[0].Record
	w1 := kv["0"]
	w2 := kv["1"]
	s1 := kv["2"]
	s2 := kv["4"]
	fmt.Println(w1, w2, s1[:5], s2[:5])
}

func TestReadCsv(t *testing.T) {
	csv := ReadCsv("C:/Users/Zhang_Li/idea-workspace/near_synonym_extraction/src/test/0_62_word_tag.csv", 1)
	fmt.Println(len(csv))
	fmt.Println(len(csv[0]))
	fmt.Println(csv[0])

}
