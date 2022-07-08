package kubernets

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"syscall"
	"testing"
)

const (
	mockMetadata = `{
  "version": "v1.19.8",
  "arch": "amd64",
  "ClusterRuntime": "k0s",
  "NydusFlag": false,
  "kubeVersion": "none",
  "variant": "none"
}
`
)

func TestLoadMetadata(t *testing.T) {
	const (
		rootfsPath       = "rootfs"
		metadataFileName = "Metadata"
	)
	type args struct {
		RuntimeMetadata []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Metadata
		wantErr bool
	}{
		{
			name: "test metadata file from rootfs",
			args: args{
				[]byte(mockMetadata),
			},
			want: &Metadata{
				Version:        "v1.19.8",
				Arch:           "amd64",
				ClusterRuntime: K0s,
				NydusFlag:      false,
				KubeVersion:    "none",
				Variant:        "none",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//prevent the read permission denied.
			oldmask := syscall.Umask(0)
			defer syscall.Umask(oldmask)

			dir, err := os.MkdirTemp("", "test-rootfs-metadata-tmp")
			if err != nil {
				t.Errorf("Make temp dir %s error = %s, wantErr %v", dir, err, tt.wantErr)
			}
			defer func() {
				err = os.RemoveAll(dir)
				if err != nil {
					t.Errorf("Remove temp dir %s error = %v, wantErr %v", dir, err, tt.wantErr)
				}
			}()

			err = os.Mkdir(filepath.Join(dir, rootfsPath), 0777)
			if err != nil {
				t.Errorf("Make dir %s error = %s, wantErr %v", dir, err, tt.wantErr)
			}

			err = ioutil.WriteFile(filepath.Join(dir, rootfsPath, metadataFileName), tt.args.RuntimeMetadata, 0777)
			if err != nil {
				t.Errorf("Write file %s error = %v, wantErr %v", filepath.Join(dir, rootfsPath, metadataFileName), err, tt.wantErr)
			}

			metadata, err := LoadMetadata(filepath.Join(dir, rootfsPath))
			if err != nil {
				t.Errorf("LoadMetadata error: %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(metadata, tt.want) {
				t.Errorf("Metadata loaded from file is not wanted! Got: %v, wanted: %v", metadata, tt.want)
			}
		})
	}
}
