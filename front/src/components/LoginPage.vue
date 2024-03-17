<template>
  <div class="ring">
    <i style="--clr: #00ff0a"></i><i style="--clr: #ff0057"></i
    ><i style="--clr: #fffd44"></i>
    <div class="login">
      <h2>登录</h2>
      <div class="inputBx">
        <input type="username" v-model="username" placeholder="用户名" />
      </div>
      <div class="inputBx">
        <input type="password" v-model="password" placeholder="密码" />
      </div>
      <div class="inputBx">
        <input type="submit" @click="login" value="确认进入" />
      </div>

      <div class="text-danger" id="">{{ errormsg }}</div>
      <!-- <div class="links"><a href="#">忘记密码</a><a href="#">注册</a></div> -->
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "LoginPage",
  data() {
    return {
      username: "",
      password: "",
      errormsg: "",
    };
  },
  methods: {
    login() {
      axios
        .post(
          process.env.VUE_APP_DOMAIN + "/api/public/login",
          {
            username: this.username,
            password: this.password,
          },
          {
            withCredentials: true,
          }
        )
        .then((response) => {
          if (response.data.code === 200) {
            // 登录成功，跳转到/chat页面
            this.$router.push("/chat");
          } else {
            console.log(response.data.msg);
            this.errormsg = response.data.msg;
          }
        })
        .catch((error) => {
          this.errormsg = "网络异常，请稍后重试";
          console.log(error);
        });
    },
  },
};
</script>

<style scoped>
@import "../assets/css/login-page.css";
</style>
