package main

import "path/filepath"

// rename base without extension
func ReBaseName (path string, change string, opt ReBaseNameOption) string {
	dir, file := filepath.Split(path)
	if !opt.IsDirectory {
		ext := filepath.Ext(file)
		file = change + ext
	} else {
		file = change
	}
	return filepath.Join(dir, file)
}

type ReBaseNameOption struct {
	IsDirectory bool // if true overwrite with ext
}

// split file name to name and extension
func SplitName(file string) (name string, ext string){
	ext = filepath.Ext(file)
	if ext == "" {
		return file, ""
	}
	return file[:len(file) - len(ext)], ext
}