<template>
        <div class="box">
          <img :src="product.src" :alt="product.name"/>
          <i class="fa fa-plus" v-on:click="addToCart(product)">+</i>
          <h2>{{product.name}}</h2>
          <p>${{product.price}}</p>
        </div>
</template>

<script>
    export default{
        name: "Product",
        props: {
            product: {
                type: Object,
                default: ()=>{}
            },
            addToCart: {
                type: Function
            }
        }
    }
</script>