package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func migrations_01_project_sql() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x74, 0xcf,
		0xcf, 0x6a, 0x83, 0x40, 0x10, 0x06, 0xf0, 0xfb, 0x3c, 0xc5, 0x1c, 0x95,
		0xd6, 0xbe, 0x80, 0xa7, 0x6d, 0x9d, 0x82, 0x74, 0xfd, 0xc3, 0x3a, 0xd2,
		0xda, 0x8b, 0x6c, 0x75, 0x29, 0x9b, 0x44, 0x5d, 0xcc, 0x86, 0xbc, 0x7e,
		0x0c, 0x89, 0x21, 0x87, 0x64, 0x4e, 0x33, 0xf0, 0x1b, 0xf8, 0xbe, 0x28,
		0xc2, 0x97, 0xc1, 0xfe, 0xcf, 0xda, 0x1b, 0xac, 0x1d, 0xc0, 0x87, 0x22,
		0xc1, 0x84, 0x2c, 0xde, 0x25, 0xa1, 0x3b, 0xfc, 0xed, 0x6c, 0xf7, 0xe6,
		0xe6, 0x69, 0x63, 0x3a, 0xbf, 0x0f, 0x00, 0xd1, 0xf6, 0xb8, 0x4e, 0x45,
		0x2a, 0x15, 0xf2, 0xb2, 0x97, 0x2a, 0xcd, 0x84, 0x6a, 0xf0, 0x8b, 0x9a,
		0xd7, 0x45, 0x8d, 0x7a, 0x30, 0x57, 0xc5, 0xf4, 0xc3, 0xeb, 0x47, 0x5e,
		0x30, 0xe6, 0xb5, 0x94, 0x67, 0xa2, 0x9d, 0x6d, 0xfd, 0xb4, 0x35, 0xe3,
		0x73, 0xd2, 0xcd, 0x66, 0x89, 0xd5, 0xb7, 0xda, 0x23, 0xa7, 0x19, 0x55,
		0x2c, 0xb2, 0x92, 0x7f, 0x6f, 0x04, 0x13, 0xfa, 0x14, 0xb5, 0x5c, 0x8e,
		0xe2, 0x3b, 0x08, 0x21, 0x8c, 0x01, 0xee, 0xdb, 0x24, 0xd3, 0x71, 0x04,
		0x48, 0x54, 0x51, 0x3e, 0x6e, 0x13, 0x9f, 0x02, 0x00, 0x00, 0xff, 0xff,
		0x4e, 0x59, 0xa0, 0x67, 0xfb, 0x00, 0x00, 0x00,
	},
		"migrations/01-project.sql",
	)
}

func migrations_02_deployments_sql() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x74, 0x91,
		0x41, 0x53, 0xbb, 0x30, 0x10, 0xc5, 0xef, 0xf9, 0x14, 0x7b, 0x6c, 0xe7,
		0xff, 0xaf, 0xe3, 0xbd, 0xa7, 0x48, 0x17, 0x65, 0x84, 0xd0, 0x09, 0x61,
		0xb4, 0x5e, 0x18, 0x0a, 0x41, 0xa3, 0x90, 0x30, 0x90, 0xda, 0xf1, 0xdb,
		0x1b, 0xec, 0x80, 0xd8, 0x69, 0xb9, 0xb0, 0x99, 0xfc, 0xde, 0xee, 0xbe,
		0xbc, 0xd5, 0x0a, 0xfe, 0x35, 0xea, 0xb5, 0xcb, 0xad, 0x84, 0xb4, 0x25,
		0xc4, 0xe3, 0x48, 0x05, 0x82, 0xa0, 0x77, 0x21, 0x42, 0x7b, 0xd8, 0xd7,
		0xaa, 0xb8, 0x29, 0x65, 0x5b, 0x9b, 0xaf, 0x46, 0x6a, 0xdb, 0xc3, 0x82,
		0x00, 0xa8, 0x12, 0xce, 0xbf, 0x04, 0x79, 0x40, 0xc3, 0x53, 0xbd, 0xe5,
		0x41, 0x44, 0xf9, 0x0e, 0x1e, 0x71, 0xf7, 0xdf, 0xd1, 0xfd, 0x5b, 0x7e,
		0x4e, 0x0b, 0x7c, 0x16, 0x63, 0xcd, 0x62, 0x01, 0x2c, 0x0d, 0xc3, 0x01,
		0x3d, 0x0d, 0x92, 0x65, 0x96, 0xdb, 0x09, 0x0d, 0x22, 0x4c, 0x04, 0x8d,
		0xb6, 0xe2, 0x65, 0x42, 0x61, 0x83, 0x3e, 0x4d, 0x43, 0x77, 0x88, 0x9f,
		0x16, 0xcb, 0x41, 0xd8, 0x76, 0xe6, 0x5d, 0x16, 0x36, 0x9b, 0x6d, 0xa6,
		0xf4, 0xd4, 0xe3, 0xcf, 0x0c, 0x2d, 0x8f, 0x59, 0x61, 0x9a, 0x46, 0x59,
		0xf7, 0x3b, 0x68, 0x2b, 0xbb, 0x8b, 0xe8, 0x34, 0xe3, 0x76, 0x10, 0x7d,
		0xca, 0x4e, 0x55, 0x4a, 0xce, 0x7d, 0xef, 0x8d, 0xa9, 0x65, 0xae, 0x2f,
		0x8b, 0xaa, 0xbc, 0xee, 0xe5, 0x5c, 0x78, 0xc5, 0xd1, 0x64, 0xe4, 0x67,
		0x39, 0xc7, 0x7b, 0x31, 0x4b, 0x04, 0xa7, 0x01, 0x13, 0xa3, 0xa7, 0xea,
		0x03, 0xfc, 0x98, 0x63, 0x70, 0xcf, 0x86, 0x07, 0x85, 0xc5, 0xaf, 0xd5,
		0x25, 0x70, 0xf4, 0x91, 0x23, 0xf3, 0x30, 0x19, 0x69, 0x17, 0xd0, 0x70,
		0x11, 0x51, 0xe1, 0x3d, 0x80, 0xef, 0xba, 0x92, 0xe5, 0x9a, 0x90, 0x79,
		0xc8, 0x1b, 0x73, 0xd4, 0x84, 0x6c, 0x78, 0xbc, 0xbd, 0x1a, 0xf2, 0xfa,
		0x3b, 0x00, 0x00, 0xff, 0xff, 0xee, 0x52, 0x5c, 0xf1, 0x15, 0x02, 0x00,
		0x00,
	},
		"migrations/02-deployments.sql",
	)
}

