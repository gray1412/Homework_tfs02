<template>
	<div class="row">
		<img :src="item.img" />
		<h4>{{ item.title }}</h4>
		<p>${{ item.price }}</p>
		<div @click="$emit('minus', item)" class="qty-minus">-</div>
		<div class="qty">{{ item.quantity }}</div>
		<div @click="$emit('plus', item)" class="qty-plus">+</div>
		<div @click="$emit('remove', item)" class="del">
			Remove
		</div>
		<div class="total-price">{{ item.total }}</div>
	</div>
</template>

<script>
export default {
	name: "CartItem",
	props: {
		item: Object,
	},
};
</script>

<style scoped>
.row {
	width: 100%;
	border-bottom: 1px solid #bfbfbf;
	overflow: hidden;
	padding: 10px 0;
	display: block;
}
.row img {
	height: 100px;
	float: left;
}
.row h4 {
	float: left;
	line-height: 100px;
	margin: 0 0 0 20px;
	width: 100px;
}
.row p {
	float: left;
	width: 80px;
	line-height: 100px;
	margin: 0 0 0 35px;
	text-align: center;
}
.row .qty-minus {
	float: left;
	width: 20px;
	line-height: 100px;
	margin-left: 60px;
	text-align: center;
	cursor: pointer;
}
.row .qty {
	float: left;
	width: 20px;
	line-height: 100px;
	margin-left: 20px;
	text-align: center;
}
.row .qty-plus {
	float: left;
	width: 20px;
	line-height: 100px;
	margin-left: 20px;
	text-align: center;
	cursor: pointer;
}
.row .del {
	float: left;
	width: 80px;
	line-height: 100px;
	margin-left: 60px;
	cursor: pointer;
	text-decoration: underline;
	color: #ed277f;
}
.row .total-price {
	float: left;
	width: 80px;
	line-height: 100px;
	margin-left: 10px;
	text-align: right;
}
</style>
