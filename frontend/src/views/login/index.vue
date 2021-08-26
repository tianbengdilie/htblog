<template>
  <v-container fuild>
    <v-row class="fill-height" justify="center" align="center">
      <v-col col="12" class="fill-height">
        <v-card flat>
          <v-card-title primary-title>
            <h4>Login</h4>
          </v-card-title>
          <v-form ref="loginForm">
            <v-text-field
              label="Username"
              v-model="username"
              :counter="10"
              :rules="nameRules"
              required
            ></v-text-field>
            <v-text-field
              label="Password"
              type="password"
              v-model="password"
              :counter="10"
              :rules="pwdRules"
              required
            ></v-text-field>
            <v-card-actions>
              <v-btn primary large block @click.stop="login">Login</v-btn>
            </v-card-actions>
          </v-form>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { setToken } from "@/utils/auth.js";

export default {
  name: "loginpage",
  components: {},
  data: () => {
    return {
      username: "",
      nameRules: [
        v => !!v || "Name is required",
        v => (v && v.length <= 10) || "Name must be less than 10 characters"
      ],
      password: "",
      pwdRules: [
        v => !!v || "Password is required",
        v => (v && v.length <= 10) || "Password must be less than 10 characters"
      ]
    };
  },
  methods: {
    login: function() {
      if (this.$refs.loginForm.validate()) {
        console.log("login succeed");
        setToken(this.username);
        // TODO 调用后台接口, 判断权限
      } else {
        console.log("login fail");
      }
    }
  }
};
</script>