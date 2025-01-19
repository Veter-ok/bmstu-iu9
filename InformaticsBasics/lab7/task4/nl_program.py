#!/usr/bin/env python3
import sys

def nl():
    for file in sys.argv[1:]:
        with open(file) as filee:
            index = 1
            for line in filee:
                if len(line.strip()) != 0:
                    print(f"{index} {line}", end='')
                    index += 1
                else:
                    print()
        print("\n----------------------------------")
    if len(sys.argv) == 1:
        program = sys.stdin.read().split('\n')
        for line in program:
            if len(line.strip()) != 0:
                print(f"{index} {line}", end='')
                index += 1
            else:
                print()
nl()