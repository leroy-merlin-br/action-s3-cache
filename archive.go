package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Zip - create .zip file and add dirs and files that match glob patterns
func Zip(patterns[]string, fileName string) (string, error) {
	outFile, err := os.Create(fmt.Sprintf("%s.zip", fileName))
	if err != nil {
		return fileName, err
	}
	defer outFile.Close()

	archive := zip.NewWriter(outFile)
	defer archive.Close()

	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			return fileName, err
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

	return fileName, nil
}

// Unzip - all files and directories inside .zip file
func Unzip(fileName string) (string, error) {
	reader, err := zip.OpenReader(fmt.Sprintf("%s.zip", fileName))
	if err != nil {
		return fileName, err
	}
	defer reader.Close()

	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(file.Name, os.ModePerm); err != nil {
				return fileName, err
			}

			continue
		}

		outFile, err := os.OpenFile(file.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return fileName, err
		}

		currentFile, err := file.Open()
		if err != nil {
			return fileName, err
		}

		_, err = io.Copy(outFile, currentFile)
		if err != nil {
			return fileName, err
		}

		outFile.Close()
		currentFile.Close()
	}

	return fileName, nil
}
