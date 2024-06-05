// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build linux

package cgroups

import "time"

// Cgroup interface abstracts actual Cgroup implementation (v1/v2)
type Cgroup interface {
	// Identifier returns the cgroup identifier that was generated by the selected `Filter` (default: folder name)
	Identifier() string
	// Inode returns the cgroup node inode
	Inode() uint64
	// GetParent returns parent Cgroup (will fail if used on root cgroup)
	GetParent() (Cgroup, error)
	// GetCPUStats returns all cgroup statistics at once. Each call triggers a read from filesystem (no cache)
	// The given CPUStats object is filled with new values. If re-using object, old values are not cleared on read failure.
	GetCPUStats(*CPUStats) error
	// GetMemoryStats returns all cgroup statistics at once. Each call triggers a read from filesystem (no cache)
	// The given MemoryStats object is filled with new values. If re-using object, old values are not cleared on read failure.
	GetMemoryStats(*MemoryStats) error
	// GetIOStats returns all cgroup statistics at once. Each call triggers a read from filesystem (no cache)
	// The given IOStats object is filled with new values. If re-using object, old values are not cleared on read failure.
	GetIOStats(*IOStats) error
	// GetPIDStats returns all cgroup statistics at once. Each call triggers a read from filesystem (no cache)
	// The given PIDStats object is filled with new values. If re-using object, old values are not cleared on read failure.
	GetPIDStats(*PIDStats) error
	// GetPIDs returns the list of pids in this cgroup. This call MAY have a caching layer as retrieving PIDs may be costly.
	// - When running in host PID namespace, no cache is used (cacheValidity is discarded)
	// - When running in a different PID namespace, cache is used
	GetPIDs(cacheValidity time.Duration) ([]int, error)
}

// GetStats allows to extract all available stats from cgroup
func GetStats(c Cgroup, stats *Stats) (allFailed bool, errs []error) {
	allFailed = true
	if stats == nil {
		return true, []error{&InvalidInputError{Desc: "input stats cannot be nil"}}
	}

	cpuStats := &CPUStats{}
	err := c.GetCPUStats(cpuStats)
	if err == nil {
		stats.CPU = cpuStats
		allFailed = false
	} else {
		errs = append(errs, err)
	}

	memoryStats := &MemoryStats{}
	err = c.GetMemoryStats(memoryStats)
	if err == nil {
		stats.Memory = memoryStats
		allFailed = false
	} else {
		errs = append(errs, err)
	}

	ioStats := &IOStats{}
	err = c.GetIOStats(ioStats)
	if err == nil {
		stats.IO = ioStats
		allFailed = false
	} else {
		errs = append(errs, err)
	}

	pidStats := &PIDStats{}
	err = c.GetPIDStats(pidStats)
	if err == nil {
		stats.PID = pidStats
		allFailed = false
	} else {
		errs = append(errs, err)
	}

	return
}
