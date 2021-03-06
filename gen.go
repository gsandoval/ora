// +build ignore

// Copyright 2016 Tamás Gulácsi. All rights reserved.
// Use of this source code is governed by The MIT License
// found in the accompanying LICENSE file.

package main

import (
	"bytes"
	"io/ioutil"
	"log"
)

func main() {
	src, err := ioutil.ReadFile("defFloat64.go")
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(
		"defFloat32.go",
		bytes.Replace(src, []byte("64"), []byte("32"), -1),
		0644,
	); err != nil {
		log.Fatal(err)
	}

	src, err = ioutil.ReadFile("defInt64.go")
	if err != nil {
		log.Fatal(err)
	}
	src = bytes.Replace(src, []byte("//go:generate "), []byte("// Generated by "), -1)
	for _, s := range []string{"8", "16", "32"} {
		if err := ioutil.WriteFile(
			"defInt"+s+".go",
			bytes.Replace(src, []byte("64"), []byte(s), -1),
			0644,
		); err != nil {
			log.Fatal(err)
		}
	}

	for _, pair := range [][2]string{
		{"OCI_NUMBER_SIGNED", "OCI_NUMBER_UNSIGNED"},
		{"int64", "uint64"},
		{"Int64", "Uint64"},
	} {
		src = bytes.Replace(src, []byte(pair[0]), []byte(pair[1]), -1)
	}
	if err := ioutil.WriteFile("defUint64.go", src, 0644); err != nil {
		log.Fatal(err)
	}
	for _, s := range []string{"8", "16", "32"} {
		if err := ioutil.WriteFile(
			"defUint"+s+".go",
			bytes.Replace(src, []byte("64"), []byte(s), -1),
			0644,
		); err != nil {
			log.Fatal(err)
		}
	}
}
