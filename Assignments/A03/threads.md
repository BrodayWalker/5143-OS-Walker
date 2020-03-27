# Threads
Author: Broday Walker
References: <br> 
   1. Abraham Silberschatz, Greg Gagne, and Peter Baer Galvin, "Operating System Concepts, Ninth Edition", Chapter 4
   2. William Stallings, "Operating Systems: Internals and Design Principles, Seventh Edition", Chapter 4


## Overview
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Threads are commonly described as a __lightweight process__. A thread is a basic unit of CPU utilization, consisting of a program counter, a stack, and a set of registers. A multi-threaded program shares code, data, and files among the threads performing work. Each thread has its own program counter, stack, and set of registers. 

## Purpose of threads
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Threads allow for more work to be done per unit of time. The classic approach to running a process followed the single-threaded model. In this approach, a process followed an execution path in a certain order. In other words, the program was executed serially.  Contemporary operating systems allow for programs in execution to follow multiple, concurrent paths through the use of threading. There can be multiple processes which have multiple threads.
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;When studying threads, it is important to understand the link between threads and processes. A __process__ is defined as the unit of resource allocation and a unit of protection, which includes a virtual address space that holds the process image, protected access to processors, other processes, files, and I/O resources. A process includes a process control block, the virtual address space, a user stack, and a kernel stack as shown below. <br>
![thread_stack](https://github.com/BrodayWalker/5143-OS-Walker/tree/master/images/thread_stack.jpg)
<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;



## Thread properties
