# jsonforms-parser

A dependency-free Go parser for [JSON Forms](https://jsonforms.io) schemas. Produces a type-safe AST for building custom renderers.

## Features

- **All Standard Elements** - Control, Layout, Group, Categorization, Category, Label
- **Custom Elements** - Unknown types preserved with parsed children
- **Rules & Conditions** - HIDE/SHOW/ENABLE/DISABLE with schema-based, leaf, AND, OR conditions
- **Visitor Pattern** - Traverse and transform the AST
- **Type-Safe** - Strongly-typed nodes, no generic maps
- **Zero Dependencies** - Pure Go stdlib

## Installation

```bash
go get github.com/tinybluerobots/jsonforms-parser
```

## Usage

```go
import "github.com/tinybluerobots/jsonforms-parser"

ast, err := jsonforms.Parse(uiSchemaJSON, schemaJSON)
if err != nil {
    log.Fatal(err)
}

layout := ast.UISchema.(*jsonforms.VerticalLayout)
control := layout.Elements[0].(*jsonforms.Control)
fmt.Println(control.Scope) // "#/properties/name"
```

## Elements

All elements implement `UISchemaElement`:
- `Control` - Binds UI to data property
- `VerticalLayout`, `HorizontalLayout` - Layout containers
- `Group` - Labeled container
- `Categorization`, `Category` - Tab navigation
- `Label` - Static text
- `CustomElement` - Unknown/custom types

## Visitor Pattern

```go
type MyVisitor struct {
    jsonforms.BaseVisitor
}

func (v *MyVisitor) VisitControl(c *jsonforms.Control) error {
    fmt.Println("Found:", c.Scope)
    return nil
}

jsonforms.Walk(ast.UISchema, &MyVisitor{})
```

## Use Cases

Build custom renderers, transform schemas, validate structures, generate docs, or convert between form systems.
