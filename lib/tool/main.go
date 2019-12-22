package main

import (
	"fmt"
	"github.com/jkisk/latency-tracker/lib/datafile"
	"io/ioutil"
	"log"
	"sort"
)

// Buckets keeps a slice of results at the 50th, 95th, and 99th percentile
//for each batch of 10k ints representing latency times, and a count of batches.
// type buckets struct {
// 	P50, P95, P99 []int
// 	BatchCount    int
// }

// func (b *buckets) fillBuckets() int {
// }

// reportOut sums the entries in the bucket slices, then divides by the total
// to track running percentiles to report out as the program runs.
// func (b *buckets) reportOut() int {
// 	total := 0
// 	for _, v := range b {
// 		total += v
// 	}
// 	return total / b.BatchCount
//}

func main() {
	files, err := ioutil.ReadDir("test-data")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		chunk, err := datafile.GetInts("test-data/" + cdfile.Name())
		if err != nil {
			log.Fatal(err)
		}
		sortAndRecord(chunk)
	}
	return
}

func sortAndRecord(chunk []int) {
	sort.Ints(chunk)
	//console log p50 p95 p99
	fmt.Printf("CURRENT BATCH: p50: %v, p95: %v, p99: %v", chunk[5000], chunk[9500], chunk[9900])
	//update batchCount, add p values to slices, console log updated average of slices
	// fmt.Printf("TOTAL RUN: p50: %v, p95: %v, p99: %v", b.P50.reportOut(), b.P95.reportOut(), b.P99.reportOut())
}

//compute percentiles and update and report them every 10k items
