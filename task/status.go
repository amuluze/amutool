// Package task
// Date: 2023/7/19 15:45
// Author: Amu
// Description:
package task

import "sync/atomic"

const (
	stopStatus = iota
	startStatus
	runStatus
	refreshStatus
)

// AtomicStatus atmoic status struct
type AtomicStatus struct {
	status uint32
}

// IsRun judge if status is runStatus
func (a *AtomicStatus) IsRun() bool {
	return atomic.LoadUint32(&a.status) == runStatus
}

// IsRefresh judge if status is refreshStatus
func (a *AtomicStatus) IsRefresh() bool {
	return atomic.LoadUint32(&a.status) == refreshStatus
}

func (a *AtomicStatus) IsStop() bool {
	return atomic.LoadUint32(&a.status) == stopStatus
}

func (a *AtomicStatus) IsStart() bool {
	return atomic.LoadUint32(&a.status) == startStatus
}

// ToStart change status to startStatus
func (a *AtomicStatus) ToStart() bool {
	return atomic.CompareAndSwapUint32(&a.status, stopStatus, startStatus)
}

// SetRunStatus set status to runStatus
func (a *AtomicStatus) SetRunStatus() {
	atomic.StoreUint32(&a.status, runStatus)
}

// SetStopStatus set status to stopStatus
func (a *AtomicStatus) SetStopStatus() {
	atomic.StoreUint32(&a.status, stopStatus)
}

// ToClose change status to stopStatus from runStatus
func (a *AtomicStatus) ToClose() bool {
	return atomic.CompareAndSwapUint32(&a.status, runStatus, stopStatus)
}

// RefreshToRun change status to runStatus from refreshStatus
func (a *AtomicStatus) RefreshToRun() bool {
	return atomic.CompareAndSwapUint32(&a.status, refreshStatus, runStatus)
}

// SetRefreshStatus set status to refreshStatus
func (a *AtomicStatus) SetRefreshStatus() {
	atomic.StoreUint32(&a.status, refreshStatus)
}
