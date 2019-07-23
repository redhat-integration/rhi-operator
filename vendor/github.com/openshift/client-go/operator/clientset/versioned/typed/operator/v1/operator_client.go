// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/operator/v1"
	"github.com/openshift/client-go/operator/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type OperatorV1Interface interface {
	RESTClient() rest.Interface
	AuthenticationsGetter
	ConsolesGetter
	DNSesGetter
	EtcdsGetter
	IngressControllersGetter
	KubeAPIServersGetter
	KubeControllerManagersGetter
	KubeSchedulersGetter
	NetworksGetter
	OpenShiftAPIServersGetter
	OpenShiftControllerManagersGetter
	ServiceCAsGetter
	ServiceCatalogAPIServersGetter
	ServiceCatalogControllerManagersGetter
}

// OperatorV1Client is used to interact with features provided by the operator.openshift.io group.
type OperatorV1Client struct {
	restClient rest.Interface
}

func (c *OperatorV1Client) Authentications() AuthenticationInterface {
	return newAuthentications(c)
}

func (c *OperatorV1Client) Consoles() ConsoleInterface {
	return newConsoles(c)
}

func (c *OperatorV1Client) DNSes() DNSInterface {
	return newDNSes(c)
}

func (c *OperatorV1Client) Etcds() EtcdInterface {
	return newEtcds(c)
}

func (c *OperatorV1Client) IngressControllers(namespace string) IngressControllerInterface {
	return newIngressControllers(c, namespace)
}

func (c *OperatorV1Client) KubeAPIServers() KubeAPIServerInterface {
	return newKubeAPIServers(c)
}

func (c *OperatorV1Client) KubeControllerManagers() KubeControllerManagerInterface {
	return newKubeControllerManagers(c)
}

func (c *OperatorV1Client) KubeSchedulers() KubeSchedulerInterface {
	return newKubeSchedulers(c)
}

func (c *OperatorV1Client) Networks() NetworkInterface {
	return newNetworks(c)
}

func (c *OperatorV1Client) OpenShiftAPIServers() OpenShiftAPIServerInterface {
	return newOpenShiftAPIServers(c)
}

func (c *OperatorV1Client) OpenShiftControllerManagers() OpenShiftControllerManagerInterface {
	return newOpenShiftControllerManagers(c)
}

func (c *OperatorV1Client) ServiceCAs() ServiceCAInterface {
	return newServiceCAs(c)
}

func (c *OperatorV1Client) ServiceCatalogAPIServers() ServiceCatalogAPIServerInterface {
	return newServiceCatalogAPIServers(c)
}

func (c *OperatorV1Client) ServiceCatalogControllerManagers() ServiceCatalogControllerManagerInterface {
	return newServiceCatalogControllerManagers(c)
}

// NewForConfig creates a new OperatorV1Client for the given config.
func NewForConfig(c *rest.Config) (*OperatorV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &OperatorV1Client{client}, nil
}

// NewForConfigOrDie creates a new OperatorV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *OperatorV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new OperatorV1Client for the given RESTClient.
func New(c rest.Interface) *OperatorV1Client {
	return &OperatorV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *OperatorV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
