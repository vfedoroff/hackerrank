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

/////////////// ignore above this line ////////////////////

function main() {
    // the number of queries
    const q = parseInt(readLine());
    for(let i = 0; i < q; i++){
        const n_temp = readLine().split(' ');
        const numCities = parseInt(n_temp[0]);
        const numRoads = parseInt(n_temp[1]);
        const libCost = parseInt(n_temp[2]);
        const roadCost = parseInt(n_temp[3]);
        if (roadCost >= libCost || numRoads === 0) {
            input_currentline += numRoads;
            console.log(numCities * libCost);
            continue;
        }
        let nodes = new Map();
        for(let i = 0; i < numRoads; i++){
            const cities = readLine().split(' ');
            const city1 = parseInt(cities[0]);
            const city2 = parseInt(cities[1]);
            if (nodes.has(city1)){
                nodes.set(city1,{links: nodes.get(city1).links.add(city2), visited: false});
            } else {
                nodes.set(city1,{links: new Set([city2]), visited: false});
            }
            if (nodes.has(city2)){
                nodes.set(city2,{links: nodes.get(city2).links.add(city1), visited: false});
            } else {
                nodes.set(city2,{links: new Set([city1]), visited: false});
            }
        }
        let cost = 0;
        let visited = 0;        
        for (var [key, value] of nodes.entries()) {
            if (! value.visited) {
                value.visited = true;
                visited++;
                cost += libCost;
                let queue = [key];
                while ( queue.length > 0 ) {
                    let city = nodes.get(queue.shift());
                    for (let link of city.links){
                        let connection = nodes.get(link);
                        if (! connection.visited) {
                            connection.visited = true;
                            visited++;
                            cost += roadCost;
                            queue.push(link);
                        }
                    }
                }        
            }
        }
        //account for cities with no roads
        const strandedCities = numCities - visited;
        console.log(cost+strandedCities*libCost);        
    }
}
