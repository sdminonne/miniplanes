package e2e

import (
	goflag "flag"
	"os"
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/spf13/pflag"
	//"k8s.io/client-go/rest"
)

func TestE2E(t *testing.T) {
	RunE2ETests(t)
}

var (
	//kubeConfig         *rest.Config
	kubeConfigFilePath string
	StorageHost        string
	StoragePort        int
)

func TestMain(m *testing.M) {

	//pflag.StringVar(&kubeConfigFilePath, "kubeconfig", "", "Path to kubeconfig")
	pflag.StringVar(&StorageHost, "storage-host", "", "storage host/service name")
	pflag.IntVar(&StoragePort, "storage-port", 0, "storage port number")
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	pflag.Parse()
	goflag.CommandLine.Parse([]string{})

	os.Exit(m.Run())
}

// RunE2ETests runs e2e test
func RunE2ETests(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "miniapp suite")
}
