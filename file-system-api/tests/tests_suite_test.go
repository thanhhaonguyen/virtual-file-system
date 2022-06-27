package tests_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestTests(t *testing.T) {
	RegisterFailHandler(Fail)
	testReporter := reporters.NewJUnitReporter("./reports/main-test-suite.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Main Test Suite", []Reporter{testReporter})
}
