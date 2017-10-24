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
    let count = new Array(100);
    count.fill(0);
    let values = Array.apply(null, Array(100));
    values = values.map(function () { return new Array(); });
    // calculate the histogram of key frequencies:    
    for (let i=0; i < n; i++){
        let ar = readLine().split(' ');
        let key = parseInt(ar[0]);
        count[key]++;
        if (i < n/2) {
            values[key].push('-');    
        } else
        {
            values[key].push(ar[1]);
        }
    }
    // calculate the starting index for each key:
    let output = new Array();
    for(let i=0; i<100; i++) {
        let cnt = count[i];
        // for the number of times the item occurs
        for(let j=0;j<cnt;j++){
            output.push(values[i].shift());
        }
    }
    console.log(output.join(' '));   
}
