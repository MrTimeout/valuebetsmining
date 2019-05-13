```golang
root := Root{
		Query: graphql.NewObject(
            grapgql.ObjectConfig{
                Name: "Query",
                Fields: graphql.Fields{
                    "key": &graphql.Field{
                        Type: graphql.NewList(User)
                    }
                },
            },
		),
    }

type ObjectConfig struct {
    Name        string      `json:"name"`
    Interfaces  interface{} `json:"interfaces"`
    Fields      interface{} `json:"fields"`
    IsTypeOf    IsTypeOfFn  `json:"isTypeOf"`
    Description string      `json:"description"`
}

type Fields map[string]*Field

type Field struct {
    Name              string              `json:"name"`
    Type              Output              `json:"type"`
    Args              FieldConfigArgument `json:"args"`
    Resolve           FieldResolveFn      `json:"-"`
    DeprecationReason string              `json:"deprecationReason"`
    Description       string              `json:"description"`
}

type IsTypeOfFn func(p IsTypeOfParams) bool

type IsTypeOfParams struct {
    // Value that needs to be resolve.
    // Use this to decide which GraphQLObject this value maps to.
    Value interface{}

    // Info is a collection of information about the current execution state.
    Info ResolveInfo

    // Context argument is a context value that is provided to every resolve function within an execution.
    // It is commonly
    // used to represent an authenticated user, or request-specific caches.
    Context context.Context
}

```