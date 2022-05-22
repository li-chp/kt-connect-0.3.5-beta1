package cluster

import (
	"context"
	"encoding/json"
	"istio.io/api/networking/v1beta1"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type patchSubsetValue struct {
	// 操作, 例如: add、remove、replace、
	Op string `json:"op"`
	// 到什么路径，列如下面的: /spec/subset 从 / 下开始这里的位置就是根据具体的配置
	Path string `json:"path"`
	// 路径所指向的结构体
	Value *v1beta1.Subset `json:"value"`
}

type patchHTTPRouteValue struct {
	// 操作, 例如: add、remove、replace、
	Op string `json:"op"`
	// 到什么路径，列如下面的: /spec/subset 从 / 下开始这里的位置就是根据具体的配置
	Path string `json:"path"`
	// 路径所指向的结构体
	Value map[string]interface{} `json:"value"`
}

// GetVirtualServiceList get all vs in specified namespace
func (k *Kubernetes) GetVirtualServiceList(namespace string) (*v1alpha3.VirtualServiceList, error) {
	return k.IstioClient.NetworkingV1alpha3().VirtualServices(namespace).List(context.TODO(), v1.ListOptions{})
}

// GetVirtualService
func (k *Kubernetes) GetVirtualService(name, namespace string) (*v1alpha3.VirtualService, error) {
	return k.IstioClient.NetworkingV1alpha3().VirtualServices(namespace).Get(context.TODO(), name, v1.GetOptions{})
}

// GetDestinationRule
func (k *Kubernetes) GetDestinationRule(name, namespace string) (*v1alpha3.DestinationRule, error) {
	return k.IstioClient.NetworkingV1alpha3().DestinationRules(namespace).Get(context.TODO(), name, v1.GetOptions{})
}

// PatchVirtualService
func (k *Kubernetes) PatchVirtualService(name, service, namespace, op, meshKey, meshVersion string) (*v1alpha3.VirtualService, error) {
	byteData := []byte(`{"match":[{"headers":{"` + meshKey + `":{"exact":"` + meshVersion +
		`"}}}],"route":[{"destination":{"host":"` + service + `","subset":"` + meshVersion + `"}}]}`)
	var jsonData map[string]interface{}
	json.Unmarshal(byteData, &jsonData)
	patch := make([]patchHTTPRouteValue, 0)
	patch = append(patch, patchHTTPRouteValue{
		Op: op,
		// /spec/subsets/- 指定插入到数组结尾 subset 是一个数组
		Path:  "/spec/http/0",
		Value: jsonData,
	})

	val, _ := json.Marshal(patch)
	return k.IstioClient.NetworkingV1alpha3().VirtualServices(namespace).Patch(context.TODO(), name, types.JSONPatchType, val, v1.PatchOptions{})
}

// PatchDestinationRule
func (k *Kubernetes) PatchDestinationRule(name, namespace, op, meshKey, meshVersion string) (*v1alpha3.DestinationRule, error) {
	patch := make([]patchSubsetValue, 0)
	patch = append(patch, patchSubsetValue{
		Op: op,
		// /spec/subsets/- 指定插入到数组结尾 subset 是一个数组
		Path: "/spec/subsets/-1",
		Value: &v1beta1.Subset{
			Name:   meshVersion,
			Labels: map[string]string{meshKey: meshVersion},
		},
	})

	val, _ := json.Marshal(patch)
	return k.IstioClient.NetworkingV1alpha3().DestinationRules(namespace).Patch(context.TODO(), name, types.JSONPatchType, val, v1.PatchOptions{})
}
