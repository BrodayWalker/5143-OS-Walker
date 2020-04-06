
# Other Classic Synchronization Problems
## Race Conditions
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;A __race condition__ is a situation in which multiple processes access and modify shared data with the outcome dependent on the relative timing of the processes. This type of problem results in non-deterministic behavior. Race conditions are serious debugging issues as they can be intermittent. The result of a computation that suffers from a race condition could be correct some of the time and incorrect at other times. Various strategies may be used to combat race conditions. The typical solution for a race condition is to control the critical section of code with a mutex or semaphore. 
#### Example
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;p1 and p2 are processes that simultaneously execute the statement x = x + 1, where x is a static global variable. Each process reads the value of x, adds 1, then writes the result to x. If both processes read x = 1, both processes will update x as x = 2. In this situation, x was intended to be incremented twice, but it was only incremented once. 

<br><br>
![race_condition](https://cs2.msutexas.edu/~opsysuser/images/race_cond.png)
<br><br>


#### Mutex Example
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;p1 and p2 are processes that simultaneously execute the statement x = x + 1, where x is a static global variable. p1 and p2 simultaneously try to acquire the mutex protecting this critical section of code. p2 obtains the lock, reads the value of x, increments x, then releases the lock. p1 obtains the lock, reads the value of x (which is now x = 2), then releases the lock. The final value of x is 3.
<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The fixed example works because mutexes can only be held by one thread or process at a time. If several threads attempt to acquire the lock at the same time, only one will win the race and be admitted into the critical section of code to complete some computation. 

