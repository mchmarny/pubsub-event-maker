package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

func makeEvent(srcID string, min, max float64) string {

	ms, err := mem.SwapMemory()
	failOnErr(err)

	d, err := time.ParseDuration("1s")
	failOnErr(err)
	cs, err := cpu.Percent(d, false)
	failOnErr(err)

	ls, err := load.Avg()
	failOnErr(err)

	event := struct {
		SourceID     string    `json:"source_id"`
		EventID      string    `json:"event_id"`
		EventTs      time.Time `json:"event_ts"`
		Label        string    `json:"label"`
		MemUsed      float64   `json:"mem_used"`
		CPUUsed      float64   `json:"cpu_used"`
		Load1        float64   `json:"load_1"`
		Load5        float64   `json:"load_5"`
		Load15       float64   `json:"load_15"`
		RandomMetric float64   `json:"random_metric"`
	}{
		SourceID:     srcID,
		EventID:      fmt.Sprintf("eid-%s", uuid.NewV4().String()),
		EventTs:      time.Now().UTC(),
		Label:        *metricLabel,
		RandomMetric: min + rand.Float64()*(max-min),
		MemUsed:      ms.UsedPercent,
		CPUUsed:      cs[0],
		Load1:        ls.Load1,
		Load5:        ls.Load5,
		Load15:       ls.Load15,
	}

	data, _ := json.Marshal(event)

	return string(data)

}
