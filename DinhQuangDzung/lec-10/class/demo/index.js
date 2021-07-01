Vue.component("Product", {
	data() {
		return {
			product: {
				title: "Beer Bottle",
				price: 25,
				img: "https://chenyiya.com/codepen/product-1.jpg",
			},
		};
	},
	template: `
  <div class="box">
		<img :src="product.img" />
		<i @click="add" class="fa fa-plus"></i>
		<h2>{{product.title}}</h2>
		<p>$ {{product.price}}</p>
	</div>`,
});

const app = new Vue({
	el: "#app",
	computed: {
		subTotal() {
			let sum = 0;
			for (let i = 0; i < this.cart.items.length; i++) {
				sum += this.cart.items[i].total;
			}
			return sum;
		},
	},
	data: {},

	methods: {
		add(e) {
			const index = e.target.parentNode.getAttribute("id");
			if (
				this.cart.items.some((item) => item.title == this.products[index].title)
			) {
				// console.log("already in cart");
				this.cart.items.map((item) => {
					if (item.title == this.products[index].title) {
						item.quantity++;
						item.total = item.price * item.quantity;
						return item;
					}
					return item;
				});

				return;
			}
			this.cart.items.push({
				title: this.products[index].title,
				price: this.products[index].price,
				img: this.products[index].img,
				quantity: 1,
				total: this.products[index].price,
			});
		},

		plus(e) {
			const index = e.target.parentNode.getAttribute("index");
			this.cart.items[index].quantity++;
			this.cart.items[index].total += this.cart.items[index].price;
		},
		minus(e) {
			const index = e.target.parentNode.getAttribute("index");
			if (this.cart.items[index].quantity == 1) {
				return;
			}
			this.cart.items[index].quantity--;
			this.cart.items[index].total -= this.cart.items[index].price;
		},
		remove(e) {
			const index = e.target.parentNode.getAttribute("index");
			this.cart.items.splice(index, 1);
		},
	},
});
