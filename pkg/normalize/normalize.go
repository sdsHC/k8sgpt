package normalize

import (
	"context"
	"fmt"

	"github.com/k8sgpt-ai/k8sgpt/pkg/kubernetes"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Normalize struct {
}

func RunNormalize(client kubernetes.Client, ns string) error {
	//kind := "Deployment"
	//fmt.Println("$$ start RunNormalize")

	fmt.Println("$$ Get Deployment")

	deployment, err := client.GetClient().AppsV1().Deployments("app").Get(context.TODO(), "tom", v1.GetOptions{})
	if err != nil {
		return fmt.Errorf("Deployment list Error %s: %v", ns, err)
	}

	fmt.Println("$$ Set imagePullSecret", deployment)
	deployment.Spec.Template.Spec.ImagePullSecrets[0].Name = "tom"

	//fmt.Println("$$ End RunNormalize")

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
