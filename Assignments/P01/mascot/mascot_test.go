// This declares the package name for the current file. The "_test" suffix is a convention 
// in Go indicating this package contains tests for the "mascot" package.
package mascot_test

// This section imports external libraries and other packages that this file depends on.
import (
    // "testing" is a package provided by Go's standard library. It offers utilities to 
    // facilitate writing unit tests for your code.
	"testing"

    // This imports the "mascot" package located at "example.com/program-1/mascot", 
    // allowing this test file to call functions, access types, etc. from that package.
	"example.com/program-1/mascot"
)

// This is a test function for TestMascot
// or feature being tested (in this case, "Mascot"). This is a naming convention for test 
// functions in Go. It takes a pointer to "testing.T", which provides methods to signal 
// testing status and log information.
func TestMascot(t *testing.T) {
    // This is an if statement that checks the return value of the "BestMascot" function 
    // from the "mascot" package. It tests if the returned mascot is not "Bevo".
	if mascot.BestMascot() != "Bevo" {
        // If the condition in the if statement is true (meaning the returned mascot isn't "Bevo"), 
        // the test fails. The "t.Fatal" method logs the provided message and stops the test 
        // immediately, indicating a failure.
		t.Fatal("Wrong Mascot :(")
	}
    // If the function passes this point without hitting the "t.Fatal", it means the test passed.
}
