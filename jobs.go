package gronjob

import (
	"time"
)

//job : lastExecutionTime
var jobs = make([]*job, 0)

type job struct {
	fun func() //Function to execute

	period int64 //Every period seconds

	lastExec int64 //Last time it got executed
	initExec bool //Execute when job is created
}

//The function you want to execute
func (j *job) Func(f func()) *job {
	j.fun = f
	return j
}

//The period you want to execute the function in seconds (e.g. every 5 seconds)
func (j *job) Period(period int64) *job {
	j.period = period
	return j
}

//Shall it execute the function at Start()?
func (j *job) InitExec() *job {
	j.initExec = true
	return j
}

func (j *job) Start() {
	if j.fun == nil {
		panic("Func is nil")
	}

	if j.period <= 0 {
		panic("Period is <= 0")
	}

	if j.lastExec == 0 {
		j.lastExec = time.Now().Unix()
	}
	jobs = append(jobs, j)

	if j.initExec {
		Runner(j.fun)
	}
}

func Create() *job {
	return &job{}
}