function inputData(data) {
	let content  = document.getElementById("content").innerHTML;
	document.getElementById("content").innerHTML = content+data+" ";	
}

function equal(){
	let data = document.getElementById("content").innerHTML;
	let res = data.split(" ");
	let a = res[0];
	let op = res[1];
	let b = res[2];
	
	switch (op){
		case '+':
		op = "sum";
		break;
		case '-':
		op = "sub";
		break;
		case '*':
		op = "mul";
		break;
		case '/':
		op = "div";
		break;
	}
	fetch("http://localhost:8000/?type="+op+"&a="+a+"&b="+b+"")
	.then(response => response.json())
	.then(data => document.getElementById("content").innerHTML = data.Result);
}


