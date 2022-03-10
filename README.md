> # Monorepo CI/CD example with testing stage

---

This is simple TODO app. Backend written in Golang and frontend in Vuejs.
There isn't used any database systems. All TODO items is stored in memory.

**Features:**

-   add task
-   update task (not implemented)
-   remove task
-   mark as done
-   mark as undone

## Run e2e tests locally

First, you need to install cypress

```
yarn install
```

Secondly, you need to pull together backend and frontend with help of docker-compose

```
docker-compose up --build
```

Now you can run e2e tests

```
yarn cy:run
```
