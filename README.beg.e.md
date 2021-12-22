# {{.Name}}

{{render "license/shields" . "License" "MIT"}}
{{template "badge/godoc" .}}
{{template "badge/goreport" .}}
[![Build Status](https://github.com/{{.User}}/{{.Name}}/actions/workflows/go-release-build.yml/badge.svg?branch=master)](https://github.com/{{.User}}/{{.Name}}/actions/workflows/go-release-build.yml)
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-B.svg)](http://godoc.org/github.com/go-easygen/wireframe)

{{pkgdoc}}

