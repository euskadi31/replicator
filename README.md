Replicator
==========

Configuration
-------------

A discovery service, https://discovery.etcd.io, is provided as a free service to
help connect Replicator instances together by storing a list of peer addresses,
metadata and the initial size of the cluster under a unique address, known as
the discovery URL. You can generate them very easily:

~~~shell
$ curl -w "\n" 'https://discovery.etcd.io/new?size=3'
https://discovery.etcd.io/6a28e078895c5ec737174db2419bb2f3
~~~

/etc/replicator.json
~~~json
{
    "discovery": "https://discovery.etcd.io/<token>",
    "commands": {
        "my_project": {
            "cmd": "rsync",
            "option": "-alWz",
            "path": "/path/to/project/",
            "user": "myuser",
            "excludes": [
                "app/cache/*",
                "app/logs/*",
                ".git",
                "node_modules",
                "bower_components",
                "_assets"
            ],
            "delete": true
        }
    }
}
~~~

Usage
-----

Join the cluster
~~~shell
$ replicator join
~~~

Leave the cluster
~~~shell
$ replicator leave
~~~

List members of the cluster
~~~shell
$ replicator list
~~~

~~~
 * my-server-web-01 : 10.0.0.4
 * my-server-web-02 : 10.0.0.5
~~~

Sync projects on the cluster
~~~shell
$ replicator sync
~~~

~~~
 * Sync my_project on my-server-web-01      [ ok ]
 * Sync my_project on my-server-web-02      [ ok ]
~~~
