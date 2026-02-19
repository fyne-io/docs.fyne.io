package main

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"codeberg.org/sdassow/atomic"
)

func main() {
	root := "../content/docs"

	files := []string{}

	pathFn := func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() && path == filepath.Join(root, "api") {
			return filepath.SkipDir
		}
		file := filepath.Base(path)
		if filepath.Ext(file) != ".md" {
			return nil
		}
		files = append(files, path)
		return nil
	}

	if err := filepath.Walk(root, pathFn); err != nil {
		log.Fatal(err)
	}

	for _, path := range files {
		if err := handleMarkdown(root, path); err != nil {
			log.Fatal(err)
		}
	}
}

func handleMarkdown(root, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	f.Close()

	// scan for things to turn into api links
	re, err := regexp.Compile("`(\\w+)\\.(\\w+)(\\(\\))?`")
	if err != nil {
		return err
	}

	c := re.ReplaceAllFunc(b, func(b []byte) []byte {
		m := re.FindStringSubmatch(string(b))
		dir := findDir(filepath.Join(root, "api"), m[1])

		file, t, target := findObj(dir, m[2])
		if file == "" || target == "" {
			return b
		}

		return []byte(fmt.Sprintf("[%s.%s%s](%s/#%s--%s)",
			m[1], m[2], m[3],
			file[len(root):], t, target,
		))
	})

	if bytes.Equal(b, c) {
		return nil
	}

	return atomic.WriteFile(path, bytes.NewReader(c))
}

func findDir(root, dir string) string {
	var r string

	if err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		base := filepath.Base(path)
		if !info.IsDir() {
			return nil
		}
		if base == dir {
			r = path
		}
		return nil
	}); err != nil {
		return ""
	}

	return r
}

func findObj(dir, funcName string) (string, string, string) {
	files, err := filepath.Glob(filepath.Join(dir, "*.md"))
	if err != nil {
		return "", "", ""
	}

	re := regexp.MustCompile(`####\s+(func|type)\s+` + funcName)

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return "", "", ""
		}
		defer f.Close()

		b, err := io.ReadAll(f)
		if err != nil {
			return "", "", ""
		}

		m := re.FindStringSubmatch(string(b))

		if len(m) > 0 {
			return file[:len(file)-3], m[1], strings.ToLower(funcName)
		}
	}

	return "", "", ""
}
