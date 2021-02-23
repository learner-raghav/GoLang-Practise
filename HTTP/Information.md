1. The onion architecture is an architectural pattern that enables maintainable and 
evolutionary enterprise systems.

2. Following are the principles of onion architectures
    
    1. Good Coupling: The primary proportion of this architecture is called good coupling. Coupling is nothing 
        but dependency of one thing over another. The higher the coupling, the lower the ability to change
       and evolve the system.
       
    2. Dependencies should always be inward and never be outward. Code should depend
        on the same layer or layers more central to itself.
       
    3. Separation of concerns: The layers are
        1. Domain Model - High level data objects we use
        2. Domain Services - Where our business logic lives
        3. Application services - Here we define what we can do through a series of contracts
        4. External services - Infrastructure, Tests and User Interface.
    
3. One of the core advantages of the architecture is to increase maintainability.

4. Introduced by Jeffrey Palermo, Onion architecture guides engineers to model their business logic in a core collection with 
    no coupling to the outer concerns, such as the Database choice or how th UI interacts.
   
5. So, typically like an onion, we walk our way into the core. 

6. "Externalize your dependencies and decouple them using contracts"
    
