package worker

import (
	"errors"
	"github.com/ChinasoftNobody/go-datacollector/collector/executor"
	"github.com/ChinasoftNobody/go-datacollector/collector/source"
)

//start a job with meta and meta's schedule
func StartJob(meta source.Meta, schedule source.Schedule) (err error) {
	if meta.ID != schedule.SourceId {
		err = errors.New("source id not matched")
		return
	}
	jobExecutor := executor.Executor{Meta: meta, Schedule: schedule}
	jobExecutor.Execute()
	return
}
