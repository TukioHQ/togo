// Code generated by go-bindata.
// sources:
// files/ddl.tmpl
// files/sql.tmpl
// DO NOT EDIT!

package template

import (
	"bytes"
	"compress/gzip"
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
	bytes []byte
	info  os.FileInfo
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

var _filesDdlTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x56\x6d\x6f\xdb\x36\x10\xfe\x2c\xfd\x8a\x9b\xb0\x06\xd2\xe6\xca\xe8\x80\x7d\xf1\x10\x0c\xa9\xab\x62\x06\x52\xa7\xb5\x9d\x62\x40\x50\x2c\xb4\x44\xa9\x42\xf4\x16\x92\xde\xb2\x09\xfa\xef\xbb\x23\x25\x8b\x9e\x1c\x2c\xc3\x3e\x18\x30\x4f\x77\xcf\x73\xef\x64\xc3\xe2\x07\x96\x71\x68\x5b\x08\x3f\xf6\xff\xbb\xce\x75\xf3\xb2\xa9\x85\x02\xdf\x75\xbc\x84\x29\xb6\x67\x92\xcf\xe5\x63\xe1\xb9\x4e\xdb\xbe\x86\x3c\x85\xf0\xba\xce\x32\x2e\xc2\xa8\x62\xfb\x82\x27\x64\xe4\x78\x45\x9d\xf5\x1a\xbc\xd2\xa2\xc0\x75\x7f\x67\x02\xca\x3c\x13\x4c\xe5\x75\x25\xe1\x12\xee\xbe\x48\x25\x0e\xb1\x82\xd6\x75\x2a\x56\x72\xc0\x63\x5e\x65\xae\x23\x55\xa9\xcc\x57\x3a\x76\xad\x41\x12\xac\x42\x9f\xc2\x0f\x23\x04\x51\xe1\x47\x6d\xbc\x20\xcf\x1b\x34\x50\x29\x78\xaf\x1e\x3d\x08\xd7\x04\xd9\x75\x33\x54\x20\xc0\xc5\x11\x91\x4c\x6c\xc4\xad\x62\x8a\x97\xbc\x52\x06\x91\xbe\x41\x8c\xc6\x45\xfe\x17\x3f\x81\xb1\x03\x72\x1c\x12\xd1\xcf\x12\x62\xbe\xe6\x73\x30\x1e\x72\x68\xb8\x48\x6b\x51\x4a\x50\x5f\x39\x0c\xc9\x1b\x53\x10\xc2\x2a\xd5\x9f\x8e\x12\x48\x59\x5e\x48\x82\x60\x88\xc7\x85\xa8\x05\xe4\x12\x04\x57\x07\x51\xf1\x24\x74\xd3\x43\x15\x0f\xf0\x7e\xb2\x87\xef\xb0\x12\xe1\xbb\xb7\x41\xaf\x8b\x81\x61\x41\xf0\x3f\x2c\x2e\x21\x16\x1c\xb5\x76\x54\x14\x54\x0d\x7e\xd2\xf2\x6f\x2e\xa1\xca\x0b\x52\x74\x0c\x2a\x49\x31\x0c\xd7\x89\xeb\xb2\x29\xb8\xe2\xc9\x6c\x00\x90\xbc\xe0\xb1\x5a\x0e\x72\x02\x39\xe2\xf7\x38\x17\x17\xc3\x89\x1c\x89\x84\x58\xd7\x9b\xfa\x0f\x79\x0e\x1f\x33\x01\xbf\xcd\xac\x58\x91\xc1\xe4\xdf\xea\x09\xb2\x43\xa5\xfa\x41\x07\x30\x30\xdf\x8d\x29\xa3\x4a\x7f\x41\x25\x74\x03\x95\x8e\x85\x7c\xa6\x0d\x1d\x27\xe1\xfb\x43\xe6\x7b\x23\xab\x7c\xc8\x9b\x86\x27\x0b\x78\x25\x3d\xcb\x1b\x0d\x1c\x4c\x4a\x8c\x59\xa9\x54\x5e\x1d\x38\x95\x1b\x7f\x7d\x10\xba\x3f\xa7\xfe\x87\x5a\xfe\xef\x4e\xe5\x55\x5a\x9f\xf8\xa4\x98\x50\xcf\xfa\xd4\x37\x24\xf9\xf4\xba\xb7\x4f\xc9\x89\xbe\x4a\xc9\x3e\x8c\x9e\x78\xec\x13\xf7\xb4\xc8\xa7\x65\x30\x41\x58\x3d\x92\x57\x92\x0b\x75\x9c\x28\x2c\xf1\x84\xff\x05\x90\xff\x39\xe0\x34\xaf\x72\xf9\xf5\xc5\x55\xe8\x74\x03\xf5\xac\xe8\x06\x8d\x99\x9e\x84\xd3\x16\x3f\x33\x0d\xd3\x34\x1d\xb9\xb4\xd5\x52\x03\x04\xae\x1d\xd2\x00\x3e\xcd\xcd\x40\x30\x03\x6b\x5b\xbd\x8c\x6d\xa5\xc1\x8c\xe1\x79\xba\xe9\xb4\x8d\xf1\xf8\x25\x6b\xee\x0c\x5d\xbf\x32\xdb\x6e\x66\x78\x03\x22\xb6\x26\x08\xc9\xcf\x28\xb7\x94\x3f\x1c\x4c\xdb\xc1\x4f\x07\x2e\xfe\x1c\x3d\xdc\x6a\xfe\xc9\x84\x5b\x93\x8c\xc7\xd9\x71\x9c\x13\x9e\x72\x01\x84\x19\x2e\x8b\x5a\x72\x3f\x30\x23\xae\x25\x6b\xfe\xa4\x7c\xed\x99\x43\x4b\xff\x64\xb7\x5b\xdd\xa7\x75\xb7\x31\xab\xfc\x8b\x67\x5a\x6d\xc2\xac\x1b\x78\x0c\xf7\x4e\xef\x03\xbc\x4a\xec\x40\xc7\x5e\x19\x15\x67\x43\xdf\xe0\x24\x9d\x6f\xd3\xf9\x9c\xf6\x2e\xde\x5b\x19\xba\x09\x54\x12\x6d\x48\x72\x53\x20\xb3\x4a\x4a\x2e\x25\x5d\x8d\x26\x9c\x19\x34\x4c\x30\xdc\xf0\x61\x18\xe2\xcd\x83\xfb\x9e\xc5\xbc\xed\x74\xe8\x08\x15\x7e\xd4\xf7\x91\xd5\xfa\x0b\x03\xb3\x00\xef\xfb\x1e\x69\x80\x40\x84\xc0\xea\x3d\x1c\x98\xff\x4d\x45\x28\xcf\x33\xf5\x3b\xa5\xd3\x77\x16\xc5\x3e\xce\xa7\xa2\xac\x40\x92\x14\xfa\x1e\xc2\x2e\xd4\x59\x38\xb9\xbf\xad\xf1\xc1\xf4\xdf\xbb\xcb\x4d\x74\xb5\x8b\x60\x77\xf5\xf6\x3a\x82\xd5\x7b\x58\xdf\xec\x20\xfa\x75\xb5\xdd\x6d\xed\xfd\xee\xbb\xa6\x19\x3e\x5f\x6d\x96\xbf\x5c\x6d\xfc\x1f\xdf\xfc\x10\xb8\xb3\xdb\xf5\xea\xd3\x6d\xe4\x9b\xd9\x08\xdc\xfb\x7f\x50\x99\xd9\xd1\x2c\xab\xf5\x36\xda\xec\x60\xb5\xde\xdd\x9c\xc0\x6a\x53\x44\xbd\xbe\x8d\xb6\xe0\x9b\x1a\xf3\x47\x08\xdf\xe5\x8c\xba\x1a\xbc\xa6\x96\x2a\x13\x5c\x7a\xb4\x43\xe1\xdb\x37\xa0\x17\x4c\x81\xd7\x31\x9d\x7f\xd6\xc7\x3e\x1d\x53\x07\xcc\x68\x68\x07\xb6\xd1\x75\xb4\xdc\x99\x20\xde\x6f\x6e\x3e\x58\x5e\x90\x19\x32\x9f\x7d\xa5\x98\x04\xd3\xd3\xaa\x7f\x4d\x90\x64\x54\x3e\x7d\x80\x68\xf2\x73\x4f\x10\xed\x01\x81\x7c\x66\xc5\x41\xa3\xdc\x5b\x55\x1c\xff\xfd\x1d\x00\x00\xff\xff\x9f\xa5\x9f\x1c\xcb\x09\x00\x00")

func filesDdlTmplBytes() ([]byte, error) {
	return bindataRead(
		_filesDdlTmpl,
		"files/ddl.tmpl",
	)
}

func filesDdlTmpl() (*asset, error) {
	bytes, err := filesDdlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/ddl.tmpl", size: 2507, mode: os.FileMode(420), modTime: time.Unix(1490514174, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesSqlTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x50\xcd\x4a\x03\x31\x10\x3e\x9b\xa7\xf8\x28\x08\x0a\x36\xbd\x0b\xbe\x81\x88\x20\x78\x29\x05\x87\xdd\x69\x5d\xba\x3b\x1b\xb3\x89\x88\x21\xef\xee\x24\x59\xbd\xd8\xd3\xfc\x7c\x3f\xf3\x25\x8e\xba\x33\x9d\x18\x29\xc1\x3e\xaf\x7d\xce\xc6\xec\x76\x78\x9c\xe7\x73\x74\xf0\x1c\xa2\x97\x05\xe1\x9d\x21\x34\x71\x8f\x25\x50\xe0\x89\x25\x58\x73\x8c\xd2\xad\xc4\x9b\x02\x2a\xe6\x07\x39\xdd\xae\x15\xc9\x5c\x35\x3d\x06\xe9\xf9\x6b\x5f\x38\x07\xa3\xfe\x9f\xe4\xdb\x0a\x0f\x98\xc8\xed\x1b\xff\xd0\x8a\xaa\x52\xda\xc2\x93\x68\x1a\xfb\xf2\x7b\x6e\xc1\x56\xa3\x29\x04\xa7\xac\x70\xc4\xe6\xfa\x63\x03\xfb\x54\xee\xe6\x7c\x5f\xde\xd0\x69\x3f\x0e\xdf\xfc\xb7\xbd\xab\x7c\x96\xbe\x6a\xf5\xb0\x4e\xff\x7d\x15\x2a\x81\x2e\x19\x68\xbe\xb7\x22\xb2\xaf\x34\xc6\xfa\x37\x75\x2c\x8e\xda\xff\x04\x00\x00\xff\xff\x92\xf4\x54\xbf\x3f\x01\x00\x00")

func filesSqlTmplBytes() ([]byte, error) {
	return bindataRead(
		_filesSqlTmpl,
		"files/sql.tmpl",
	)
}

func filesSqlTmpl() (*asset, error) {
	bytes, err := filesSqlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/sql.tmpl", size: 319, mode: os.FileMode(420), modTime: time.Unix(1490514861, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
	"files/ddl.tmpl": filesDdlTmpl,
	"files/sql.tmpl": filesSqlTmpl,
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
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"files": &bintree{nil, map[string]*bintree{
		"ddl.tmpl": &bintree{filesDdlTmpl, map[string]*bintree{}},
		"sql.tmpl": &bintree{filesSqlTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

