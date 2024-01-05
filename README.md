# Coffee Shop Management

## Contributors

| ID                                           | Fullname                                               | Github                                                    |
| -------------------------------------------- | ------------------------------------------------------ | --------------------------------------------------------- |
| &nbsp;&nbsp;&nbsp;21520095&nbsp;&nbsp;&nbsp; | &nbsp;&nbsp;&nbsp;Bùi Vĩ Quốc&nbsp;&nbsp;&nbsp;        | &nbsp;&nbsp;&nbsp;[bvquoc](https://github.com/bvquoc)     |
| &nbsp;&nbsp;&nbsp;21520339&nbsp;&nbsp;&nbsp; | &nbsp;&nbsp;&nbsp;Nguyễn Lê Ngọc Mai&nbsp;&nbsp;&nbsp; | &nbsp;&nbsp;&nbsp;[NLNM-0-0](https://github.com/NLNM-0-0) |
| &nbsp;&nbsp;&nbsp;21521495&nbsp;&nbsp;&nbsp; | &nbsp;&nbsp;&nbsp;Nguyễn Kim Anh Thư&nbsp;&nbsp;&nbsp; | &nbsp;&nbsp;&nbsp;[kimthu09](https://github.com/kimthu09) |

### Supervisors

**Nguyen Trinh Dong**

## I. How to run project

### a. Frontend

```bash
cd frontend/
npm install
npm run build
npm start
```

### b. Backend

```bash
docker-compose up
```

## II. Run test

```bash
cd backend/
```

### a. Unit-test

##### x86 CPU

```bash
sh ./backend/run_test_x86.sh
```

##### arm64 CPU

```bash
sh ./run_test_arm64.sh
```

### b. Mutation test

##### x86 CPU

```bash
sh ./run_mutation_test_x86.sh
```

##### arm64 CPU

```bash
sh ./run_mutation_test_arm64.sh
```

### c. Unit test coverage

##### x86 CPU

```bash
sh ./run_coverage_x86.sh
```

##### arm64 CPU

```bash
sh ./run_coverage_arm64.sh
```
