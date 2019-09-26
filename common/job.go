package common

import (
	"time"

	"github.com/pborman/uuid"
)

// Job is the structure used for any process to be worked by the Queue
type Job struct {
	UUID             string            // UUID generated by the Queue
	ToolUUID         string            // ID of the tool to use with this job
	Name             string            // Name of the job
	Status           string            // Status of the job
	Error            string            // Last returned error from the tool
	StartTime        time.Time         // Start time of the job
	PurgeTime        time.Time         // Time to remove the job from the queue during a Queue.keeper()
	RunTime          time.Time         // Time running
	ETC              string            // The estimated time of completion
	Owner            string            // Owner provided by the web frontend
	TeamVisible      []string          // Name of teams that have visibility to this job
	ResAssigned      string            // Resource this job is assinged to if any
	CrackedHashes    int64             // # of hashes cracked
	TotalHashes      int64             // # of hashes provided
	Progress         float64           // # % of cracked/provided
	Parameters       map[string]string // Parameters returned to the tool
	PerformanceData  map[string]string // Some performance status map[timestamp]perf#
	PerformanceTitle string            // Title of the perf #
	OutputData       [][]string        // A 2D array of rows for output values
	OutputTitles     []string          // The headers for the 2D array of rows above
}

// EmptyJob returns an initialized but empty job struct
func EmptyJob() Job {
	var j Job
	j.Parameters = map[string]string{}

	return j
}

// NewJob creates a new job with the provided parameters
func NewJob(tooluuid string, name string, owner string, params map[string]string) Job {
	return Job{
		UUID:            uuid.New(),
		ToolUUID:        tooluuid,
		Name:            name,
		Status:          STATUS_CREATED,
		Owner:           owner,
		Parameters:      params,
		PerformanceData: make(map[string]string),
	}
}

// CleanJobParamsForLogging takes a job and returns a map of parameters than can be used in loggin without sensitive information
func CleanJobParamsForLogging(j Job) map[string]string {
	logParam := make(map[string]string)

	for k, v := range j.Parameters {
		if k != "hashes_multiline" && k != "hashes_file_upload" && k != "dict_use_custom_prepend" && k != "dict_rules_use_custom" && k != "dict_rules_custom_file" {
			logParam[k] = v
		}
	}

	return logParam
}
