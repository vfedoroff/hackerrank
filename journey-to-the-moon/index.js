process.stdin.resume();
process.stdin.setEncoding('ascii');

var input_stdin = '';
var input_stdin_array = '';
var input_currentline = 0;

process.stdin.on('data', function (data) {
    input_stdin += data;
});

process.stdin.on('end', function () {
    input_stdin_array = input_stdin.split('\n');
    main();    
});

function readLine() {
    return input_stdin_array[input_currentline++];
}

function buildGraph(p) {
    let graph = new Map();
    for(let i = 0; i < p; i++){
        let pair = readLine().split(' ');
        if (graph.has(pair[0])){
            graph.set(pair[0],{links: graph.get(pair[0]).links.add(pair[1]), visited: false});
        } else {
            graph.set(pair[0],{links: new Set([pair[1]]), visited: false});
        }
        if (graph.has(pair[1])){
            graph.set(pair[1],{links: graph.get(pair[1]).links.add(pair[0]), visited: false});
        } else {
            graph.set(pair[1],{links: new Set([pair[0]]), visited: false});
        } 
    }
    return graph;
}


function main() {
    let line = readLine();
    let pair = line.split(' ');
    let n = pair[0];
    let p = pair[1];
    let graph = buildGraph(p);
    let combinations = 0;
    let astronauts = 0;
    for (var [key, value] of graph.entries()) {
        if (! value.visited) {
            value.visited = true;
            let queue = [key];
            let country_size = 1;
            // dfs
            while ( queue.length > 0 ) {
                let astronaut = graph.get(queue.shift());
                for (let link of astronaut.links){
                    let connection = graph.get(link);
                    if (! connection.visited) {
                        connection.visited = true;
                        country_size++;
                        queue.push(link);
                    }
                }
            }
            combinations += country_size * astronauts;
            astronauts += country_size;
        }
    }
    combinations += (astronauts + n - 1) * 0.5 * (n - astronauts);
    console.log(combinations);
} 
