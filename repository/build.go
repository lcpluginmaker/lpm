package repository

import (
	"os"
	"path"

	builder "github.com/alexcoder04/LeoConsole-apkg-builder/pkg"
	"github.com/alexcoder04/lpm/utils"
)

func BuildFolder(folder string) (string, error) {
	tempFolder := path.Join(os.TempDir(), "lpm-build")
	if utils.IsDir(tempFolder) {
		err := os.RemoveAll(tempFolder)
		if err != nil {
			return "", err
		}
	}
	err := os.MkdirAll(tempFolder, 0700)
	if err != nil {
		return "", err
	}

	manifest := builder.LoadManifest(folder)
	builder.Compile(folder, manifest)
	builder.PreparePackage(folder, tempFolder, manifest)
	builder.GenPkgInfo(tempFolder, manifest)
	outputFile := path.Join(folder, manifest.PackageName+".lcp")
	builder.Compress(tempFolder, outputFile)

	return outputFile, nil
}
