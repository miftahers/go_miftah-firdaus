# Clean and Hexagonal Architecture
"The separation of concerns using layers"

Mission statement: Good architecture is a separation of concern using layer to build a modular, scalable, maintainable and testable application.

MVC:

![](https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftas-dp-prod-media.s3.amazonaws.com%2Fblog%2Fmvc.jpg&f=1&nofb=1&ipt=f26a30e96f0f154232e62f16ff19611451222a2bab10288ca3f5b92adbec457e&ipo=images)

MVVM:

![](https://cdn-ak.f.st-hatena.com/images/fotolife/l/li-yunjie/20190611/20190611111055.png)


If every team has his own structure
- How to maintain this?
- Mobility issue -> It's hard to transfer knowledge on business domain and plus with unique code structure we have extra time and brain to learn that from top to down.
- Another issue -> ex: it's hard to implement the unit test especially test that has connection with database and then we hardly choose the integrated test instead.


History

over the last several years we've seen a whole range of ideas regarding the architecture of systems.
- hexagonal architecture
- onion architecture
- screaming architecture
- DCI from agile development
- BCE from Object Oriented Software
- Clean architecture

Though these architectures all vary somewhat in their details, theyre very similar. they all have the same objectives, which is the separation of concerns. They all achieve this separation by dividing the software into layers. Each has at least one layer for business rules, and another for interfaces.


The constraint before designing the clean architecture are:
- Independent of Framework: framework is tool.
- Testable: The business rules can be tested without any other stuff.
- Independent of UI: The UI can change easily, without changing the rest of the system.
- Independent of Database: Your business rules are not bound to the database, you can swap out the database easily.
- Independent of any external: In fact of business rules simply doesn't know anything about that business. 


Benefit of clean architecture:
- A standard structure, so it's easy to find your way in the project,
- Faster development in the long term,
- Mocking dependenciew becomes trivial in unit tests,
- Easy switching from prototypers to proper solutions (e.g., changing in-memory storage to an SQL database).


Clean Architecture Layer:
(could have more layers)
- (optional) Entities Layer - Business objects as they reflect the concepts that your app manages
- Use Case (Domain Layer)/services - Contains business logic
- Controller (Presentation Layer)/handler - Presents data to a screen and handle use interactions (framework only exist in presentation layer)
- Drivers (Data Layer) - Manages Application data eg. Retrieve data from the network, manage data cache.


Clean Architecture is a softwre architecture

Domain driven design is software design technique

