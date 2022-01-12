<template>
    <div class="container1">
      <div>
        <div class="form_imput">
          <form action="#" method="POST">
            <input type="file" @change="loadTextFromFile" accept=".xml">
            <p>Nhấp chuột để chọn hóa đơn.</p>
          </form>
        </div>
        <div>
            <b>Lưu ý:</b> Trường hợp có vướng mắc về nội dung đọc hóa đơn hoặc hiển thị thông tin hóa đơn, Anh/Chị vui lòng gửi thông tin qua Email: <b>trutc@smartsign.com.vn</b> hoặc liên hệ qua số Hotline: <b>0909052248</b> để được hỗ trợ.
        </div>
        <div class="verify_info">
          <div>
            <p v-if="isValid"><i class="fa fa-check-circle text-green" aria-hidden="true"></i> <b>Nội dung hóa đơn còn nguyên vẹn</b></p>
            <p v-if="isNotValid"><i class="fa fa-exclamation-triangle text-orange" aria-hidden="true"></i> <b>Hóa đơn đã bị chỉnh sửa</b></p>
            <p v-if="isNotPubicKey"><i class="fa fa-exclamation-triangle text-orange" aria-hidden="true"></i> <b>Không tìm thấy chữ ký số trên hóa đơn</b></p>
          </div>
          <div>
            <p v-html="KHMSHDon"></p>
          </div>
        </div>
        <div class="about_us">
          <h2>Giới thiệu chương trình đọc hóa đơn điện tử xml</h2>
          <p>Hiện nay trên thị trường có rất nhiều nhà cung cấp hóa đơn điện tử khác nhau, theo quy định của bộ thông tin truyền thông thì hóa đơn điện tử được lưu dưới định dạng xml chỉ có giá trị khi được ký bằng chữ ký số của doanh nghiệp phát hành hóa đơn
            . Hóa đơn được lưu dưới dạng xml để đảm bảo dung lượng lưu trữ (chiếm ít dung lương bộ nhớ) và tính xác thực (Có thể ký bằng chữ ký số).</p>
          <p>Tuy nhiên hóa đơn điện tử xml cũng gây khó khăn về xem hóa đơn và xác thực hóa đơn trên file hóa đơn xml. Chương trình đọc hóa đơn xml của SmartSign sẽ giúp bạn có thể kiểm tra chữ ký số và kiểm tra hóa đơn 1 cách dễ dàng hơn</p>
        </div>
      </div>
    </div>
</template>
<script>
</script>