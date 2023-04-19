package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Zip - Create .zip file and add dirs and files that match glob patterns
func Zip(filename string, artifacts []string) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("unable to create output file %s: %w", filename, err)
	}
	defer outFile.Close()

	archive := zip.NewWriter(outFile)
	defer archive.Close()

	for _, pattern := range artifacts {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			return fmt.Errorf("cannot match pattern %s: %w", pattern, err)
		}

		for _, match := range matches {
			filepath.Walk(match, func(path string, info os.FileInfo, err error) error {
				header, err := zip.FileInfoHeader(info)
				if err != nil {
					return fmt.Errorf("cannot create zip file header from %s: %w", info.Name(), err)
				}

				header.Name = path
				header.Method = zip.Deflate

				dst, err := archive.CreateHeader(header)
				if err != nil {
					return fmt.Errorf("cannot add file %s to the archive: %w", header.Name, err)
				}

				if info.IsDir() {
					return nil
				}

				src, err := os.Open(path)
				if err != nil {
					return fmt.Errorf("cannot open source file %s: %w", path, err)
				}
				defer src.Close()

				_, err = io.Copy(dst, src)
				if err != nil {
					return fmt.Errorf("cannot copy file from %s to %s: %w", header.Name, path, err)
				}
				return nil
			})
		}
	}

	return nil
}

// Unzip - Unzip all files and directories inside .zip file
func Unzip(filename string) error {
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		if err := os.MkdirAll(filepath.Dir(file.Name), os.ModePerm); err != nil {
			return fmt.Errorf("unable to create directory: %w", err)
		}

		if file.FileInfo().IsDir() {
			continue
		}

		dst, err := os.OpenFile(file.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return fmt.Errorf("unable to open destination file %s: %w", file.Name, err)
		}

		src, err := file.Open()
		if err != nil {
			return fmt.Errorf("unable to open source file %s: %w", file.Name, err)
		}

		if _, err = io.Copy(dst, src); err != nil {
			return fmt.Errorf("unable to copy from %s to %s: %w", file.Name, dst.Name(), err)
		}

		dst.Close()
		src.Close()
	}

	return nil
}
