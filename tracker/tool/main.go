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
type Buckets struct {
	CountsByRange          map[int]int
	SampleCount, Size, Max int
}

// Create map with keys that will represent range of results tallied there.
func (b *Buckets) makeCountsByRange() {
	b.CountsByRange = make(map[int]int)
	i := b.Max / b.Size
	current := b.Max
	for i > 0 {
		b.CountsByRange[current] = 0
		current -= b.Size
		i--
	}
}

// Take a slice of int and increase the counts in the appropriate buckets.
func (b *Buckets) fillBuckets(chunk []int) {
	for _, item := range chunk {
		index := item - (item % b.Size) + b.Size
		b.CountsByRange[index]++
	}
	b.SampleCount += len(chunk)
}

// Calculate running percentile ranges.
func (b *Buckets) rangePercentile(p int) int {
	target := p * b.SampleCount / 100
	current := b.Size
	count := b.CountsByRange[current]

	for current <= b.Max {
		if count > target {
			return current
		}
		current += b.Size
		count += b.CountsByRange[current]
	}
	return -1
}

// Report percentile ranges.
func (b *Buckets) reportRunningPercentiles() {
	p50 := b.rangePercentile(50)
	p95 := b.rangePercentile(95)
	p99 := b.rangePercentile(99)
	fmt.Printf("cumulative ms ranges:\n P50: %v-%v\n P95: %v-%v\n P99: %v-%v\n", p50-b.Size, p50, p95-b.Size, p95, p99-b.Size, p99)
}

func main() {
	// Set Max on response time before time out, and choose a bucket size below.
	// Max should be evenly divisible by size. Smaller size will be slower and more accurate and vice versa.
	b := &Buckets{
		Size:        1000,
		Max:         30000,
		SampleCount: 0,
	}

	b.makeCountsByRange()

	files, err := ioutil.ReadDir("test-data")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		chunk, err := datafile.GetInts("test-data/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}

		reportBatchPercentiles(chunk)
		b.fillBuckets(chunk)
		b.reportRunningPercentiles()
	}
	return
}

// ReportBatchPercentiles takes slice of int and reports various percentile values from current batch.
func reportBatchPercentiles(chunk []int) {
	sort.Ints(chunk)
	p50 := chunk[5000]
	p95 := chunk[9500]
	p99 := chunk[9900]
	// Output results for current batch.
	fmt.Printf("CURRENT BATCH:\n P50: %v\n P95: %v\n P99: %v\n", p50, p95, p99)
}
