# REST vs. GraphQL

1. REST and GraphQL are two wasy of sending data over HTTP. The REST based approach is traditional way of doing so and is used a lot in many application stacks in the previous years.

2. GraphQL is presented as a revolutionary way to think about APIs

3. REST
    
    1. REST(Representational state transfer) is an API design architecture used to implement web services. REST-compliant web services allow the requesting
        systems to access and manipulate textual representations of web resources by sharing a uniform and predefined set of stateless operations. When we use HTTP, the most common available are 
       GET,POST,PUT,DELETE,PATCH
       
    2. Limitations of REST
        
        1. Multiple round trips to fetch related resources.
        2. Over fetching/Under fetching.
    
4. GraphQL - A different approach
    
    1. As REST, GraphQL is an API design architecture, but with a different approach which is much flexible. The 
        main and the most important difference is that, GraphQL is not dealing with dedicated resources, Instead everything is regarded as a graph and therefore is connected.
       
    2. This means we can tailor the request to the exact needs by using the GraphQL query language and describing what you would like to get as an answer. 
    3. Both REST and GraphQL are API design architectures, The RESTFul approach is limited to deal with single reosurces. The GraphQL approach is much more flexible.