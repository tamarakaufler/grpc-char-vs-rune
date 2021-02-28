// +build acceptance_tests

package acceptance

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// TODO
// test setup:
//		setup REST service (communicating with grpc-char-vs-rune) client ... using lib-service-run

//		prepare test scenarios ... using lib-service-acceptance-testing
//		run the tests and compile them

func TestMe(t *testing.T) {
	setup()
	require.NotEqual(t, "Hello", "World")
	fmt.Println("All works fine")
}

func setup() {
	fmt.Println("Getting everyting ready")
}
