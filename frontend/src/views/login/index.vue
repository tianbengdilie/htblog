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
              label="Account"
              v-model="loginForm.account"
              :counter="10"
            ></v-text-field>
            <v-text-field
              label="Password"
              type="password"
              v-model="loginForm.password"
              :counter="10"
            ></v-text-field>
            <v-card-actions>
              <v-btn primary large @click.stop="login">Login</v-btn>
            </v-card-actions>
          </v-form>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { mapActions } from "vuex";

export default {
  name: "loginpage",
  components: {},
  data: () => {
    return {
      loginForm: {
        account: "",
        password: ""
      }
    };
  },
  methods: {
    ...mapActions("user", ["LoginIn"]),
    login() {
      console.log("here");
      this.LoginIn(this.loginForm).then(data => {
        console.log("login success : " + data);
      });
    }
  }
};
</script>