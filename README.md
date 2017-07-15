```
.
├── delivery            // Serve content via HTTP? CLI? everything related to that should be here
|
├── usecases            // (service layer???) The glue between our delivery layer and our domain layer. Different
|                       // Delivery mechanisms probably will have the same use cases or very similiar
|                       // use cases, this allows you to use the same code for different mechanisms by
|                       // using the same use case interactors in different mechanisms
|
├── domain              // Where we have our domain logic
|
└── infrastructure      // Where we have our implementation details (Database connections, Queues, External services)

```

```
# dependencies flow in one direction
delivery (depends on) usecases (depends on) domain (depends on infrastructure)
```