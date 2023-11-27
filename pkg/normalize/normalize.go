package normalize

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/fatih/color"
	openapi_v2 "github.com/google/gnostic/openapiv2"
	"github.com/k8sgpt-ai/k8sgpt/pkg/ai"
	"github.com/k8sgpt-ai/k8sgpt/pkg/analyzer"
	"github.com/k8sgpt-ai/k8sgpt/pkg/cache"
	"github.com/k8sgpt-ai/k8sgpt/pkg/common"
	"github.com/k8sgpt-ai/k8sgpt/pkg/kubernetes"
	"github.com/k8sgpt-ai/k8sgpt/pkg/util"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/viper"
)


type Normalize struct {
}

func (Normalize) Normalize(a common.Analyzer) ([]common.Result, error) {
  kind := "Deployment"
	apiDoc := kubernetes.K8sApiReference{
		Kind: kind,
		ApiVersion: schema.GroupVersion{
			Group:   "apps",
			Version: "v1",
		},
		OpenapiSchema: a.OpenapiSchema,
	}
  
  deployments, err := a.Client.GetClient().AppsV1().Deployments(a.Namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		return nil, err
	}

  for _, deployment := range deployments.Items {
		var failures []common.Failure
    var dname []common.Name
		if *deployment.Spec.Replicas != deployment.Status.Replicas {
			doc := apiDoc.GetApiDocV2("spec.replicas")

      fmt.Println("# Deployment : ", name)
		
	}


  return a.Results, nil
}
