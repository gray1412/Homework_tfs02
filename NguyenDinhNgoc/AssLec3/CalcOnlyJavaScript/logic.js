let buts = document.getElementsByClassName("button");
let datashow = document.getElementById("screen");

let theFirstInput = true;

for (let i = 0; i < buts.length; i++ ){
    if (buts[i].id == "AC"){
        theFirstInput = true;
        buts[i].onclick = function(){
            datashow.innerHTML = "0"
        }
    } else{
        if (buts[i].id == "="){
            buts[i].onclick = function(){
                try{
                    datashow.innerHTML = eval(datashow.innerHTML);                   
                } catch (ex){
                    datashow.innerHTML = "Error";
                }
                theFirstInput = true;
            } 
        } else{
            buts[i].onclick = function(){
                if (theFirstInput){
                    datashow.innerHTML = buts[i].id;
                    theFirstInput = false;
                } else{
                    datashow.innerHTML = datashow.innerHTML + buts[i].id
                }               
            }
        }
    }
}