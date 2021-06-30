<template>
	<div class="product-container">
		<img :src="product.img" />
		<i @click="$emit('add-cart', product.id)" class="fa fa-plus"></i>
		<h2>{{ product.title }}</h2>
		<p>$ {{ product.price }}</p>
	</div>
</template>

<script>
export default {
	name: "Product",
	props: {
		product: Object,
	},
};
</script>

<style scoped>
.product-container {
	width: 230px;
	background-color: #fff;
	box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
	display: inline-block;
	margin: 0 10px;
	position: relative;
}
.product-container img {
	width: 230px;
}
.product-container i {
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
.product-container i:hover {
	transform: scale(1.05);
}
.product-container h2 {
	margin-left: 20px;
}
.product-container p {
	margin-left: 20px;
}
</style>
