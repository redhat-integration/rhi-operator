package config

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"

	integreatlyv1alpha1 "github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

const (
	mockProductName   = "mock"
	mockConfigMapName = "test"
	mockNamespaceName = "test"
)

type ProductConfigWrapper struct {
	config ProductConfig
}

func TestReadCamelK(t *testing.T) {
	testKey := "testKey"
	testVal := "testVal"
	testMap := fmt.Sprintf("%s: %s", testKey, testVal)
	existingResources := []runtime.Object{&corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      mockConfigMapName,
			Namespace: mockNamespaceName,
		},
		Data: map[string]string{
			string(integreatlyv1alpha1.Product3Scale):              testMap,
			string(integreatlyv1alpha1.ProductAMQOnline):           testMap,
			string(integreatlyv1alpha1.ProductAMQStreams):          testMap,
			string(integreatlyv1alpha1.ProductApicurito):           testMap,
			string(integreatlyv1alpha1.ProductCamelK):              testMap,
			string(integreatlyv1alpha1.ProductCloudResources):      testMap,
			string(integreatlyv1alpha1.ProductCodeReadyWorkspaces): testMap,
			string(integreatlyv1alpha1.ProductDataSync):            testMap,
			string(integreatlyv1alpha1.ProductFuse):                testMap,
			string(integreatlyv1alpha1.ProductFuseOnOpenshift):     testMap,
			string(integreatlyv1alpha1.ProductMonitoring):          testMap,
			string(integreatlyv1alpha1.ProductRHSSO):               testMap,
			string(integreatlyv1alpha1.ProductRHSSOUser):           testMap,
			string(integreatlyv1alpha1.ProductSolutionExplorer):    testMap,
			string(integreatlyv1alpha1.ProductUps):                 testMap,
		},
	}}
	fakeClient := fake.NewFakeClient(existingResources...)
	fakeInst := &integreatlyv1alpha1.RHMI{}

	mgr, err := NewManager(context.TODO(), fakeClient, mockNamespaceName, mockConfigMapName, fakeInst)
	if err != nil {
		t.Fatalf("could not create manager %v", err)
	}

	configs := make([]ProductConfig, 0)

	amqOnline, _ := mgr.ReadAMQOnline()
	if err != nil {
		t.Fatalf("could not read AMQ Online config %v", err)
	}
	configs = append(configs, amqOnline.Read())

	amqStreams, err := mgr.ReadAMQStreams()
	if err != nil {
		t.Fatalf("could not read AMQ Streams config %v", err)
	}
	configs = append(configs, amqStreams.Read())

	apicurito, err := mgr.ReadApicurito()
	if err != nil {
		t.Fatalf("could not read apicurito config %v", err)
	}
	configs = append(configs, apicurito.Read())

	camelK, err := mgr.ReadCamelK()
	if err != nil {
		t.Fatalf("could not read camelK config %v", err)
	}
	configs = append(configs, camelK.Read())

	cloudResources, err := mgr.ReadCloudResources()
	if err != nil {
		t.Fatalf("could not read could resources config %v", err)
	}
	configs = append(configs, cloudResources.Read())

	codeReady, err := mgr.ReadCodeReady()
	if err != nil {
		t.Fatalf("could not read CodeReady config %v", err)
	}
	configs = append(configs, codeReady.Read())

	dataSync, err := mgr.ReadDataSync()
	if err != nil {
		t.Fatalf("could not read DataSync config %v", err)
	}
	configs = append(configs, dataSync.Read())

	fuse, err := mgr.ReadFuse()
	if err != nil {
		t.Fatalf("could not read Fuse config %v", err)
	}
	configs = append(configs, fuse.Read())

	fuseOnOpenshift, err := mgr.ReadFuseOnOpenshift()
	if err != nil {
		t.Fatalf("could not read Fuse on OpenShift config %v", err)
	}
	configs = append(configs, fuseOnOpenshift.Read())

	monitoring, err := mgr.ReadMonitoring()
	if err != nil {
		t.Fatalf("could not read monitoring config %v", err)
	}
	configs = append(configs, monitoring.Read())

	rhsso, err := mgr.ReadRHSSO()
	if err != nil {
		t.Fatalf("could not read RH-SSO config %v", err)
	}
	configs = append(configs, rhsso.Read())

	rhssoUser, err := mgr.ReadRHSSOUser()
	if err != nil {
		t.Fatalf("could not read RH-SSO User config %v", err)
	}
	configs = append(configs, rhssoUser.Read())

	solutionExplorer, err := mgr.ReadSolutionExplorer()
	if err != nil {
		t.Fatalf("could not read Solution Explorer config %v", err)
	}
	configs = append(configs, solutionExplorer.Read())

	threeScale, _ := mgr.ReadThreeScale()
	if err != nil {
		t.Fatalf("could not read 3scale config %v", err)
	}
	configs = append(configs, threeScale.Read())

	ups, err := mgr.ReadUps()
	if err != nil {
		t.Fatalf("could not read ups config %v", err)
	}
	configs = append(configs, ups.Read())

	for _, config := range configs {
		if config[testKey] != testVal {
			t.Fatalf("expected '%s' but got '%s' for key '%s'", testVal, config[testKey], testKey)
		}
	}
}

