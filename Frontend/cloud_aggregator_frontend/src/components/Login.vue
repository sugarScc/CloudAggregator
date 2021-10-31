<template>
  <el-dialog
    title="提示"
    :visible.sync="EnableLogin"
    width="60%"
    :before-close="handleClose">
    <el-form ref="form" :model="form" size="small" label-width="100px">
      <el-row>
        <el-form-item label="Address">
          <el-input class="search-input" placeholder="请输入钱包地址" v-model="form.id"/>
        </el-form-item>
      </el-row>
    </el-form>
    <span slot="footer" class="dialog-footer">
    <el-button @click="EnableLogin = false">取 消</el-button>
    <el-button type="primary" @click="EnableLogin = false">确 定</el-button>
  </span>
  </el-dialog>
</template>

<script>

export default {
  name: 'Login',
  data () {
    return {
      EnableLogin: null,
      form: {
        Address: ''
      }
    }
  },
  methods: {
    handleClose (done) {
      this.$confirm('确认关闭？')
        .then(_ => {
          done()
        })
        .catch(_ => {
        })
    }
  }
}
</script>
