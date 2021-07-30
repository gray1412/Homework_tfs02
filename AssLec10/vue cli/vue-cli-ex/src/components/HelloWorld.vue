<template>
  <div class="hello">
    <chart :chart-data="datacollection"></chart>
  </div>
</template>

<script>
import axios from "axios";
// import VueAxios from "vue-axios";
// Vue.use(VueAxios, axios);

export default {
  name: 'HelloWorld',
  components:{
    Chart
  },
  data() {
    return {
      datacollection: null,
    };
  },
  methods: {
    fillData() {
          this.datacollection = {
            labels: [
              "January",
              "February",
              "March",
              "April",
              "May",
              "June",
              "July",
              "August",
              "September",
              "October",
              "November",
              "December"
            ],
            datasets: [
              {
                label: "GitHub Commits",
                backgroundColor: "#f87979",
                data: [70, 20, 12, 39, 100, 40, 95, 80, 80, 20, 12, 101]
              },
              {
                label: "Monthly incomes",
                backgroundColor: "#A5CC82",
                data: [205, 408, 188, 190, 58, 200, 190, 400, 164, 254, 290, 201]
              }
            ]
          };
    }
  },
  mounted() {
    this.fillData();
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
