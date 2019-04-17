/*Before you execute the program, Launch `cqlsh` and execute:
create keyspace example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
create table example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));
create index on example.tweet(timeline);
*/
package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func main() {
	// connect to the cluster
	// cluster := gocql.NewCluster("192.168.1.1", "192.168.1.2", "192.168.1.3")
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	// insert a tweet
	if err := session.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
		"we", gocql.TimeUUID(), "ARDIC").Exec(); err != nil {
		log.Fatal(err)
	}

	var tline string
	var id gocql.UUID
	var text string

	/* Search for a specific set of records whose 'timeline' column matches
	 * the value 'me'. The secondary index that we created earlier will be
	 * used for optimizing the search */
	if err := session.Query(`SELECT timeline, id, text FROM tweet WHERE timeline = ? LIMIT 1`,
		"me").Consistency(gocql.One).Scan(&tline, &id, &text); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tweet:", tline, id, text)

	// list all tweets
	//	iter := session.Query(`SELECT timeline, id, text FROM tweet WHERE timeline = ?`, "me").Iter()
	iter := session.Query(`SELECT timeline, id, text FROM tweet`).Iter()
	for iter.Scan(&tline, &id, &text) {
		fmt.Println("Tweet:", tline, id, text)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}
