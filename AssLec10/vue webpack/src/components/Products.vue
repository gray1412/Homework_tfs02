<template>
  <div class="box">
   
      <img :src="product.src" />
      <i class="fa fa-flus" v-on:click="addToCart(product)">+</i>
      <h2>{{ product.name }}</h2>
      <p>$ {{ product.price }}</p>
    
  </div>
</template>

<script>
export default {
  name: "Products",
  props: {
    product: {
      type: Object,
      default: () => {}
    },
    addToCart:{
      type: Function,
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
