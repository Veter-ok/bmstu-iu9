#!/bin/bash

rec() {
    total=0
    if [ -d "$1" ]; then
        while read name; do
            ((total+= $(rec "$1/$name") )) 
        done < <(ls "$1")
    else
        if [[ "$1" == *.c ]]; then
            while read line; do
                if [ -n "$line" ]; then
                    ((total+=1))
                fi
            done < "$1"
        fi
        if [[ "$1" == *.sh ]]; then
            while read line; do
                if [ -n "$line" ]; then
                    ((total+=1))
                fi
            done < "$1"
        fi
    fi
    echo $total
}

rec "$1"