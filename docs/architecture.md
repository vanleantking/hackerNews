1. `clean architecture` vs `hexagonal architecture`
   https://www.youtube.com/watch?v=gVZM61e-uJw
2. `clean architecture` vs `domain driven design`
3. https://github.com/amitshekhariitbhu/go-backend-clean-architecture
4. https://github.com/bxcodec/go-clean-arch
5. https://github.com/khannedy/golang-clean-architecture
6. https://blog.scalablebackend.com/understand-the-theory-behind-clean-architecture
7. https://github.com/golang-standards/project-layout
8. https://eltonminetto.dev/en/post/2020-07-06-clean-architecture-2years-later/
9. https://ccd-akademie.de/en/clean-architecture-vs-onion-architecture-vs-hexagonale-architektur/
10. https://medium.com/@ebubekiryigit/hexagonal-architecture-a-golang-perspective-7eb3cb6117e7
11. https://eltonminetto.dev/en/post/2020-07-06-clean-architecture-2years-later/
12. https://threedots.tech/post/ddd-cqrs-clean-architecture-combined/


**PROJECT ARCHITECTURE**
 - backend
    -- cmd
        ---- api
    -- db
    -- docker
    -- docs
    -- internal
        ---- components
        ---- delivery
        ---- entity
        ---- model
        ---- repository
        ---- storage
        ---- usecase
    -- pkg
- `cmd`: expose application into runable format => expose `api` into `main` package for run api instance
- `db`: run on migrations purpose
- `docker`: include `api` and `postgre` service for run docker service
- `internal`: core folder for our application:
   + `components`: declare all of external package use in app
   + `delivery`:  exposing your application's use cases to the outside world, it handles how requests come in and responses go out
   + `entity`: table database
   + `model`: object && their method
   + `repository`: implement function to interact with database (crud operations)
   + `storage`
   + `usecase`: higher level with `repository` level, interaction between entities and repositories to fulfill the use case.
   + `service`: Often more granular than use cases, providing specific functionality that may be reused across multiple use cases. Handle tasks that might not directly relate to a single user action, such as sending notifications, processing payments, or interacting with external systems.
    => key difference between `usecase` vs `service`
    ![`usecase` vs `service`](./usecase_service.png)
    Use Cases: When you need to model a specific user interaction or a complete business process that involves multiple steps. Use cases help you encapsulate complex business logic and workflows.
    Services:
        Domain Services: When you have complex business logic that doesn't naturally belong to an entity or a value object.
        Infrastructure Services: When you need to interact with external systems, libraries, or resources (e.g., databases, email providers, payment gateways).
    + `domain` layer = `usecase` + `repository`
