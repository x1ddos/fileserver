Simplest file server to (temporarily) serve static content from a local filesystem.

Usage:

    fileserver [-a host:port] [pattern:]dir ...

The program serves files from a local dist using dir argument as root.
An optional pattern specified the URL base path. For instance,
to serve /tmp dir at /temporary, one can use the following:

    fileserver /temporary:/tmp

The dir argument can be specified multiple times.
If no dir argument is provided, current directory is used.

  -a string
    	Address (host:port) to listen on (default "localhost:8000")
