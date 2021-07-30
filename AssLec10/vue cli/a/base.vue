<template>
  <div class="hello">
    {{msg}}
    <button @click="getData">get</button>
  </div>
</template>

<script>
import axios from "axios";
// import VueAxios from "vue-axios";
// Vue.use(VueAxios, axios);

export default {
  name: 'HelloWorld',
  data() {
    return {
      msg: 'chua fetch dc',
    };
  },
  methods: {
    async getData() {
      try {
				const res = await axios({
					method: "GET",
					url: "http://localhost:8080/statistic",
					data: {
						email: "nguyendinhhpv3@gmail.com",
					},
					headers: {
						"Access-Control-Allow-Origin": "*",
            "Content-Type": "application/json",
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjcyMzc3OTYsImlzQWRtaW4iOnRydWUsInVzZXJJZCI6MSwidXNlcm5hbWUiOiJuZ29jaGQifQ.eRoyIuo-ycXTX8hZ44IJ0dzI5OLETxAvNRxLQzXVpx8"
					},
				});
				if (res) {
					this.msg = res.data;
          console.log("try", res.data)
				};
			} catch (err) {
        this.msg = err.response.data;
        console.log("catch", err.response.data)
				// alert(err.response.data.message);
			};
    }
  },
  mounted() {
    // this.getData()
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
