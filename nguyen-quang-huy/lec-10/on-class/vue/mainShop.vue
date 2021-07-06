<template>
  <div id="app">
    <div class="main">
      <div class="product">
        <div class="box" v-for="item in list_pro" v-bind:key="item.name">
          <img v-bind:src="item.src" />
          <i class="fa fa-flus" v-on:click="addItem(item.id)">+</i>
          <h2>{{ item.name }}</h2>
          <p>$ {{ item.price }}</p>
        </div>
      </div>

      <div class="cart">
        <div class="head">
          <h3>Shopping cart</h3>
          <div class="price">Price</div>
          <div class="quantity">Quantity</div>
          <div class="total">Total</div>
        </div>
        <div class="row" v-for="item in itemAdded" v-bind:key="item.name">
          <img v-bind:src="item.src" />
          <h4>{{ item.name }}</h4>
          <p>{{ item.price }}</p>
          <div class="qty-minus" v-on:click="minus(item.id)">-</div>
          <div class="qty">{{ item.quantity }}</div>
          <div class="qty-plus" v-on:click="plus(item.id)">+</div>
          <div class="del" v-on:click="del(item.id)">Remove</div>
          <div class="total-price">$ {{ item.total }}</div>
        </div>
        <h5>Subtotal amount: $ {{ totalPrice }}</h5>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      list_pro: [
        {
          id: 0,
          name: "Beer Bottle",
          price: 25,
          src: "https://chenyiya.com/codepen/product-1.jpg",
        },
        {
          id: 1,
          name: "Eco Bag",
          price: 73,
          src: "https://chenyiya.com/codepen/product-2.jpg",
        },
        {
          id: 2,
          name: "Paper Bag",
          price: 35,
          src: "https://chenyiya.com/codepen/product-3.jpg",
        },
      ],
      itemAdded: [],
      totalPrice: 0,
    };
  },
  methods: {
    addItem(id) {
      let check = 0;
      for (var i = 0; i < this.itemAdded.length; i++) {
        if (this.itemAdded[i].id == id) {
          this.itemAdded[i].quantity++;
          this.itemAdded[i].total =
            this.itemAdded[i].quantity * this.itemAdded[i].price;
          check = 1;
        }
      }
      if (check == 0) {
        let newItem = {
          id: id,
          name: this.list_pro[id].name,
          price: this.list_pro[id].price,
          src: this.list_pro[id].src,
          quantity: 1,
          total: this.list_pro[id].price,
        };
        this.itemAdded.push(newItem);
      }
      this.totalPrice = 0
      for (var j = 0; j < this.itemAdded.length; j++) {
        this.totalPrice += this.itemAdded[j].total
      }
    },
    minus(id) {
      for (let i = 0; i < this.itemAdded.length; i++) {
        if (this.itemAdded[i].id == id) {
          if (this.itemAdded[i].quantity > 1) {
            this.itemAdded[i].quantity--;
            this.itemAdded[i].total =
              this.itemAdded[i].quantity * this.itemAdded[i].price;
          }
        }
      }
      this.totalPrice = 0
      for (var j = 0; j < this.itemAdded.length; j++) {
        this.totalPrice += this.itemAdded[j].total
      }
    },
    plus(id) {
      for (let i = 0; i < this.itemAdded.length; i++) {
        if (this.itemAdded[i].id == id) {
          this.itemAdded[i].quantity++;
          this.itemAdded[i].total =
            this.itemAdded[i].quantity * this.itemAdded[i].price;
        }
      }
      this.totalPrice = 0
      for (var j = 0; j < this.itemAdded.length; j++) {
        this.totalPrice += this.itemAdded[j].total
      }
    },
    del(id) {
      for (let i = 0; i < this.itemAdded.length; i++) {
        if (this.itemAdded[i].id == id) {
          this.itemAdded.splice(i, 1);
        }
      }
      this.totalPrice = 0
      for (var j = 0; j < this.itemAdded.length; j++) {
        this.totalPrice += this.itemAdded[j].total
      }
    },
  },
};
</script>

<style scoped>
#app {
  background-color: #eee;
  font-family: calibri, sans-serif;
}
.main {
  width: 760px;
  margin: 20px auto;
}
.box {
  width: 230px;
  background-color: #fff;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
  display: inline-block;
  margin: 0 10px;
  position: relative;
}
.box > img {
  width: 230px;
}
.box > i {
  width: 50px;
  height: 50px;
  background: #ed277f;
  color: #ffffff;
  border-radius: 25px;
  text-align: center;
  line-height: 50px;
  font-size: 1.4rem;
  position: absolute;
  right: 20px;
  top: 150px;
  box-shadow: 0 0 4px 2px rgba(80, 80, 80, 0.1);
  cursor: pointer;
  transition: all 0.3s;
}
.box > h2 {
  margin-left: 20px;
}
.box > p {
  margin-left: 20px;
}
.cart {
  margin-top: 50px;
  overflow: hidden;
}
.head {
  width: 100%;
  border-bottom: 1px solid #bfbfbf;
  height: 40px;
  display: block;
}
.head > h3 {
  display: inline-block;
  line-height: 40px;
  margin: 0;
}
.price {
  display: inline-block;
  color: #777777;
  margin-left: 200px;
  line-height: 40px;
}
.quantity {
  display: inline-block;
  color: #777777;
  margin-left: 100px;
  line-height: 40px;
}
.total {
  display: block;
  color: #777777;
  line-height: 40px;
  float: right;
}
.row {
  width: 100%;
  border-bottom: 1px solid #bfbfbf;
  overflow: hidden;
  padding: 10px 0;
  display: block;
}
.row > img {
  height: 100px;
  float: left;
}
.row > h4 {
  float: left;
  line-height: 100px;
  margin: 0 0 0 20px;
  width: 100px;
}
.row > p {
  float: left;
  width: 80px;
  line-height: 100px;
  margin: 0 0 0 35px;
  text-align: center;
}
.qty-minus {
  float: left;
  width: 20px;
  line-height: 100px;
  margin-left: 60px;
  text-align: center;
  cursor: pointer;
}
.qty {
  float: left;
  width: 20px;
  line-height: 100px;
  margin-left: 20px;
  text-align: center;
}
.qty-plus {
  float: left;
  width: 20px;
  line-height: 100px;
  margin-left: 20px;
  text-align: center;
  cursor: pointer;
}
.del {
  float: left;
  width: 80px;
  line-height: 100px;
  margin-left: 60px;
  cursor: pointer;
  text-decoration: underline;
  color: #ed277f;
}
.total-price {
  float: left;
  width: 80px;
  line-height: 100px;
  margin-left: 10px;
  text-align: right;
}
h5 {
  font-size: 1.2rem;
  text-align: right;
}
</style>