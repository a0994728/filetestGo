package create

import (
	"encoding/csv"
	"fmt"
	"main/entity"
	"os"
)

func CreateFile(data entity.WorldList, fileName string) error {

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err, "failed to create file")
	}
	defer func(f *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err, "failed to close file")
		}
	}(file)

	cw := csv.NewWriter(file)
	defer cw.Flush()

	for i, world := range data {
		col := []string{world.CountryCode.String, world.CountryName.String, world.IndepYear.Decimal.String()}
		err = cw.Write(col)
		if err != nil {
			fmt.Println(err, "could not write", i, "th record")
		}
	}

	return err

}
