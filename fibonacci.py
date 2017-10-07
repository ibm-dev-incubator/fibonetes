import sys

def F(n):
    if n == 1:
        return 1
    elif n == 0:
        return 0
    else:
        return F(n - 2) + F(n - 1)


if __name__ == '__main__':
    print(F(int(sys.argv[1])))
