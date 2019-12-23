package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetInts(filename string) ([]int, error) {
	numbers := make([]int, 10000)
	file, err := os.Open(filename)
	if err != nil {
		return numbers, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.Atoi(scanner.Text())
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}
