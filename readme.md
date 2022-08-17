### Assemblyscript

Let's write our Assemblyscript

```asc
export function sum(a: i32, b: i32): i32 {
  return a + b;
}
```

Compile with

```bash
asc exports.ts -o exports.wasm
```

### Go

```bash
go run .
```

navigate to `http://localhost:3000/10/12`, you will get `32`
