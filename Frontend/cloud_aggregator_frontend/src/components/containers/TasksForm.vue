<template>
  <el-form ref="form" :model="form" label-width="200px" style="height: 100%">
    <el-form-item label="image">
      <el-input v-model="form.DockerImage"></el-input>
    </el-form-item>
    <el-form-item label="livenessPort">
      <el-input v-model="form.LivenessProbePort"></el-input>
    </el-form-item>
    <el-form-item label="LivenessResult">
      <el-input v-model="form.LivenessProbeResult"></el-input>
    </el-form-item>
    <el-form-item label="Path">
      <el-input v-model="form.Path"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submitForm()">submit</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import axios from 'axios'

export default {
  data () {
    return {
      form: {
        DockerImage: '',
        LivenessProbePort: '',
        LivenessProbeResult: '',
        Path: ''
      }
    }
  },
  methods: {
    submitForm () {
      let val = this.decimalToHexString(10000000000000000)
      console.log(val)
      window.ethereum.request({
        method: 'eth_sendTransaction',
        params: [
          {
            from: this.Global.AccountAddress,
            // Aggregator Address
            to: '0xb72493Fe5b655882d355Ba548F6Cf56BF279062B',
            value: val, // 0.01ETH
            gasPrice: this.decimalToHexString(1000000000),
            gas: this.decimalToHexString(4200000)
          }
        ]
      })
        .then((txHash) => {
          console.log(txHash)
          // 调用后端接口创建新的任务
          let body = {
            'user_address': this.Global.AccountAddress,
            'docker_image': this.form.DockerImage,
            'port': this.form.LivenessProbePort,
            'flag_message': this.form.LivenessProbeResult,
            'url': this.form.Path
          }
          axios.post('/api/v1/transaction', body)
        })
        .catch((error) => console.log(error))
    },
    decimalToHexString (number) {
      if (number < 0) {
        number = 0xFFFFFFFF + number + 1
      }

      return number.toString(16).toUpperCase()
    }
  }
}
</script>
