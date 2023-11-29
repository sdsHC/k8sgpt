package normalize

import (
	"context"
	"fmt"
	"os"

	"github.com/k8sgpt-ai/k8sgpt/pkg/analysis"
	"github.com/k8sgpt-ai/k8sgpt/pkg/cache"
	"github.com/k8sgpt-ai/k8sgpt/pkg/common"
	"github.com/k8sgpt-ai/k8sgpt/pkg/kubernetes"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
)

type Normalize struct {
	Context   context.Context
	Filters   []string
	Client    *kubernetes.Client
	Results   []common.Result
	Errors    []string
	Namespace string
	Cache     cache.ICache
}

func NewNormalize(a analysis.Analysis) (*Normalize, error) {
	return &Normalize{
		Context:   a.Context,
		Filters:   a.Filters,
		Client:    a.Client,
		Results:   a.Results,
		Errors:    a.Errors,
		Namespace: a.Namespace,
		Cache:     a.Cache,
	}, nil
}

func (n *Normalize) RunNormalize() error {
	//kind := "Deployment"
	//fmt.Println("$$ start RunNormalize")
	client := n.Client

	//fmt.Println("$$ Get Deployment")

	//get Deployment Client
	deploymentClient := client.GetClient().AppsV1().Deployments("app")

	deployment, err := deploymentClient.Get(context.TODO(), "tom", v1.GetOptions{})
	if err != nil {
		return fmt.Errorf("Deployment list Error %s: %v", "app", err)
	}

	// fmt.Println("$$ Set imagePullSecret", deployment)
	//fmt.Println("$$ Cache data : ", n.Cache.)

	// deployment.Spec.Template.Spec.ImagePullSecrets[0].Name = "tom"
	deployment.Spec.Template.Spec.ImagePullSecrets = append(deployment.Spec.Template.Spec.ImagePullSecrets, corev1.LocalObjectReference{Name: "tom"})

	// Deployment를 업데이트합니다.
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		_, updateErr := deploymentClient.Update(context.TODO(), deployment, v1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		fmt.Fprintf(os.Stderr, "Error updating Deployment: %v\n", retryErr)
	}
	//fmt.Println("$$ End RunNormalize")
	fmt.Println("$$ Deployment Updated!")

	return nil
}

func TestNormalize(client kubernetes.Client, ns string) error {
	//kind := "Deployment"
	//fmt.Println("$$ start RunNormalize")

	deployments, err := client.GetClient().AppsV1().Deployments("app").List(context.Background(), v1.ListOptions{})
	if err != nil {
		return fmt.Errorf("Deployment list Error %s: %v", ns, err)
	}

	for _, deployment := range deployments.Items {
		//var failures []common.Failure
		var dname = deployment.Name
		var dnamespace = deployment.Namespace

		fmt.Println("# Namespace : ", dnamespace, ", # Deployment : ", dname)
		// fmt.Println("# Deployment : ", dname)

	}

	//fmt.Println("$$ End RunNormalize")

	return nil
}
