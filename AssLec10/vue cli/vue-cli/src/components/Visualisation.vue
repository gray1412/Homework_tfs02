<template>
  <div>
    <div class="data">
      <div class="top">
        <div class="revenue"> Doanh thu: {{revenue}} vnđ</div>
        <div class="totalOrder"> Tổng số đơn hàng: {{totalorder}}</div>
        <div class="totalProduct"> Tổng số sản phẩm bán ra: {{totalproduct}}</div>
      </div>
      <div class="botton">
        <button @click="fillData">Xem chi tiết</button>
      </div>
    </div>
    <div :class="{chart: chartShow, topProduct: !chartShow}">
      <div class="product">
        <div class="order">Top</div>
        <div class="nameProduct">Tên sản Phẩm</div>
        <div class="numberProduct">Số lượng bán ra</div>
      </div>
      <div class="product a">
        <div class="order">1</div>
        <div class="nameProduct">{{topProduct[0].product.name}}</div>
        <div class="numberProduct">{{topProduct[0].product_count}}</div>
      </div>
      <div class="product b">
        <div class="order">2</div>
        <div class="nameProduct">{{topProduct[1].product.name}}</div>
        <div class="numberProduct">{{topProduct[1].product_count}}</div>
      </div>
      <div class="product c">
        <div class="order">3</div>
        <div class="nameProduct">{{topProduct[2].product.name}}</div>
        <div class="numberProduct">{{topProduct[2].product_count}}</div>
      </div>
    </div>
    <div :class="{chart: chartShow, h: !chartShow}">
      <chart :chart-data="chartData"></chart>
    </div>
  </div>
</template>
    <script>
import Chart from "./../Chart.js";
import axios from "axios";
export default {
  components: {
    Chart,
  },
  data() {
    return {
      chartData: null,
      revenue: 0,
      totalorder: 0,
      totalproduct: 0,
      time: [],
      listOrderCount: [],
      listProductCount: [],
      topProduct: [],
      chartShow: true,
    };
  },
  mounted() {
    this.getData();
    // this.fillData();
  },
  methods: {
    async getData() {
      try {
        const res = await axios({
          method: "GET",
          url: "http://localhost:8080/statistic",
          data: {
          },
          headers: {
            "Access-Control-Allow-Origin": "*",
            "Content-Type": "application/json",
            token:
              "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjcyMzc3OTYsImlzQWRtaW4iOnRydWUsInVzZXJJZCI6MSwidXNlcm5hbWUiOiJuZ29jaGQifQ.eRoyIuo-ycXTX8hZ44IJ0dzI5OLETxAvNRxLQzXVpx8",
          },
        });
        if (res) {
          this.revenue = res.data.revenue
          this.totalorder = res.data.total_order
          this.totalproduct = res.data.total_product
          this.time = res.data.list_times
          this.listOrderCount = res.data.list_orders
          this.listProductCount = res.data.list_products
          this.topProduct = res.data.top_sell
          console.log(this.time);
        }
      } catch (err) {
        this.msg = err.response.data;
        console.log("catch", err.response.data);
        // alert(err.response.data.message);
      }
    },
    fillData() {
      this.chartShow = false
      this.chartData = {
        labels: this.time,
        datasets: [
          {
            label: "Số lượng đơn hàng",
            backgroundColor: "#A5CC82",
            data: this.listOrderCount,
          },
          {
            label: "Số sản phẩm",
            backgroundColor: "#008aff",
            data: this.listProductCount,
          },
        ],
      };
    },
  },
};
</script>
<style>
  .data{
    display: flex;
    flex-direction: column;
  }
  .top{
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }
  .botton{
    margin-top: 50px;
    display: flex;
    flex-direction: row;
    justify-content: center;
    
  }
  .total{
    display: flex;
    justify-content: center;
    align-items: center;
    color: #000;
    font-size: 30;
  }
  button{
  align-items: center;
   background-image: linear-gradient(135deg, #008aff, #86d472);
   border-radius: 6px;
   display: flex;
   justify-content: center;
   height: 30px;
   width: 150px;
  }
  button:hover {
    background: transparent;
  }
  .topProduct{
    margin-top: 30px ;
    margin-bottom: 30px;
    display: flex;
    flex-direction: column;
  }
  .product{
    height: auto;
    background-color: #c6d127;
    margin: 5px;
    display: flex;
    flex-direction: row;
    align-items: center;
    height: 60px;
    width: 100%;
  }
  .a{
    background-color: rgba(134, 212, 114, 1);
  }
  .b{
    background-color: rgba(134, 212, 114, 0.8);
  }
  .c{
    background-color: rgba(134, 212, 114, 0.6);
  }
  .order{
    width: 10%;
  }
  
  .nameProduct{
    width: 80%;
  }
  .numberProduct{
    width: 10%;
  }
  .chart{
    display: none;
  }
  .h{
    max-height: 500px;
  }
</style>