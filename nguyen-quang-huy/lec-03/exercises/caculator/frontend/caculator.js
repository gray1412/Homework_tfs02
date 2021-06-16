let datashow = document.getElementById("display");
function clickNumber(input){
    calculator.display.value+=input
}

displayCaculator = (rawData)=>{
    let data = rawData.split(/(\+|-|\*|\/)/)
    let query
    let url = "http://localhost:5500/calculate"
    if(data.length == 1 ){
        if(typeof +data[0] === "number") return data[0]
        return "error"
    } else if(data.length == 3) {
        if(typeof +data[0] === "number" && typeof +data[2] === "number"){

            switch(data[1]) {
                case "+":
                    op = "add"
                    break;
                case "-":
                    op = "sub"
                    break;
                case "*":
                    op = "mul"
                    break;
                case "/":
                    op = "div"
                    break;
                default:
                  return "error"
              }
            query="?op="+ op + "&a=" + data[0] + "&b=" +data[2];
            console.log(url+query)
            let res = 0
            return fetch(url+query)
                .then(response => response.json())
                .then((data) => {
                    res = data.result
                    console.log(res)
                    return res;
                });


        } else return "error"

    } else{
        return "error"
    }
}


