import random
import sys
import math
from datetime import datetime

# seed value
random.seed(datetime.now())

# Max value
MAX = (2**53) -1

# Always guess max so I can cut down 50% of the set of possibilities
# everytime it's higher
guess = MAX

# Create a range, and pick a random number from within it
L = random.randint(1, MAX-1) # Low end of range
H = random.randint(L + 1, MAX) # High end of range
num = random.randint(L, H) # The actual random number to look for

count = 0 # debugging

# Returns 1 if the guess is too high
# 0 If the guess is correct
# -1 If the guess is too low
def guessCheck(guess, num):
    if guess > num:
        return 1
    elif guess < num:
        return -1
    else:
        return 0
"""
$$\      $$\            $$\     $$\                       $$\         $$\   
$$$\    $$$ |           $$ |    $$ |                      $$ |      $$$$ |  
$$$$\  $$$$ | $$$$$$\ $$$$$$\   $$$$$$$\   $$$$$$\   $$$$$$$ |      \_$$ |  
$$\$$\$$ $$ |$$  __$$\\_$$  _|  $$  __$$\ $$  __$$\ $$  __$$ |        $$ |  
$$ \$$$  $$ |$$$$$$$$ | $$ |    $$ |  $$ |$$ /  $$ |$$ /  $$ |        $$ |  
$$ |\$  /$$ |$$   ____| $$ |$$\ $$ |  $$ |$$ |  $$ |$$ |  $$ |        $$ |  
$$ | \_/ $$ |\$$$$$$$\  \$$$$  |$$ |  $$ |\$$$$$$  |\$$$$$$$ |      $$$$$$\ 
\__|     \__| \_______|  \____/ \__|  \__| \______/  \_______|      \______|
"""
# Binary search
def search1():
    count = 0
    high = MAX -1
    low = 0
    while low <= high:
        guess = math.floor((low+high)/2)
        response = guessCheck(guess, num)
        if response == -1:
            low = guess + 1
            count+=1
        elif response == 1:
            count+=1
            high = guess - 1
        else:
            print("\n\nBinary Search")
            print(f"Found the number ({num}) in {count} guesses. Last guess:{guess}")
            print("====================")
            return guess
"""


$$\      $$\            $$\     $$\                       $$\        $$$$$$\  
$$$\    $$$ |           $$ |    $$ |                      $$ |      $$  __$$\ 
$$$$\  $$$$ | $$$$$$\ $$$$$$\   $$$$$$$\   $$$$$$\   $$$$$$$ |      \__/  $$ |
$$\$$\$$ $$ |$$  __$$\\_$$  _|  $$  __$$\ $$  __$$\ $$  __$$ |       $$$$$$  |
$$ \$$$  $$ |$$$$$$$$ | $$ |    $$ |  $$ |$$ /  $$ |$$ /  $$ |      $$  ____/ 
$$ |\$  /$$ |$$   ____| $$ |$$\ $$ |  $$ |$$ |  $$ |$$ |  $$ |      $$ |      
$$ | \_/ $$ |\$$$$$$$\  \$$$$  |$$ |  $$ |\$$$$$$  |\$$$$$$$ |      $$$$$$$$\ 
\__|     \__| \_______|  \____/ \__|  \__| \______/  \_______|      \________|
"""
#Linear Search 
def search2():
    print("\nLinear Searching. Fair warning, this will take forever.")
    count = 0
    guess = MAX
    response = guessCheck(guess,num)
    while response !=0:
        if response == 1:
            guess -= 1
            count += 1
        elif response == -1:
            guess +=1
            count += 1
        else:
            print("\nLinear Search")
            print(f"Found the number ({num}) in {count} guesses. Last guess:{guess}")
            print("====================")
            return guess


"""
$$\      $$\            $$\     $$\                       $$\        $$$$$$\  
$$$\    $$$ |           $$ |    $$ |                      $$ |      $$ ___$$\ 
$$$$\  $$$$ | $$$$$$\ $$$$$$\   $$$$$$$\   $$$$$$\   $$$$$$$ |      \_/   $$ |
$$\$$\$$ $$ |$$  __$$\\_$$  _|  $$  __$$\ $$  __$$\ $$  __$$ |        $$$$$ / 
$$ \$$$  $$ |$$$$$$$$ | $$ |    $$ |  $$ |$$ /  $$ |$$ /  $$ |        \___$$\ 
$$ |\$  /$$ |$$   ____| $$ |$$\ $$ |  $$ |$$ |  $$ |$$ |  $$ |      $$\   $$ |
$$ | \_/ $$ |\$$$$$$$\  \$$$$  |$$ |  $$ |\$$$$$$  |\$$$$$$$ |      \$$$$$$  |
\__|     \__| \_______|  \____/ \__|  \__| \______/  \_______|       \______/ 
"""

# Original Algorithm 
def search3():
    guess = MAX
    response = guessCheck(guess, num)
    count = 0 # debugging

    while response != 0:
        if response == 1:
            lastHigh = guess
            guess = math.floor(guess/2)

        elif response == -1:
            guess = math.floor((guess + lastHigh)/2)
        
        count+=1
        response = guessCheck(guess, num)

    # debugging
    print("\nOriginal Alg")
    print(f"Found the number ({num}) in {count} guesses. Last guess:{guess}")
    print("====================")




print(f"Looking for {num}")
#Linear Search takes forever so I run it last
search1()
search3()
search2()




"""
$$$$$$$\            $$\                                  $$$$$$\                  $$\           
$$  __$$\           $$ |                                $$  __$$\                 $$ |          
$$ |  $$ | $$$$$$\  $$$$$$$\  $$\   $$\  $$$$$$\        $$ /  \__| $$$$$$\   $$$$$$$ | $$$$$$\  
$$ |  $$ |$$  __$$\ $$  __$$\ $$ |  $$ |$$  __$$\       $$ |      $$  __$$\ $$  __$$ |$$  __$$\ 
$$ |  $$ |$$$$$$$$ |$$ |  $$ |$$ |  $$ |$$ /  $$ |      $$ |      $$ /  $$ |$$ /  $$ |$$$$$$$$ |
$$ |  $$ |$$   ____|$$ |  $$ |$$ |  $$ |$$ |  $$ |      $$ |  $$\ $$ |  $$ |$$ |  $$ |$$   ____|
$$$$$$$  |\$$$$$$$\ $$$$$$$  |\$$$$$$  |\$$$$$$$ |      \$$$$$$  |\$$$$$$  |\$$$$$$$ |\$$$$$$$\ 
\_______/  \_______|\_______/  \______/  \____$$ |       \______/  \______/  \_______| \_______|
                                        $$\   $$ |                                              
                                        \$$$$$$  |                                              
                                         \______/    
"""
# with open("test.txt","w+") as f:        #debugging
#     f.write(f"number to find: {num}\n")
#     while response != 0:
#         if response == 1:
#             lastHigh = guess
#             guess = math.floor(guess/2)
#             # print(guess)

#         elif response == -1:
#             guess = math.floor((guess + lastHigh)/2)
#             # print(guess)
#         count+=1

#         #debugging stuff
#         diff = num - guess
#         f.write(f"{diff}   guess {guess}     last high: {lastHigh}")
#         f.write("\n")

#         response = guessCheck(guess, num)
# # debugging
# print(f"Found the number ({num}) in {count} guesses. Last guess:{guess}")







