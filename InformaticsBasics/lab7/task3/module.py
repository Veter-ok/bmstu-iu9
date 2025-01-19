#!/usr/bin/env python3
import random, string

def generate_strings(a, b):
    letters = string.ascii_lowercase
    for _ in range(b): 
        s = ''.join(random.choice(letters) for i in range(a))
        print(s)