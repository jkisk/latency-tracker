### Latency Tracker

This tool is designed to take in large batches of server latency times and compute running percentiles.

## Tech

Go Version 1.13.4

## Data Generation

I generated sample batches of data in text files with `generate-data/main.go`

## Data input

I formatted the data into slices of int with `data-input/ints.go`
