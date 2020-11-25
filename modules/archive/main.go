package archive

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CreateZip(patterns[]string, fileName string) (string, error) {
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

func Unzip() {

}