func TestWriteConfig(t *testing.T) {
	defaultProductConfig := ProductConfig{"testKey": "testVal"}
	defaultConfigReadable := &ConfigReadableMock{
		GetProductNameFunc: func() integreatlyv1alpha1.ProductName {
			return mockProductName
		},
		ReadFunc: func() ProductConfig {
			return defaultProductConfig
		},
	}

	fakeInst := &integreatlyv1alpha1.RHMI{}

	tests := []struct {
		productName       string
		existingResources []runtime.Object
		toWrite           ConfigReadable
		expected          ProductConfig
	}{
		// Test basic adding config
		{
			productName: mockProductName,
			existingResources: []runtime.Object{&corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      mockConfigMapName,
					Namespace: mockNamespaceName,
				},
			}},
			toWrite:  defaultConfigReadable,
			expected: defaultProductConfig,
		},
		// Test overwrite config
		{
			productName: mockProductName,
			existingResources: []runtime.Object{&corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      mockConfigMapName,
					Namespace: mockNamespaceName,
				},
				Data: map[string]string{
					"testKey1": "testVal1",
					"testKey2": "testVal2",
				},
			}},
			toWrite:  defaultConfigReadable,
			expected: defaultProductConfig,
		},
		// Test create configmap if one doesn't exist
		{
			productName:       mockProductName,
			existingResources: []runtime.Object{},
			toWrite:           defaultConfigReadable,
			expected:          defaultProductConfig,
		},
	}
	for _, test := range tests {
		fakeClient := fake.NewFakeClient(test.existingResources...)

		mgr, err := NewManager(context.TODO(), fakeClient, mockNamespaceName, mockConfigMapName, fakeInst)
		if err != nil {
			t.Fatalf("could not create manager %v", err)
		}
		if err = mgr.WriteConfig(test.toWrite); err != nil {
			t.Fatalf("could not write config %v", err)
		}
		readCfgMap := &corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      mockConfigMapName,
				Namespace: mockNamespaceName,
			},
		}
		fakeClient.Get(context.TODO(), k8sclient.ObjectKey{Name: mockConfigMapName, Namespace: mockNamespaceName}, readCfgMap)

		decoder := yaml.NewDecoder(strings.NewReader(readCfgMap.Data[test.productName]))
		testCfg := map[string]string{}
		decoder.Decode(testCfg)

		for key, value := range test.expected {
			if strings.Compare(testCfg[key], value) != 0 {
				t.Fatalf("expected %s but got %s for key %s", value, testCfg[key], key)
			}
		}
	}
}

func TestReadConfigForProduct(t *testing.T) {
	fakeInst := &integreatlyv1alpha1.RHMI{}

	tests := []struct {
		productName       string
		existingResources []runtime.Object
		expected          ProductConfig
	}{
		{
			productName: mockProductName,
			existingResources: []runtime.Object{&corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      mockConfigMapName,
					Namespace: mockNamespaceName,
				},
				Data: map[string]string{
					"mock": "testKey: testVal",
				},
			}},
			expected: ProductConfig{"testKey": "testVal"},
		},
		{
			productName:       mockProductName,
			existingResources: []runtime.Object{},
			expected:          map[string]string{},
		},
	}

	for _, test := range tests {
		fakeClient := fake.NewFakeClient(test.existingResources...)
		mgr, err := NewManager(context.TODO(), fakeClient, mockNamespaceName, mockConfigMapName, fakeInst)
		if err != nil {
			t.Fatalf("could not create manager %v", err)
		}
		config, err := mgr.readConfigForProduct(mockProductName)
		if err != nil {
			t.Fatalf("could not read config %v", err)
		}
		for key, value := range test.expected {
			if strings.Compare(config[key], value) != 0 {
				t.Fatalf("expected %s but got %s for key %s", value, config[key], key)
			}
		}
	}

}
