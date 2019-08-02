package cluster

import (
	"github.com/apache/dubbo-go/common"
	"github.com/apache/dubbo-go/common/extension"
	"github.com/apache/dubbo-go/protocol"
	"sort"
)

type RouterChain struct {
	routers        []Router
	builtinRouters []Router
	Invokers       []protocol.Invoker
}

func NewRouterChain(url common.URL) *RouterChain {
	var builtinRouters []Router
	factories := extension.GetRouterFactorys()
	for _, factory := range factories {
		router, _ := factory.Router(&url)
		builtinRouters = append(builtinRouters, router)
	}
	var routers []Router
	copy(routers, builtinRouters)
	sort.Slice(routers, func(i, j int) bool {
		return routers[i].Priority() < routers[j].Priority()
	})
	return &RouterChain{
		builtinRouters: routers,
		routers:        routers,
	}
}

func (r RouterChain) AddRouters(routers []Router) {
	r.routers = append(r.routers, routers...)
	sort.Slice(r.routers, func(i, j int) bool {
		return routers[i].Priority() < routers[j].Priority()
	})
}

func (r RouterChain) SetInvokers(invokers []protocol.Invoker) {
	r.Invokers = invokers
	/*for _, _ := range r.routers {
		//router.Notify(r.Invokers)
	}*/
}
