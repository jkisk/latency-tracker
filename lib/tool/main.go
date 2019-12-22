package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jkisk/latency-tracker/lib/datafile"
)

//Buckets keeps a slice of results at the 50th, 95th, and 99th percentile
//for each batch of 10k ints representing latency times, and a count of batches.
type buckets struct {
	Mapsy       map[int]int
	SampleCount int
}

// func (b *buckets) reportPercentile(p int) int {
// 	for _, item := range Mapsy
// }

func (b *buckets) makeMapsyBuckets() {
	b.Mapsy = make(map[int]int)
	b.Mapsy[500] = 0
	b.Mapsy[1000] = 0
	b.Mapsy[1500] = 0
	b.Mapsy[2000] = 0
}

// reports out which bucket a given percentile would fall in
// func (b *buckets) reportPercentile(int) int {

// }

func main() {
	b := new(buckets)
	b.makeMapsyBuckets()
	b.SampleCount = 0

	files, err := ioutil.ReadDir("test-data")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		chunk, err := datafile.GetInts("test-data/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		sortAndRecord(chunk, b)
	}
	return
}

func sortAndRecord(chunk []int, b *buckets) {
	// move all below to method?
	// sort.Ints(chunk)

	// P50 := chunk[5000]
	// P95 := chunk[9500]
	// P99 := chunk[9900]
	// //console log p50 p95 p99
	// fmt.Printf("CURRENT BATCH:\n P50: %v\n P95: %v\n P99: %v\n", P50, P95, P99)

	// prepare counts for buckets
	one := 0
	two := 0
	three := 0
	four := 0

	for _, item := range chunk {
		if item <= 500 {
			one += 1
		} else if item <= 1000 {
			two += 1
		} else if item <= 1500 {
			three += 1
		} else {
			four += 1
		}
	}

	//update buckets
	b.Mapsy[500] += one
	b.Mapsy[1000] += two
	b.Mapsy[1500] += three
	b.Mapsy[2000] += four

	//update batchCount
	b.SampleCount += len(chunk)
	//, console log updated average of slices
	fmt.Println(b.SampleCount)
	fmt.Println(b.Mapsy[2000])
	// fmt.Printf("TOTAL RUN:\n P99: %v\n", b.report99())
}

//compute percentiles and update and report them every 10k items
