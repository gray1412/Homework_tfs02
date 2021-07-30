<template>
  <div id="app">
    <div class="main">
      <div class="product">
        <Products
          v-for="item in products"
          :product="item"
          :key="item.id"
          
          :addToCart="addToCart"
        />
      </div>
      
      <div class="cart">
          <div class="head">
        <h3>Shopping cart</h3>
        <div class="price">Price</div>
        <div class="quantity">Quantity</div>
        <div class="total">Total</div>
      </div>
        <Cart
          v-for="item in cart"
          :product="item"
          :key="item.id"
          
          :remove="remove"
          :minus="minus"
          :plus="plus"
        />
      </div>
      <h5>Subtotal amount: $ {{ subTotal }}</h5>
    </div>
  </div>
</template>

<script>
import Products from "./components/Products.vue";
import Cart from "./components/Cart.vue";

export default {
  name: "App",
  components: {
    Products,
    Cart,
  },
  data() {
    return {
      products: [
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
      cart: [],
    //   totalPrice: 0,
    };
  },
  computed: {
            subTotal() {
                let total = 0;
                this.cart.forEach((item) => {
                    total += (item.price * item.quantity);
                })
                return total;
            }
        },
  methods: {
    addToCart(item) {
      const index = this.cart.findIndex(({ id }) => item.id === id);
      if (index === -1) {
        this.cart.push(Object.assign({}, item, { quantity: 1, total: item.price }));
      } else {
        this.cart[index].quantity++;
        this.cart[index].total = this.cart[index].quantity * this.cart[index].price; 
      }
    },
    minus(item) {
      const index = this.cart.findIndex(({ id }) => item.id === id);
      if (this.cart[index].quantity > 1) {
        this.cart[index].quantity--;
        this.cart[index].total = this.cart[index].quantity * this.cart[index].price; 
      }
    },
    plus(item) {
      const index = this.cart.findIndex(({ id }) => item.id === id);
      this.cart[index].quantity++;
      this.cart[index].total = this.cart[index].quantity * this.cart[index].price; 
    },
    remove(item) {
      const index = this.cart.findIndex(({ id }) => item.id === id);
      this.cart.splice(index, 1);
    },
  }
};
</script>

<style>
html,
body {
  background-color: #eee;
  font-family: calibri, sans-serif;
}
.test {
  color: red;
}
.main {
  width: 760px;
  margin: 20px auto;
}
.main .product .box {
  width: 230px;
  background-color: #fff;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
  display: inline-block;
  margin: 0 10px;
  position: relative;
}
.main .product .box img {
  width: 230px;
}
.main .product .box i {
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
.main .product .box i:hover {
  transform: scale(1.05);
}
.main .product .box h2 {
  margin-left: 20px;
}
.main .product .box p {
  margin-left: 20px;
}
.main .cart {
  margin-top: 50px;
  overflow: hidden;
}
.main .cart .head {
  width: 100%;
  border-bottom: 1px solid #bfbfbf;
  height: 40px;
  display: block;
}
.main .cart .head h3 {
  display: inline-block;
  line-height: 40px;
  margin: 0;
}
.main .cart .head .price {
  display: inline-block;
  color: #777777;
  margin-left: 200px;
  line-height: 40px;
}
.main .cart .head .quantity {
  display: inline-block;
  color: #777777;
  margin-left: 100px;
  line-height: 40px;
}
.main .cart .head .total {
  display: inline-block;
  color: #777777;
  line-height: 40px;
  float: right;
}
.main .cart .row {
  width: 100%;
  border-bottom: 1px solid #bfbfbf;
  overflow: hidden;
  padding: 10px 0;
  display: block;
}
.main .cart .row img {
  height: 100px;
  float: left;
}
.main .cart .row h4 {
  float: left;
  line-height: 100px;
  margin: 0 0 0 20px;
  width: 100px;
}
.main .cart .row p {
  float: left;
  width: 80px;
  line-height: 100px;
  margin: 0 0 0 35px;
  text-align: center;
}
.main .cart .row .qty-minus {
  float: left;
  width: 20px;
  line-height: 100px;
  margin-left: 60px;
  text-align: center;
  cursor: pointer;
}
.main .cart .row .qty {
  float: left;
  width: 20px;
  line-height: 100px;
  margin-left: 20px;
  text-align: center;
}
.main .cart .row .qty-plus {
  float: left;
  width: 20px;
  line-height: 100px;
  margin-left: 20px;
  text-align: center;
  cursor: pointer;
}
.main .cart .row .del {
  float: left;
  width: 80px;
  line-height: 100px;
  margin-left: 60px;
  cursor: pointer;
  text-decoration: underline;
  color: #ed277f;
}
.main .cart .row .total-price {
  float: left;
  width: 80px;
  line-height: 100px;
  margin-left: 10px;
  text-align: right;
}
.main .cart h5 {
  font-size: 1.2rem;
  text-align: right;
}
</style>
