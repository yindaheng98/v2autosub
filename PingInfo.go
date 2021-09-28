package main

import (
	"errors"
	"github.com/remeh/sizedwaitgroup"
	"iochen.com/v2gen/v2"
	"iochen.com/v2gen/v2/common/mean"
	"iochen.com/v2gen/v2/ping"
	"log"
	"sort"
)

type PingInfo struct {
	Status   *ping.Status
	Duration ping.Duration
	Link     v2gen.Link
	Err      error
}

type PingInfoList []*PingInfo

func (pf *PingInfoList) Len() int {
	return len(*pf)
}

func (pf *PingInfoList) Less(i, j int) bool {
	if (*pf)[i].Err != nil {
		return false
	} else if (*pf)[j].Err != nil {
		return true
	}

	if len((*pf)[i].Status.Errors) != len((*pf)[j].Status.Errors) {
		return len((*pf)[i].Status.Errors) < len((*pf)[j].Status.Errors)
	}

	return (*pf)[i].Duration < (*pf)[j].Duration
}

func (pf *PingInfoList) Swap(i, j int) {
	(*pf)[i], (*pf)[j] = (*pf)[j], (*pf)[i]
}

func Ping(link v2gen.Link, count int, dest string) *PingInfo {
	pi := &PingInfo{
		Link: link,
	}
	status, err := link.Ping(count, dest)
	if status.Durations == nil || len(*status.Durations) == 0 {
		pi.Err = errors.New("all error")
		status.Durations = &ping.DurationList{-1}
	}
	if err != nil {
		pi.Err = err
		pi.Status = &ping.Status{
			Durations: &ping.DurationList{},
		}
	} else {
		pi.Status = &status
	}
	return pi
}

func PingAndSort(links []v2gen.Link, count int, dest string, threads int) PingInfoList {
	// make ping info list
	piList := make(PingInfoList, len(links))
	wg := sizedwaitgroup.New(threads)
	for i := range links {
		wg.Add()
		go func(i int) {
			log.Printf("[%d/%d]Pinging %s\n", i, len(links)-1, links[i].Safe())
			defer func() {
				wg.Done()
			}()
			piList[i] = Ping(links[i], count, dest)
		}(i)
	}
	wg.Wait()

	for i := range piList {
		var ok bool
		piList[i].Duration, ok = mean.ArithmeticMean(piList[i].Status.Durations).(ping.Duration)
		if !ok {
			piList[i].Duration = 0
		}
	}
	sort.Sort(&piList)
	return piList
}