func migrations_03_messages_sql() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x6c, 0x90,
		0xd1, 0x4e, 0xc2, 0x30, 0x14, 0x86, 0xef, 0xfb, 0x14, 0xff, 0x1d, 0x10,
		0xc5, 0x17, 0x20, 0x5e, 0xd4, 0x71, 0xa6, 0x8b, 0xa3, 0x23, 0x67, 0x87,
		0x44, 0xae, 0xc8, 0x84, 0xba, 0x34, 0x6e, 0x65, 0x61, 0x18, 0xe2, 0xdb,
		0xdb, 0x46, 0xc4, 0x91, 0xd0, 0xab, 0x26, 0xe7, 0xfb, 0xfe, 0xf6, 0xfc,
		0xd3, 0x29, 0xee, 0x5a, 0x57, 0x1f, 0xaa, 0xa3, 0xc5, 0xaa, 0x53, 0xaa,
		0x24, 0xc1, 0xb6, 0x71, 0xd6, 0x1f, 0x37, 0xad, 0xf3, 0x9b, 0xd6, 0xf6,
		0x7d, 0x55, 0xdb, 0x1e, 0x8f, 0x18, 0x9d, 0xaa, 0x83, 0x77, 0xbe, 0x1e,
		0xcd, 0x94, 0x4a, 0x98, 0xb4, 0x10, 0x44, 0x3f, 0xe5, 0x84, 0xee, 0xeb,
		0xbd, 0x71, 0xdb, 0x87, 0x0b, 0x3a, 0x56, 0x80, 0xdb, 0x61, 0x70, 0x4a,
		0xe2, 0x4c, 0xe7, 0xbf, 0xf7, 0x25, 0x67, 0x0b, 0xcd, 0x6b, 0xbc, 0xd2,
		0xfa, 0x3e, 0x80, 0x67, 0xeb, 0x0c, 0x0a, 0xbd, 0xc9, 0x9f, 0x64, 0x0a,
		0x81, 0x59, 0xe5, 0x79, 0xa4, 0x76, 0xb6, 0x6b, 0xf6, 0xdf, 0x6d, 0xfc,
		0x55, 0x48, 0xce, 0xcc, 0x05, 0xba, 0xa2, 0x92, 0xc2, 0x94, 0xc2, 0x3a,
		0x8e, 0xff, 0x85, 0x8f, 0x4f, 0xa4, 0x05, 0x53, 0xf6, 0x6c, 0xe2, 0x93,
		0x18, 0x5f, 0x45, 0x4d, 0xc0, 0x94, 0x12, 0x93, 0x49, 0xa8, 0x1c, 0x38,
		0x61, 0x87, 0x38, 0x5b, 0x68, 0x49, 0x5e, 0x90, 0x86, 0x74, 0x35, 0x09,
		0x4b, 0x0f, 0xab, 0x9a, 0xef, 0x4f, 0x5e, 0xa9, 0x39, 0x17, 0xcb, 0xdb,
		0x25, 0xcc, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xa7, 0x4c, 0x4c,
		0x58, 0x01, 0x00, 0x00,
	},
		"migrations/03-messages.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"migrations/01-project.sql": migrations_01_project_sql,
	"migrations/02-deployments.sql": migrations_02_deployments_sql,
	"migrations/03-messages.sql": migrations_03_messages_sql,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"migrations": &_bintree_t{nil, map[string]*_bintree_t{
		"02-deployments.sql": &_bintree_t{migrations_02_deployments_sql, map[string]*_bintree_t{
		}},
		"03-messages.sql": &_bintree_t{migrations_03_messages_sql, map[string]*_bintree_t{
		}},
		"01-project.sql": &_bintree_t{migrations_01_project_sql, map[string]*_bintree_t{
		}},
	}},
}}
