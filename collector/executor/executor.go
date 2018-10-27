package executor

import "github.com/ChinasoftNobody/go-datacollector/collector/source"

type Collector interface {
	collect()
	Execute()
}

type Executor struct {
	Meta     source.Meta
	Schedule source.Schedule
}

func (executor Executor) collect() {
	// 采集
	switch executor.Meta.Type {
	case source.MetaTypeWebSite:
		//TODO reptile
	}
}

func (executor Executor) Execute() {
	switch executor.Schedule.Mode {
	case source.ScheduleModeManual:
		// schedule manage
		go func() {
			executor.collect()
		}()
	}
}
