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
    const n = parseInt(readLine());
    let ar = readLine().split(' ');
    let count = new Array(100);
    count.fill(0);
    // calculate the histogram of key frequencies:    
    for (let i=0; i < n; i++){
        let key = parseInt(ar[i]);
        count[key]++;
    }
    // calculate the starting index for each key:
    let output = new Array();
    for(let i=0; i<100; i++) {
        let cnt = count[i];
        // for the number of times the item occurs
        for(let j=0;j<cnt;j++){
            output.push(i);
        }
    }
    console.log(output.join(' '));   
}
