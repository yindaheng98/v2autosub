package main

import (
	"errors"
	"iochen.com/v2gen/v2"
	"iochen.com/v2gen/v2/ping"
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
