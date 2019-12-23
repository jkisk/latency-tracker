# Latency Tracker

This tool is designed to take in large batches of server latency times and compute running percentiles.

### Tech

Go Version 1.13.4

### Data Generation

You can generate sample batches of data in text files with `generate-data/main.go` adjusting min and max times according to reflect different scenarios.

### Data input

Is handled separately in `data-input/ints.go`, this could be expanded to support other data sources which could improve the performance of this tracker.
