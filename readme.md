# HTTP server benchmarks

- Apache Bench: https://httpd.apache.org/docs/2.4/programs/ab.html

ab -c 10 -n 30000 http://localhost:{{port}}

# Languages / Runtimes

- Go
- Elixir
- Python
- Nodejs
  - express
  - fastify
  - koa
- Bun
- Deno

Get the best candidate of each language and test it behind ngnix and caddy

- Caddy/Go
- Caddy/Elixir
- Caddy/Nodejs
- Caddy/Bun
- Caddy/Deno

- Ngnix/Go
- Ngnix/Elixir
- Ngnix/Nodejs
- Ngnix/Bun
- Ngnix/Deno

# System specs

<img src="./.github/systemspec.png" alt="System specs">
