#!/bin/bash

# usage
usage() {
cat << EOF
Usage: ${0##*/} [-hv] [-a ADDRESS] [-k K_STUFF] [-n N_STUFF]
                    [-o O_STUFF] [-s S_STUFF] [-S SS_STUFF]
                    [-u U_STUFF] [-x X_STUFF]
This script would normall do some cool stuff, but really it just outputs its
usage. This file exists merely to show how to do usage in code.

    -h              displays the usage and exits
    -a ADDRESS      an ipv4 address, for some use
    -k K_STUFF      do some stuff, probably with a value starting with K, yada
                      yada yada yada yada
    -n N_STUFF      do some stuff, probably with a value starting with N
    -o O_STUFF      do some stuff, probably with a value starting with O
    -s S_STUFF      do some stuff, probably with a value starting with S
    -S SS_STUFF     do some stuff, probably with a value starting with Capital S
    -u U_STUFF      do some stuff, probably with a value starting with U
    -v              noisy
    -x X_STUFF      do some stuff, probably with a value starting with X
EOF
}

# IPv4 address validation
IPV4="^((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9][0-9]|[0-9])(\.|$)){4}$"
validIP() {
	if [[ "$1" =~ $IPV4 ]]; then
		return 0
	else
		return 1
	fi
}

# reporting stuff
report() {
    if [ "$1" -eq "1" ]; then
        echo $2 >&2
        usage >&2
    else
        if [ "$VERBOSE" != "" ]; then
            echo $2
        fi
    fi
}

# parse arguments
while getopts "hva:k:n:o:s:S:u:x:" opt; do
    case "$opt" in
        h)
            usage
            exit 0
            ;;
        a)
            ADDRESS=$OPTARG
            ;;
        k)
            K_STUFF=$OPTARG
            ;;
        n)
            N_STUFF=$OPTARG
            ;;
        o)
            O_STUFF=$OPTARG
            ;;
        s)
            S_STUFF=$OPTARG
            ;;
        S)
            SS_STUFF=$OPTARG
            ;;
        u)
            U_STUFF=$OPTARG
            ;;
        v)
            VERBOSE=" -v"
            ;;
        x)
            X_STUFF=$OPTARG
            ;;
        '?')
            usage >&2
            exit 100
            ;;
    esac
done

# check two required variables and validate all the IP addresses
report 0 "Checking Arguments"

if [[ -z "$ADDRESS" ]]; then
    report 1 "Address not provided!"
    exit 101
fi

# from here on exit on any error
set -e
