<template>
  <div class="main">
    <div class="product">
      <Product
        v-for="item in products"
        :product="item"
        :key="item.id"
        :addToCart="addToCart"
      >
      </Product>
    </div>
    <div class="cart">
      <div class="head">
        <h3>Shopping Cart</h3>
        <div class="price">Price</div>
        <div class="quantity">Quantity</div>
        <div class="total">Total</div>
      </div>

      <Cart
        v-for="item in cart"
        :cart="item"
        :key="item.id"
        :minus="minus"
        :plus="plus"
        :remove="remove"
      ></Cart>
      <h5>Subtotal amount: $ {{ total }}</h5>
    </div>
  </div>
</template>
<script>
import Product from "../components/common/Product";
import Cart from "../components/common/Cart";
export default {
  name: "Shopping",
  components: {
    Product,
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
    };
  },
  methods: {
    addToCart(item) {
      const index = this.cart.findIndex(({ id }) => item.id === id);
      if (index === -1) {
        this.cart.push(Object.assign({}, item, { quantity: 1 }));
      } else {
        this.cart[index].quantity++;
      }
      console.log(this.cart);
    },
    minus(item){
        const index = this.cart.findIndex(({ id }) => item.id === id);
        if(this.cart[index].quantity>1){
            this.cart[index].quantity--;
        }
    },
    plus(item){
        const index = this.cart.findIndex(({ id }) => item.id === id);
        this.cart[index].quantity++;
    },
    remove(item){
        const index = this.cart.findIndex(({ id }) => item.id === id);
         this.cart.splice(index, 1);
    }
  },
  computed: {
    total() {
      let total = 0;
      // Reduce
      this.cart.forEach((item) => {
        total += item.price * item.quantity;
      });
  
      return total;
    },
  },
};
</script>
<style lang="scss">
@import "../assets/scss/shopping.scss";

</style>
