package services

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/do"
)

type UnzipArgs struct {
	ZipFile     string
	Destination string
}

type RemoveArgs struct {
	File string
}

type IFileService interface {
	Unzip(args UnzipArgs) (err error)
	Remove(args RemoveArgs) (err error)
}

type FileService struct{}

func (fs *FileService) Unzip(args UnzipArgs) (err error) {
	r, err := zip.OpenReader(args.ZipFile)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			return
		}
	}()

	os.MkdirAll(args.Destination, 0755)

	extractAndWriteFile := func(f *zip.File) error {
		// Skip macOS system files
		if strings.HasPrefix(f.Name, "__MACOSX") || strings.HasPrefix(filepath.Base(f.Name), ".") {
			return nil
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				return
			}
		}()

		// Remove the top-level directory from the path
		relPath := f.Name
		parts := strings.Split(relPath, string(os.PathSeparator))
		if len(parts) > 1 {
			relPath = strings.Join(parts[1:], string(os.PathSeparator))
		}

		path := filepath.Join(args.Destination, relPath)

		// Modified path validation to handle absolute paths
		// Resolve any symlinks in the destination path
		destPath, err := filepath.EvalSymlinks(args.Destination)
		if err != nil {
			return fmt.Errorf("failed to resolve destination path: %w", err)
		}
		destPath = filepath.Clean(destPath)

		// Resolve and validate the target path
		path = filepath.Clean(path)
		targetPath, err := filepath.EvalSymlinks(filepath.Dir(path))
		if err != nil {
			return fmt.Errorf("failed to resolve target path: %w", err)
		}

		if !strings.HasPrefix(targetPath, destPath) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					return
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func (fs *FileService) Remove(args RemoveArgs) (err error) {
	return os.RemoveAll(args.File)
}

func NewFileService(i *do.Injector) (IFileService, error) {
	return &FileService{}, nil
}
