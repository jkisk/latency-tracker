package main

import (
	"fmt"
	datafile "github.com/jkisk/latency-tracker/tracker/data-input"
	"io/ioutil"
	"log"
	"sort"
)

//Buckets keeps a count of results that fall within a range of times, with equal size
//and is updated with each batch of times.
type buckets struct {
	Mapsy                    map[int]int
	SampleCount, Size, Limit int
}

func (b *buckets) makeMapsyBucketsBetter() {
	b.Mapsy = make(map[int]int)
	i := b.Limit/b.Size + 1
	for i > 1 {
		b.Mapsy[b.Limit] = 0
		b.Limit -= b.Size
		i--
	}
	fmt.Println(len(b.Mapsy))
}

// fillBuckets take a sorted slice of int and increases the count in the appropriate buckets.
func (b *buckets) fillBuckets(chunk []int) {
	current := b.Size
	for _, item := range chunk {
		if item <= current {
			b.Mapsy[current]++
		} else {
			current += b.Size
		}
	}
	b.SampleCount += len(chunk)
}

func (b *buckets) rangePercentile(p int) {
	target := (p / 100) * b.SampleCount
	current := b.Size
	//
}

func main() {
	// Set limit on response time before time out, and choose a bucket size below.
	// Smaller size will be slower and more accurate and vice versa.
	b := new(buckets)

	b.SampleCount = 0
	b.Limit = 30000
	b.Size = 1000
	//(add error message if limit % size != 0?)
	b.makeMapsyBucketsBetter()

	files, err := ioutil.ReadDir("test-data")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		chunk, err := datafile.GetInts("test-data/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		sort.Ints(chunk)
		ReportBatchPercentiles(chunk)
		b.fillBuckets(chunk)
	}
	fmt.Println(b.Mapsy)
	return
}

// ReportBatches takes sorted slice of int and reports various percentile values from current batch.
func ReportBatchPercentiles(chunk []int) {
	P50 := chunk[5000]
	P95 := chunk[9500]
	P99 := chunk[9900]
	//console log p50 p95 p99
	fmt.Printf("CURRENT BATCH:\n P50: %v\n P95: %v\n P99: %v\n", P50, P95, P99)
}

//
//update batchCount

// fmt.Printf("TOTAL RUN:\n P99: %v\n", b.report99())
