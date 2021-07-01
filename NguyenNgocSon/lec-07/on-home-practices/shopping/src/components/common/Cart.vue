<template>
  <div class="row">
    <img :src="cart.src" :alt="cart.name" />
    <h4>{{ cart.name}}</h4>
    <p>${{ cart.price }}</p>
    <div class="qty-minus" v-on:click="minus(cart)">-</div>
    <div class="qty">{{ cart.quantity }}</div>
    <div class="qty-plus" v-on:click="plus(cart)">+</div>
    <div class="del" v-on:click="remove(cart)">Remove</div>
    <div class="total-price">${{ cart.price * cart.quantity }}</div>
  </div>
</template>
<script>
export default{
    name: "Cart",
    props: {
        cart: {
            type: Object,
            default: ()=>{}
        },
        minus: {
            type: Function
        },
        plus: {
            type: Function
        },
        remove: {
            type: Function
        }
    }
}
</script>
<style></style>
