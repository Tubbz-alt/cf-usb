// Code generated by go-bindata.
// sources:
// schemas/config.json
// schemas/dials.json
// DO NOT EDIT!

package driverdata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _schemasConfigJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x92\xc1\xce\xc2\x20\x0c\xc7\xcf\xe3\x29\x48\xcf\x7b\x82\xef\xfa\x99\x78\xf5\x6e\x3c\x6c\xae\x99\xa8\xa3\xd8\x32\x8d\x31\x7b\x77\x19\x46\xc3\x92\xa9\x73\x17\x48\xfe\xe5\xd7\xfe\x02\xdc\x94\xd6\x1a\xfc\xd5\x21\xfc\x69\xa0\x72\x8f\x5b\x0f\x79\x08\xc1\x31\x39\x64\x6f\x50\x42\x25\x1e\xd3\x50\x32\x1d\x90\x97\xb4\x22\xf1\x35\xa3\x2c\xd8\x9c\x91\x5f\xf5\xa4\x93\x78\x36\xb6\x86\x18\x77\x79\x4a\x3f\xd9\x7f\xb2\x36\x0c\x33\x64\xc7\xf8\xc4\x24\xe6\x23\x36\x31\xdf\x85\x6e\x83\xe4\x8d\x43\xe2\xf1\xe8\x47\x3c\x8b\xab\x4a\x5b\x34\x38\x83\x54\x59\x06\xad\x0c\x2e\x6b\xba\x6c\x21\x72\x21\xae\x7e\x60\xc3\x34\x91\x63\x43\xd5\x74\x57\x95\xee\xfd\x1a\x15\x80\xf1\xd4\x1a\xc6\x7e\xf8\xfa\xe3\x2f\xf8\xf6\xca\xa1\xbc\x51\xdd\x3d\x00\x00\xff\xff\xb1\x28\x79\xc3\x71\x02\x00\x00")

func schemasConfigJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemasConfigJson,
		"schemas/config.json",
	)
}

func schemasConfigJson() (*asset, error) {
	bytes, err := schemasConfigJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schemas/config.json", size: 625, mode: os.FileMode(420), modTime: time.Unix(1444830897, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _schemasDialsJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x05\x04\x00\x00\xff\xff\x43\xbf\xa6\xa3\x02\x00\x00\x00")

func schemasDialsJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemasDialsJson,
		"schemas/dials.json",
	)
}

func schemasDialsJson() (*asset, error) {
	bytes, err := schemasDialsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schemas/dials.json", size: 2, mode: os.FileMode(420), modTime: time.Unix(1444830312, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"schemas/config.json": schemasConfigJson,
	"schemas/dials.json": schemasDialsJson,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"schemas": &bintree{nil, map[string]*bintree{
		"config.json": &bintree{schemasConfigJson, map[string]*bintree{
		}},
		"dials.json": &bintree{schemasDialsJson, map[string]*bintree{
		}},
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

