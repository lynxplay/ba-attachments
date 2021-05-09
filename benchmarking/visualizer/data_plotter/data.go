package data_plotter

import (
	"fmt"
	"strconv"
	"time"
)

type Platform = string
type Framework = string

type ClientDataPoints = []ClientDataPoint

type MemoryDataPoints = []MemoryDataPoint
type ResultDataPoints struct {
	MemoryDataPoints MemoryDataPoints
	ClientDataPoints ClientDataPoints
}
type FrameworkResultDataPoints = map[Platform]ResultDataPoints

type ResultDataPointMap = map[Framework]FrameworkResultDataPoints

type MemoryDataPoint struct {
	Timestamp float64
	CPUTime   float64
	CPUUsage  float64
	RSS       float64
}

const (
	FrameworkQuarkus    Framework = "quarkus"
	FrameworkMicronaut  Framework = "micronaut"
	FrameworkHelidon    Framework = "helidon"
	FrameworkSpringBoot Framework = "springboot"

	PlatformHotSpot Platform = "hotspot"
	PlatformOpenJ9  Platform = "openj9"
	PlatformNative  Platform = "native"
)

func ParseMemoryDataPoint(timestamp string, cpuTime string, cpuUsage string, rss string) (MemoryDataPoint, error) {
	point := MemoryDataPoint{}
	if parseResult, err := strconv.Atoi(timestamp); err == nil {
		point.Timestamp = float64(parseResult)
	} else {
		return point, fmt.Errorf("failed to parse %s to memory timestamp: %s", timestamp, err)
	}

	if parseResult, err := time.Parse("4:05", cpuTime); err == nil {
		point.CPUTime = float64(parseResult.Second() + parseResult.Minute()*60)
	} else {
		return point, fmt.Errorf("failed to parse %s to cpu time: %s", cpuTime, err)
	}

	if parseResult, err := strconv.ParseFloat(cpuUsage, 64); err == nil {
		point.CPUUsage = parseResult
	} else {
		return point, fmt.Errorf("failed to parse %s to memory cpu usage: %s", cpuUsage, err)
	}

	if parseResult, err := strconv.Atoi(rss); err != nil {
		return point, fmt.Errorf("failed to parse %s to memory rss: %s", rss, err)
	} else {
		point.RSS = float64(parseResult)
	}
	return point, nil
}

type ClientDataPoint struct {
	Timestamp       float64
	Elapsed         float64
	Latency         float64
	ConnectDuration float64
}

func ParseClientDataPoint(timestamp string, elapsed string, latency string, connected string) (ClientDataPoint, error) {
	point := ClientDataPoint{}
	if parseResult, err := strconv.Atoi(timestamp); err == nil {
		point.Timestamp = float64(parseResult)
	} else {
		return point, fmt.Errorf("failed to parse %s to client timestamp: %s", timestamp, err)
	}

	if parseResult, err := strconv.Atoi(elapsed); err == nil {
		point.Elapsed = float64(parseResult)
	} else {
		return point, fmt.Errorf("failed to parse %s to client elapsed[ms]: %s", elapsed, err)
	}

	if parseResult, err := strconv.Atoi(latency); err == nil {
		point.Latency = float64(parseResult)
	} else {
		return point, fmt.Errorf("failed to parse %s to client latency: %s", latency, err)
	}

	if parseResult, err := strconv.Atoi(connected); err != nil {
		return point, fmt.Errorf("failed to parse %s to client connect duration: %s", connected, err)
	} else {
		point.ConnectDuration = float64(parseResult)
	}
	return point, nil
}
