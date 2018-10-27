package source

import "github.com/jinzhu/gorm"

const MetaTypeWebSite = iota

const ScheduleModeManual = iota

type MetaType int

type ScheduleMode int

type Meta struct {
	gorm.Model
	Name string   `json:"name"`
	Type MetaType `json:"type"`
	Url  string   `json:"url"`
}

type Schedule struct {
	gorm.Model
	SourceId uint         `json:"source_id"`
	Name     string       `json:"name"`
	Mode     ScheduleMode `json:"mode"`
	Content  string       `json:"content"`
}
