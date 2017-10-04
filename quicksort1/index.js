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
    let left = [];
    let right = [];
    let p = parseInt(ar[0]); 
    for (let i=0; i < n; i++){
        let val = parseInt(ar[i]);
        if (val >= p){
            right.push(val);
            continue;
        }
        if (val < p){
            left.push(val);
            continue;
        }
    }
    let res = left.concat(right);
    console.log(res.join(' '));
}
