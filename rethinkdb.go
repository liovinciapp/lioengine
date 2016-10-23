package lioengine

import (
	r "gopkg.in/dancannon/gorethink.v2"
)

var rethink *rethinkdb

// rethinkdb contains info of the rethinkdb server.
type rethinkdb struct {
	session  *r.Session
	hosts    []string
	database string
	table    string
}

// configRethinkdb configures the rethinkdb session.
func configRethinkdb(hosts []string, database string, table string) (rethink *rethinkdb, err error) {
	rethink = new(rethinkdb) // sets non nil value
	rethink.session, err = r.Connect(r.ConnectOpts{
		Addresses: hosts,
		Database:  database,
	}) // We don't check the err because we'll return anyways
	rethink.table = table
	return
}
