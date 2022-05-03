// FizzBuzz = Sisa bagi 3 dan 5
// Fizz = Sisa bagi 3
// Buzz = Sisa bagi 5

for (i = 0; i <= 100; i++) {
    if (i % 15 == 0) {
        console.log("FizzBuzz")
    } else if (i % 3 == 0) {
        console.log("Fizz")
    } else if (i % 5 == 0) {
        console.log("Buzz")
    } else {
        console.log(i)
    }
}

// Run
// node fizzbuzz.js