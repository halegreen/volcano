package nodeaffinitywithcache

import "volcano.sh/volcano/pkg/scheduler/framework"

const (
	PluginName = "NodeAffinityWithCache"
)

// NodeAffinityWithCachePlugin implements data affinity scheduling
type NodeAffinityWithCachePlugin struct {
	// Arguments given for the plugin
	pluginArguments framework.Arguments
}

func New(arguments framework.Arguments) framework.Plugin {
	return &NodeAffinityWithCachePlugin{pluginArguments: arguments}
}

func (f *NodeAffinityWithCachePlugin) Name() string {
	return PluginName
}

func (f *NodeAffinityWithCachePlugin) OnSessionOpen(ssn *framework.Session) {
	//  NodeAffinityWithCache : inject scheduling to pod

}

func (f *NodeAffinityWithCachePlugin) OnSessionClose(ssn *framework.Session) {

}
