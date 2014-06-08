import sys

rows, cols, count = map(int,sys.stdin.readline().split())
population = ''.join([a.replace(' ','').replace('\n','') for a in sys.stdin.readlines()])
from collections import defaultdict
f = defaultdict(lambda  : 0)
hoods = [''.join([population[k] for k in [index, index + 1, index + cols, index + cols + 1]]) for index in range((rows-1) * (cols)) if not (index+1) % cols == 0]

for hood in hoods :
    if hood in ['o##o','#oo#']:
        f['2d'] += 1
    else:
        c = hood.split("#")
        if len(c) == 3:
            f['2a'] += 1
        else :
            f[str(len(c) - 1)] += 1

j = list(f.items())
j.sort()

for a in j:
    if a[0] == '2a':
        a = ('2s', a[1])
    print ("%s: %s"% a)

