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
    for (let i=0; i < n; i++){
        let val = parseInt(ar[i]);
        count[val]++;
    }
    console.log(count.join(' '));
}
