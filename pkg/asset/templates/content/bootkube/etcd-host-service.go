package bootkube

import (
	"os"
	"path/filepath"

	"github.com/openshift-metalkube/kni-installer/pkg/asset"
	"github.com/openshift-metalkube/kni-installer/pkg/asset/templates/content"
)

const (
	etcdHostServiceFileName = "etcd-host-service.yaml"
)

var _ asset.WritableAsset = (*EtcdHostService)(nil)

// EtcdHostService is an asset for the etcd host-network service
type EtcdHostService struct {
	FileList []*asset.File
}

// Dependencies returns all of the dependencies directly needed by the asset
func (t *EtcdHostService) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Name returns the human-friendly name of the asset.
func (t *EtcdHostService) Name() string {
	return "EtcdHostService"
}

// Generate generates the actual files by this asset
func (t *EtcdHostService) Generate(parents asset.Parents) error {
	fileName := etcdHostServiceFileName
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
func (t *EtcdHostService) Files() []*asset.File {
	return t.FileList
}

// Load returns the asset from disk.
func (t *EtcdHostService) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, etcdHostServiceFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
