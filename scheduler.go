package gronjob

import "time"

//You can add a custom runner if you want so. It is used to executed a job's method
var Runner = func(f func()) { f() }

func schedulerFunc() {
	for _, job := range jobs {
		println("Checking job")
		if job.lastExec > time.Now().Unix() { //Not ready yet.
			continue
		}

		checkPeriodic(job)

		time.Sleep(time.Second * 1) //1 is the minimum schedule time
	}
}

func checkPeriodic(job *job) {
	nowSec := time.Now().Unix()

	if job.lastExec + job.period > nowSec {
		return
	}

	Runner(job.fun)

	job.lastExec = nowSec
}

func StartScheduler() {
	go func() {
		for {
			schedulerFunc()
		}
	}()
}
