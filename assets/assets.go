// Code generated by go-bindata. DO NOT EDIT.
// sources:
// mock.tpl

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _mockTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x57\x41\x6f\xdc\x36\x13\xbd\xeb\x57\x4c\x84\x0f\x81\x14\xc8\xcc\x77\xe9\x45\xf5\x1e\x0a\x37\x69\x0d\x34\x4e\x91\xb8\xe8\x99\xa5\x66\x65\xc1\x5a\x4a\x25\xb9\xb1\x03\x81\xff\xbd\x18\x92\x92\x28\xad\x76\xeb\x14\x45\xf7\xe0\xd5\x52\x9c\x99\x37\x6f\x86\x8f\xe3\x9e\x8b\x47\x5e\x23\x0c\x03\xb0\x5f\xc3\xb3\xb5\x49\xd2\x1c\xfa\x4e\x19\xc8\x12\x00\x80\x54\x1d\xa5\x69\x0e\x98\xfa\x5f\x12\x4d\x78\x8a\x16\x45\x27\x0d\x3e\x9b\x34\xf1\x3f\xeb\xae\xab\x5b\x64\x75\xd7\x72\x59\xb3\x4e\xd5\x6f\x6b\xd5\x8b\xf4\xd2\xcb\xb7\xa2\xab\x50\x5f\xde\xa2\x0d\x37\x47\x1d\x82\x0c\x83\xe2\xb2\x46\x60\xb7\x0e\xac\xb6\x36\x1d\x06\xe6\xfe\xa2\xac\xac\x4d\xf2\x24\x31\x5f\x7b\x04\x32\xfd\x99\xcb\xaa\x45\x05\xfb\xa3\x14\x59\x00\xcb\x6e\xfc\x77\x01\x8d\x34\xa8\xf6\x5c\xe0\x60\x73\xc8\xa2\x5f\x05\xa0\x52\x9d\xca\x67\x47\x9f\x50\xf7\xa0\x8d\x3a\x0a\x03\x83\xc3\x71\xd0\x75\xec\xc0\xad\xa1\x52\xde\x32\xb1\x49\x42\x31\xe1\xc1\x01\xc8\x84\x79\x86\x93\xf0\x0a\xff\x84\x45\xd0\x87\x80\x36\x42\xbe\x8d\x2b\x40\x50\xa8\x8f\xad\xb9\x81\x72\x07\x07\xfe\x88\x99\x78\xe0\x72\x42\x9b\x7b\x44\xcf\xcd\x7a\x83\xcf\x62\xb0\x7e\x43\xdd\x79\x72\x46\x9f\xf4\xa9\x70\x3f\x72\x96\xc3\x10\x7c\x5c\x5f\x4d\x96\x83\x05\x9b\xe5\xd3\x76\x85\xba\x77\xc8\x28\x4e\xc8\x81\x32\x76\x19\x2e\xb6\x39\xb4\xd7\x57\x13\xc6\xe1\xa0\xeb\x72\xb6\x2f\xe9\x8f\x67\x92\xfc\xbb\x07\x8d\x2d\x4e\x9c\xd3\x47\x70\x8d\xe0\x42\x5d\x5f\x05\x97\xe5\xf4\xd2\xc7\x31\x47\x25\x41\xb1\x83\xae\x0b\x50\x0c\x95\x5a\x1a\x5f\x5f\xb9\x8c\x36\xad\x64\xd3\x16\xe0\xfb\x8d\xdd\xe1\x53\xe6\xba\x93\xdd\x52\x09\x24\x6f\x0b\x48\x0d\x6a\x03\x7b\xde\xb4\x58\xa5\x39\x7b\xa7\x54\x44\x44\xf0\x4e\xe7\x83\xdd\x37\xe2\x31\x73\x4f\x9f\x51\x74\xb2\x82\x37\xf0\x5d\xfe\x0d\x21\x7f\x93\xfc\x0b\x6f\x5a\xfe\x47\x8b\x63\x54\xf2\xd6\x1d\xcd\x22\xac\x9d\x3a\xcd\x7b\xbb\xdd\xdf\xa3\x36\xef\x1d\xbe\x6c\x6a\xc7\xb1\xb8\xcd\xde\x95\x69\xb7\xa3\xa0\x11\xa7\xde\x36\x38\x74\xac\x9b\x02\xba\x47\x22\x39\x00\x7b\xaf\xba\xc3\x3b\x72\x45\x4e\xf3\xd1\xd9\xab\xee\xf1\x92\x97\x66\x0f\xda\x78\xb0\xcc\x1b\xe7\x14\x3b\x55\xbd\xf0\xb8\x4a\xa0\x64\x61\x07\x23\xc1\x50\xa1\x16\xb0\x83\x98\xe5\x38\x82\x57\x24\xf6\x53\x47\x15\x8c\x28\x18\x06\x08\xaa\xf0\x19\xd5\x97\x46\xa0\x26\x3d\xf3\x72\x01\xff\x0b\x6b\x77\xfc\x80\x94\x12\xa3\x07\x6b\x3d\x6d\x77\xf8\x44\x2a\xe8\xde\x59\x7b\xd3\x36\x28\xcd\x87\x8e\x6a\xc7\x55\x8d\x86\x7a\xbe\x91\x75\x0e\x27\x9b\x02\x2c\xd1\x49\x39\x35\x3f\xf5\x35\xfb\xb1\xe1\x6d\xb0\x2e\xfc\xca\xef\x8d\x79\xb8\x95\x1a\xc5\x51\x61\x96\xe7\x71\x29\x5e\xad\x4b\xd1\x73\xd9\x88\x99\x64\x9b\x44\x9d\xf2\x7a\x04\x11\x03\x8d\x4e\x86\x28\x43\x3a\x8b\x8c\xc7\xcd\xa4\x7f\x32\x2f\x56\x6d\xb3\x69\x40\xbf\x50\x39\x1a\x72\x78\xb3\xfd\xbe\x11\x48\x1b\x02\xf4\x76\x26\x41\xa2\x61\xbf\x34\xda\xa0\xcc\x52\x23\xfa\xb4\x80\xb4\xfc\x7f\xfa\x8d\x59\xfb\x2e\x9c\x38\xbd\xc3\x27\x8f\x69\x94\x85\x03\x85\x2e\x77\x8e\x92\x18\xdc\x02\xdb\xec\x9f\x57\x95\x2a\xa1\x95\xec\x87\xaa\xa2\x6e\xfc\xec\xaa\x9a\x05\x36\xbc\xce\x90\xfb\x12\x5e\x3b\xc3\x33\x8c\x0c\xc9\x30\x5c\x8d\x9d\xf6\x01\xcd\x43\x57\x4d\x8d\x36\x7e\xa2\x46\x09\x1a\x5e\x46\xe2\xeb\x97\xa2\x3d\x79\xe1\x7c\xa2\xac\x62\x47\xb6\x88\x79\xf8\x84\x35\xf1\xa9\xce\xc1\xca\x74\xe1\x08\x61\x3e\x89\xf3\xc2\x1e\xe8\xa7\x63\xed\x0e\x0a\x66\xad\xcc\xbf\xdf\x2e\xc9\x46\x59\xe6\x86\x1c\xe5\x3f\x34\x26\x05\xa7\x8e\x72\xd7\xe5\x25\x02\xe7\x2b\xf4\x3c\x91\xa7\x04\xc2\x06\x75\xd6\xc6\xbc\xd9\x24\x92\x80\xc8\x9f\x03\xb4\x36\x0c\xd3\xc0\xb9\x2b\x99\xfa\x9d\xdd\xca\xfe\x68\xee\xc9\xda\xd2\x6c\xe0\xd6\x3e\x1e\x4d\x58\xb4\xf3\x84\xe0\x4f\x51\xa6\xe1\xcd\x46\xde\x63\xda\x24\x1f\x63\xf0\x7f\x29\x70\xa8\xd4\xf6\xf5\x38\x0e\x12\xee\x92\xd4\x6c\xa3\x25\x17\x65\xde\x4b\xa0\xad\x97\x69\xf9\xdb\x61\x69\xd5\x3b\x51\x7b\xac\x67\x02\x46\x69\xcd\x99\x5a\x9b\xe7\x0b\x53\xbb\xba\x26\x4f\xc6\x8b\xc9\x53\x01\x7b\xb9\xb4\x3d\xa7\x30\x2b\x48\xee\xd2\x8d\x27\x83\xad\xb8\x7e\x8e\x40\xdd\xb3\xd3\x3a\xe4\x05\xf9\xf8\x0f\xae\xfe\xa8\x78\x67\x27\x80\xe8\x24\xf8\x96\xbf\x2c\xd9\x8b\x39\x96\xa4\xd1\xdd\xf8\xa4\x88\xa1\xa7\xdc\x51\xdd\x6a\xe8\xe9\x20\x4f\xf7\x07\x75\xfe\xc5\x68\x39\x78\xcd\x0d\x11\xa6\xd1\xd5\x11\xa0\x19\x85\x3f\x7b\x7c\x5f\x1a\x21\xe2\x28\xdb\x6f\x89\x6c\x88\xba\xa5\x89\x27\x27\x68\x2a\xa7\x0e\x7a\xba\x71\x7e\x68\x88\xdd\xcb\x72\x6b\xf4\xf9\x07\x1d\xb1\xb2\x9e\x25\x96\x2a\xeb\xff\xa9\x39\x57\xd6\x79\x0e\x58\x56\x55\x88\xf5\xe6\x71\xef\x62\x5a\x7a\x19\xd7\x73\x90\x25\xd5\x9b\x5a\xd1\xc8\x53\x21\x2b\xa0\xeb\x8d\x06\xc6\x98\xbb\xcd\x6f\x78\xdb\x7e\xec\x4d\xd3\xc9\x17\x28\xdc\xe2\xf8\x6b\x26\x04\x5b\x41\xa0\x90\xe3\xa0\xc5\x1b\xf3\xbe\x53\x9f\x90\x57\x5f\x33\xa3\x8e\x98\xc7\xd7\xd4\x7a\x36\x5e\x5c\x61\x53\x98\x88\xf4\xf1\xfb\xaf\x00\x00\x00\xff\xff\xa3\xde\x5e\x21\x2f\x0f\x00\x00")

func mockTplBytes() ([]byte, error) {
	return bindataRead(
		_mockTpl,
		"mock.tpl",
	)
}

func mockTpl() (*asset, error) {
	bytes, err := mockTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mock.tpl", size: 3887, mode: os.FileMode(420), modTime: time.Unix(1613411744, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x78, 0xfa, 0x50, 0xef, 0xcb, 0x79, 0x22, 0x85, 0xf3, 0xcb, 0x36, 0x2, 0xc7, 0xb2, 0xfa, 0x65, 0xb9, 0x13, 0xd6, 0x4a, 0xda, 0x91, 0x23, 0x4c, 0x95, 0x41, 0xb, 0xef, 0x1a, 0x65, 0x3f, 0xc2}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"mock.tpl": mockTpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"mock.tpl": &bintree{mockTpl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
