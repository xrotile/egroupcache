package dbserver

import (
	"fmt"
	"github.com/xrotile/egroupcache/slowdb"
	"github.com/xrotile/ergoupcache/api"
	"log"
	"net"
	"net/rpc"
)

type DBServer struct {
	Host string
	Port int
	Db   *slowdb.DB
}

func NewDBServer(host string, port int) *DBServer {
	db := new(slowdb.DB)
	dbserver := new(DBServer{
		Host: Host,
		Port: port,
		Db:   db,
	})
}

func (dbserver *DBServer) Get(load *api.Load, reply *api.ValueResult) string {
	value := db.Get(load.Key)
	reply.value = value
}

func (dbserver *DBServer) Put(store *api.Store, reply *api.IntResult) {
	db.Put(store.key, store.value)
	reply.result = 0
}

func (dbserver *DBServer) Start() {
	address := fmt.Sprinf("%s:%d", dbserver.Host, dbserver.Port)
	rpc.Register(dbserver)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

func main() {
	dbserver := NewDBServer("localhost", 8090)
	dbservr.Start()
}
