package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	// CLI flags
	dir := flag.String("path", "./files", "Path to the directory to organize")
	mode := flag.String("mode", "extension", "Organize mode: extension | date")
	dryRun := flag.Bool("dry-run", false, "Preview changes without moving files")
	flag.Parse()

	files, err := os.ReadDir(*dir)
	if err != nil {
		panic(err)
	}

	//collect file info
	var infos []os.FileInfo
	for _, file := range files {
		if !file.IsDir() {
			fullPath := filepath.Join(*dir, file.Name())
			info, err := os.Stat(fullPath) // <- get full info
			//info, err := file.Info()
			if err != nil {
				continue
			}
			infos = append(infos, info)
		}
	}

	//sort based on mode
	switch *mode {
	case "extension":
		sort.Slice(infos, func(i, j int) bool {
			return strings.ToLower(filepath.Ext(infos[i].Name())) < strings.ToLower(filepath.Ext(infos[j].Name()))
		})
		organizeByExtension(*dir, infos, *dryRun)
	case "date":
		sort.Slice(infos, func(i, j int) bool {
			return infos[i].ModTime().Before(infos[j].ModTime())
		})
		organizeByDate(*dir, infos, *dryRun)
	default:
		fmt.Println("unknown mode:", mode)
	}
}

func organizeByExtension(base string, files []os.FileInfo, dryRun bool) {
	for _, file := range files {
		ext := strings.TrimPrefix(filepath.Ext(file.Name()), ".")
		if ext == "" {
			ext = "unknown"
		}
		destDir := filepath.Join(base, ext)
		src := filepath.Join(base, file.Name())
		dst := filepath.Join(destDir, file.Name())
		if dryRun {
			fmt.Printf("[Dry Run] Would move %s → %s\n", src, dst)
			continue
		}

		os.MkdirAll(destDir, os.ModePerm)
		err := os.Rename(src, dst)
		if err != nil {
			fmt.Println("error moving", file.Name(), err)
		}
	}
}

func organizeByDate(base string, files []os.FileInfo, dryRun bool) {
	for _, file := range files {
		modTime := file.ModTime()
		//group by year/month
		destDir := filepath.Join(
			base,
			fmt.Sprintf("%d", modTime.Year()),
			modTime.Month().String(),
		)
		src := filepath.Join(base, file.Name())
		dst := filepath.Join(destDir, file.Name())
		if dryRun {
			fmt.Printf("[Dry Run] Would move %s → %s\n", src, dst)
			continue
		}

		os.MkdirAll(destDir, os.ModePerm)
		err := os.Rename(src, dst)
		if err != nil {
			fmt.Println("error moving", file.Name(), err)
		}
	}
}
