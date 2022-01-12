<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div>
            <div class="panel-body_modal">
              <button aria-label="Dismiss" class="ngdialog-close" @click="$emit('close')"></button>
            </div>
            <div class="form-group">
              <label class="col-sm-3 control-label">Họ và Tên</label>
              <div class="col-sm-9">
                <input type="text" class="form-control" v-model="name">
              </div>
            </div>
            <div class="form-group">
              <label class="col-sm-3 control-label">Tên đăng nhập</label>
              <div class="col-sm-9">
                <input type="text" class="form-control" v-model="username">
              </div>
            </div>
            <div class="form-group">
              <label class="col-sm-3 control-label">Mật khẩu</label>
              <div class="col-sm-9">
                <input type="password" class="form-control" v-model="password">
              </div>
            </div>
            <div class="form-group">
              <label class="col-sm-3 control-label">Nhập lại mật khẩu</label>
              <div class="col-sm-9">
                <input type="password" class="form-control" v-model="repassword">
              </div>
            </div>
            <div class="form-group">
              <label class="col-sm-3 control-label">Email</label>
              <div class="col-sm-9">
                <input type="text" class="form-control" v-model="email">
              </div>
            </div>
            <div class="form-group">
              <label class="col-sm-3 control-label">Điện thoại</label>
              <div class="col-sm-9">
                <input type="text" class="form-control" v-model="phone">
              </div>
            </div>
            <div class="form-group">
              <label class="col-sm-3 control-label">Địa chỉ</label>
              <div class="col-sm-9">
                <input type="text" class="form-control" v-model="address">
              </div>
            </div>
            <div class="form-group">
              <label class="col-sm-3 control-label">Phân quyền</label>
              <div class="col-sm-9">
                <select v-model="type" class="form-control">
                  <option v-for="option in options" v-bind:value="option.value">
                    {{ option.text }}
                  </option>
                </select>
              </div>
            </div> 
            <div class="panel-body_modal">
              <button class="btn btn-primary m-r-5 m-b-5" @click="handleSubmit">Lưu</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>
<script>
  export default {
    name: 'adduser',
    data(){
      return {
          name : "",
          username : "",
          password: "",
          repassword: "",
          email: "",
          phone: "",
          address: "",
          type: 4,
          options: [
            { text: 'Nhân viên', value: 3 },
            { text: 'Người quản lý', value: 2 },
            { text: 'Quản trị viên', value: 1 }
          ],
      }
    },
    methods: {
      close() {
        this.$emit('close');
      },
      handleSubmit(e){
                e.preventDefault()
                if (this.username.length > 6) {
                  if (this.password === this.repassword) {
                    this.$emit('close');
                    this.$http.post('api/user/adduser', {
                      name : this.name,
                      username : this.username,
                      password: this.password,
                      email: this.email,
                      phone: this.phone,
                      address: this.address,
                      type: this.type
                    },
                    {headers: {'Authorization': `Basic ${localStorage.getItem('token')}` }}
                    )
                    .then(response => {
                      this.name = "";
                      this.username = "";
                      this.password = "";
                      this.repassword = "";
                      this.email = "";
                      this.phone = "";
                      this.address = "";
                      this.$http.get('api/user/getallusers?page='+this.$route.query.page,{headers: {'Authorization': `Basic ${localStorage.getItem('token')}` }})
                        .then(response => (this.$emit("inputData", response.data)))
                    })
                    .catch(function (error) {
                        console.error(error.response);
                    });
                  }
                  else{
                    alert("Mật khẩu phải giống nhau!")
                  }
 
                }
                else{
                  alert("Tên đăng nhập quá ngắn!")
                }
            }
    },
    mounted () {
      this.$http.get('api/projects/getallprojectnotpage',{headers: {'Authorization': `Basic ${localStorage.getItem('token')}` }})
        .then(response => (this.listprojects = response.data));
    },
  };
</script>