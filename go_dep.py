import os
from subprocess import call
import time

def replace_all(path='.', old="github.com/ethereum", new="github.com/project-douglas"):
    in_this_dir = os.listdir(path)
    for f in in_this_dir:
        f = os.path.join(path, f)
        if os.path.isdir(f):
            replace_all(path=f, old=old, new=new)
        elif f[-3:] == ".go":
            d = open(f)
            src = d.readlines()
            d.close()
            d = open(f, "w")
            for s in src:
                s = s.replace(old, new)
                d.write(s)
            d.close()


replace_all(old="github.com/ethereum/eth-go", new="github.com/eris-ltd/eth-go-mods")

