<template>
	<div class="cart">
		<div class="head">
			<h3>Shopping Cart</h3>
			<div class="price">Price</div>
			<div class="quantity">Quantity</div>
			<div class="total">Total</div>
		</div>
		<CartItem
			v-for="item in items"
			:key="item.id"
			:item="item"
			@minus="$emit('minus', items.indexOf(item))"
			@plus="$emit('plus', items.indexOf(item))"
			@remove="$emit('remove', items.indexOf(item))"
		/>
	</div>
</template>

<script>
import CartItem from "./CartItem.vue";

export default {
	name: "Cart",
	props: {
		items: Array,
	},
	components: {
		CartItem,
	},
};
</script>

<style scoped>
.cart {
	margin-top: 50px;
	overflow: hidden;
}
.cart .head {
	width: 100%;
	border-bottom: 1px solid #bfbfbf;
	height: 40px;
	display: block;
}
.cart .head h3 {
	display: inline-block;
	line-height: 40px;
	margin: 0;
}
.cart .head .price {
	display: inline-block;
	color: #777777;
	margin-left: 200px;
	line-height: 40px;
}
.cart .head .quantity {
	display: inline-block;
	color: #777777;
	margin-left: 100px;
	line-height: 40px;
}
.cart .head .total {
	display: inline-block;
	color: #777777;
	line-height: 40px;
	float: right;
}
</style>
