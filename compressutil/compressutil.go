package compressutil

import (
	"archive/zip"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// zip压缩
//
//	dest:  压缩后的文件路径
//	paths: 需要压缩的文件路径
func Zip(dest string, paths ...string) error {
	zfile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer zfile.Close()
	zipWriter := zip.NewWriter(zfile)
	defer zipWriter.Close()
	for _, src := range paths {
		// remove the trailing path sepeartor if it is a directory
		src := strings.TrimSuffix(src, string(os.PathSeparator))
		err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// create local file header
			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}
			// set compression method to deflate
			header.Method = zip.Deflate
			// set relative path of file in zip archive
			header.Name, err = filepath.Rel(filepath.Dir(src), path)
			if err != nil {
				return err
			}
			if info.IsDir() {
				header.Name += string(os.PathSeparator)
			}
			// create writer for writing header
			headerWriter, err := zipWriter.CreateHeader(header)
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = io.Copy(headerWriter, f)
			return err
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// zip解压缩
//
//	src:  压缩文件路径
//	dest: 解压到的目录
func Unzip(src string, dest string) error {
	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		filePath := path.Join(dest, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
				return err
			}
			inFile, err := file.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()
			outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()
			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
