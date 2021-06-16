let buts = document.getElementsByClassName("button");
let datashow = document.getElementById("screen");

let theFirstInput = true;
let url = "http://localhost:5500/calc";


function ExpressionToQuery(expression){
    let query = "";
    let op = "";
    let indexOfOp = 0;
    for (let i = 0; i < expression.length; i++){
        let a = expression.charCodeAt(i);
        if ( a < 48 || a > 57){
            indexOfOp = i;
            break;
        }
    }
    switch (expression[indexOfOp]){
        case '+':
            op = "add";
            break;
        case '-':
            op = 'sub';
            break;
        case '*':
            op = 'mul';
            break;
        case '/':
            op = 'div';
            break;
        default:
            op = 'none'
    }
    query = "?s="+ op + "&a=" + expression.slice(0, indexOfOp) + "&b=" +expression.slice(indexOfOp+1, expression.length);
    return query;
}

function CallApi(query){
    return fetch(url+query).then(r => r.json());
}

for (let i = 0; i < buts.length; i++ ){
    if (buts[i].id == "AC"){
        theFirstInput = true;
        buts[i].onclick = function(){
            datashow.innerHTML = "0";
        }
    } else{
        if (buts[i].id == "="){
            buts[i].onclick = function(){
                theFirstInput = true;
                console.log(url+ExpressionToQuery(datashow.innerHTML));
                
                let r = CallApi(ExpressionToQuery(datashow.innerHTML));
                r.then(data => {
                    console.log(data.Result);
                    datashow.innerHTML = data.Result;
                })
            } 
        } else{
            buts[i].onclick = function(){
                if (theFirstInput){
                    datashow.innerHTML = buts[i].id;
                    theFirstInput = false;
                } else{
                    datashow.innerHTML = datashow.innerHTML + buts[i].id;
                }               
            }
        }
    }
}
