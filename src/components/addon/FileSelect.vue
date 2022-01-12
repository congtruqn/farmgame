<template>
  <VueFileAgent
    ref="vueFileAgent"
    :theme="'list'"
    :multiple="true"
    :deletable="true"
    :meta="true"
    :accept="'image/*,.xml'"
    :maxSize="'10MB'"
    :maxFiles="14"
    :helpText="'Choose images or zip files'"
    :errorText="{
      type: 'Invalid file type. Only images or zip Allowed',
      size: 'Files should not exceed 10MB in size',
    }"
    @select="filesSelected($event)"
    @beforedelete="onBeforeDelete($event)"
    @delete="fileDeleted($event)"
    v-model="fileRecords"
  ></VueFileAgent>
</template>
<script>
  export default {
    data: function () {
      return {
        fileRecords: [],
        uploadUrl: 'https://www.mocky.io/v2/5d4fb20b3000005c111099e3',
        uploadHeaders: { 'X-Test-Header': 'vue-file-agent' },
        fileRecordsForUpload: [], // maintain an upload queue
      };
    },
    methods: {
      uploadFiles: function () {
        // Using the default uploader. You may use another uploader instead.
        this.$refs.vueFileAgent.upload(this.uploadUrl, this.uploadHeaders, this.fileRecordsForUpload);
        this.fileRecordsForUpload = [];
      },
      deleteUploadedFile: function (fileRecord) {
        // Using the default uploader. You may use another uploader instead.
        this.$refs.vueFileAgent.deleteUpload(this.uploadUrl, this.uploadHeaders, fileRecord);
      },
      filesSelected: function (fileRecordsNewlySelected) {
          
        var validFileRecords = fileRecordsNewlySelected.filter((fileRecord) => !fileRecord.error);
        console.log(validFileRecords);
        this.fileRecordsForUpload = this.fileRecordsForUpload.concat(validFileRecords);
      },
      onBeforeDelete: function (fileRecord) {
        var i = this.fileRecordsForUpload.indexOf(fileRecord);
        if (i !== -1) {
        // queued file, not yet uploaded. Just remove from the arrays
          this.fileRecordsForUpload.splice(i, 1);
          var k = this.fileRecords.indexOf(fileRecord);
          if (k !== -1) this.fileRecords.splice(k, 1);
        } else {
          if (confirm('Are you sure you want to delete?')) {
            this.$refs.vueFileAgent.deleteFileRecord(fileRecord); // will trigger 'delete' event
          }
        }
      },
      fileDeleted: function (fileRecord) {
        var i = this.fileRecordsForUpload.indexOf(fileRecord);
        if (i !== -1) {
          this.fileRecordsForUpload.splice(i, 1);
        } else {
          this.deleteUploadedFile(fileRecord);
        }
      },
    },
  };
</script>