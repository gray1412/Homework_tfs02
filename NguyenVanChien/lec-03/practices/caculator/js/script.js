function show_pre_result(){
    return document.getElementById("pre_result").innerText;
}

function print_pre_result(so) {
    document.getElementById("pre_result").innerText = so;
}


function show_cur_result() {
    return document.getElementById("cur_result").innerText;
}
function print_cur_resulta(so) {
    if(so=="") {
        document.getElementById("cur_result").innerText = so;
    } else {
        document.getElementById("cur_result").innerText = format_str(so);
    }
 
}
function format_str(so) {
    if(so == "-") {
        return "";
    }
    var n = Number(so);
    var gia_tri = n.toLocaleString("en");
    return gia_tri;
}

function de_format_str(so) {
    return Number(so.replace(/,/g, ''))
} 


var he_thong = document.getElementsByClassName('operator');
for(var i=0; i < he_thong.length; i++) {
    he_thong[i].addEventListener('click', function() {
        if(this.id == "xoa_tat_ca") {
            print_cur_resulta("");
            print_pre_result("");
        } 
        else {
            var result = show_cur_result();
            var last_result = show_pre_result();
            if(result != "") {
                result = de_format_str(result);
                last_result =last_result + result;
                if(this.id == "=") {
                    var last_result = eval(last_result);
                    print_cur_resulta(last_result)
                    print_pre_result("")
                } else {
                    last_result = last_result + this.id;
                    print_pre_result(last_result)
                    print_cur_resulta("")
                }
            }
        }
    })
}

var chu_so = document.getElementsByClassName('number');
for(var i=0; i < chu_so.length; i++) {
    chu_so[i].addEventListener('click', function() {
        var result = de_format_str(show_cur_result())
        if(result != NaN) {
            result = result + this.id;
            print_cur_resulta(result)
        }
    })
}