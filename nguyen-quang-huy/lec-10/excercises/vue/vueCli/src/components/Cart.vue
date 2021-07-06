<template>
  
    <div class="row">
      <img :src="product.src" />
      <h4>{{ product.name }}</h4>
      <p>{{ product.price }}</p>
      <div class="qty-minus" v-on:click="minus(product)">-</div>
      <div class="qty">{{ product.quantity }}</div>
      <div class="qty-plus" v-on:click="plus(product)">+</div>
      <div class="del" v-on:click="remove(product)">Remove</div>
      <div class="total-price">$ {{ product.total }}</div>
    </div>
</template>

<script>
export default {
  name: "Cart",
  props: {
    product: {
      type: Object,
      default: () => {}
    },
    minus: {
      type: Function
    },
    plus: {
      type: Function
    },
    remove:{
      type: Function
    }
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
