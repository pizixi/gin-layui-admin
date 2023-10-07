#!/bin/bash

#go build -ldflags "-w -s"
#go build -v -x

zip -vr gin-layui-admin.zip gin-layui-admin conf.yaml.tpl views/ static/ *.sql -x "conf.yaml"
