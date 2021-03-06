package bootkube

import (
	"os"
	"path/filepath"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	etcdHostServiceEndpointsFileName = "etcd-host-service-endpoints.yaml.template"
)

var _ asset.WritableAsset = (*EtcdHostServiceEndpoints)(nil)

// EtcdHostServiceEndpoints is an asset for the etcd host network service endpoints
type EtcdHostServiceEndpoints struct {
	FileList []*asset.File
}

// Dependencies returns all of the dependencies directly needed by the asset
func (t *EtcdHostServiceEndpoints) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Name returns the human-friendly name of the asset.
func (t *EtcdHostServiceEndpoints) Name() string {
	return "EtcdHostServiceEndpoints"
}

// Generate generates the actual files by this asset
func (t *EtcdHostServiceEndpoints) Generate(parents asset.Parents) error {
	fileName := etcdHostServiceEndpointsFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{
		{
			Filename: filepath.Join(content.TemplateDir, fileName),
			Data:     []byte(data),
		},
	}
	return nil
}

// Files returns the files generated by the asset.
func (t *EtcdHostServiceEndpoints) Files() []*asset.File {
	return t.FileList
}

// Load returns the asset from disk.
func (t *EtcdHostServiceEndpoints) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, etcdHostServiceEndpointsFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
