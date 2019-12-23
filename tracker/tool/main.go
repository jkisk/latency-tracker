package main

import (
	"fmt"
	datafile "github.com/jkisk/latency-tracker/tracker/data-input"
	"io/ioutil"
	"log"
	"sort"
)

//buckets keeps a count of results that fall within a range of times, with equal size
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
}

func (b *buckets) config(size int, limit int, samplecount int) {
	b.Size = size
	b.Limit = limit
	b.SampleCount = samplecount
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

func (b *buckets) rangePercentile(p int) int {
	target := p * b.SampleCount / 100
	current := b.Size
	count := b.Mapsy[current]
	fmt.Println(b)
	for current <= b.Limit {
		if count > target {
			return current
		}
		current += b.Size
		count += b.Mapsy[current]
	}
	return -1
}

func (b *buckets) reportRunningPercentiles() {
	p50 := b.rangePercentile(50)
	p95 := b.rangePercentile(95)
	p99 := b.rangePercentile(99)
	fmt.Printf("cumulative ms ranges:\n P50: %v-%v\n P95: %v-%v\n P99: %v-%v\n", p50, p50-b.Size, p95, p95-b.Size, p99, p99-b.Size)
}

func main() {
	// Set limit on response time before time out, and choose a bucket size below.
	// Limit should be evenly divisible by size. Smaller size will be slower and more accurate and vice versa.
	b := new(buckets)
	b.config(1000, 30000, 0)
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
		b.reportRunningPercentiles()
	}
	return
}

// ReportBatchPercentiles takes sorted slice of int and reports various percentile values from current batch.
func ReportBatchPercentiles(chunk []int) {
	p50 := chunk[5000]
	p95 := chunk[9500]
	p99 := chunk[9900]
	//console log p50 p95 p99
	fmt.Printf("CURRENT BATCH:\n P50: %v\n P95: %v\n P99: %v\n", p50, p95, p99)
}
