package main

import (
	"github.com/corrots/go-demo/crawler/engine"
	"github.com/corrots/go-demo/crawler/persist"
	"github.com/corrots/go-demo/crawler/scheduler"
	"github.com/corrots/go-demo/crawler/zhenai/parser"
)

const URL = "http://www.zhenai.com/zhenghun"

func main() {
	//e := &engine.ConcurrentEngine{
	//	Scheduler: &scheduler.SimpleScheduler{},
	//	ChanCount: 10,
	//}
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := &engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		ChanCount: 50,
		ItemChan:  itemChan,
	}
	e.Run(engine.Request{
		URL:        URL,
		ParserFunc: parser.ParseCityList,
	})
}
