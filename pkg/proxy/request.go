// Copyright 2016 CodisLabs. All Rights Reserved.
// Licensed under the MIT (MIT-LICENSE.txt) license.

package proxy

import (
	"sync"

	"github.com/CodisLabs/codis/pkg/proxy/redis"
	"github.com/CodisLabs/codis/pkg/utils/sync2/atomic2"
)

type Request struct {
	Multi []*redis.Resp
	Start int64
	Batch *sync.WaitGroup
	Group *sync.WaitGroup

	OpStr string
	OpFlag
	Broken *atomic2.Bool

	*redis.Resp
	Err error

	Coalesce func() error
}

func (r *Request) IsBroken() bool {
	return r.Broken != nil && r.Broken.Get()
}

func (r *Request) MakeSubRequest(n int) []Request {
	var sub = make([]Request, n)
	for i := range sub {
		x := &sub[i]
		x.Start = r.Start
		x.Batch = r.Batch
		x.OpStr = r.OpStr
		x.OpFlag = r.OpFlag
		x.Broken = r.Broken
	}
	return sub
}
