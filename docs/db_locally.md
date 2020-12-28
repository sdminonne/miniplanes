#Db locally


#MongoDB

###Two instances of mongo on same machine

#### As a Docker(s)


#Neo4J

##As a Docker
```shell
docker run \
    --publish=7474:7474 --publish=7687:7687 \
    --volume=/tmp/neo4j/data:/data \
    --env=NEO4J_AUTH=none \
    neo4j
```

```shell
$ cypher-shell
```


in shell one may want to do something similar:

```
neo4j> CREATE (:Person {name : 'Dick Grayson', alias : 'Robin' });
neo4j> CREATE (:Person {name : 'Dick Grayson', alias : 'Batman' });
neo4j> CREATE (:Person {name : 'Bruce Wayne', alias : 'Batman' });
```

now
```
neo4j> MATCH(n) RETURN n;
+---------------------------------------------------+
| n                                                 |
+---------------------------------------------------+
| (:Person {name: "Dick Grayson", alias: "Robin"})  |
| (:Person {name: "Dick Grayson", alias: "Batman"}) |
| (:Person {name: "Bruce Wayne", alias: "Batman"})  |
+---------------------------------------------------+
```

```
eo4j> MATCH(n { name: 'Dick Grayson', alias : 'Batman'}) delete n;
0 rows available after 55 ms, consumed after another 0 ms
Deleted 1 nodes
```
```
neo4j> match (n) RETURN n;
+--------------------------------------------------+
| n                                                |
+--------------------------------------------------+
| (:Person {name: "Dick Grayson", alias: "Robin"}) |
| (:Person {name: "Bruce Wayne", alias: "Batman"}) |
+--------------------------------------------------+
```



go get github.com/neo4j/neo4j-go-driver/neo4j

In golang
```go
func helloWorld(uri, username, password string) (string, error) {
    var (
        err      error
        driver   neo4j.Driver
        session  neo4j.Session
        result   neo4j.Result
        greeting interface{}
    )

    driver, err = neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
    if err != nil {
        return "", err
    }
    defer driver.Close()

    session, err = driver.Session(neo4j.AccessModeWrite)
    if err != nil {
        return "", err
    }
    defer session.Close()

    greeting, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
        result, err = transaction.Run(
            "CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
            map[string]interface{}{"message": "hello, world"})
        if err != nil {
            return nil, err
        }

        if result.Next() {
            return result.Record().GetByIndex(0), nil
        }

        return nil, result.Err()
    })
    if err != nil {
        return "", err
    }

    return greeting.(string), nil
}
```
