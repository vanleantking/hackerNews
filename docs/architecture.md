1. `clean architecture` vs `hexagonal architecture`
   https://www.youtube.com/watch?v=gVZM61e-uJw
2. `clean architecture` vs `domain driven design`
3. https://github.com/amitshekhariitbhu/go-backend-clean-architecture
4. https://github.com/bxcodec/go-clean-arch
5. https://github.com/khannedy/golang-clean-architecture
6. https://github.com/harmannkibue/golang-gin-clean-architecture
7. https://blog.scalablebackend.com/understand-the-theory-behind-clean-architecture
8. https://github.com/golang-standards/project-layout
9. https://eltonminetto.dev/en/post/2020-07-06-clean-architecture-2years-later/
10. https://ccd-akademie.de/en/clean-architecture-vs-onion-architecture-vs-hexagonale-architektur/
11. https://medium.com/@ebubekiryigit/hexagonal-architecture-a-golang-perspective-7eb3cb6117e7
12. https://eltonminetto.dev/en/post/2020-07-06-clean-architecture-2years-later/
13. https://threedots.tech/post/ddd-cqrs-clean-architecture-combined/
14. https://medium.com/@shershnev/layered-architecture-implementation-in-golang-6318a72c1e10
15. https://medium.com/@kyodo-tech/layered-go-applications-simple-testable-design-1648c7e44b18
16. https://github.com/AleksK1NG/Go-Clean-Architecture-REST-API/tree/master
17. https://dev.to/bagashiz/building-restful-api-with-hexagonal-architecture-in-go-1mij
18. https://dev.to/dyarleniber/hexagonal-architecture-and-clean-architecture-with-examples-48oi
19. https://medium.com/@rayato159/how-to-implement-clean-architecture-in-golang-en-f50d66378ebf
20. https://medium.com/@sadensmol/my-clean-architecture-go-application-e4611b1754cb
21. https://www.youtube.com/watch?v=eUW2CYAT1Nk
22. https://github.com/resotto/goilerplate/tree/master
23. https://outcomeschool.com/blog/go-backend-clean-architecture
24. https://vidhyanshu.medium.com/implementing-clean-architecture-in-go-56aca59311b3
25. https://www.reddit.com/r/golang/comments/1bzg8gi/what_architecture_has_been_the_best_for_you/
26. https://github.com/herryg91/go-clean-architecture
27. https://www.jamesmichaelhickey.com/domain-driven-design-aggregates/
28. https://www.plainionist.net/Implementing-Clean-Architecture-UseCases/
29. https://www.reddit.com/r/golang/comments/10020js/domaindriven_design_framework_for_go_developers/
30. https://englyk.com/book2/
31. https://www.alexedwards.net/blog/11-tips-for-structuring-your-go-projects


**PROJECT ARCHITECTURE**
 - backend
    -- cmd
        ---- api
        ---- migrations
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

- backend
    -- cmd
        ---- api                   # Main API entry point
        ---- migrations            # Database migrations entry point
    -- db                          # Database migrations, seeds, and scripts
    -- docker                      # Docker configurations (API, PostgreSQL, etc.)
    -- docs                        # Documentation for the project
    -- internal
        ---- domain                 # Core business logic (DDD principles)
             ---- entity             # Business entities (e.g., User, Order)
             ---- service            # Domain services (complex business rules)
             ---- factory            # Factories for creating domain objects
             ---- event              # Domain events (e.g., OrderCreatedEvent)
        ---- delivery               # Interface Adapters
             ---- http               # HTTP handlers (REST API)
             ---- grpc               # gRPC handlers (if needed)
             ---- cli                # Command-line interface handlers
        ---- repository             # Data access layer (implements interfaces)
             ---- postgresql         # PostgreSQL repository implementation
             ---- mongodb            # MongoDB repository (if applicable)
             ---- cache              # Redis or in-memory caching
        ---- storage                # File storage (S3, local, etc.)
        ---- usecase                # Application layer (business logic orchestration)
        ---- components              # External libraries & integrations
    -- pkg                          # Shared utility functions & helpers
    -- config                       # Configuration files (env, yaml, etc.)
- `cmd`: expose application into runable format =>
    + expose `api` into `main` package for run api instance
    + expose `migrate` on `migrations` for construct database hierachy
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
    + `domain` layer = `usecase` + `repository` => in a simple application, it's includes `usecase + repository`, for other complex applications, to know how many domains you need for your application, dive into `DDD`
      * `domain layer` is responsible for anything that has to do with `business logic` + `business decisions` + `business terminology`
=> `DDD(domain driven design)` => Domain design => domain expert => expose `problem space` for domain (banking, real estate) => build common language between business domain expert + developer `before` start building the system

**clean architecture**
    a `software design philosophy` proposed by Robert C. Martin (Uncle Bob), `emphasizes the separation of concerns and the independence of frameworks, databases, and user interfaces`
    I. principle of clean architecture:
    - `Separation of Concerns`: Dividing the system into layers, each with a specific responsibility.
    - `Dependency Rule`: Code dependencies can only point inward. `High-level modules should not depend on low-level modules`.
    - `Encapsulation`: Each layer hides its internal workings from others, exposing only what is necessary.
