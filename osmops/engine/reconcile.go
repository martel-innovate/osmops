package engine

import (
	"context"

	"github.com/go-logr/logr"

	"github.com/martel-innovate/osmops/osmops/cfg"
	"github.com/martel-innovate/osmops/osmops/nbic"
	u "github.com/martel-innovate/osmops/osmops/util"
	"github.com/martel-innovate/osmops/osmops/util/file"
)

type Engine struct {
	ctx       context.Context
	opsConfig *cfg.Store
	nbic      nbic.Workflow
}

func newNbic(opsConfig *cfg.OsmConnection) (nbic.Workflow, error) {
	hp, err := u.ParseHostAndPort(opsConfig.Hostname)
	if err != nil {
		return nil, err
	}

	conn := nbic.Connection{
		Address: *hp,
		Secure:  false,
	}
	usrCreds := nbic.UserCredentials{
		Username: opsConfig.User,
		Password: opsConfig.Password,
		Project:  opsConfig.Project,
	}

	return nbic.New(conn, usrCreds)
}

func newProcessor(ctx context.Context, repoRootDir string) (*Engine, error) {
	rootDir, err := file.ParseAbsPath(repoRootDir)
	if err != nil {
		return nil, err
	}

	store, err := cfg.NewStore(rootDir)
	if err != nil {
		return nil, err
	}

	client, err := newNbic(store.OsmConnection())
	return &Engine{
		ctx:       ctx,
		opsConfig: store,
		nbic:      client,
	}, err
}

func log(ctx context.Context) logr.Logger {
	return logr.FromContext(ctx)
}

func (p *Engine) log() logr.Logger {
	return log(p.ctx)
}

func (p *Engine) repoScanner() *cfg.KduNsActionRepoScanner {
	return cfg.NewKduNsActionRepoScanner(p.opsConfig)
}

const (
	processingMsg    = "processing"
	packageLogKey    = "osm package"
	fileLogKey       = "file"
	engineInitErrMsg = "can't initialize reconcile engine"
	processingErrMsg = "processing errors"
	errorLogKey      = "error"
)

func (p *Engine) processPackages() []error {
	es := []error{}
	pkgs, err := p.opsConfig.RepoPkgDirectories()
	if err != nil {
		es = append(es, err)
		return es
	}
	for _, pkgPath := range pkgs {
		p.log().Info(processingMsg, packageLogKey, pkgPath.Value())

		err = p.nbic.CreateOrUpdatePackage(pkgPath)
		if err != nil {
			es = append(es, err)
		}
	}
	return es
}

func (p *Engine) Process(file *cfg.KduNsActionFile) error {
	p.log().Info(processingMsg, fileLogKey, file.FilePath.Value())

	data := nbic.NsInstanceContent{
		Name:           file.Content.Name,
		Description:    file.Content.Description,
		NsdName:        file.Content.NsdName,
		VnfName:        file.Content.VnfName,
		VimAccountName: file.Content.VimAccountName,
		KduName:        file.Content.Kdu.Name,
		KduParams:      file.Content.Kdu.Params,
	}
	return p.nbic.CreateOrUpdateNsInstance(&data)
}

// New instantiates an Engine to reconcile the state of the OSM deployment
// with that declared in the OSM GitOps files found in the specified repo.
func New(ctx context.Context, repoRootDir string) (*Engine, error) {
	engine, err := newProcessor(ctx, repoRootDir)
	if err != nil {
		log(ctx).Error(err, engineInitErrMsg)
		return nil, err
	}
	return engine, nil
}

// Reconcile looks for OSM GitOps files in the repo and, for each file
// found, it calls OSM NBI to reach the deployment state declared in the
// file.
//
// Additionally, if there's an OSM package root directory (see: Store),
// Reconcile creates or updates any OSM packages found in there. Reconcile
// blindly assumes that any sub-directory p of the OSM package root directory
// contains the source files of an OSM package. It reads p's contents to
// create a gzipped tar archive in the OSM format (including creating the
// "checksums.txt" file) and then streams it to OSM NBI to create or update
// the package in OSM. (See: nbic.CreateOrUpdatePackage)
//
// Notice at the moment OsmOps does **not explicitly** handle dependencies
// among OSM packages. But it does process sub-directories of the OSM package
// root directory in alphabetical order. This way, the operator can name
// package directories in such a way that if package p2 depends on p1,
// p2's name comes before p1's in alphabetical order. For example, say
// you want to deploy a KNF using two packages: one, p1, contains the actual
// KNF definition whereas the other, p2, contains an NS definition referencing
// p1. Then you could use the following naming scheme:
//
//     my-gitops-repo
//     |
//     + -- deployment-target-dir
//          + -- osm-pkgs
//               + -- my-service_knf           (<- p1's contents)
//                  | - my-service_vnfd.yaml
//               + -- my-service_ns            (<- p2's contents)
//                  | - my-service_nsd.yaml
//                  | - README.md
//
// Because my-service_knf < my-service_ns (alphabetical order), Reconcile
// will first process my-service_knf and then my-service_ns.
//
// Surely this is a stopgap solution. Eventually we'll implement proper
// handling of package dependencies. (Solution: parse OSM package definitions,
// build dependency graph, extract DAG d[k] for each graph component g[k],
// topologically sort d[k] ~~> s[k]; process s[k] sequences in parallel.)
func (p *Engine) Reconcile() {
	errors := p.processPackages()
	if len(errors) == 0 {
		errors = p.repoScanner().Visit(p)
	}
	// else stop there since KDU ops might fail b/c referenced packages
	// weren't created or updated.

	if len(errors) > 0 {
		for k, e := range errors {
			p.log().Error(e, processingErrMsg, errorLogKey, k)
		}
	}
}
