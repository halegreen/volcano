package fluid

import "volcano.sh/volcano/pkg/scheduler/framework"

const (
	PluginName = "fluidPlugin"
)

// fluidPlugin implements data affinity scheduling
type fluidPlugin struct {
	// Arguments given for the plugin
	pluginArguments framework.Arguments
}

func New(arguments framework.Arguments) framework.Plugin {
	return &fluidPlugin{pluginArguments: arguments}
}

func (f *fluidPlugin) Name() string {
	return PluginName
}

func (f *fluidPlugin) OnSessionOpen(ssn *framework.Session) {
	//  NodeAffinityWithCache : inject scheduling to pod

}

func (f *fluidPlugin) OnSessionClose(ssn *framework.Session) {

}
