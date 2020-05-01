package config

import (
	"k8s.io/apimachinery/pkg/runtime"

	integreatlyv1alpha1 "github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1"
)

func NewCamelK(config ProductConfig) *CamelK {
	return &CamelK{config: config}
}

type CamelK struct {
	config ProductConfig
}

func (c *CamelK) GetWatchableCRDs() []runtime.Object {
	return []runtime.Object{}
}

func (c *CamelK) GetNamespace() string {
	return c.config["NAMESPACE"]
}

func (c *CamelK) SetNamespace(newNamespace string) {
	c.config["NAMESPACE"] = newNamespace
}

func (c *CamelK) GetOperatorNamespace() string {
	return c.config["OPERATOR_NAMESPACE"]
}

func (c *CamelK) SetOperatorNamespace(newNamespace string) {
	c.config["OPERATOR_NAMESPACE"] = newNamespace
}

func (c *CamelK) GetHost() string {
	return ""
}

func (c *CamelK) Read() ProductConfig {
	return c.config
}

func (c *CamelK) GetProductName() integreatlyv1alpha1.ProductName {
	return integreatlyv1alpha1.ProductCamelK
}

func (c *CamelK) GetProductVersion() integreatlyv1alpha1.ProductVersion {
	return integreatlyv1alpha1.VersionCamelK
}

func (c *CamelK) GetOperatorVersion() integreatlyv1alpha1.OperatorVersion {
	return integreatlyv1alpha1.OperatorVersionCamelK
}
