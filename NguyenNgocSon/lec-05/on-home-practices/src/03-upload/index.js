document.addEventListener('DOMContentLoaded', () => {
    // Init http instance
    const http = utils.createHttpInstance({ endpoint: 'https://jsonplaceholder.typicode.com' });
    const button = document.getElementById('btnRequest');
    const result = document.getElementById('result');
  
    button.addEventListener('click', () => {
      // Reset result
      result.innerHTML = 'Loading...';
  
      const image = document.getElementById('image');
      if (!image || image.files.length === 0) {
        result.innerHTML = 'Please select a file';
        return;
      }
  
      const body = new FormData();
      body.append('image', image.files[0]);
  
      // Request
      http.post('/posts', { body }).then((response) => {
        result.innerHTML = `<p>Result success: ${JSON.stringify(response)}`;
      }).catch((e) => {
        result.innerHTML = `<p>Result error: ${JSON.stringify(e)}`;
      });
    });
  });