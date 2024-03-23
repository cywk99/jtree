utility tool to prettify json in cli

1. go build -o jtree
2. `./jtree  '{"name": "John", "age": 30, "city": "New York"}'`
3. `curl --location 'https://www.arbeitnow.com/api/job-board-api' | ./jtree -`