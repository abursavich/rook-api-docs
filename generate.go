package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"

	"github.com/machinezone/configmapsecrets/pkg/genapi"
	cassandrav1alpha1 "github.com/rook/rook/pkg/apis/cassandra.rook.io/v1alpha1"
	cephv1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	nfsv1alpha1 "github.com/rook/rook/pkg/apis/nfs.rook.io/v1alpha1"
	rookv1 "github.com/rook/rook/pkg/apis/rook.io/v1"
	rookv1alpha2 "github.com/rook/rook/pkg/apis/rook.io/v1alpha2"
	"golang.org/x/sync/errgroup"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var linkRE = regexp.MustCompile(`\(https://pkg.go.dev/github.com/rook/rook/pkg/apis/([^/]+)/([^/]+)#(.+)\)`)

func main() {

	var g errgroup.Group
	for _, api := range []struct {
		PackagePath  string
		AddToScheme  func(*runtime.Scheme) error
		GroupVersion schema.GroupVersion
	}{
		{
			PackagePath:  "github.com/rook/rook/pkg/apis/cassandra.rook.io/v1alpha1",
			AddToScheme:  cassandrav1alpha1.AddToScheme,
			GroupVersion: cassandrav1alpha1.SchemeGroupVersion,
		},
		{
			PackagePath:  "github.com/rook/rook/pkg/apis/ceph.rook.io/v1",
			AddToScheme:  cephv1.AddToScheme,
			GroupVersion: cephv1.SchemeGroupVersion,
		},
		{
			PackagePath:  "github.com/rook/rook/pkg/apis/nfs.rook.io/v1alpha1",
			AddToScheme:  nfsv1alpha1.AddToScheme,
			GroupVersion: nfsv1alpha1.SchemeGroupVersion,
		},
		{
			PackagePath:  "github.com/rook/rook/pkg/apis/rook.io/v1",
			AddToScheme:  rookv1.AddToScheme,
			GroupVersion: rookv1.SchemeGroupVersion,
		},
		{
			PackagePath:  "github.com/rook/rook/pkg/apis/rook.io/v1alpha2",
			AddToScheme:  rookv1alpha2.AddToScheme,
			GroupVersion: rookv1alpha2.SchemeGroupVersion,
		},
	} {
		api := api
		g.Go(func() error {
			scheme := runtime.NewScheme()
			if err := api.AddToScheme(scheme); err != nil {
				return err
			}
			pkg, err := genapi.ParsePackage(api.PackagePath)
			if err != nil {
				return err
			}
			buf := bytes.NewBuffer(nil)
			err = genapi.WriteMarkdown(buf, pkg, genapi.WithScheme(scheme), genapi.WithGroupVersion(api.GroupVersion))
			if err != nil {
				return err
			}
			// Convert links
			// from (https://pkg.go.dev/github.com/rook/rook/pkg/apis/rook.io/v1#Annotations)
			//   to (../rook.io/v1.md#Annotations)
			data := linkRE.ReplaceAll(buf.Bytes(), []byte("(../$1/$2.md#$3)"))

			filename := filepath.Join(
				"docs",
				path.Base(path.Dir(api.PackagePath)),
				path.Base(api.PackagePath)+".md",
			)
			if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
				return err
			}
			return ioutil.WriteFile(filename, data, 0644)
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
