package main

import (
	"fmt"

	"github.com/nitram509/lib-bpmn-engine/pkg/bpmn_engine"
)

func main() {
	// create a new named engine
	engineName := "myBPMNEngine"
	bpmnEngine := bpmn_engine.New(engineName)
	// basic example loading a BPMN from file,
	process, err := bpmnEngine.LoadFromFile("./simple_task.bpmn")
	if err != nil {
		panic("file \"simple_task.bpmn\" can't be read.")
	}
	fmt.Println("process==", process.ProcessKey)
	// 注册任务处理器
	taskId := "hello-world"
	bpmnEngine.AddTaskHandler(taskId, printContextHandler)
	// setup some variables
	variables := map[string]interface{}{}
	variables["foo"] = "bar"
	// and execute the process
	instance, err := bpmnEngine.CreateInstance(process.ProcessKey, variables)
	fmt.Println("instance===", instance)
	continueInstance, err := bpmnEngine.RunOrContinueInstance(instance.GetInstanceKey())
	if err != nil {
		return
	}
	fmt.Println("continueInstance===", continueInstance)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println(instance)
}

func printContextHandler(job bpmn_engine.ActivatedJob) {
	println("< Hello World >")
	println(fmt.Sprintf("job=====           = %s", job))
	println(fmt.Sprintf("job=====           = %s", job.GetBpmnProcessId()))
	println(fmt.Sprintf("job=====           = %s", job.GetInstanceKey()))
	println(fmt.Sprintf("job=====           = %s", job.GetElementId()))
	println(fmt.Sprintf("job=====           = %s", job.GetBpmnProcessId()))
	println(fmt.Sprintf("job=====           = %s", job.GetBpmnProcessId()))
	println(fmt.Sprintf("job=====           = %s", job.GetBpmnProcessId()))
	job.Complete() // don't forget this one, or job.Fail("foobar")
}
