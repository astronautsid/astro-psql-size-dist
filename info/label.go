package info

import (
	"encoding/json"
	"os"
)

func (info *Label) CountSize() uint64 {
	var count uint64 = 0
	for _, r := range info.Tables {
		count += r.Size
	}

	return count
}

func WriteToFile(filename string, labels []Label) error {
	res, err := json.Marshal(labels)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(res)
	file.Sync()
	return nil
}
