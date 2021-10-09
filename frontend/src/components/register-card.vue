<template>
  <v-card flat>
    <v-card-title primary-title>
      <h4>Register</h4>
    </v-card-title>
    <v-card-text>
      <v-form ref="loginForm">
        <v-text-field
          label="Account"
          v-model="loginForm.account"
          :counter="10"
        ></v-text-field>
        <v-text-field
          label="NickName"
          v-model="loginForm.nickname"
          :counter="10"
        ></v-text-field>
        <v-text-field
          label="Password"
          type="password"
          v-model="loginForm.password"
          :counter="10"
        ></v-text-field>
        <v-text-field
          label="Repeat Password"
          type="password"
          v-model="loginForm.password2"
          :counter="10"
        ></v-text-field>
      </v-form>
    </v-card-text>
    <v-card-actions class="justify-space-around">
      <v-btn large @click.stop="onClickRegister">Register</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import { register } from "@/api/user";

export default {
  name: "register-card",
  components: {},
  props: {},
  data: () => {
    return {
      loginForm: {
        nickname: "dev",
        account: "dev",
        password: "dev",
        password2: "dev"
      }
    };
  },
  methods: {
    onClickRegister() {
      register(this.loginForm).then(data => {
        console.log("register success : " + data);
        this.$emit("completed");
      });
    }
  }
};
</script>