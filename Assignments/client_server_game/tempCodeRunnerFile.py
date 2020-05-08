import random
import sys
import math
from datetime import datetime

def guessCheck(guess, num):
    if guess > num:
        return 1
    elif guess < num:
        return -1
    else:
        return 0

random.seed(datetime.now())

low = False
high = False

MAX = (2**54)-1 


L = random.randint(1, MAX-1)
H = random.randint(L + 1, MAX)
num = random.randint(L, H)
input(num)

# Always guess max in
guess = MAX
high = sys.maxsize -1
low = 0

response = guessCheck(guess, num)

count = 0
# while high == False or low == False:
#     if response == 1:
#         high = True
#         lastHigh = guess
#         guess = math.floor(guess / 2)
#     elif response == -1:
#         low = True
#         guess *= 2
    
#     count += 1
#     response = guessCheck(guess, num)



# with open("test.txt","w+") as f:
#     f.write(f"number to find: {num}\n")
#     while response != 0:
#         if response == 1:
#             lastHigh = guess
#             guess = math.floor(guess/2)
#             print(guess)

#         elif response == -1:
#             guess = math.floor((guess + lastHigh)/2)
#             print(guess)
#         count+=1
#         diff = num - guess
#         f.write(f"{diff}   guess {guess}     last high: {lastHigh}")
    

#         f.write("\n")
#         response = guessCheck(guess, num)
#     print(count)

# print(f"found the number ({num}) in {count} guesses. Last guess:{guess}")


with open("test.txt","w+") as f:

    while low != 0:
        guess = math.floor((high + low) /2)
        if response == 1:
            high = guess -1
        elif response == -1:

            low = guess + 1

        count+=1
        diff = num - guess
        print(f"{diff}   guess {guess}     last high: {high}")
    
        f.write("\n")
        
        response = guessCheck(guess, num)
    print(count)

print(f"found the number!{guess}")