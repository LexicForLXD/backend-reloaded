## Places to modify for another graphql resolver

- schema in coresponding folder
- main.go (Routes()): add usecase to resolvers
- resolvers/resolver.go: add usecase to new...() function
- resolvers/query.go: add usecase to new...() function
- resolvers/mutation.go: add usecase to new...() function


## Places to modify for another rest endpoint

- main.go mount Handler with usecase on api endpoint
- 