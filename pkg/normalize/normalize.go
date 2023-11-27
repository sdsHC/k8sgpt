package normalize

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/k8sgpt-ai/k8sgpt/pkg/common"
)


type Normalize struct {
}

func RunNormalize(a common.Analyzer) ([]common.Result, error) {
  //kind := "Deployment"
  fmt.Println("$$ start RunNormalize")
  
  deployments, err := a.Client.GetClient().AppsV1().Deployments(a.Namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		return nil, err
	}

  for _, deployment := range deployments.Items {
    //var failures []common.Failure
    var dname = deployment.Name
    var dnamespace = deployment.Namespace

    fmt.Println("# Namespace : ", dnamespace)
    fmt.Println("# Deployment : ", dname)
		
  }

  fmt.Println("$$ End RunNormalize")

  return a.Results, nil
}
