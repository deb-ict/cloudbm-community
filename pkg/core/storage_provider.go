package core

import (
	"context"
	"os/user"
	"path/filepath"
	"strings"
)

type StorageProvider interface {
	GetRootFolder(ctx context.Context) string
	GetPath(ctx context.Context, elem ...string) string
}

type defaultStorageProvider struct {
	rootFolder string
}

func DefaultStorageProvider(rootFolder string) StorageProvider {
	if strings.HasPrefix(rootFolder, "~") {
		usr, _ := user.Current()
		rootFolder = filepath.Join(usr.HomeDir, rootFolder[2:])
	}
	return &defaultStorageProvider{
		rootFolder: rootFolder,
	}
}

func (p *defaultStorageProvider) GetRootFolder(ctx context.Context) string {
	return strings.ToLower(p.rootFolder)
}

func (p *defaultStorageProvider) GetPath(ctx context.Context, elem ...string) string {
	path := []string{p.GetRootFolder(ctx)}
	path = append(path, elem...)
	return filepath.Join(path...)
}
