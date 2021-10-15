<template>
  <el-form ref="form" :model="form" label-width="80px" style="height: 100%">
    <el-form-item label="镜像地址">
      <el-input v-model="form.DockerImage"></el-input>
    </el-form-item>
    <el-form-item label="探测端口">
      <el-input v-model="form.LivenessProbePort"></el-input>
    </el-form-item>
    <el-form-item label="正确结果">
      <el-input v-model="form.LivenessProbeResult"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submitForm()">立即创建</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
export default {
  data () {
    return {
      form: {
        DockerImage: '',
        LivenessProbePort: '',
        LivenessProbeResult: '',
        date2: '',
        delivery: false,
        type: [],
        resource: '',
        desc: ''
      }
    }
  },
  methods: {
    onSubmit () {
      console.log('submit!')
    }
  }
}
</script>
