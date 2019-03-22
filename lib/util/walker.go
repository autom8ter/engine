package util

import "path/filepath"

type Walker struct {
	Path     string
	WalkFunc filepath.WalkFunc
}

func NewWalker(path string, walkFunc filepath.WalkFunc) *Walker {
	return &Walker{Path: path, WalkFunc: walkFunc}
}

func (w *Walker) Walk() error {
	if w.Path == "" {
		w.Path = "."
	}
	return filepath.Walk(w.Path, w.WalkFunc)
}
