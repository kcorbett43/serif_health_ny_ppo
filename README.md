Project Summary:
Parse and analyze a large dataset (“anthem index”) containing healthcare provider and pricing information, and identify those entries that are relevant to the NY PPO context.

Output file is in the ny_ppo_output file.

Executing Code (Will create txt file in the ny_ppo_output file):
- Docker
  - If you have Docker installed you can run the following commands:
    - Build image and run server: "docker-compose up"
    - Find the container id: "docker ps"
    - enter container: "docker exec -it container-id bash"
    - run main.go in the container: "go run main.go"
- If Go is installed:
  - run from command line: "go run main.go"
 
Solution Overview:
- Download anthem index and read from gzip.NewReader
- Create structs representing the price transparency schema
- Iterate over json array element by element to minimize memory overhead
- Check if modifier "_39" is present as an indication of NY PPO present in the location
- Write location to a file in ny_ppo_output
