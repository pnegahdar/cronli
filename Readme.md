## cronli: a cli cron tool

Lets you run crons with second granularity as a standard cli proc

Very light wrapper around: https://github.com/robfig/cron

#### Installation:

Grab the right precompiled bin from github releases and put it in your path. Don't forget to `chmod +x` the bin.

OSX:

    curl -SL https://github.com/pnegahdar/cronli/releases/download/0.1.0/cronli_0.1.0_darwin_amd64.tar.gz \
        | tar -xzC /usr/local/bin --strip 1 && chmod +x /usr/local/bin/cronli

Nix:

    curl -SL https://github.com/pnegahdar/cronli/releases/download/0.1.0/cronli_0.1.0_linux_amd64.tar.gz \
        | tar -xzC /usr/local/bin --strip 1 && chmod +x /usr/local/bin/cronli

#### Usage:

    cronli "* * * * *" "echo hello"

Arg 1: Schedule see: [schedule format](https://godoc.org/github.com/robfig/cron#hdr-CRON_Expression_Format)

Arg 2: The command you would put into /bin/bash -c "THIS COMMAND"
