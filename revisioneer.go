package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/user"
	"syscall"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
	"github.com/splicers/jet"
)

func setup() *jet.Db {
	var revDsn = os.Getenv("REV_DSN")
	if revDsn == "" {
		user, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		revDsn = "user=" + user.Username + " dbname=revisioneer sslmode=disable"
	}

	db, err := jet.Open("postgres", revDsn)
	if err != nil {
		log.Fatal("failed to connect to postgres", err)
	}
	db.SetMaxIdleConns(100)

	return db
}

var _db *jet.Db

func runMigrations(db *sql.DB) {
	migrations := &migrate.AssetMigrationSource{
		Asset:    Asset,
		AssetDir: AssetDir,
		Dir:      "migrations",
	}

	if n, err := migrate.Exec(db, "postgres", migrations, migrate.Up); err != nil {
		log.Printf("unable to migrate: %v", err)
	} else {
		log.Printf("Applied %d migrations!\n", n)
	}
}

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	log.SetPrefix(fmt.Sprintf("pid:%d ", syscall.Getpid()))

	_db = setup()
	runMigrations(_db.DB)
}

func main() {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Listen on a TCP or a UNIX domain socket (TCP here).
	l, err := net.Listen("tcp", "0.0.0.0:"+port)
	if nil != err {
		log.Fatalln(err)
	}
	log.Printf("listening on %v", l.Addr())

	writePid()

	defer _db.Close()

	deployments := NewDeploymentsController(_db)
	projects := NewProjectsController(_db)

	r := mux.NewRouter()
	r.HandleFunc("/deployments", deployments.WithValidProject(deployments.ListDeployments)).
		Methods("GET")
	r.HandleFunc("/deployments", deployments.WithValidProject(deployments.CreateDeployment)).
		Methods("POST")
	r.HandleFunc("/deployments/{sha}/verify", deployments.WithValidProjectAndParams(deployments.VerifyDeployment)).
		Methods("POST")
	r.HandleFunc("/projects", projects.CreateProject).
		Methods("POST")
	http.Handle("/", r)

	http.Serve(l, r)
}

func writePid() {
	var file, error = os.OpenFile("tmp/rev.pid", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if error == nil {
		var line = fmt.Sprintf("%v", os.Getpid())
		file.WriteString(line)
		file.Close()
	}
}
