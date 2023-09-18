# eth-parser
Small information about service:
1. HTTP app, implemented with classic onion-style architecture;
2. Application uses own bootstrapper (I use code generation tool to create base for proto file generation), other code written from scratch;
3. Supports inmemory cache, but does not support multi-instance workload deployment;
4. Allows to manipulate internal worker count and manipulate fairness of internal workers serving outbound requests;
5. Since there was no previous experience, transaction parsing done via `eth_GetLogs`, perhaps adjustments required;
6. No excess/useless code written, every method used;
7. Every method implemented with http-endpoints, take a look at `api` folder, proto and swagger present there;
8. Helm charts & pipeline jobs are not defined, since app published in github. In gitlab instead, I could have implemented CI/CD;
9. No external libraries used to perform any sort of tasks (except oklog/run to assign http server and any custom pool to run)

## Installation & startup

1. clone repository
2. `make run`
3. (optional docker-way) `docker build -t . ethparser && docker run ethparser`
4. (optional) check current dependency graph by running `make arch-graph`, validate with `make arch-lint`