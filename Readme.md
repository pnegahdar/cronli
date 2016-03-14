## cronli: a cli cron tool

Lets you run crons with second granularity as a standard cli proc

Very light wrapper around: https://github.com/robfig/cron

Installation:

Grab the right precompiled bin from github releases and put it in your path. Don't forget to `chmod +x` the bin.

For example:

    curl -SL https://github.com/grammarly/rocker-compose/releases/download/0.1.3/rocker-compose-0.1.3_darwin_amd64.tar.gz | tar -xzC /usr/local/bin && chmod +x /usr/local/bin/rocker-compose

Usage:

    cronli "* * * * *" "echo hello"

Arg 1: Schedule see: [schedule format](https://godoc.org/github.com/robfig/cron#hdr-CRON_Expression_Format)

Arg 2: The command you would put into /bin/bash -c "THIS COMMAND"