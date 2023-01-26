package maybebig

import (
	"fmt"
	"math"
	"runtime"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
)

// Desribes the big float precision that is used for computations.
// Only gets evaluated/called when comparisons are closer than
// "check threshold".
const defaultPrecision uint = 1000

var precision uint = defaultPrecision

// Any comparison that is within "CheckThreshold" will do the full
// computation for the value using the "big" precision.
var defaultCheckThreshold float64 = 0.0000000001
var negCheckThreshold float64 = -1 * defaultCheckThreshold
var posCheckThreshold float64 = defaultCheckThreshold

// If double check mode is enabled, all computations are checked
// for accuracy, and if incorrect, they will be flagged for review.
// This allows you to run a slower binary to check that it generally
// works, then speend the binary up by turning off double check mode.
var doubleCheckMode bool = false
var precisionLock sync.RWMutex

var numDoubleChecksMade uint64
var numHardRecomputes uint64

func SetPrecision(p uint) {
	precisionLock.Lock()
	precision = p
	precisionLock.Unlock()
}

func SetCheckThreshold(f float64) {
	precisionLock.Lock()
	negCheckThreshold = math.Abs(f) * -1
	posCheckThreshold = math.Abs(f) * -1
	precisionLock.Unlock()
}

func EnableDoubleCheck() {
	precisionLock.Lock()
	doubleCheckMode = true
	precisionLock.Unlock()
}

func GetPrecision() uint {
	return precision
}

var precisionLosses = make(map[string]int)
var precisionMutex = sync.Mutex{}

func recordDoubleCheckError(lossyOps uint) {
	precisionMutex.Lock()
	defer precisionMutex.Unlock()
	i := 1
	for {
		_, file, no, ok := runtime.Caller(i)
		if !ok {
			return
		}
		if !strings.Contains(file, "maybebig") {
			precisionLosses[fmt.Sprintf("%s#%d\n", file, no)]++
			return
		}
		i++
	}
}

func recordDoubleCheckCount() {
	atomic.AddUint64(&numDoubleChecksMade, 1)
}

func recordHardRecompute() {
	atomic.AddUint64(&numHardRecomputes, 1)
}

func DoubleCheckErrors() string {
	precisionMutex.Lock()
	defer precisionMutex.Unlock()
	results := []string{}
	for key, losses := range precisionLosses {
		results = append(results, fmt.Sprintf("%s = %+v", key, losses))
	}
	sort.Strings(results)
	return strings.Join(results, "\n") + fmt.Sprintf("\nTotal Double Checks = %d\nTotal Hard Computes = %d\n", numDoubleChecksMade, numHardRecomputes)
}
