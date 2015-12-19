var x = 100;
document.write(x);

document.write("\n\r");
var price = 44.44;
document.write(price);

var name = "liwei";
var test = "My name is liwei";

document.write(name);
document.write("\n\r");
document.write(test);;
document.write("\n\r");

var n1 = 7;
var n2 = 10;
if (n1 < n2) {
	alert("n1 is less than n2");
}

var course = 1;
if (course == 1) {
	document.write("<h1> HTML Tutorial</h1>")
} else if (course == 2) {
	document.write("<h1> CSS Tutorial</h1>")
} else {
	document.write("<h1> JavaScript Tutorial</h1>")
}

var day = 2;
switch(day) {
	case 1:
		document.write("Monday");
		break;
	case 2:
		document.write("Tusday");
		break;
	case 3:
		document.write("Wednesday");
		break;
	default:
		document.write("Another Day");
}

for (i = 1; i < 5; i++) {
	document.write(i + "<br />");
}

var j = 0;
while (j<=10) {
	document.write(j + "<br />");
	j++;
}

function myFunction() {
	document.write("<h1> Calling my function </h1>")
}
myFunction();

function sayHello(name) {
	document.write("Hello " + name)
}

sayHello("liwei")

function person(name, age) {
	this.name = name;
	this.age = age;
}

var John = new person("John", 25)
var James = new person("James", 33)

function printTime() {
	var d = new Date();
	var hours = d.getHours();
	var mins = d.getMinutes();
	var secs = d.getSeconds();
	document.body.innerHTML = hours + ":" + mins + ":" + secs;
}

printTime()
