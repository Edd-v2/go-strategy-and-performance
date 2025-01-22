# Strategy Pattern Application: Demonstrating Task Execution Approaches

This application illustrates the use of the **Strategy Pattern** to demonstrate performance differences between various task execution methods. Specifically, it compares:

1. **Sequential Execution**
2. **Concurrent Execution with Goroutines and Channels**
3. **Concurrent Execution with Goroutines and Mutexes**

---

## Project Goal

The primary goal is to showcase how the chosen execution approach affects:
- **Efficiency**: The speed of task completion.
- **Parallelism**: The system's ability to handle multiple tasks simultaneously.
- **Overhead**: The cost of synchronization or managing shared resources.

---

## Strategies Implemented

### 1. **Sequential Execution**
- **Description**: Tasks are executed one after another without concurrency.
- **Characteristics**:
  - Each task waits for the previous one to finish.
  - Does not utilize multicore hardware effectively.
- **Real-World Use**: Suitable for workflows requiring strict task ordering.

### 2. **Goroutines with Channels**
- **Description**: Tasks run in separate goroutines, and **channels** coordinate their execution.
- **Characteristics**:
  - Full parallel execution.
  - Channels regulate task communication and workflow.
- **Real-World Use**: Perfect for independent tasks, such as handling API requests or batch processing.

### 3. **Goroutines with Mutexes**
- **Description**: Tasks execute concurrently but access shared resources via a **mutex**.
- **Characteristics**:
  - Prevents race conditions by ensuring safe access to shared data.
  - Mutex locking introduces potential bottlenecks.
- **Real-World Use**: Essential for operations involving shared resources (e.g., caching).

---

## Performance Analysis

Below is a comparison of execution times based on a simulated workload of 10 tasks:

| Strategy               | Workload | Execution Time  |
|------------------------|----------|-----------------|
| **Sequential**         | 10       | 104.7116 ms     |
| **Goroutines + Channel**| 10       | 10.0814 ms      |
| **Goroutines + Mutex**  | 10       | 100.3438 ms     |

---

## Conclusions

1. **Sequential Execution**:
   - Slowest approach; avoid for independent tasks.
   - Best suited for ordered workflows.

2. **Goroutines + Channels**:
   - Most efficient for parallel, independent tasks.
   - Highly recommended for high-concurrency systems.

3. **Goroutines + Mutex**:
   - Necessary for safe shared resource access.
   - Optimize to reduce mutex contention wherever possible.

---

## Real-World Applications

- **Sequential Execution**: Processing files line by line in strict order.
- **Goroutines + Channel**: Parallel task handling like API requests or pipeline processing.
- **Goroutines + Mutex**: Managing shared counters, caches, or in-memory data structures.

---

By applying the **Strategy Pattern**, this project highlights how to choose the optimal task execution method depending on the requirements and context.
