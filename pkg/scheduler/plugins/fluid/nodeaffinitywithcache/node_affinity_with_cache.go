package nodeaffinitywithcache

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
	"volcano.sh/volcano/pkg/scheduler/framework"
)

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
	//  NodeAffinityWithCache : inject scheduling info to pod

	requiredSchedulingTerms := []corev1.NodeSelectorTerm{}
	for _, runtimeInfo := range runtimeInfos {
		requiredSchedulingTerm, err := getRequiredSchedulingTerms(runtimeInfo)
		if err != nil {
			klog.Error("")
			return
		}
		if requiredSchedulingTerms != nil {
			requiredSchedulingTerms = append(requiredSchedulingTerms, *requiredSchedulingTerm)
		}
	}
	InjectNodeSelectorTerms(requiredSchedulingTerms, pod)
}

// InjectRequiredSchedulingTerms inject the NodeSelectorTerms into a pod
func InjectNodeSelectorTerms(requiredSchedulingTerms []corev1.NodeSelectorTerm, pod *corev1.Pod) {
	if len(requiredSchedulingTerms) == 0 {
		return
	}

	if pod.Spec.Affinity == nil {
		pod.Spec.Affinity = &corev1.Affinity{}
	}

	if pod.Spec.Affinity.NodeAffinity == nil {
		pod.Spec.Affinity.NodeAffinity = &corev1.NodeAffinity{}
	}

	if pod.Spec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution == nil {
		pod.Spec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution = &corev1.NodeSelector{}
	}

	if len(pod.Spec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms) == 0 {
		pod.Spec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms = requiredSchedulingTerms
	} else {
		for i := 0; i < len(requiredSchedulingTerms); i++ {
			pod.Spec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms[0].MatchExpressions = append(pod.Spec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms[0].MatchExpressions, requiredSchedulingTerms[i].MatchExpressions...)
		}
	}

}

func (f *NodeAffinityWithCachePlugin) OnSessionClose(ssn *framework.Session) {

}
