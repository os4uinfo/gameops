# gameops
## Usage
Install go-bindata
> apt-get install go-bindata

cd front
go-bindata -pkg=front -nocompress=true -debug=true html/...
go-bindata -pkg=front -nocompress=true html/...
go install gameops


