# SINGLETON

- A component which is instantiated only once.

## Thread safety

- Use either sync.Once or a package lever .init()

## Laziness

- Only load the data whenever somebody asks for it (not guaranteed in .init(), but is guaranteed in sync.Once )
