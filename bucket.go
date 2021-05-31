package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

// A Bucket represents the Destination folder for files
// to be copied
type Bucket struct {
	Name     string   // name of the bucket
	Path     string   // full path to the bucket
	Patterns []string //
}

// An Array of Buckets
type Buckets []Bucket

// check if bucket exists
func (bucket *Bucket) Exists() bool {
	_, err := os.Stat(bucket.Path)
	return !os.IsNotExist(err)
}

// create bucket directory
func (bucket *Bucket) CreateDir() error {
	err := os.MkdirAll(bucket.Path, os.ModeDir)
	log.Printf("Creating bucket directory for %s", bucket.Name)
	return err
}

// load buckets from the config
func LoadBucketsFromConfig(bucketConfig map[string][]string, dir string) Buckets {
	buckets := Buckets{}
	for bucketName, bucketPatterns := range bucketConfig {
		bucket := Bucket{
			Name:     bucketName,
			Path:     path.Join(dir, bucketName),
			Patterns: bucketPatterns,
		}
		buckets = append(buckets, bucket)
	}
	return buckets
}

// Get all files belonging to a bucket
func (bucket *Bucket) FindAllFiles(searchDir string) []string {
	unsortedFiles := []string{}
	for _, pattern := range bucket.Patterns {
		matches, _ := filepath.Glob(fmt.Sprintf("%s/%s", searchDir, pattern))
		unsortedFiles = append(unsortedFiles, matches...)
	}
	return unsortedFiles
}
