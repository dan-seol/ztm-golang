//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) AverageBandwidth() int {
	sum := 0

	for _, bandwidth := range b.amount {
		sum += int(bandwidth)
	}

	return sum / len(b.amount)
}

func (c *CpuTemp) AverageCpuTemp() float64 {
	sum := 0.0

	for _, cpuTemp := range c.temp {
		sum += float64(cpuTemp)
	}

	return sum / float64(len(c.temp))
}

func (m *MemoryUsage) AverageMemoryUsage() int {
	sum := 0

	for _, memoryUsage := range m.amount {
		sum += int(memoryUsage)
	}

	return sum / len(m.amount)
}

type BandwidthComponent struct {
	averageBandwidth int
}

type CpuTempComponent struct {
	averageCpuTemp float64
}

type MemoryUsageComponent struct {
	averageMemoryUsage int
}

type Dashboard struct {
	BandwidthComponent
	CpuTempComponent
	MemoryUsageComponent
	seconds int
}

func (d *Dashboard) loadComponents(
	b BandwidthUsage,
	c CpuTemp,
	m MemoryUsage,
	s int) {
	d.averageBandwidth = b.AverageBandwidth()
	d.averageCpuTemp = c.AverageCpuTemp()
	d.averageMemoryUsage = m.AverageMemoryUsage()
	d.seconds = s
}

func (d *Dashboard) display() {
	fmt.Printf("Dashboard showing %v-second averages\n", d.seconds)
	fmt.Println("Average Bandwidth:", d.averageBandwidth)
	fmt.Println("Average CpuTemp:", d.averageCpuTemp)
	fmt.Println("Average MemoryUsage:", d.averageMemoryUsage)
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}
	dashboard := Dashboard{}
	dashboard.loadComponents(bandwidth, temp, memory, 5)
	dashboard.display()
}
