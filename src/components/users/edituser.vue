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
                <input type="text" class="form-control" v-model="editid.name">
              </div>
            </div>
            <div class="form-group">
              <label class="col-sm-3 control-label">Tên đăng nhập</label>
              <div class="col-sm-9">
                <input type="text" class="form-control" v-model="editid.username">
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
                <input type="text" class="form-control" v-model="editid.email">
              </div>
            </div>
            <div class="form-group">
              <label class="col-sm-3 control-label">Điện thoại</label>
              <div class="col-sm-9">
                <input type="text" class="form-control" v-model="editid.phone">
              </div>
            </div>
            <div class="form-group">
              <label class="col-sm-3 control-label">Địa chỉ</label>
              <div class="col-sm-9">
                <input type="text" class="form-control" v-model="editid.address">
              </div>
            </div>
            <div class="form-group">
              <label class="col-sm-3 control-label">Phân quyền</label>
              <div class="col-sm-9">
                <select v-model="editid.type" class="form-control">
                  <option v-for="option in options" v-bind:value="option.value">
                    {{ option.text }}
                  </option>
                </select>
              </div>
            </div>
            <div class="form-group">
              <label class="col-sm-3 control-label">Trạng thái</label>
              <div class="col-sm-9">
                <select v-model="editid.status" class="form-control">
                  <option v-for="status in statuss" v-bind:value="status.value">
                    {{ status.text }}
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
    name: 'EditInvoiceType',
    data(){
      return {
          options: [
            { text: 'Nhân viên', value: 3 },
            { text: 'Người quản lý', value: 2 },
            { text: 'Quản trị viên', value: 1 }
          ],
          statuss: [
            { text: 'Hoạt động', value: 1 },
            { text: 'Tạm dừng', value: 0 }
          ]
      }
    },
    methods: {
        handleSubmit() {
                if (this.editid.username.length > 6) {
                  if (this.password === this.repassword) {
                      this.$http.post('api/user/edituser', {
                        name : this.editid.name,
                        username : this.editid.username,
                        password: this.password,
                        email: this.editid.email,
                        phone: this.editid.phone,
                        address: this.editid.address,
                        type: this.editid.type,
                        status: this.editid.status,
                        id:this.editid._id,
                      },
                      {headers: {'Authorization': `Basic ${localStorage.getItem('token')}` }}
                      )
                      .then(response => {
                        this.$emit('close');
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
    props:['editid'],
    
  };
</script>