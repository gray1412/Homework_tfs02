document.addEventListener('DOMContentLoaded', () => {
    // Init http instance
    const http = utils.createHttpInstance({ endpoint: 'http://localhost:8000/api' });
    const button = document.getElementById('btnRequest');
    const result = document.getElementById('result');
  
    button.addEventListener('click', () => {
      // Reset result
      result.innerHTML = 'Loading...';
      const body = {
       name: "Kiiwi",
       age: 22,
       phone: 123,
       gender: "Nam"
      };
      // Request
      http.post(`/members?name=${body.name}&age=${body.age}&phone=${body.phone}&gender=${body.gender}`).then((response) => {
        console.log(`res`, JSON.stringify(response));
      }).catch((e) => {
        console.log(`error`, JSON.stringify(e));
      });
    });
  });

