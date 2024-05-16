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
- Iterate over json array element by element to minimize memory needed
- Check if modifier "_39" is present as an indication of NY PPO present in the location
- Write location to a file in ny_ppo_output

Time To Complete: 2.5 hours
- Approx. 1hr to write json parser
- Approx. 1 hr to analyze json and research potential indicators of NY PPO and finalize code
- Approx. 0.5 hr of running different solutions
  
Time To Run Final Solution: 12 minutes

Trade Offs:
- Simple check for NY PPO. Didn't utilize description as most NY PPO I saw in the description were not Anthem related.
- Potentially could utilize Anthem's MRF EIN lookup for more robust confirmation of NY PPO. Perhaps running multiple operations in parrallel as each loop is independent.


