package main

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/klauspost/compress/zip"
)

// Zip - Create .zip file and add dirs and files that match glob patterns
func Zip(filename string, artifacts []string) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	archive := zip.NewWriter(outFile)
	defer archive.Close()

	for _, pattern := range artifacts {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			return err
		}

		for _, match := range matches {
			if err := filepath.Walk(match, func(path string, info os.FileInfo, walkErr error) error {
				if walkErr != nil {
					return walkErr
				}

				header, err := zip.FileInfoHeader(info)
				if err != nil {
					return err
				}

				header.Name = path
				header.Method = zip.Deflate

				writer, err := archive.CreateHeader(header)
				if err != nil {
					return err
				}

				if info.IsDir() {
					return nil
				}

				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()

				_, err = io.Copy(writer, file)
				return err
			}); err != nil {
				return err
			}
		}
	}

	return nil
}

// Unzip - Unzip all files and directories inside .zip file using a worker pool.
func Unzip(filename string) error {
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer reader.Close()

	// Create all directories up-front to avoid races between goroutines.
	for _, file := range reader.File {
		if err := os.MkdirAll(filepath.Dir(file.Name), os.ModePerm); err != nil {
			return err
		}
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(file.Name, os.ModePerm); err != nil {
				return err
			}
		}
	}

	sem := make(chan struct{}, runtime.NumCPU())
	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	for _, f := range reader.File {
		if f.FileInfo().IsDir() {
			continue
		}
		wg.Add(1)
		sem <- struct{}{}
		go func(zf *zip.File) {
			defer wg.Done()
			defer func() { <-sem }()
			if err := extractZipFile(zf); err != nil {
				select {
				case errCh <- err:
				default:
				}
			}
		}(f)
	}

	wg.Wait()
	select {
	case err := <-errCh:
		return err
	default:
		return nil
	}
}

func extractZipFile(f *zip.File) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	out, err := os.OpenFile(f.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, rc)
	return err
}
