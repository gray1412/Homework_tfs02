const app = new Vue({
	el: "#app",
	data: {
		products: [
			{
				title: "Beer Bottle",
				price: 25,
				img: "https://chenyiya.com/codepen/product-1.jpg",
			},
			{
				title: "Eco Bag",
				price: 73,
				img: "https://chenyiya.com/codepen/product-2.jpg",
			},
			{
				title: "Paper Bag",
				price: 35,
				img: "https://chenyiya.com/codepen/product-3.jpg",
			},
		],
		cart: {
			items: [],
			subTotal: 0,
		},
	},

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
						this.cart.subTotal = this.cal();
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
			this.cart.subTotal = this.cal();
		},

		plus(e) {
			const index = e.target.parentNode.getAttribute("index");
			this.cart.items[index].quantity++;
			this.cart.items[index].total += this.cart.items[index].price;
			this.cart.subTotal = this.cal();
		},
		minus(e) {
			const index = e.target.parentNode.getAttribute("index");
			if (this.cart.items[index].quantity == 1) {
				return;
			}
			this.cart.items[index].quantity--;
			this.cart.items[index].total -= this.cart.items[index].price;
			this.cart.subTotal = this.cal();
		},
		remove(e) {
			const index = e.target.parentNode.getAttribute("index");
			this.cart.items.splice(index, 1);
			this.cart.subTotal = this.cal();
		},
		cal() {
			let sum = 0;
			for (let i = 0; i < this.cart.items.length; i++) {
				sum += this.cart.items[i].total;
			}
			return sum;
		},
	},
});
