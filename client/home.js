let x = 5;
let y = 3;

function add(x, y) {
	return x + y;
}

console.log(add(x, y));


const car = {
	model: "honda",
	year: 2018,
	make: "car",
	get model1() { return this.model; },
	get make1() { return this.make; },
	get year1() { return this.year; },

};
document.getElementById("demo").innerHTML = `Model: ${car.model1}` +  " " +  `Year: ${car.year1}`;
