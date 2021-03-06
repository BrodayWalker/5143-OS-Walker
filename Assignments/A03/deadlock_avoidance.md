## Avoidance
#### Updated: 04 Apr 2020 @ 2320
<sup>Updated By: Broday Walker</sup>
## Deadlock Avoidance
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;For deadlock to occur, the following four conditions must be met:
+ __Mutual Exclusion__ <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Mutual exclusion refers to the idea that only one process may use a resource at a time. When the process is using the resource, no other resources are able to access a resource unit that has been allocated to another process.
+ __Hold and Wait__ <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;A process may hold an allocated resource while awaiting assignment of other resources required to complete a task. For example, a process needing to copy files from one hard drive to another hard drive may hold hard drive A while hard drive B is in use. Once hard drive B is available, the resource will be allocated and the process can complete the file copying.
+ __No Preemption__ <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Preemption is the act of stopping or pausing the use of a resource so another (usually higher priority) process may use it. Once a process holds a resource, it cannot be taken away by another process or the kernel. No resource can be forcibly removed from a process holding it.
+ __Circular Wait__ <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;A circular wait occurs when a closed chain of processes exists such that each process holds at least one resource needed by the next process in the chain. Each process continues to wait for the next, resulting in a situation where no process is able to complete its task and free their claimed resources.

<br><br>
![deadlock](https://cs2.msutexas.edu/~opsysuser/images/deadlock2.png)
<br><br>

### Avoiding Deadlock
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;One way to avoid deadlock is through resource denial: only grant resources if granting them cannot result in a future deadlock situation. This is not an easy solution, as the system would need to be able to predict or know the requests for resources that will be made in the future by other processes. One proposed solution to this problem is the __Banker's algorithm__. The Banker's algorithm allows three of the four deadlock conditions to occur: mutual exclusion, hold and wait, and no preemption. This algorithm avoids the fourth (and perhaps most significant) necessary condition: circular wait. 
<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The Bankers algorithm operates in the context of a system with a fixed number or processes and a fixed number of resources. In this system, a process may hold zero or more resources at any given time. The Banker's algorithm determines whether or not the current state of the system is safe or not. A __safe state__ is one in which there is at least one sequence of resource allocation to processes that does not result in a deadlock. This state ensures that all processes will be able to run to completion. The opposite occurs when the system is in an __unsafe state__. To avoid deadlock, the "banker" will "loan" the resources requested by a process only if there are enough resources available to ensure that at least one process can terminate, freeing it's allocated resources.

