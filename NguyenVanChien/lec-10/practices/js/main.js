var vm = new Vue ({
    el: '#app',
    // computed: {
    //     subTotal() {
    //         let total = 0
    //         this.itemAdded.forEach((item) => {
    //             total += (item.quantity)*(item.price);
    //         });
    //         return total
    //     }
    // },

    data: {
        list_pro: [
            {
              id: 0,
              name: "Beer Bottle",
              price: 25,
              src: "https://chenyiya.com/codepen/product-1.jpg",
            },
            {
              id: 1,
              name: "Eco Bag",
              price: 73,
              src: "https://chenyiya.com/codepen/product-2.jpg",
            },
            {
              id: 2,
              name: "Paper Bag",
              price: 35,
              src: "https://chenyiya.com/codepen/product-3.jpg",
            },
          ],
          itemAdded: [],
          totalPrice: 0,
        },


    methods: {
        addItem(id) {
          let check = 0;
          for (var i = 0; i < this.itemAdded.length; i++) {
            if (this.itemAdded[i].id == id) {
              this.itemAdded[i].quantity++;
              this.itemAdded[i].total =
                this.itemAdded[i].quantity * this.itemAdded[i].price;
              check = 1;
            }
          }
          if (check == 0) {
            let newItem = {
              id: id,
              name: this.list_pro[id].name,
              price: this.list_pro[id].price,
              src: this.list_pro[id].src,
              quantity: 1,
              total: this.list_pro[id].price,
            };
            this.itemAdded.push(newItem);
          }
          this.totalPrice = 0
          for (var j = 0; j < this.itemAdded.length; j++) {
            this.totalPrice += this.itemAdded[j].total
          }
        },
        minus(id) {
          for (let i = 0; i < this.itemAdded.length; i++) {
            if (this.itemAdded[i].id == id) {
              if (this.itemAdded[i].quantity > 1) {
                this.itemAdded[i].quantity--;
                this.itemAdded[i].total =
                  this.itemAdded[i].quantity * this.itemAdded[i].price;
              }
            }
          }
          this.totalPrice = 0
          for (var j = 0; j < this.itemAdded.length; j++) {
            this.totalPrice += this.itemAdded[j].total
          }
        },
        plus(id) {
          for (let i = 0; i < this.itemAdded.length; i++) {
            if (this.itemAdded[i].id == id) {
              this.itemAdded[i].quantity++;
              this.itemAdded[i].total =
                this.itemAdded[i].quantity * this.itemAdded[i].price;
            }
          }
          this.totalPrice = 0
          for (var j = 0; j < this.itemAdded.length; j++) {
            this.totalPrice += this.itemAdded[j].total
          }
        },
        del(id) {
          for (let i = 0; i < this.itemAdded.length; i++) {
            if (this.itemAdded[i].id == id) {
              this.itemAdded.splice(i, 1);
            }
          }
          this.totalPrice = 0
          for (var j = 0; j < this.itemAdded.length; j++) {
            this.totalPrice += this.itemAdded[j].total
          }
        },
      },
})

// Vue.component('Product', {
//   data() {
//     return {
//       product: {
//         id: 0,

//       }
//     }
//   }
// })