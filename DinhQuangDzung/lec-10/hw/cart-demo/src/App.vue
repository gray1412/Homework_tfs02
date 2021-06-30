<template>
	<div id="app">
		<div class="main">
			<div class="product">
				<Product
					v-for="product in products"
					:key="product.id"
					:product="product"
					@add-cart="addCart"
				></Product>
			</div>
			<Cart :items="cart.items" @minus="minus" @plus="plus" @remove="remove" />
			<h5>Subtotal amount: ${{ subTotal }}</h5>
		</div>
	</div>
</template>

<script>
import Product from "./components/Product.vue";
import Cart from "./components/Cart.vue";

export default {
	name: "App",
	components: {
		Product,
		Cart,
	},
	computed: {
		subTotal() {
			let sum = 0;
			for (let i = 0; i < this.cart.items.length; i++) {
				sum += this.cart.items[i].total;
			}
			return sum;
		},
	},
	data() {
		return {
			products: [
				{
					id: 1,
					title: "Beer Bottle",
					price: 25,
					img: "https://chenyiya.com/codepen/product-1.jpg",
				},
				{
					id: 2,
					title: "Eco Bag",
					price: 73,
					img: "https://chenyiya.com/codepen/product-2.jpg",
				},
				{
					id: 3,
					title: "Paper Bag",
					price: 35,
					img: "https://chenyiya.com/codepen/product-3.jpg",
				},
			],
			cart: {
				items: [],
			},
		};
	},
	methods: {
		addCart(id) {
			const product = this.products.find((product) => product.id == id),
				cartItems = this.cart.items;

			// Check if item is already in cart
			if (cartItems.some((item) => item.id == id)) {
				// Increase qty by 1 and add to total
				cartItems.map((item) => {
					if (item.id == product.id) {
						item.quantity++;
						item.total = item.price * item.quantity;
						return item;
					}
					return item;
				});
				// Return if item is not in cart
				return;
			}
			// Add new item to cart
			const item = Object.assign({}, product, {
				quantity: 1,
				total: product.price,
			});
			cartItems.push(item);
		},

		plus(index) {
			const item = this.cart.items[index];
			item.quantity++;
			item.total += item.price;
		},

		minus(index) {
			const item = this.cart.items[index];
			if (item.quantity == 1) {
				return;
			}
			item.quantity--;
			item.total -= item.price;
		},

		remove(index) {
			this.cart.items.splice(index, 1);
		},
	},
};
</script>

<style scoped>
@media (max-width: 800px) {
	.product {
		display: grid;
		justify-content: center;
		gap: 20px;
	}
}

.main {
	width: 760px;
	margin: 20px auto;
}

h5 {
	font-size: 1.2rem;
	text-align: right;
}
</style>
