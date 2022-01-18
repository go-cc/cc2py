 pwd=`pwd`
 cd $GOPATH/src/github.com/go-easygen/easygen/test
 easygen commandlineGoFlags.header,commandlineGoFlags.ityped.tmpl,commandlineGoFlags "$pwd/cc2py"_cli | gofmt > "$pwd/cc2py"_cliDef.go
