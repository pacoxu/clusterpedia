// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	"context"

	v1alpha2 "github.com/clusterpedia-io/api/cluster/v1alpha2"
	scheme "github.com/clusterpedia-io/clusterpedia/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// PediaClustersGetter has a method to return a PediaClusterInterface.
// A group's client should implement this interface.
type PediaClustersGetter interface {
	PediaClusters() PediaClusterInterface
}

// PediaClusterInterface has methods to work with PediaCluster resources.
type PediaClusterInterface interface {
	Create(ctx context.Context, pediaCluster *v1alpha2.PediaCluster, opts v1.CreateOptions) (*v1alpha2.PediaCluster, error)
	Update(ctx context.Context, pediaCluster *v1alpha2.PediaCluster, opts v1.UpdateOptions) (*v1alpha2.PediaCluster, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, pediaCluster *v1alpha2.PediaCluster, opts v1.UpdateOptions) (*v1alpha2.PediaCluster, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.PediaCluster, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.PediaClusterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.PediaCluster, err error)
	PediaClusterExpansion
}

// pediaClusters implements PediaClusterInterface
type pediaClusters struct {
	*gentype.ClientWithList[*v1alpha2.PediaCluster, *v1alpha2.PediaClusterList]
}

// newPediaClusters returns a PediaClusters
func newPediaClusters(c *ClusterV1alpha2Client) *pediaClusters {
	return &pediaClusters{
		gentype.NewClientWithList[*v1alpha2.PediaCluster, *v1alpha2.PediaClusterList](
			"pediaclusters",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha2.PediaCluster { return &v1alpha2.PediaCluster{} },
			func() *v1alpha2.PediaClusterList { return &v1alpha2.PediaClusterList{} }),
	}
}
