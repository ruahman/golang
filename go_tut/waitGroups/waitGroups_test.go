package waitgroups

import "testing"

func TestWaitGroups(t *testing.T) {
	Run()
}

func TestAtomicCounters(t *testing.T) {
	RunAtomicCounters()
}

func TestMutex(t *testing.T) {
	RunMutex()
}
