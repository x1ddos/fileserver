Simplest file server to (temporarily) serve static content from a local filesystem.

go get github.com/crhym3/fileserver
fileserver -d /tmp -a 0.0.0.0 -p 8001

Usage of fileserver:
  -a="localhost": Address to listen on
  -d=".": Root directory
  -p=8000: Port to listen at
 
