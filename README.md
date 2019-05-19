# Activity Scheduler

## Technical description

This problem was solved using Genetic Algorithm.

This decision was made considering the different constraints in the problem, such as time window, budget limit,
etc.

Genetic algorithms don't provide an exact solution all the time, but they are fast enough for a quick response
in complex scenarios, such the one presented here.

The algorithm works as below:
1. A initial population of 16K individuals is generated
2. Every generation is crossed-over among them, generating the next population
3. On each generation, the two individuals with the worst fitness are sacrificed
4. The evolution continues until a single individual is left alive, being the one with the highest fitness score
and, thus, the best result.

The process above is repeated for each day of trip of the customer.

The budget constraint is divided equaly among the days.

In the end, the last individual is translated into the JSON structure presented on the terminal output.

## How to run the binary?

This project is built on Go, so the Go compiler is a requirement for the project.

Instructions on how to install Go can be found [here](https://golang.org/doc/install#install).

After Go compiler is installed:

```bash
# Change to the project path
$ cd [project path]

# Build the project
$ go build -o activity-scheduler

# Run it
$ ./activity-scheduler [path to json file] [budget] [days]
```

## How to run the tests?

```bash
$ go test -v ./...
```

## Limitations

Given the short amount of time to implement a full solution, some aspects were postponed to a future version:

1. Results tends to be close to the given arguments, but aren't guarantee. This is a characteristic of the algorithm used (genetic algorithm). Thus, some results with a total budget above the informed are possible. This can be solved implementing filtering on the generations, killing the individuals that overflow the budget, and increasing the size of the initial population

2. Commutation time isn't considered in this version. So, activities starts as soon the previous one is finished. This can be solved including this parameter into the fitness function, penalizing the results that have a very short window between activities

3. Individuals are limitated to 3 activities fixed. This was done considering the complexity to generate random
activity scheduling on the initial population and the complexity on the crossover function. In the next iteration, support to non-rigid activities size is planned

4. The system contains basic tests, specially for the chromossome methods, which are the core of the algorithm. Further tests are necessary, specially integration tests.

5. This system can be easily containerized using Docker, which allows us to deploy the solution using container managers, such as Kubernetes, for example. The Dockerfile implementation is planned for the next iteration as well.
