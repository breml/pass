package pass

import (
	"reflect"
	"runtime"
	"testing"
	"unsafe"
)

// Now marks the test as being successful (PASS:) and stops its execution by
// calling runtime.Goexit.
// If a test fails (see testing.Error, testing.Errorf, testing.Fail) and is then
// ended using this function, it is still considered to have failed.
// Execution will continue at the next test.
// Now must be called from the goroutine running the test, not from other
// goroutines created during the test. Calling Now does not stop those other
// goroutines.
func Now(t *testing.T) {
	rt := reflect.ValueOf(t).Elem()
	rf := rt.FieldByName("finished")
	rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
	rf.SetBool(true)

	runtime.Goexit()
}
