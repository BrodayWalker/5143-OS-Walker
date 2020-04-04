## Readers-Writers Problem
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The __Readers-Writers Problem__ is an example of a synchronization problem in which many processes or threads are attempting to access and use a shared resource. Processes must be controlled in such a way to ensure data is not overwritten as it is being read by another process and that data is not overwritten while another process is writing. There are several variations of this problem, but a general explanation involves the following context to set up the problem:
<br>
+ There is a shared object (in our example, some text file) used by many threads or processes
+ Threads/processes belong to one of the following categories:
	+ __Readers__: those threads or processes that read the data from the shared object, but do not modify it
	+ __Writers__: those threads or processes that read data and modify it
+ Using a single lock to control the whole file is too restrictive 
+ Many readers should be allowed to access the file concurrently as they do not make any modifications.
+ Only one writer should be able to modify the data at a time.
+ The reads and writes happen in a __critical section__, which is a part of the program that cannot but executed simultaneously with another asynchronous procedure. As stated before, reads may occur concurrently with other reads in the critical section because they do not modify the shared object.
+ In the __critical section__, writers must be guaranteed __mutual exclusion__, which is to say that the writer is the only process given access to the resource (the shared object text file) when it is its turn to write to the file.
+ Control of the __critical section__ is accomplished by a __mutex__.

### Read/Write Locks
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Typically, the Readers-Writers problem discussed above is solved using a __read/write lock__. Many readers can gain access to the file concurrently by acquiring the mutex protecting the critical section of code. The last reader to leave the critical section will signal a waiting writer that they can enter the critical section. The write will obtain the mutex and proceed to write (no other processes are in the critical section at this point). If another writer is waiting, it can obtain the lock and proceed to write. The readers and writers alternate control of the critical section of code using this locking mutex strategy.

