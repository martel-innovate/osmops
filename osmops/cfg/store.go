// Access to the OSM Ops program configuration.
//
package cfg

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	u "github.com/martel-innovate/osmops/osmops/util"
	"github.com/martel-innovate/osmops/osmops/util/file"
)

// Store holds the OSM Ops program configuration read from the OSM Ops
// config and credentials files.
type Store struct {
	rootDir   file.AbsPath
	targetDir file.AbsPath
	fileExt   []u.NonEmptyStr
	osmCreds  *OsmConnection
}

// NewStore reads the program configuration and credentials files, validates
// their content, and packs the content in a Store. If an I/O error happens
// when reading the files or some of the YAML content isn't valid, NewStore
// returns a error. Each YAML type documents what a valid instance is.
func NewStore(repoRootDir file.AbsPath) (*Store, error) {
	var err error
	var cfg *OpsConfig
	s := Store{rootDir: repoRootDir}

	if err = s.rootDir.IsDir(); err != nil {
		return nil, err
	}

	if cfg, err = readConfig(s.rootDir); err != nil {
		return nil, err
	}

	s.fileExt = getFileExtensions(cfg)

	if s.targetDir, err = buildTargetDirPath(s.rootDir, cfg); err != nil {
		return nil, err
	}
	if s.osmCreds, err = readCreds(s.rootDir, cfg); err != nil {
		return nil, err
	}

	return &s, nil
}

// The name of the YAML file containing the OpsConfig. For now this name is
// hardcoded to "osm_ops_config.yaml" and the file is expected to be in the
// repo root directory.
const OpsConfigFileName = "osm_ops_config.yaml"

// The name of the sub-directory of RepoTargetDirectory where to look for
// OSM source package directories. For now this is not configurable, if
// there are any OSM package source directories they should be in "t/osm-pkgs"
// where t is the absolute path returned by RepoTargetDirectory.
const OsmPackagesDirName = "osm-pkgs"

func readConfig(rootDir file.AbsPath) (*OpsConfig, error) {
	file := rootDir.Join(OpsConfigFileName)
	if fileData, err := ioutil.ReadFile(file.Value()); err != nil {
		return nil, err
	} else {
		return readOpsConfig(fileData)
	}
}

func buildTargetDirPath(rootDir file.AbsPath, cfg *OpsConfig) (file.AbsPath, error) {
	target := rootDir.Join(cfg.TargetDir)
	if err := target.IsDir(); err != nil {
		return target, err
	}
	return target, nil
}

func buildCredsDirPath(rootDir file.AbsPath, cfg *OpsConfig) (file.AbsPath, error) {
	if filepath.IsAbs(cfg.ConnectionFile) {
		return file.ParseAbsPath(cfg.ConnectionFile)
	}
	return rootDir.Join(cfg.ConnectionFile), nil
}

func readCreds(rootDir file.AbsPath, cfg *OpsConfig) (*OsmConnection, error) {
	var fileData []byte
	if credsFile, err := buildCredsDirPath(rootDir, cfg); err != nil {
		return nil, err
	} else {
		if fileData, err = ioutil.ReadFile(credsFile.Value()); err != nil {
			return nil, err
		}
		return readOsmConnection(fileData)
	}
}

// DefaultOpsFileExtensions returns the default file extensions used to filter
// OSM GitOps files: ".osmops.yaml" and ".osmops.yml".
func DefaultOpsFileExtensions() []u.NonEmptyStr {
	y1, _ := u.NewNonEmptyStr(".osmops.yaml")
	y2, _ := u.NewNonEmptyStr(".osmops.yml")
	return []u.NonEmptyStr{y1, y2}
}

func getFileExtensions(cfg *OpsConfig) []u.NonEmptyStr {
	nonEmpty := []u.NonEmptyStr{}
	for _, x := range cfg.FileExtensions {
		if y, err := u.NewNonEmptyStr(strings.TrimSpace(x)); err == nil {
			nonEmpty = append(nonEmpty, y)
		}
	}

	if len(nonEmpty) > 0 {
		return nonEmpty
	}
	return DefaultOpsFileExtensions()
}

// RepoRootDirectory returns the absolute path to the repo root directory.
func (s *Store) RepoRootDirectory() file.AbsPath {
	return s.rootDir
}

// RepoTargetDirectory returns the absolute path to the directory within
// the repo where to find OSM Git Ops files.
func (s *Store) RepoTargetDirectory() file.AbsPath {
	return s.targetDir
}

// RepoPkgDirectories lists, in alphabetical order, the sub-directories
// of the OSM package root directory. If there's no OSM package directory,
// RepoPkgDirectories returns an empty list.
// See also: OsmPackagesDirName.
func (s *Store) RepoPkgDirectories() ([]file.AbsPath, error) {
	dirs := []file.AbsPath{}

	pkgsDir := s.targetDir.Join(OsmPackagesDirName)
	if err := pkgsDir.IsDir(); err != nil {
		return dirs, nil
	}

	sortedDirNames, err := file.ListSubDirectoryNames(pkgsDir.Value())
	for _, name := range sortedDirNames { // (*)
		dirPath := pkgsDir.Join(name)
		dirs = append(dirs, dirPath)
	}
	return dirs, err

	// (*) sortedDirNames is empty if err but never nil.
}

// OpsFileExtensions returns the file extensions used to filter OSM GitOps
// files within the RepoTargetDirectory.
// If the OpsConfig YAML file contains no extensions field, then the extensions
// will be the DefaultOpsFileExtensions.
// OSM Ops will only look for OSM Git Ops files in the RepoTargetDirectory
// that have one of these extensions.
func (s *Store) OpsFileExtensions() []u.NonEmptyStr {
	return s.fileExt
}

// OsmCredentials returns the OSM connection and credentials details to
// connect to the OSM north-bound interface.
func (s *Store) OsmConnection() *OsmConnection {
	return s.osmCreds
}
