
<template>
    <div>
    <section class="content-header">
            <h1>Quản lý người dùng</h1>
            <ol class="breadcrumb">
                <li><a><i class="fa fa-dashboard"></i>Trang chủ</a></li>
                <li><a>Danh sách người dùng</a></li>
            </ol>
        </section>
        <section class="panel panel-inverse">
			<div class="row">
				<div class="panel-body">
					<button class="btn btn-primary m-r-5 m-b-5"  @click="showModal">Thêm người dùng</button>
				</div>
				
					<table class="table">
						<thead>
		    				<tr>
		        				<th width="25px">STT</th>
		            			<th>Tên người dùng</th>
		            			<th>Username</th>
		            			<th>Email</th>
                      <th>Điện thoại</th>
                      <th>Phân quyền</th>
                      <th width="15px" class="functionicon"></th>
                      <th width="15px" class="functionicon"></th>
		        			</tr>
						</thead>
						<tbody>
		    			<tr v-for="(invoicetype, index) in listinvoicetypes" :key="index">
                <td>{{((page-1)*20)+index+1}}</td>
		        				<td>{{invoicetype.name}}</td>
		        				<td>{{invoicetype.username}}</td>
		        				<td>{{invoicetype.email}}</td>	
                    <td>{{invoicetype.phone}}</td>
                    <td>{{displayUserType(invoicetype.type)}}</td>
								<td>
									<span>
		                <a href="#" @click="showisModalEditVisible(invoicetype)">
		                    <i class="fa fa-edit"></i>
		                </a>
									</span>
		        		</td>
								<td>
		            	<a href="#" @click="dellInvoiceType(invoicetype._id,invoicetype.toquote_name)">
		                <i class="fa fa-trash"></i>
		            	</a>
		        		</td>
		    			</tr>
						</tbody>
    				</table>
			</div>
			<div id="example2_paginate" class="dataTables_paginate paging_simple_numbers">
  <paginate
    v-model="page"
    :page-count="countinvoicetype/20"
    :page-range="3"
    :margin-pages="2"
    :click-handler="clickCallback"
    :prev-text="'Prev'"
    :next-text="'Next'"
    :container-class="'pagination'"
    :page-class="'page-item'">
  </paginate>
    	</div>
    <adduser @inputData="updateMessage"
      v-show="isModalVisible"
      @close="closeModal"
    />
    <edituser
      :editid="editid"
      v-show="isModalEditVisible"
      @close="closeEditModal"
    />
		</section>	
    <vue-confirm-dialog></vue-confirm-dialog>
    </div>
    
    		
</template>
<script>
import adduser from '@/components/users/adduser';
import edituser from '@/components/users/edituser';

export default {
  data() {
    return {
      isModalVisible: false,
      isModalEditVisible:false,
      listinvoicetypes:'',
      countinvoicetype:'',
      page:Number(this.$route.query.page),
      editid: '',
    };
  },
	components: {
        adduser,
        edituser,
  },
  methods: {
      showModal() {
        this.isModalVisible = true;
      },
      showisModalEditVisible(id) {
        this.editid = id;
        this.isModalEditVisible = true;
      },
      closeModal() {
        this.isModalVisible = false;
      },
      displayUserType(id) {
        switch(id) {
          case 1:
            return 'Quản trị viên'
            break;
          case 2:
            return 'Người quản lý'
            break;
          default:
            return 'Nhân viên'
        }
      },
      closeEditModal() {
        this.isModalEditVisible = false;
      },
      clickCallback (pageNum){
        this.$http.get('api/user/getalluser?page='+pageNum,{headers: {'Authorization': `Basic ${localStorage.getItem('token')}` }})
        .then(response => (this.listinvoicetypes = response.data))
      },
      updateMessage(variable) {
        this.listinvoicetypes = variable;
      },
      dellInvoiceType(id,name){
        this.$confirm(
        {
          message:'Bạn có muốn xóa loại hóa đơn ' + name,
          button: {
            yes: 'Đồng ý'
          },
          callback: confirm => {
            if (confirm) {
              this.$http.post('api/toquote/delltoquote', {
                id:id,
              },
              {headers: {'Authorization': `Basic ${localStorage.getItem('token')}` }}
              ).then(response => {
                this.$http.get('api/user/getalluser?page='+this.$route.query.page,{headers: {'Authorization': `Basic ${localStorage.getItem('token')}` }})
                .then(response => (this.listinvoicetypes = response.data));
                    this.$http.get('api/user/getcountuser',{headers: {'Authorization': `Basic ${localStorage.getItem('token')}` }})
                .then(response => (this.countinvoicetype = response.data));
              }).catch(function (error) {
                console.error(error.response);
              })
            }
          }
        })

      }
  },
  mounted () {
    if(this.$route.query.page===undefined){
      this.page =1;
    }
    this.$http.get('api/user/getalluser?page='+this.$route.query.page,{headers: {'Authorization': `Basic ${localStorage.getItem('token')}` }})
      .then(response => (this.listinvoicetypes = response.data));
      this.$http.get('api/user/getcountuser',{headers: {'Authorization': `Basic ${localStorage.getItem('token')}` }})
        .then(response => (this.countinvoicetype = response.data.numofcompany));
  },
};
</script>