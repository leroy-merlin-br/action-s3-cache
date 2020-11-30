package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
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
			filepath.Walk(match, func(path string, info os.FileInfo, err error) error {
				header, err := zip.FileInfoHeader(info)
				if err != nil {
					return err
				}

				header.Name = path
				header.Method = zip.Deflate

				writter, err := archive.CreateHeader(header)
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

				_, err = io.Copy(writter, file)

				return err
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
			return err
		}

		if file.FileInfo().IsDir() {
			continue
		}

		outFile, err := os.OpenFile(file.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		currentFile, err := file.Open()
		if err != nil {
			return err
		}

		if _, err = io.Copy(outFile, currentFile); err != nil {
			return err
		}

		outFile.Close()
		currentFile.Close()
	}

	return nil
}
