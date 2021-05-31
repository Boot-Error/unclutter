package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/superhawk610/bar"
)

// copy a file to bucket
func copyFileToBucket(filePath string, bucket *Bucket) error {
	sourceFile, err := os.Open(filePath)
	if err != nil {
		log.Printf("Failed to open file %s due to %s", filePath, err)
		return err
	}
	defer sourceFile.Close()

	destPath := path.Join(bucket.Path, filepath.Base(filePath))
	destFile, err := os.Create(destPath)
	if err != nil {
		log.Printf("Failed to create file at %s due to %s", destPath, err)
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		log.Printf("Failed to copy %s to bucket %s", filePath, bucket.Name)
		return err
	}

	return nil
}

// find files belonging to a bucket
func SortFilesIntoBuckets(buckets *Buckets, dir string, doClean bool) {
	var wg sync.WaitGroup
	for _, bucket := range *buckets {
		// log.Println("Sorting bucket", bucket.Name)
		currentBucket := bucket
		wg.Add(1)
		go func(wg *sync.WaitGroup, bucket *Bucket, dir string) {
			defer wg.Done()
			filesToCopy := bucket.FindAllFiles(dir)
			filesToCopyCount := len(filesToCopy)
			// log.Printf("Found %d files to be copied to %s", filesToCopyCount, bucket.Name)
			if filesToCopyCount == 0 {
				return
			}

			if !bucket.Exists() {
				err := bucket.CreateDir()
				if err != nil {
					log.Fatalf("Failed to create %s due to %s", bucket.Name, err)
				}
			}

			copyBar := bar.NewWithOpts(
				bar.WithDimensions(filesToCopyCount, 30),
				bar.WithFormat(fmt.Sprintf("Copying to [:bucket] :bar (:percent) of %d files", filesToCopyCount)),
			)

			for _, filePath := range filesToCopy {
				err := copyFileToBucket(filePath, bucket)
				if err == nil && doClean {
					os.Remove(filePath)
				}
				copyBar.TickAndUpdate(bar.Context{
					bar.Ctx("bucket", bucket.Name),
				})
			}
			copyBar.Done()
		}(&wg, &currentBucket, dir)
	}
	wg.Wait()
}
