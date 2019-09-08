package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(t *testing.T) {

}

func TestParseEnv(t *testing.T) {
	// mock out log.Fatalf
	origLogFatalf := logFatalf
	defer func() { logFatalf = origLogFatalf }()
	errors := []string{}
	logFatalf = func(format string, args ...interface{}) {
		if len(args) > 0 {
			errors = append(errors, fmt.Sprintf(format, args))
		} else {
			errors = append(errors, format)
		}
	}

	originLogPrintf := logMsg
	defer func() { logMsg = originLogPrintf }()
	logs := []string{}
	logMsg = func(format string, args ...interface{}) {
		if len(args) > 0 {
			logs = append(logs, fmt.Sprintf(format, args))
		} else {
			logs = append(logs, format)
		}
	}

	requiredEnvs := []string{
		"GRAPH_DB_ENDPOINT",
		"TWO_WAY_KV_ENDPOINT",
		"METRICS_PORT",
		"PARALLELISM",
	}

	for _, v := range requiredEnvs {
		os.Setenv(v, "5")
	}
	// positive test
	parseEnv()
	assert.Equal(t, len(errors), 0)

	for _, v := range requiredEnvs {
		t.Run("it validates "+v, func(t *testing.T) {
			errors = []string{}
			os.Unsetenv(v)
			parseEnv()
			assert.Equal(t, len(errors) > 0, true)
			// cleanup
			os.Setenv(v, "5")
		})
	}

	t.Run("fails if PARALLELISM is not valid int", func(t *testing.T) {
		errors = []string{}
		os.Setenv("PARALLELISM", "f232")
		parseEnv()
		assert.Equal(t, 2, len(errors))
	})
	t.Run("fails if PARALLELISM is not a positive int", func(t *testing.T) {
		errors = []string{}
		os.Setenv("PARALLELISM", "-253")
		parseEnv()
		assert.Equal(t, 1, len(errors))
	})
	t.Run("throws no errors if PARALLELISM is '-1'", func(t *testing.T) {
		errors = []string{}
		os.Setenv("PARALLELISM", "-1")
		parseEnv()
		assert.Equal(t, 0, len(errors))
	})
}
