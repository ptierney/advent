(() => {
    let aoc = require('../lib/aoc.js');
    let links = aoc.inputfile('./day05.txt', true).map( Number );

    let current = 0;
    let instruction = 0;
    for (let x = 1 ; x < 30000000 ; x++ ) {
        instruction = links[current];
        links[current] += (instruction >= 3)? -1 : 1;
        current += instruction;
        if (current < 0 || current >= (links.length)) {
            console.log(x);
            break;
        }
    }
})()