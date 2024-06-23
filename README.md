# Balance

A website that can compliment and insult you, at the same time.

## Installation

1. Populate the environment variables defined at `.env.example` and rename it to `.env`

2. Regenerate the proto codes
   _(Optional step, generated codes are checked into git)_

   ```bash
   cd proto
   ./generate.sh
   ```

3. Run the http server and the grpc server

   ```bash
   # App
   make run
   # Classification grpc
   make py-server
   ```
