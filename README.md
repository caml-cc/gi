# GI

A simple HTTP service that generates .gitignore files.

## Usage

```bash
# get default .gitignore
curl gi.caml.cc > .gitignore

# list available templates  
curl gi.caml.cc/list

# get specific template
curl gi.caml.cc/go > .gitignore

# combine templates
curl gi.caml.cc/go,macos > .gitignore
```

## Development

```bash
git clone https://github.com/CircuitCamel/gi
cd gi
make full
```

> G.I., your government has abandoned you. They have ordered you to die. Don’t trust them. They lied to you, G.I.s, you know you cannot win this war.
