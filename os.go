package go_utils

import (
	"log"
	"os"
)

// CreateDir creates a directory.
func CreateDir(dir string) error {
	// 0755 is the permission for the directory.
	// The difference between 0755 and 0777 is that 0755 is read and write for the owner and read for everyone else.
	err := os.Mkdir(dir, 0755)
	if err != nil {
		return err
	}

	return nil
}

// CreateFile creates a file.
func CreateFile(path string) error {
	var err error

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()
	if err != nil {
		return err
	}

	return nil
}

// DeleteFile deletes a file.
func DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

// FileExists checks if a file exists.
func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// FileSize returns the size of a file.
func FileSize(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		log.Fatalln(err)
	}

	return info.Size()
}

// ReadFile reads a file.
func ReadFile(path string) ([]byte, error) {
	var err error

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = f.Close()
	}()
	if err != nil {
		return nil, err
	}

	data := make([]byte, FileSize(path))
	if _, err := f.Read(data); err != nil {
		return nil, err
	}

	return data, nil
}

// WriteFile writes to a file.
func WriteFile(path string, data []byte) error {
	var err error

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()
	if err != nil {
		return err
	}

	if _, err := f.Write(data); err != nil {
		return err
	}

	if err := f.Sync(); err != nil {
		return err
	}

	return nil
}
