# Golang Concurrency and Datastore Modeling

## Why This Article/Project :

```
The Motivation Behind this project is to touch up on what I've covered about web-development so far since the beginning. When you begin learning stuff in web-development, concepts start out incredibly abstract so its very hard to gain an intuitive understanding of what is happening. In this project I will try to solidify my understanding of concepts as well as  learn Go. I will be remaking my Shopping List API and touch on both programming skills by creating in-memory managed structures to build a non-persistent non-optimized API, as well as touch up on web-development specfic concepts of deployment, managed datastores, and cache systems with Docker/Containerd/Kubernetes, Postgres, and Redis.
```

---

## The Main Learning Objectives Of This Project for Myself are

- The Go Language
- Comparisons Between a Program's Memory View of Data, A File System's View of Data, and a Database's view of Data.

---

## When Learning Golang and Application Development, I will be Covering

1. HTTP Servers, Request Parameters, Request Parsing, and Sending Requests (formatted HTML or JSON API For an SPA)
2. Concurrency (GoRoutines, Channels, Mutexes/Semaphores/Locks, and the Share By Memory GoLang Model)
3. Database Connections and Querying (Undecided whether Vanilla SQL, Query-Builder, or ORM)
4. Application-Level Security (Refresh Tokens, Access Tokens, Salting/Hashing, OpenIDConnect and 0Auth, HTTPS)
5. Containers and Container Orchestrator Services

---

## By Comparing Different Datstore Techniques, I Will Look At

1. Comparison of what a RDBMS has to offer in comparison to handling flat files of a particular format with our own simple program.
2. Comparison of Data Structures optimized to handle Querying from Disk with a Memorry Cache vs A Memmory Only Programming View.
3. Concurrency for dealing with updating Data Structures (File Systems, Indexes, Tuples (WALs and Locks) vs. In-Memory Golang Structures by using our own Synchronization Primitives)
4. Datbase Design Decisions and Performances (example Indexes Pointing to PKs vs Tuples, and what this even means)
5. In-Memory Datastructures with what an In-Memory Cache (Redis) has to offer.
