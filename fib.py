import sys

number1 = int(sys.argv[1])
number2 = int(sys.argv[2])
count = int(sys.argv[3])
stop_count = int(sys.argv[4])

if count < stop_count:
    print(number2, number1 + number2, count + 1, stop_count)


