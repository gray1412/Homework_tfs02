<template>
  <main>
    <div class="add-multi-product">
      <div>
        <h1>Thêm nhiều sản phẩm</h1>
        <br />
        <label for="myfile">Chọn tệp csv: </label>
        <input type="file" @change="readFileUpload" />
      </div>
      <div class="btn-box">
        <button @click="submitFile">Gửi</button>
      </div>
    </div>
    <div class="add-a-product">
      <h1>Thêm một sản phẩm</h1>
      <div class="input-box">
        <label for="name">Tên sản phẩm</label>
        <br />
        <input v-model="name" type="text" id="name" placeholder="Nhập tên" />
      </div>
      <div class="detail-info">
        <div class="input-box e">
          <label for="technology">Công nghệ</label>
          <br />
          <select v-model="technology" id="technology">
            <option value="LED">LED</option>
            <option value="OLED">OLED</option>
          </select>
        </div>
        <div class="input-box e">
          <label for="resolution">Độ phân giải</label>
          <br />
          <select v-model="resolution" id="type">
            <option value="Full HD">Full HD</option>
            <option value="4k">4k</option>
            <option value="8k">8k</option>
          </select>
        </div>
        <div class="input-box e">
          <label for="type">Loại</label>
          <br />
          <select v-model="type" id="type">
            <option value="Smart tv">Smart tv</option>
            <option value="Android tv">Android tv</option>
            <option value="Google tv">Google tv</option>
          </select>
        </div>
      </div>
      <div class="detail-info-1">
        <div class="fixed">
          <div class="input-box a">
            <label for="newDescription">Chi tiết mô tả</label>
            <br />
            <input
              v-model="newDescription.content"
              type="text"
              placeholder="Thêm chi tiết"
            />
          </div>
          <div class="b">
            <button @click="ListDescriptions.push(newDescription)">Thêm</button>
          </div>
        </div>
        <div class="dynamic">
          <div
            class="view"
            v-for="(detail, index) in ListDescriptions"
            :key="index"
          >
            <span>{{ detail.content }}</span>
          </div>
        </div>
      </div>
      <div class="detail-info-1">
        <div class="fixed">
          <div class="input-box a">
            <label for="newImage">Thêm đường dẫn ảnh</label>
            <br />
            <input v-model="newImage.link" type="text" placeholder="Link" />
          </div>
          <div class="b">
            <button @click="ListImages.push(newImage)">Thêm</button>
          </div>
        </div>
        <div class="dynamic">
          <div
            class="view"
            v-for="(linkImage, index) in ListImages"
            :key="index"
          >
            <span>{{ linkImage.link }}</span>
          </div>
        </div>
      </div>
      <div class="detail-info-2">
        <div class="fixed">
          <div class="input-box d">
            <label for="newSize">Kích thước</label>
            <br />
            <input v-model="newSize" type="text" placeholder="0” (0 cm)" />
          </div>
          <div class="input-box d">
            <label for="newPrice">Giá</label>
            <br />
            <input v-model="newPrice" type="text" placeholder="0" />
          </div>
          <div class="input-box d">
            <label for="newQuantity">Số lượng</label>
            <br />
            <input v-model="newQuantity" type="text" placeholder="0" />
          </div>
          <div class="b">
            <button @click="getNewOption">Thêm</button>
          </div>
        </div>
        <div class="dynamic">
          <div class="view" v-for="(op, index) in ListOptions" :key="index">
            {{ op.size }} -- {{ op.price }} vnđ -- {{ op.quantity }} chiếc
          </div>
        </div>
      </div>
      <div class="btn-box">
        <button type="submit" @click="sendProduct">Gửi</button>
      </div>
    </div>
  </main>
