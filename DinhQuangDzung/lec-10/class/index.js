const app = new Vue({
	el: "#app",
	data: {
		product: {
			title: "A new Product",
			price: 100,
			desc: "<p><strong>lorem ipsum</p></strong>",
			quantity: 4,
			sale: 50,
			tags: ["new", "sale", "2021"],
			options: {
				option1: 123,
				option2: 456,
				option3: 789,
			},
		},
		cart: {
			quantity: 1,
		},
	},
	methods: {
		plus() {
			this.cart.quantity++;
		},
		minus() {
			if (this.cart.quantity < 1) {
				return;
			}
			this.cart.quantity--;
		},
	},
});
