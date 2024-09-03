# compac
Simple and powerful data types, based on generics

## 🧰 How to install
```
go get github.com/go-compac/compac
```

## Generic nullable data-types

`compac.Nl[T]` - provides fast and easy to use sql.Null-like abstraction, which allocates memory on stack. Also removes checks on nil from code.

### 🛠 How to use

<table>
<thead><tr><th>StdLib</th><th>compac.Nl[T]</th></tr></thead>
<tbody>
<tr><td>

```go
// Optional struct fields in std lib with pointers
type personInfo struct {
    ID        int
    FirstName *string
    LastName  *string
    Age       *int
}
```

</td><td>

```go
// Optional struct fields in compac with compac.Nl
type personInfo struct {
    ID        int
    FirstName compac.Nl[string]
    LastName  compac.Nl[string]
    Age       compac.Nl[int]
}
```

</td></tr>

<tr><td>

```go
// Optional(not functional option) func params in stdLib
func CountHumansInCity(cityID int, onlyWithName *string) int {
    query := `select count(*) from humans`
    var filter []string
    filter = append(filter, fmt.Sprintf("city_id = %v", cityID))
    if onlyWithName != nil {
        filter = append(filter, fmt.Sprintf("name = '%v'", *onlyWithName))
    }

    if len(filter) > 0 {
        query += " WHERE " + strings.Join(filter, " AND ")
    }
    // just an abstraction
    return db.Query(query)
}
```

</td><td>

```go
// Optional(not functional option) func params in stdLib
func CountHumansInCity(cityID int, onlyWithName compac.Nl[string]) int {
    query := `select count(*) from humans`
    var filter []string
    filter = append(filter, fmt.Sprintf("city_id = %v", cityID))
    if onlyWithName.Valid {
        filter = append(filter, fmt.Sprintf("name = '%v'", onlyWithName.Data))
    }
    
    if len(filter) > 0 {
        query += " WHERE " + strings.Join(filter, " AND ")
    }
    // just an abstraction
    return db.Query(query)
}
```

</td></tr>
</tbody></table>

[Example of sql query builder with Nl[T]](example/nl_test.go)

