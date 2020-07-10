package common

import (
	"fmt"
	"testing"
	"time"

	v1 "github.com/apache/camel-k/pkg/apis/camel/v1"
	camelkclientversioned "github.com/apache/camel-k/pkg/client/camel/clientset/versioned"
	integreatlyv1alpha1 "github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1"
	"github.com/operator-framework/operator-sdk/pkg/test/e2eutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TestCamelKCreateAndDelete tests the creation and deletion of a Camel K Integration
func TestCamelKCreateAndDelete(t *testing.T, ctx *TestingContext) {
	name := "dummy"
	ns, _ := getCamelKNamespace(t, ctx)
	camelClient, _ := camelkclientversioned.NewForConfig(ctx.KubeConfig)

	// Create Integration
	_, err := camelClient.CamelV1().Integrations(ns).Create(&v1.Integration{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	})
	if err != nil {
		t.Fatalf("error creating integration: %v", err)
	}

	// Check Deployment was created
	err = e2eutil.WaitForDeployment(t, ctx.KubeClient, ns, name, 1, time.Second*10, time.Minute*10)
	if err != nil {
		t.Fatalf("error deploying integration: %v", err)
	}
	_, err = ctx.KubeClient.AppsV1().Deployments(ns).Get(name, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("error getting integration deployment: %v", err)
	}

	// Delete Integration
	err = camelClient.CamelV1().Integrations(ns).Delete(name, nil)
	if err != nil {
		t.Fatalf("error deleting integration: %v", err)
	}
}

func getCamelKNamespace(t *testing.T, ctx *TestingContext) (string, error) {
	rhmi, err := getRHMI(ctx.Client)
	if err != nil {
		t.Fatal(err.Error())
	}
	return fmt.Sprintf("%v%v", rhmi.Spec.NamespacePrefix, integreatlyv1alpha1.ProductCamelK), nil
}