</template>
<script>
import axios from "axios";
export default {
  name: "Upload",
  data() {
    return {
      filedata: null,
      name: "",
      linkDetail: "",
      technology: "",
      resolution: "",
      type: "",
      newDescription: {},
      ListDescriptions: [],
      newSize: "",
      newPrice: 0,
      newQuantity: 0,
      newImage: {},
      ListImages: [],
      newOption: {},
      ListOptions: [],
    };
  },
  methods: {
    getNewOption() {
      this.newOption.size = this.newSize;
      this.newOption.price = parseInt(this.newPrice);
      this.newOption.sale_price = parseInt(this.newPrice);
      this.newOption.quantity = parseInt(this.newQuantity);
      this.ListOptions.push(this.newOption);
    },
    readFileUpload(e) {
      const file = e.target.files[0];
      const reader = new FileReader();
      reader.onload = (e) => 
      {
        this.filedata = e.target.result;
      };
      reader.readAsText(file);
    },
     
    async submitFile() {
      // console.log(this.file)
      try {
        const res = await axios({
          method: "POST",
          url: "http://localhost:8080/products/importfile",
          data: {
            file: this.filedata
          },
          headers: {
            "Access-Control-Allow-Origin": "*",
            "Content-Type": "application/json",
            token:
              "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjczNDU4MzIsImlzQWRtaW4iOnRydWUsInVzZXJJZCI6MSwidXNlcm5hbWUiOiJuZ29jaGQifQ.ZKLPL5SxCB0DRGroMnccoPJIV4l1YqmQsRRZkVmOvz0" 
          },
        });
        if (res) {
          alert(res.status);
        }
      } catch (err) {
        alert(err.response.data.message);
      }
    },
    async sendProduct() {
      try {
        const res = await axios({
          method: "POST",
          url: "http://localhost:8080/products",
          data: {
            name: this.name,
            link_detail: "",
            technology: this.technology,
            resolution: this.resolution,
            type: this.type,
            ListDescriptions: this.ListDescriptions,
            ListImages: this.ListImages,
            ListOptions: this.ListOptions,
          },
          headers: {
            "Access-Control-Allow-Origin": "*",
            "Content-Type": "application/json",
            token:
              "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjczNDU4MzIsImlzQWRtaW4iOnRydWUsInVzZXJJZCI6MSwidXNlcm5hbWUiOiJuZ29jaGQifQ.ZKLPL5SxCB0DRGroMnccoPJIV4l1YqmQsRRZkVmOvz0" 
          },
        });
        if (res) {
          alert(res.status);
        }
      } catch (err) {
        alert(err.response.data.message);
      }
    },
  },
};
</script>
<style>
* {
  padding: 0px;
  margin: 0px;
  font-family: sans-serif;
  box-sizing: border-box;
}
main {
  margin-top: 2rem;
  /* background-image: url(https://cdn.pixabay.com/photo/2015/09/28/21/32/the-palm-962785_1280.jpg); */
  background-repeat: no-repeat;
  background-size: cover;
  background-position: center;
  padding: 7.5px 15px;
}
.add-multi-product {
  background-color: rgba(0, 0, 0, 0.5);
  padding: 15px;
  border: none;
  border-radius: 10px;
  display: flex;
  flex-direction: column;
}

h1 {
  color: #da9d18;
  font-size: 20px;
}
.add-a-product {
  margin: 20px 0px;
  color: white;
  background-color: rgba(0, 0, 0, 0.5);
  padding: 15px;
  border: none;
  border-radius: 10px;
  display: flex;
  flex-direction: column;
}
.general-info {
  background-color: #da9d18;
  margin-top: 20px;
}
.input-box {
  /* background-color: aquamarine; */
  margin-bottom: 10px;
  margin-top: 10px;
  padding: 10px;

  display: flex;
  flex-direction: column;
  align-items: flex-start;
}
.add-a-product input {
  height: 33px;
  width: 100%;
  padding: 7.5px 15px;
  background-color: rgba(116, 113, 113, 0.5);
  border: 1px solid #cccccc;
  outline: none;
  font-size: 16px;
  color: white;
}
.add-a-product select {
  height: 33px;
  width: 100%;
  padding: 7.5px 15px;
  background-color: rgba(116, 113, 113, 0.5);
  border: 1px solid #cccccc;
  outline: none;
  font-size: 16px;
  color: white;
}
.fixed {
  display: flex;
  flex-direction: row;
}
.dynamic {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 0px 10px;
}
.view {
  height: 33px;
  width: 100%;
  padding: 7.5px 15px;
  background-color: rgba(172, 170, 170, 0.3);
  border: 1px solid #cccccc;
  outline: none;
  font-size: 16px;
  color: white;
  text-align: start;
  border-radius: 6px;
  margin: 4px 0px;
}
.detail-info {
  display: flex;
  flex-direction: row;
}
.btn-box {
  text-align: center;
  margin-top: 10px;
}
button {
  padding: 7.5px 15px;
  width: 150px;
  height: 33px;
  border-radius: 2px;
  background-color: #009999;
  color: #ffffff;
  border: none;
  outline: none;
  max-width: 4px 0px;
  border-radius: 6px;
}
button:hover {
  background-color: aqua;
  color: black;
}
.a {
  width: 87%;
}
.b {
  width: 13%;
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  align-items: center;
  padding: 20px;
}
.d {
  width: 29%;
}
.e {
  width: 140px;
}
</style>