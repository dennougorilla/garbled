import random
import string

def random_str(n):
   return ''.join([random.choice(string.ascii_letters + string.digits) for i in range(n)])
while True:
    l = input()
    for s in l:
        if 0.25 > random.random():
            print(int(random.randrange(0,65535)).to_bytes(2,byteorder='little').decode('shift-jis','replace'), end='')
        elif 0.25 > random.random():
            print(random_str(1), end='')
        else:
            print(s, end='')
    print()

