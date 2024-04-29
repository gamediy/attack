package model

import (
	"context"
	"time"
)

const (
	TaskCmdSourceIP = "1000"
)
const (
	EventTypeDelete = 1
	EventTypeSave   = 2
)

// 任务执行结果
type JobExecuteResult struct {
	ExecuteInfo *JobExecuteInfo // 执行状态
	Output      []byte          // 脚本输出
	Err         error           // 脚本错误原因
	StartTime   time.Time       // 启动时间
	EndTime     time.Time       // 结束时间
}

// 任务执行状态
type JobExecuteInfo struct {
	Job        *Job               // 任务信息
	PlanTime   time.Time          // 理论上的调度时间
	RealTime   time.Time          // 实际的调度时间
	CancelCtx  context.Context    // 任务command的context
	CancelFunc context.CancelFunc //  用于取消command执行的cancel函数
}
type JobEvent struct {
	EventType int //  SAVE, DELETE
	Job       *Job
}
type Job struct {
	Cmd        string //
	Playload   string
	AssignWork string //指定work
	Obj        string //
}
