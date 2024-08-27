// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/clusterpedia-io/api/policy/v1alpha1"
	scheme "github.com/clusterpedia-io/clusterpedia/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ClusterImportPoliciesGetter has a method to return a ClusterImportPolicyInterface.
// A group's client should implement this interface.
type ClusterImportPoliciesGetter interface {
	ClusterImportPolicies() ClusterImportPolicyInterface
}

// ClusterImportPolicyInterface has methods to work with ClusterImportPolicy resources.
type ClusterImportPolicyInterface interface {
	Create(ctx context.Context, clusterImportPolicy *v1alpha1.ClusterImportPolicy, opts v1.CreateOptions) (*v1alpha1.ClusterImportPolicy, error)
	Update(ctx context.Context, clusterImportPolicy *v1alpha1.ClusterImportPolicy, opts v1.UpdateOptions) (*v1alpha1.ClusterImportPolicy, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, clusterImportPolicy *v1alpha1.ClusterImportPolicy, opts v1.UpdateOptions) (*v1alpha1.ClusterImportPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterImportPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterImportPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterImportPolicy, err error)
	ClusterImportPolicyExpansion
}

// clusterImportPolicies implements ClusterImportPolicyInterface
type clusterImportPolicies struct {
	*gentype.ClientWithList[*v1alpha1.ClusterImportPolicy, *v1alpha1.ClusterImportPolicyList]
}

// newClusterImportPolicies returns a ClusterImportPolicies
func newClusterImportPolicies(c *PolicyV1alpha1Client) *clusterImportPolicies {
	return &clusterImportPolicies{
		gentype.NewClientWithList[*v1alpha1.ClusterImportPolicy, *v1alpha1.ClusterImportPolicyList](
			"clusterimportpolicies",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha1.ClusterImportPolicy { return &v1alpha1.ClusterImportPolicy{} },
			func() *v1alpha1.ClusterImportPolicyList { return &v1alpha1.ClusterImportPolicyList{} }),
	}
}
