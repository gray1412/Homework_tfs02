import { add, sub } from "./modules/lib.js";
import mul from "./modules/lib.js";

const num1 = 1;
const num2 = 2;

let sum = add(num1, num2);
let res = sub(num1, num2);

console.log(`The sum of ${num1} and ${num2} is ${sum}.`);
console.log(`The sub of ${num1} and ${num2} is ${res}.`);
console.log(`The mul of ${num1} and ${num2} is ${mul(num1, num2)}.`);
