const METHOD ={
    Get: "GET",
    Post: "POST",
    Put: "PUT",
    Delete: "DELETE",
}
const ContentType ={
    Json: 'application/json',
    Text: 'text',
    Abc: 'abc'
}
const ResponseTypee = {
    Json: 'json',
    Text: 'test',
    Blob: 'blob',
}

function parseRequestOptions(rawOptions = {}, content_type = ContentType.Json) {
    const options = {...rawOptions};
    if (!options.headers){
        options.headers = {
            'Content-Type': content_type,
        }
    }
    if (options.body){
        options.body = JSON.stringify(options.body);
    }
    return options;
}
function parseResponseOptions(response, resolve, responseType = ResponseTypee.Json) {
    if (response.ok) {
        switch (responseType) {
            case ResponseTypee.Json:
                response.json().then((data) => {
                    resolve(data);
                });
                break;
            case ResponseTypee.Blob:
                response.blob().then((data) => {
                    resolve(data);
                });
                break;
            default:
                response.text().then((data) => {
                    resolve(data);
                });   
        }
        return;
    }
    return;
}



function createHttp(){
    const http = {
        request(url, option = {}) {
            return new Promise((resolve, reject) => {
                fetch(url, option)
                .then((response) => parseResponseOptions(response, resolve, ResponseTypee.Text))
                    // if (response.ok){
                    //     response.json().then((data) => {
                    //         resolve(data)
                    //     })
                    // }
                
                .catch((e) => reject(e))
            })
        },
        get(url, option ={}){
            return http.request(url, Object.assign(option, { method: METHOD.Get}))
        },
        post(url, option ={}){
            return http.request(url, Object.assign(option, { method: METHOD.Post}))
        },
        put(url, option = {}){
            return http.request(url, Object.assign(option, { method: METHOD.Put}))
        },
        delete(url, option = {}){
            return http.request(url, Object.assign(option, { method: METHOD.Delete}))
        }
    };
    return http;
}
const subButton = document.getElementById('submit');
const findButton = document.getElementById('find');
const uploadButton = document.getElementById('upload');
const result = document.getElementById('response');


findButton.onclick = function() {
    const http = createHttp();
    http.get('http://localhost:8080/students', 
        parseRequestOptions({}, ContentType.Json)
    )
    .then((data) => {
        result.innerHTML = JSON.stringify(data)
    }).catch(e => {
        result.innerHTML = e
    });
}

subButton.onclick = function(){
    result.innerHTML = "loading...";
    const http = createHttp();

    const body = {
        "id": parseInt(document.getElementsByTagName("input")[0].value),
        "name": document.getElementsByTagName("input")[1].value,
        "age": parseInt(document.getElementsByTagName("input")[2].value),
    };
    const header = {
        'Content-Type': 'application/json',
    }
    const op = {
        header,
        body,
    }
    http.post('http://localhost:8080/students', 
        parseRequestOptions(op)
    )
    .then((data) => {
        result.innerHTML = JSON.stringify(data)
    }).catch(e => {
        result.innerHTML = e
    })
}
uploadButton.onclick = function() {
    const formData = new FormData();
    formData.append('image', image.files[0]);
    const http = createHttp();
    http.post('http://localhost:8080/students', {
        body: formData
    })
    .then((data) => {
        result.innerHTML = JSON.stringify(data)
    }).catch(e => {
        result.innerHTML = e
    })
}