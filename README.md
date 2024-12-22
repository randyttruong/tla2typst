# tla2typst (in progress) 
Convert TLA+ Specifications to Typst code with customizable profiles. 

inspired by [tla2tex](https://github.com/hengxin/tla2tex)

# Cloning
1. Clone the repository
2. From the `/tlatypst` directory, run any of the following `make` targets: 
```bash
# deletes the /build folder (if exists), then builds the project, storing it in the /tla2typst/build folder 
make clean build

# build the project to the /tla2typst/build folder
make build
```

# Usage 
```bash
./tla2typst <tlaFilename> 
```

# requirements
- the tla2typst typst package, which contains all necessary scripting for 
properly formatting a typst file containing formatted tla++ code.

- the tla2typst program, which is a CLI that ingests `.tla` specifications and outputs them as valid `.typ` file.


